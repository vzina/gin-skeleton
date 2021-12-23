package config

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// ConfigFile is the default config file
var ConfigFile = "./config.yml"

// GlobalConfig is the global config
type GlobalConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	LogConfig LogConfig `yaml:"logger"`
}

// ServerConfig is the server config
type ServerConfig struct {
	Addr               string
	Mode               string
	Version            string
	StaticDir          string `yaml:"static_dir"`
	ViewDir            string `yaml:"view_dir"`
	UploadDir          string `yaml:"upload_dir"`
	MaxMultipartMemory int64  `yaml:"max_multipart_memory"`
}

// DatabaseConfig is the database config
type DatabaseConfig struct {
	DSN          string `yaml:"datasource"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}


type LogConfig struct {
	Level string `yaml:"level"`
	Filename string `yaml:"filename"`
	MaxSize int `yaml:"maxsize"`
	MaxAge int `yaml:"max_age"`
	MaxBackups int `yaml:"max_backups"`
}

// global configs
var (
	Global   GlobalConfig
	Server   ServerConfig
	Database DatabaseConfig
	LoggerConfig LogConfig
)

// Load config from file
func Load(file string) (GlobalConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("%v", err)
		return Global, err
	}

	err = yaml.Unmarshal(data, &Global)
	if err != nil {
		log.Printf("%v", err)
		return Global, err
	}

	Server = Global.Server
	Database = Global.Database
	LoggerConfig = Global.LogConfig

	return Global, nil
}

// loads configs
func init() {
	if os.Getenv("config") != "" {
		ConfigFile = os.Getenv("config")
	}
	Load(ConfigFile)
}
