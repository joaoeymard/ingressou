package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/JoaoEymard/ingressoscariri/api/utils/logger"
)

type Settings struct {
	Database struct {
		ConnectionRw Connection `json:"connectionRw"`
		ConnectionRo Connection `json:"connectionRo"`
	} `json:"database"`
	Listen     string `json:"listen"`
	CookieName string `json:"cookieName"`
	HashKey    string `json:"hashKey"`
	BlockKey   string `json:"blockKey"`
}

type Connection struct {
	Host        string `json:"host"`
	User        string `json:"user"`
	Pass        string `json:"pass"`
	Database    string `json:"database"`
	MaxOpenConn int    `json:"maxOpenConn"`
	MaxIdleConn int    `json:"maxIdleConn"`
}

var (
	environments = map[string]string{"production": "api/utils/settings/prod.json", "development": "api/utils/settings/dev.json"}
	settings     Settings
	env          string
)

func init() {
	env = os.Getenv("GO_UTILS")
	if env == "" {
		if GoDetails, _ := strconv.ParseBool(os.Getenv("GO_DETAILS")); GoDetails {
			logger.Warnln("Setting development environment due to lack of GO_UTILS value")
		} else {
			logger.Warnln("GO_UTILS")
		}
		env = "development"
	}
	loadSettingsByEnv(env)
}

// loadSettingsByEnv Receber as configurações do json, correspondente ao env, e setar no struct
func loadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		if GoDetails, _ := strconv.ParseBool(os.Getenv("GO_DETAILS")); GoDetails {
			logger.Fatalf("While reading config file %v", err)
		} else {
			logger.Fatalln("ReadFile environments[env]")
		}
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		if GoDetails, _ := strconv.ParseBool(os.Getenv("GO_DETAILS")); GoDetails {
			logger.Fatalf("While parsing config file %v", jsonErr)
		} else {
			logger.Fatalln("Unmarshal settings")
		}
	}
}

// GetSettings Retorna as configurações
func GetSettings() Settings {
	return settings
}
