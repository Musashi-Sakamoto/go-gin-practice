package setting

import (
	"time"
	"github.com/kelseyhightower/envconfig"
	"github.com/joho/godotenv"
	"log"
)

type DBConfig struct {
	DBType      string `envconfig:"DBType"`
	DBHost      string `envconfig:"DBHost"`
	DBPort      string `envconfig:"DBPort"`
	DBUser      string `envconfig:"DBUser"`
	DBPassword  string `envconfig:"DBPassword"`
	DBName      string `envconfig:"DBName"`
	TablePrefix string `envconfig:"TABLE_PREFIX"`
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