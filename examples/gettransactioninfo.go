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

package main

import(
    "fmt"
    "log"
    "context"
    "github.com/fatih/color"

    "github.com/ethereum/go-ethereum/ethclient"
  	"github.com/ethereum/go-ethereum/common"
    // "github.com/ethereum/go-ethereum/accounts/keystore"
    // "github.com/ethereum/go-ethereum/crypto"

    "github.com/charonne/goethapi/config"
    // "github.com/charonne/goethapi/converter"
)

func getClient() (client *ethclient.Client, err error) {
    log.Println("Connect to: " + config.Config.Blockchain.Rawurl)
    client, err = ethclient.Dial(config.Config.Blockchain.Rawurl)
    return
}


// Get account list
func getTransactionInfo() {
  color.Blue("Get accounts:")
  // Connect to node
  client, err := getClient()
  if err != nil {
    log.Fatalf("could not create ipc client: %v", err)
  }
  // Get context
  ctx := context.Background()

  addr := common.HexToHash("0xedab7a87b16d0215607be2262f3eb48762c8475ef405e34e72565680ed109c81")
  infos, isPending, err := client.TransactionByHash(ctx, addr)
  if err != nil {
    log.Fatalf("Error: %v", err)
  }
  fmt.Println(infos)
  fmt.Println("Is pending:", isPending)

  if (isPending == false) {
    receipt, err := client.TransactionReceipt(ctx, addr)
    if err != nil {
      log.Fatalf("Error: %v", err)
    }
    fmt.Println(receipt)
  }
}


func main() {
  line := "---------------------------------------------------------------"
  color.Green("%s v%s", config.Config.App.Name, config.Config.App.Version)
  fmt.Println(line)

  // Get block info
  getTransactionInfo()
  fmt.Println(line)


}
