package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
)

var once sync.Once
var instance *Config

// GetConfig get config defined in config.json
func GetConfig() *Config {
	once.Do(func() {
		env := os.Getenv("DF_ENVIROMENT")
		if env == "" {
			env = "dev"
		}

		var config *Config
		var filepath string
		pwd, _ := os.Getwd()

		if flag.Lookup("test.v") == nil {
			// normal run
			filepath = path.Join(pwd, fmt.Sprintf("config.%s.json", strings.ToLower(env)))
		} else {
			// under go test
			filepath = path.Join(pwd, "testdata", "config.unit.test.json")
		}

		configFile, err := os.Open(filepath)
		defer configFile.Close()
		if err != nil {
			panic(err)
		}

		jsonParser := json.NewDecoder(configFile)
		err = jsonParser.Decode(&config)
		if err != nil {
			panic(err)
		}

		instance = config
	})

	return instance
}

// Auth auth config
type Auth struct {
	JWKS string `json:"jwks"`
}

// WebSocket websocket config
type WebSocket struct {
	Host string `json:"host"`
}

// Nats nats config
type Nats struct {
	Host   string `json:"host"`
	Size   int    `json:"maxconnectioncount"`
	Topics string `json:"topics"`
}

// Services external services like Mysql
type Services struct {
	WebSocket `json:"websocket"`
	Nats      `json:"nats"`
}

// Config config entry
type Config struct {
	Auth     `json:"auth"`
	Services `json:"services"`
}

// GetNatsTopics convert config string to topic array
func GetNatsTopics(topic string) []string {
	return strings.Split(topic, ",")
}

// GetMaxConnectionCount get nats max connection count
func (n *Nats) GetMaxConnectionCount() int {
	if n.Size == 0 {
		return 100
	}
	return n.Size
}
