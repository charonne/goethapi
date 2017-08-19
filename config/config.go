package config


import(
  "github.com/jinzhu/configor"
)

var Config = struct {
	App struct {
		Name  string
		Version  string
	}

	Blockchain struct {
		Rawurl        string
		Keystorepath  string
	}

	Account struct {
		Key         string
		Passphrase  string
	}
}{}

func init() {
  // Config
  configor.Load(&Config, "src/github.com/charonne/goethapi/config/app.yml")
}
