package whitelist

import (
	"github.com/xcareteam/xci/accounts/abi/bind"
	"github.com/xcareteam/xci/common"
	"github.com/xcareteam/xci/eth"
	"github.com/xcareteam/xci/contracts/whitelist/contract"
	"github.com/xcareteam/xci/core/types"
	"github.com/xcareteam/xci/node"
	"github.com/xcareteam/xci/internal/ethapi"
	"github.com/xcareteam/xci/les"
)

var (
	MainNetAddress = common.HexToAddress("0x314159265dD8dbb310642f98f50C066173C1259b")
	TestNetAddress = common.HexToAddress("0x112234455c3a32fd11230c42e7bccd4a84e02010")
)

type WhiteList struct {
	*contract.WhitelistSession
	contractBackend bind.ContractBackend
}

func NewWhiteList(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*WhiteList, error) {
	whitelist, err := contract.NewWhitelist(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &WhiteList{
		&contract.WhitelistSession{
			Contract:     whitelist,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func (self *WhiteList) AddNewNode(enode string, DIDJson string) (*types.Transaction, error) {
	return self.AddNewNode(enode, DIDJson)
}

func (self *WhiteList) GetDID(enode string) (string, error) {
	return self.GetDID(enode);
}

