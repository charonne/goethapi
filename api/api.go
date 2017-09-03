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

package api

import (
	"log"

  "github.com/charonne/goethapi/database"
  "github.com/charonne/goethapi/ethereum"
)


type CreateContractData struct {
  Name string
  Source string
}

type DeployContractData struct {
  Id string
  Params []string
}

type ExecContractData struct {
  Address string
  Method string
  Params []string
}


// Create a smart contract
func CreateContract(cdData CreateContractData) map[string]interface{} {
  // Connect to database
  db := database.DatabaseConnection()

  // Add contract
  contract := database.Contract{Idkey: GenerateIdKey(8), Name: cdData.Name, Source: cdData.Source, Bytecode: "", Interface: ""}
  db.Save(&contract)
  log.Println("Contract added, id:", contract.ID)

  // Close db
  defer db.Close()

  // Response data
  responseData := map[string]interface{}{
    "status": "success",
    "id": contract.Idkey,
  }
  return responseData
}

// Deploy a smart contract
func DeployContract(deployData DeployContractData) map[string]interface{} {
  log.Printf("Deploy contract id: %s", deployData.Id)
  // Connect to database
  db := database.DatabaseConnection()

  // Get contract
  var contract database.Contract
  db.First(&contract, "idkey = ?", deployData.Id)

  // Deploy contract
  txhash, address := ethereum.Deploy(contract)
	log.Println("Txhash: ", txhash, ", Address: ", address)

	// Add transaction
  transaction := database.Transaction{Txhash: txhash, Type: "contract", Contract: address, Status: 0}
  db.Save(&transaction)
  log.Println("Transaction added, id:", transaction.ID)

  // Response data
  responseData := map[string]interface{}{
    "status": "success",
    "id": deployData.Id,
    "txhash": txhash,
  }
  return responseData
}

// Exec a smart contract
func ExecContract(execData ExecContractData) map[string]interface{} {
  log.Println("Exec contract id: %s", execData.Address)

  log.Printf("method: %s", execData.Method)

	// Get contract
	// var contract database.Contract
	// db.First(&contract, "idkey = ?", deployData.Id)

	// Exec contract
	txhash := ethereum.Exec()
	// txhash := ethereum.Get()
	log.Println("Txhash: ", txhash)
	
	// Add transaction
	transaction := database.Transaction{Txhash: txhash, Type: "contract", Contract: address, Status: 0}
	db.Save(&transaction)
	log.Println("Transaction added, id:", transaction.ID)

  // Response data
  responseData := map[string]interface{}{
    "status": "success",
    "id": "7ed4944c788",
    "txhash": "0xee197ddbd7912e1d98b5c9b81b51aded6d39e0fbd4eb36e3485ee0a5d4d1feb9",
  }
  return responseData
}
