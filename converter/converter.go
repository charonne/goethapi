package converter

import(
  "fmt"
  "math/big"
)


func Convert(value *big.Int, from string, to string) *big.Float {
  // Choose convertion unit
  switch unit := from; unit {
    case "ether":
      fmt.Println("OS X.")
    case "wei":
      return ConvertWei(value, to)
  }
  return new(big.Float).SetInt(value)
}

func ConvertWei(value *big.Int, to string) *big.Float {
  // Default coef
  var coef = big.NewInt(1)
  // Choose convertion unit
  switch unit := to; unit {
    case "ether":
      coef = big.NewInt(1000000000000000000)
  }
  // Make conversion
  v := new(big.Float).SetInt(value)
  c := new(big.Float).SetInt(coef)
  converted := new(big.Float).Quo(v, c)

  // Round
  //converted = Round(converted, 5)
  // Return results
  return converted
}
