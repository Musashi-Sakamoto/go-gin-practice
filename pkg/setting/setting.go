package setting

import (
	"time"
	"github.com/kelseyhightower/envconfig"
	"github.com/joho/godotenv"
	"log"
)

type DBConfig struct {
	DBHost   string `envconfig:"DBHost"`
	Username string `envconfig:"Username"`
	Password string `envconfig:"Password"`
	Database string `envconfig:"Database"`
}

type ServerConfig struct {
	HTTPPort     int           `envconfig:"PORT"`
	ReadTimeout  time.Duration `envconfig:"READTIMEOUT"`
	WriteTimeout time.Duration `envconfig:"WRITETIMEOUT"`
}

type AppConfig struct {
	RunMode   string `envconfig:"RUNMODE"`
	PageSize  int    `envconfig:"PAGESIZE"`
	JWTSecret string `envconfig:"JWTSECRET`
}

type config struct {
	DBConfig
	ServerConfig
	AppConfig
}

var Config = config{}

func init() {
	EnvLoad()
	err := envconfig.Process("", &Config)
	if err != nil {
		log.Fatal("Fail to load config with env: %v", err)
	}
}

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}