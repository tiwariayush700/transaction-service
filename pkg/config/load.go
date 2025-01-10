package config

import (
	"log"
	"os"
	"sync"
)

const defaultLogLevel = "info"

type Config struct {
	Port     string `json:"port"`
	LogLevel string `json:"log_level"`
	PGConfig
}

type PGConfig struct {
	PostgresUser     string `json:"postgres_user"`
	PostgresPassword string `json:"postgres_password"`
	PostgresServer   string `json:"postgres_server"`
	PostgresPort     string `json:"postgres_port"`
	PostgresDB       string `json:"postgres_db"`
	TestPostgresDB   string `json:"test_postgres_db"`
}

var (
	configuration *Config
	once          sync.Once
)

func ResetConfiguration() {
	configuration = nil
}

func LoadAppConfiguration() {
	configuration = &Config{}
	loadEnvData()
}

func GetAppConfiguration() *Config {
	once.Do(func() {
		if configuration == nil {
			log.Println("Unable to get the app configuration. Loading freshly.")
			LoadAppConfiguration()
		}
	})

	log.Printf("App config ==>> %v", *configuration)

	return configuration
}

func StopService(message string) {
	p, _ := os.FindProcess(os.Getpid())
	if err := p.Signal(os.Kill); err != nil {
		log.Fatal("error killing the process while stopping the service")
	}

	log.Fatal(message)
}

func loadEnvData() {
	configuration.Port = getEnvOrDefault("PORT", "8000")
	configuration.LogLevel = getEnvOrDefault("LOG_LEVEL", defaultLogLevel)
	configuration.PostgresUser = getEnvOrDefault("POSTGRES_USER", "transaction_user")
	configuration.PostgresPassword = getEnvOrDefault("POSTGRES_PASSWORD", "defaultpassword")
	configuration.PostgresServer = getEnvOrDefault("POSTGRES_SERVER", "localhost")
	configuration.PostgresPort = getEnvOrDefault("POSTGRES_PORT", "5432")
	configuration.PostgresDB = getEnvOrDefault("POSTGRES_DB", "transaction_db")
	configuration.TestPostgresDB = getEnvOrDefault("TEST_POSTGRES_DB", "transaction_db_test")
}

func getEnvOrDefault(envKey, defaultValue string) string {
	if value, exists := os.LookupEnv(envKey); exists {
		return value
	}
	return defaultValue
}
