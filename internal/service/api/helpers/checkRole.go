package helpers

import (
	"github.com/dl-nft-books/core-svc/solidity/generated/contractsregistry"
	"github.com/dl-nft-books/core-svc/solidity/generated/rolemanager"
	networker "github.com/dl-nft-books/network-svc/connector"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func CheckMarketplacePerrmision(networker networker.Connector, address string) (bool, error) {
	networks, err := networker.GetNetworksDetailed()
	if err != nil {
		return false, errors.Wrap(err, "failed to get networks")
	}
	for _, net := range networks.Data {
		contractRegistry, err := contractsregistry.NewContractsregistry(common.HexToAddress(net.FactoryAddress), net.RpcUrl)
		if err != nil {
			return false, errors.Wrap(err, "failed to create new contracts registry")
		}
		roleManagerContract, err := contractRegistry.GetRoleManagerContract(nil)
		if err != nil {
			return false, errors.Wrap(err, "failed to create get role manager contract")
		}
		roleManager, err := rolemanager.NewRolemanager(roleManagerContract, net.RpcUrl)
		if err != nil {
			return false, errors.Wrap(err, "failed to create new role manager")
		}
		isManager, err := roleManager.RolemanagerCaller.IsMarketplaceManager(nil, common.HexToAddress(address))
		if err != nil {
			return false, errors.Wrap(err, "failed to create get is marketplace manager")
		}
		if !isManager {
			return isManager, nil
		}
	}
	return true, nil
}
