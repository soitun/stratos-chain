package mempool

import (
	"context"
	crand "crypto/rand" // #nosec // crypto/rand is used for seed generation
	"encoding/binary"
	"fmt"
	"math/rand" // #nosec // math/rand is used for random selection and seeded from crypto/rand

	"github.com/huandu/skiplist"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/cosmos/cosmos-sdk/types/mempool"

	evmtypes "github.com/stratosnet/stratos-chain/x/evm/types"
)

var (
	_ mempool.Mempool  = (*EvmSenderNonceMempool)(nil)
	_ mempool.Iterator = (*evmSenderNonceMempoolIterator)(nil)
)

var DefaultMaxTx = 0

// EvmSenderNonceMempool is a mempool that prioritizes transactions within a sender
// by nonce, the lowest first, but selects a random sender on each iteration.
// Additionality it supports evm custom message handling.
// The mempool is iterated by:
//
// 1) Maintaining a separate list of nonce ordered txs per sender
// 2) For each select iteration, randomly choose a sender and pick the next nonce ordered tx from their list
// 3) Repeat 1,2 until the mempool is exhausted
//
// Note that PrepareProposal could choose to stop iteration before reaching the
// end if maxBytes is reached.
type EvmSenderNonceMempool struct {
	senders    map[string]*skiplist.SkipList
	rnd        *rand.Rand
	maxTx      int
	existingTx map[txKey]bool
}

type EvmSenderNonceOptions func(mp *EvmSenderNonceMempool)

type txKey struct {
	address string
	nonce   uint64
}

// NewEvmSenderNonceMempool creates a new mempool that prioritizes transactions by
// nonce, the lowest first, picking a random sender on each iteration.
func NewEvmSenderNonceMempool(opts ...EvmSenderNonceOptions) *EvmSenderNonceMempool {
	senderMap := make(map[string]*skiplist.SkipList)
	existingTx := make(map[txKey]bool)
	snp := &EvmSenderNonceMempool{
		senders:    senderMap,
		maxTx:      DefaultMaxTx,
		existingTx: existingTx,
	}

	var seed int64
	err := binary.Read(crand.Reader, binary.BigEndian, &seed)
	if err != nil {
		panic(err)
	}

	snp.setSeed(seed)

	for _, opt := range opts {
		opt(snp)
	}

	return snp
}

// SenderNonceSeedOpt Option To add a Seed for random type when calling the
// constructor NewEvmSenderNonceMempool.
//
// Example:
//
//	random_seed := int64(1000)
//	NewEvmSenderNonceMempool(SenderNonceSeedTxOpt(random_seed))
func SenderNonceSeedOpt(seed int64) EvmSenderNonceOptions {
	return func(snp *EvmSenderNonceMempool) {
		snp.setSeed(seed)
	}
}

// SenderNonceMaxTxOpt Option To set limit of max tx when calling the constructor
// NewEvmSenderNonceMempool.
//
// Example:
//
//	NewEvmSenderNonceMempool(SenderNonceMaxTxOpt(100))
func SenderNonceMaxTxOpt(maxTx int) EvmSenderNonceOptions {
	return func(snp *EvmSenderNonceMempool) {
		snp.maxTx = maxTx
	}
}

func (snm *EvmSenderNonceMempool) setSeed(seed int64) {
	s1 := rand.NewSource(seed)
	snm.rnd = rand.New(s1) //#nosec // math/rand is seeded from crypto/rand by default
}

// NextSenderTx returns the next transaction for a given sender by nonce order,
// i.e. the next valid transaction for the sender. If no such transaction exists,
// nil will be returned.
func (mp *EvmSenderNonceMempool) NextSenderTx(sender string) sdk.Tx {
	senderIndex, ok := mp.senders[sender]
	if !ok {
		return nil
	}

	cursor := senderIndex.Front()
	return cursor.Value.(sdk.Tx)
}

// Insert adds a tx to the mempool. It returns an error if the tx does not have
// at least one signer. Note, priority is ignored.
func (snm *EvmSenderNonceMempool) Insert(_ context.Context, tx sdk.Tx) error {
	if snm.maxTx > 0 && snm.CountTx() >= snm.maxTx {
		return mempool.ErrMempoolTxMaxCapacity
	}
	if snm.maxTx < 0 {
		return nil
	}

	msgs := tx.GetMsgs()
	if len(msgs) == 0 {
		return fmt.Errorf("tx should have at least one msg")
	}

	var (
		sender string
		nonce  uint64
	)

	if ethMsg, ok := msgs[0].(*evmtypes.MsgEthereumTx); ok {
		ethTx := ethMsg.AsTransaction()
		sender = ethMsg.From
		nonce = ethTx.Nonce()
	} else {
		sigs, err := tx.(signing.SigVerifiableTx).GetSignaturesV2()
		if err != nil {
			return err
		}
		if len(sigs) == 0 {
			return fmt.Errorf("tx must have at least one signer")
		}

		sig := sigs[0]
		sender = sdk.AccAddress(sig.PubKey.Address()).String()
		nonce = sig.Sequence
	}

	senderTxs, found := snm.senders[sender]
	if !found {
		senderTxs = skiplist.New(skiplist.Uint64)
		snm.senders[sender] = senderTxs
	}

	senderTxs.Set(nonce, tx)

	key := txKey{nonce: nonce, address: sender}
	snm.existingTx[key] = true

	return nil
}

// Select returns an iterator ordering transactions the mempool with the lowest
// nonce of a random selected sender first.
//
// NOTE: It is not safe to use this iterator while removing transactions from
// the underlying mempool.
func (snm *EvmSenderNonceMempool) Select(_ context.Context, _ [][]byte) mempool.Iterator {
	var senders []string

	senderCursors := make(map[string]*skiplist.Element)
	orderedSenders := skiplist.New(skiplist.String)

	// #nosec
	for s := range snm.senders {
		orderedSenders.Set(s, s)
	}

	s := orderedSenders.Front()
	for s != nil {
		sender := s.Value.(string)
		senders = append(senders, sender)
		senderCursors[sender] = snm.senders[sender].Front()
		s = s.Next()
	}

	iter := &evmSenderNonceMempoolIterator{
		senders:       senders,
		rnd:           snm.rnd,
		senderCursors: senderCursors,
	}

	return iter.Next()
}

// CountTx returns the total count of txs in the mempool.
func (snm *EvmSenderNonceMempool) CountTx() int {
	return len(snm.existingTx)
}

// Remove removes a tx from the mempool. It returns an error if the tx does not
// have at least one signer or the tx was not found in the pool.
func (snm *EvmSenderNonceMempool) Remove(tx sdk.Tx) error {
	msgs := tx.GetMsgs()
	if len(msgs) == 0 {
		return fmt.Errorf("tx should have at least one msg")
	}

	var (
		sender string
		nonce  uint64
	)

	if ethMsg, ok := msgs[0].(*evmtypes.MsgEthereumTx); ok {
		ethTx := ethMsg.AsTransaction()
		sender = ethMsg.From
		nonce = ethTx.Nonce()
	} else {
		sigs, err := tx.(signing.SigVerifiableTx).GetSignaturesV2()
		if err != nil {
			return err
		}
		if len(sigs) == 0 {
			return fmt.Errorf("tx must have at least one signer")
		}

		sig := sigs[0]
		sender = sdk.AccAddress(sig.PubKey.Address()).String()
		nonce = sig.Sequence
	}

	senderTxs, found := snm.senders[sender]
	if !found {
		return mempool.ErrTxNotFound
	}

	res := senderTxs.Remove(nonce)
	if res == nil {
		return mempool.ErrTxNotFound
	}

	if senderTxs.Len() == 0 {
		delete(snm.senders, sender)
	}

	key := txKey{nonce: nonce, address: sender}
	delete(snm.existingTx, key)

	return nil
}

type evmSenderNonceMempoolIterator struct {
	rnd           *rand.Rand
	currentTx     *skiplist.Element
	senders       []string
	senderCursors map[string]*skiplist.Element
}

// Next returns the next iterator state which will contain a tx with the next
// smallest nonce of a randomly selected sender.
func (i *evmSenderNonceMempoolIterator) Next() mempool.Iterator {
	for len(i.senders) > 0 {
		senderIndex := i.rnd.Intn(len(i.senders))
		sender := i.senders[senderIndex]
		senderCursor, found := i.senderCursors[sender]
		if !found {
			i.senders = removeAtIndex(i.senders, senderIndex)
			continue
		}

		if nextCursor := senderCursor.Next(); nextCursor != nil {
			i.senderCursors[sender] = nextCursor
		} else {
			i.senders = removeAtIndex(i.senders, senderIndex)
		}

		return &evmSenderNonceMempoolIterator{
			senders:       i.senders,
			currentTx:     senderCursor,
			rnd:           i.rnd,
			senderCursors: i.senderCursors,
		}
	}

	return nil
}

func (i *evmSenderNonceMempoolIterator) Tx() sdk.Tx {
	return i.currentTx.Value.(sdk.Tx)
}

func removeAtIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
