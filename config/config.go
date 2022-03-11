package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	FsLocalPath  string `envconfig:"FS_LOCAL_PATH"`
	FsMediaPath  string `envconfig:"FS_MEDIA_PATH"`
	FsPostDBPath string `envconfig:"FS_POSTDB_PATH"`

	PgUrl   string `envconfig:"DATABASE_URL"`
	PgReset bool   `envconfig:"PG_RESET"`
	Port    string `envconfig:"PORT"`
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}

		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Configuration: ", string(configBytes))
	})
	return &config
}
