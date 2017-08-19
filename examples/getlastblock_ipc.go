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

  "github.com/ethereum/go-ethereum/rpc"

  "github.com/charonne/goethapi/config"
)

type Block struct {
  Number string
}

func main() {
  // Connect the client
  client, err := rpc.Dial(config.Config.Blockchain.Rawurl)
  if err != nil {
    log.Fatalf("could not create ipc client: %v", err)
  }

  // Main rpc call
  var lastBlock Block
  err = client.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
  if err != nil {
    fmt.Println("can't get latest block:", err)
    return
  }

  // Print events from the subscription as they arrive.
  fmt.Printf("latest block: %v\n", lastBlock.Number)
}
