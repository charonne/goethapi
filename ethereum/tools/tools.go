package tools

import (
    "log"

    "regexp"
)

/**
 Check Ethereum address
 */
func CheckAddress(address string) bool {
    matched, err := regexp.MatchString("^0x[a-fA-F0-9]{40}$", address)
    if err != nil {
      log.Printf("Failed to match address: %v", err)
    }
    return matched
}

func GetBalance() {
  /*
  balance, err := client.BalanceAt(ctx, ks.Address, nil)
  if err != nil {
    log.Fatalf("could not get balance: %v", err)
  }
  balanceEther := converter.Convert(balance, "wei", "ether")
  fmt.Printf("%d: %s, balance: %v ether,  %v wei\n", i + 1, ks.Address.String(), balanceEther, balance)
  */
}
