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

package transfer

import (
	"fmt"
	"log"
	"math/big"
  "io/ioutil"
  "context"

  "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/charonne/goethapi/ethereum/tools"
  "github.com/charonne/goethapi/config"
)

var apiconfig = config.Config()

func Transfer(to common.Address, value *big.Int) {
	//to := common.HexToAddress("0x56a4683218dcf19d5b62b667c4348e2ac06d174b")
	// Value
  // value := big.NewInt(1000000000000000000) // 1 ether
	fmt.Printf("%s v%s\n", apiconfig.GetString("app.name"), apiconfig.GetString("app.version"))

  endPoint := apiconfig.GetString("blockchain.rawurl")
	log.Println("endPoint: ", endPoint)
  client, err := ethclient.Dial(endPoint)
  if err != nil {
			log.Fatalf("Dial error: %s", err)
  }
	log.Println("Connected to: ", apiconfig.GetString("blockchain.rawurl"))
  keystoreFile := apiconfig.GetString("account.keystore")
  password := apiconfig.GetString("account.passphrase")

  keyjson, err := ioutil.ReadFile(keystoreFile)
  if err != nil {
      log.Fatalf("Keystore error: %s", err)
  }

  key, err := keystore.DecryptKey(keyjson, password)
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

  // Figure out the gas limit
  from := common.HexToAddress(apiconfig.GetString("account.address"))

  input := []byte(apiconfig.GetString("account.passphrase"))
  msg := ethereum.CallMsg{From: from, To: &to, Value: value, Data: input}
  gasLimit, err := client.EstimateGas(ctx, msg)
  if err != nil {
    log.Fatal(err)
  }
	log.Println("Gas limit: ", gasLimit)

  // Get nonce
	nonce :=  tools.GetNewNonce(from).Number
	log.Println("Nonce: ", nonce)

  // Create the transaction
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, input)

  // Sign
  transact_opt := bind.NewKeyedTransactor(key.PrivateKey)
  //log.Printf("transact_opt: %T", transact_opt)
  tx, err = transact_opt.Signer(types.HomesteadSigner{}, common.HexToAddress(apiconfig.GetString("account.address")), tx)
  if err != nil {
      log.Fatal(err)
  }
  log.Printf("Tx: %v", tx)

  err = client.SendTransaction(ctx, tx)
  if err != nil {
      log.Fatalf("Transaction failed: %v", err)
  }
}
