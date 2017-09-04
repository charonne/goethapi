// Copyright 2017 Charonne https://charonne.com
// This file is part of the goethapi library.
//
// The goethapi library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gethitihteg library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gethitihteg library. If not, see <http://www.gnu.org/licenses/>.

package ethereum

import (
	"fmt"
	"log"
	"math/big"
	"io/ioutil"
  "context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/charonne/goethapi/examples/contracts"
  "github.com/charonne/goethapi/database"
	"github.com/charonne/goethapi/config"
)



func Create() {
  fmt.Println("Create...")

}

// Deploy a smart contract
func Deploy(contract database.Contract) (string, string) {
	// Dial Blockchain
	endPoint := config.Config.Blockchain.Rawurl
	client, err := ethclient.Dial(endPoint)
	if err != nil {
			log.Fatal(err)
	}

	// Get context
	// ctx := context.Background()

	// Figure out the gas price
	// gasPrice, err := client.SuggestGasPrice(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Gas price: ", gasPrice)

	// 	Get transact opt
	keystoreFile := config.Config.Account.Keystore
  password := config.Config.Account.Passphrase
	keyjson, err := ioutil.ReadFile(keystoreFile)
  if err != nil {
      log.Fatal(err)
  }
	key, err := keystore.DecryptKey(keyjson, password)
  if err != nil {
      log.Fatal(err)
  }
	auth := bind.NewKeyedTransactor(key.PrivateKey)

	// Deploy contract
	addr, tx, _, err := contracts.DeploySimpleStorage(auth, client)
	// addr, _, contract, err := contracts.DeploySimpleStorage(auth, client)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}
	// Return txhash and contract address
	return tx.Hash().String(), addr.String();
}

// Exec a smart contract
func Exec() string {
  fmt.Println("Exec...")

	// Dial Blockchain
	endPoint := config.Config.Blockchain.Rawurl
	client, err := ethclient.Dial(endPoint)
	if err != nil {
			log.Fatal(err)
	}

	// Get context
	ctx := context.Background()

	// Figure out the gas price
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Gas price: ", gasPrice)

	// 	Get transact opt
	keystoreFile := config.Config.Account.Keystore
  password := config.Config.Account.Passphrase
	keyjson, err := ioutil.ReadFile(keystoreFile)
  if err != nil {
      log.Fatal(err)
  }
	key, err := keystore.DecryptKey(keyjson, password)
  if err != nil {
      log.Fatal(err)
  }
	auth := bind.NewKeyedTransactor(key.PrivateKey)

	// Change value
	fmt.Println("Adding new value...")
	addr := common.HexToAddress("0xf3d8ade61d450d264eb0dc592944db7d21bbdeef")
  contract, err := contracts.NewSimpleStorage(addr, client)
  if err != nil {
    log.Fatalf("could not call contract: %v", err)
  }
	tx, err := contract.Set(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(2381623),
		Value:    big.NewInt(0),
	}, big.NewInt(7890))
  if err != nil {
		log.Fatalf("could not execute contract: %v", err)
	}
	log.Printf("Contract executed, txhash: %s\n", tx.Hash().String())

	return tx.Hash().String()
}

/*
func Get() string {
  fmt.Println("Exec...")

	// Dial Blockchain
	endPoint := config.Config.Blockchain.Rawurl
	client, err := ethclient.Dial(endPoint)
	if err != nil {
			log.Fatal(err)
	}

	addr := common.HexToAddress("0x3330b99803931327cfcd92382189a38923817a66")
  contract, err := contracts.NewSimpleStorage(addr, client)
  if err != nil {
    log.Fatalf("could not call contract: %v", err)
  }
  info, _ := contract.Get(nil)
	fmt.Printf("Get data: %d\n", info)

	return "0xee197ddbd7912e1d98b5c9b81b51aded6d39e0fbd4eb36e3485ee0a5d4d1feb9"
}
*/
