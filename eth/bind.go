package eth

import (
	"context"
	"math/big"

	"github.com/xcareteam/xci"
	"github.com/xcareteam/xci/common"
	"github.com/xcareteam/xci/common/hexutil"
	"github.com/xcareteam/xci/core/types"
	"github.com/xcareteam/xci/internal/ethapi"
	"github.com/xcareteam/xci/rlp"
	"github.com/xcareteam/xci/rpc"
)

// ContractBackend implements bind.ContractBackend with direct calls to Ethereum
// internals to support operating on contracts within subprotocols like eth and
// swarm.
//
// Internally this backend uses the already exposed API endpoints of the Ethereum
// object. These should be rewritten to internal Go method calls when the Go API
// is refactored to support a clean library use.
type ContractBackend struct {
	eapi  *ethapi.PublicEthereumAPI        // Wrapper around the Ethereum object to access metadata
	bcapi *ethapi.PublicBlockChainAPI      // Wrapper around the blockchain to access chain data
	txapi *ethapi.PublicTransactionPoolAPI // Wrapper around the transaction pool to access transaction data
}

// NewContractBackend creates a new native contract backend using an existing
// Ethereum object.
func NewContractBackend(apiBackend ethapi.Backend) ContractBackend {
	return ContractBackend{
		eapi:  ethapi.NewPublicEthereumAPI(apiBackend),
		bcapi: ethapi.NewPublicBlockChainAPI(apiBackend),
		txapi: ethapi.NewPublicTransactionPoolAPI(apiBackend, new(ethapi.AddrLocker)),
	}
}

// CodeAt retrieves any code associated with the contract from the local API.
func (b *ContractBackend) CodeAt(ctx context.Context, contract common.Address, blockNum *big.Int) ([]byte, error) {
	return b.bcapi.GetCode(ctx, contract, toBlockNumber(blockNum))
}

// CodeAt retrieves any code associated with the contract from the local API.
func (b *ContractBackend) PendingCodeAt(ctx context.Context, contract common.Address) ([]byte, error) {
	return b.bcapi.GetCode(ctx, contract, rpc.PendingBlockNumber)
}

// ContractCall implements bind.ContractCaller executing an Ethereum contract
// call with the specified data as the input. The pending flag requests execution
// against the pending block, not the stable head of the chain.
func (b *ContractBackend) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNum *big.Int) ([]byte, error) {
	out, err := b.bcapi.Call(ctx, toCallArgs(msg), toBlockNumber(blockNum))
	return out, err
}

// ContractCall implements bind.ContractCaller executing an Ethereum contract
// call with the specified data as the input. The pending flag requests execution
// against the pending block, not the stable head of the chain.
func (b *ContractBackend) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	out, err := b.bcapi.Call(ctx, toCallArgs(msg), rpc.PendingBlockNumber)
	return out, err
}

func toCallArgs(msg ethereum.CallMsg) ethapi.CallArgs {
	args := ethapi.CallArgs{
		To:   msg.To,
		From: msg.From,
		Data: msg.Data,
	}
	if msg.Gas != 0 {
		args.Gas = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		args.GasPrice = hexutil.Big(*msg.GasPrice)
	}
	if msg.Value != nil {
		args.Value = hexutil.Big(*msg.Value)
	}
	return args
}

func toBlockNumber(num *big.Int) rpc.BlockNumber {
	if num == nil {
		return rpc.LatestBlockNumber
	}
	return rpc.BlockNumber(num.Int64())
}

// PendingAccountNonce implements bind.ContractTransactor retrieving the current
// pending nonce associated with an account.
func (b *ContractBackend) PendingNonceAt(ctx context.Context, account common.Address) (nonce uint64, err error) {
	out, err := b.txapi.GetTransactionCount(ctx, account, rpc.PendingBlockNumber)
	if out != nil {
		nonce = uint64(*out)
	}
	return nonce, err
}

// SuggestGasPrice implements bind.ContractTransactor retrieving the currently
// suggested gas price to allow a timely execution of a transaction.
func (b *ContractBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return b.eapi.GasPrice(ctx)
}

// EstimateGasLimit implements bind.ContractTransactor triing to estimate the gas
// needed to execute a specific transaction based on the current pending state of
// the backend blockchain. There is no guarantee that this is the true gas limit
// requirement as other transactions may be added or removed by miners, but it
// should provide a basis for setting a reasonable default.
func (b *ContractBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (*big.Int, error) {
	out, err := b.bcapi.EstimateGas(ctx, toCallArgs(msg))
	return out.ToInt(), err
}

// SendTransaction implements bind.ContractTransactor injects the transaction
// into the pending pool for execution.
func (b *ContractBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	raw, _ := rlp.EncodeToBytes(tx)
	_, err := b.txapi.SendRawTransaction(ctx, raw)
	return err
}


// FilterLogs executes a log filter operation, blocking during execution and
// returning all the results in one batch.
func (b *ContractBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	//// Initialize unset filter boundaried to run from genesis to chain head
	//from := int64(0)
	//if query.FromBlock != nil {
	//	from = query.FromBlock.Int64()
	//}
	//to := int64(-1)
	//if query.ToBlock != nil {
	//	to = query.ToBlock.Int64()
	//}
	//// Construct and execute the filter
	//filter := filters.New(&filterBackend{b.database, b.blockchain}, from, to, query.Addresses, query.Topics)
	//
	//logs, err := filter.Logs(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//res := make([]types.Log, len(logs))
	//for i, log := range logs {
	//	res[i] = *log
	//}
	return nil, nil
}

// SubscribeFilterLogs creates a background log filtering operation, returning a
// subscription immediately, which can be used to stream the found events.
func (b *ContractBackend) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	//// Subscribe to contract events
	//sink := make(chan []*types.Log)
	//
	//sub, err := b.events.SubscribeLogs(query, sink)
	//if err != nil {
	//	return nil, err
	//}
	//// Since we're getting logs in batches, we need to flatten them into a plain stream
	//return event.NewSubscription(func(quit <-chan struct{}) error {
	//	defer sub.Unsubscribe()
	//	for {
	//		select {
	//		case logs := <-sink:
	//			for _, log := range logs {
	//				select {
	//				case ch <- *log:
	//				case err := <-sub.Err():
	//					return err
	//				case <-quit:
	//					return nil
	//				}
	//			}
	//		case err := <-sub.Err():
	//			return err
	//		case <-quit:
	//			return nil
	//		}
	//	}
	//}), nil
	return nil, nil
}
