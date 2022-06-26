package config

import (
	"encoding/json"
	"github.com/saratchandra13/sampleProject/third_party/assetmnger"
	"log"
	"os"
)

const (
	activeEnvKey    = "ACTIVE_ENV"
	production      = "PRODUCTION"
	staging         = "STAGING"
	prodFilePath    = "config/prod.json"
	stagFilePath    = "config/stag.json"
	deploymentIdKey = "DEPLOYMENT_ID"
	hostNameKey     = "HOSTNAME"
)

type Store struct {
	Env struct {
		ActiveEnv    string `json:"active_env"`
		DeploymentId string `json:"deployment_id"`
		HostName     string `json:"host_name"`
	} `json:"env"`

	DataSources struct {
		UserSvc struct {
			HttpEndpoint struct {
				Url     string `json:"url"`
				Timeout int    `json:"default_timeout"`
			} `json:"http_endpoint"`
		} `json:"user_svc"`
		Database struct {
			Mysql struct {
				Host   string `json:"host"`
				Pass   string `json:"pass"`
				DbName string `json:"db_name"`
			} `json:"mysql"`
		} `json:"database"`
	} `json:"data_sources"`
	ProjectId string `json:"project_id"`
}

func loadEnvConfig(config *Store) {
	val, ok := os.LookupEnv(activeEnvKey)
	if !ok {
		val = staging
	}
	val, ok = os.LookupEnv(deploymentIdKey)
	if !ok {
		val = ""
	}
	config.Env.DeploymentId = val

	val, ok = os.LookupEnv(hostNameKey)
	if !ok {
		val = ""
	}
	config.Env.HostName = val
	config.Env.ActiveEnv = val
}

func NewConfig(am *assetmnger.Manager) *Store {
	var config Store
	loadEnvConfig(&config)
	fileToOpen := stagFilePath
	if config.Env.ActiveEnv == production {
		fileToOpen = prodFilePath
	}

	byteValue, err := am.Get(fileToOpen)
	if err != nil {
		log.Panic("Failed to load config")
	}
	json.Unmarshal(byteValue, &config)
	return &config
}
