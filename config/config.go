package config


import(
  "github.com/jinzhu/configor"
)

var Config = struct {
	App struct {
		Name  string
	}

	Blockchain struct {
		Rawurl  string
	}

	Account struct {
		Key         string
		Passphrase  string
	}
}{}

func init() {
  // Config
  //configor.Load(&Config, "./app.yml")
  configor.Load(&Config, "src/github.com/charonne/goethapi/config/app.yml")
}
