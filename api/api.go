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


type ContractCreateData struct {
  Name string
  Source string
}

type ContractDeployData struct {
  Id string
  Params []string
}

type ContractExecData struct {
  Address string
  Method string
  Params []string
}

type ContractGetData struct {
  Address string
  Method string
  Params []string
}

// Create a smart contract
func ContractCreate(cdData ContractCreateData) map[string]interface{} {
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
func ContractDeploy(deployData ContractDeployData) map[string]interface{} {
  log.Printf("Deploy contract id: %s", deployData.Id)
  // Connect to database
  db := database.DatabaseConnection()

  // Get contract
  var contract database.Contract
  db.First(&contract, "idkey = ?", deployData.Id)

  // Deploy contract
  txhash, address := ethereum.Deploy(contract)
	log.Println("Txhash: ", txhash, ", Address: ", address)

	// Add contract deployed
	contractDeployed := database.ContractDeployed{Address: address, Txhash: txhash}
  db.Save(&contractDeployed)
	// Set contract to deployed
	contract.Deployed = true
	db.Save(&contract)

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
func ContractExec(execData ContractExecData) map[string]interface{} {
  log.Printf("Exec contract address: %s\n", execData.Address)
  log.Printf("Method: %s, params: %s\n", execData.Method, execData.Params)

	// Connect to database
	db := database.DatabaseConnection()

	// Get contract
	var contractDeployed database.ContractDeployed
	if err := db.First(&contractDeployed, "Address = ?", execData.Address).Error; err != nil {
		// Response data
		responseData := map[string]interface{}{
			"status": "error",
			"address": execData.Address,
			"message": "contract not found",
		}
		return responseData
  }
	
	// Exec contract
	txhash := ethereum.Exec(execData.Address, execData.Method, execData.Params)

	// Add transaction
	transaction := database.Transaction{Txhash: txhash, Type: "exec", Contract: "", Status: 0}
	db.Save(&transaction)
	log.Println("Transaction added, id:", transaction.ID)

  // Response data
  responseData := map[string]interface{}{
    "status": "success",
    "address": execData.Address,
    "txhash": txhash,
  }
  return responseData
}

// Get a smart contract
func ContractGet(getData ContractGetData) map[string]interface{} {
  log.Printf("Get contract address: %s\n", getData.Address)
  log.Printf("Method: %s, params: %s\n", getData.Method, getData.Params)

	// Connect to database
	db := database.DatabaseConnection()

	// Get contract
	var contractDeployed database.ContractDeployed
	db.First(&contractDeployed, "Address = ?", getData.Address)

	// Exec contract
	data := ethereum.Get(getData.Address, getData.Method, getData.Params)

	// Response data
  responseData := map[string]interface{}{
    "message": data.Text,
    "author": data.Author,
  }
  return responseData
	/*
  // Response data
  responseData := map[string]interface{}{
    "status": "success",
    "address": getData.Address,
    "txhash": txhash,
  }
  return responseData
	*/
}
