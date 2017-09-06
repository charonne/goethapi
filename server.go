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

import (
  "log"
  "net/http"
  "encoding/json"

  "github.com/charonne/goethapi/database"
  "github.com/charonne/goethapi/api"

	"github.com/charonne/goethapi/config"
)

// Send response
func response(w http.ResponseWriter, responseData map[string]interface{}) {
  js, err := json.Marshal(responseData)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

// Create a smart contract
func createHandler(w http.ResponseWriter, req *http.Request) {

  // Decode json
  decoder := json.NewDecoder(req.Body)
  var createData api.ContractCreateData
  err := decoder.Decode(&createData)
  if err != nil {
      log.Println(err)
  }
  defer req.Body.Close()

  // Smart contract
  responseData := api.ContractCreate(createData)

  // Response
  response(w, responseData)
}

// Deploy a smart contract
func deployHandler(w http.ResponseWriter, req *http.Request) {

  // Decode json
  decoder := json.NewDecoder(req.Body)
  var deployData api.ContractDeployData
  err := decoder.Decode(&deployData)
  if err != nil {
      log.Println(err)
  }
  defer req.Body.Close()

  // Smart contract
  responseData := api.ContractDeploy(deployData)

  // Response
  response(w, responseData)
}

// Exec a smart contract
func execHandler(w http.ResponseWriter, req *http.Request) {

  // Decode json
  decoder := json.NewDecoder(req.Body)
  var execData api.ContractExecData
  err := decoder.Decode(&execData)
  if err != nil {
      log.Println(err)
  }
  defer req.Body.Close()

  // Smart contract
  responseData := api.ContractExec(execData)

  // Response
  response(w, responseData)
}

// Get a smart contract
func getHandler(w http.ResponseWriter, req *http.Request) {

  // Decode json
  decoder := json.NewDecoder(req.Body)
  var getData api.ContractGetData
  err := decoder.Decode(&getData)
  if err != nil {
      log.Println(err)
  }
  defer req.Body.Close()

  // Smart contract
  responseData := api.ContractGet(getData)

  // Response
  response(w, responseData)
}

// Start server
func main() {
  database.InitialiseDb()

  log.Println("Server starting on port:", config.Config.App.Port)

  http.HandleFunc("/contract/create/", createHandler)
  http.HandleFunc("/contract/deploy/", deployHandler)
  http.HandleFunc("/contract/exec/", execHandler)
  http.HandleFunc("/contract/get/", getHandler)

  //
  err := http.ListenAndServe(":" + config.Config.App.Port, nil)
  if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
