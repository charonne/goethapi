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

package ticker

import (
  "time"
  "fmt"
  "net/http"
  "io/ioutil"
  "bytes"

  "github.com/charonne/goethapi/database"
  "github.com/charonne/goethapi/ethereum"
  "github.com/charonne/goethapi/config"
)

// Ticker period
var period = config.Config.Ticker.Period

// Send confirmation to callback
func SendToCallback(txhash string) {
  url := config.Config.Callback
  //fmt.Println("URL:>", url)

  var jsonStr = []byte(`{"txhash": "` + txhash + `", "confirmed": true}`)
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
  req.Header.Set("X-Custom-Header", "myvalue")
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()

  // fmt.Println("response Status:", resp.Status)
  // fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))
}

// Start ticker
func Start() {
  // Ticker every period
  tickTxConfirmation := time.NewTicker(time.Second * time.Duration(period)).C

  // Start infinite loop
  for {
    select {
      // Ticker confirmation
      case <- tickTxConfirmation:
        fmt.Println("Ticker ticked")

        // Connect to database
        db := database.DatabaseConnection()

        // Get not confirmed transactions
        rows, _ := db.Model(&database.Transaction{}).Where("Confirmed = ?", false).Rows()
        defer rows.Close()
        for rows.Next() {
          var transaction database.Transaction
          db.ScanRows(rows, &transaction)
          // Check confirmation
          if (ethereum.GetTransactionConfirmation(transaction.Txhash) == true) {
            fmt.Println("Transaction confirmed: ", transaction.Txhash)
            // Set confirmed
            transaction.Confirmed = true
            db.Save(&transaction)
            // Send to Callback
            SendToCallback(transaction.Txhash)
          } else {
            fmt.Println("Transaction not confirmed: ", transaction.Txhash)
          }
        }
        db.Close()
    }
  }
  
}
