package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App    AppConfig    `yaml:"app"`
	DB     DBConfig     `yaml:"db"`
	Server ServerConfig `yaml:"server"`
	Logger LoggerConfig `yaml:"logger"`
	Web    WebConfig    `yaml:"web"`
	Cache  CacheConfig  `yaml:"cache"`
	Kafka  KafkaConfig  `yaml:"kafka"`
}

type AppConfig struct {
	Name    string `yaml:"name" env:"APP_NAME"`
	Version string `yaml:"version" env:"APP_VERSION"`
}

type DBConfig struct {
	MaxPoolSize int    `yaml:"max_pool_size" env:"DB_MAX_POOL_SIZE"`
	PgUrl       string `env:"DB_PG_URL,required"`
	Migrate     bool   `yaml:"migrate" env:"DB_MIGRATE"`
}

type ServerConfig struct {
	Port            string        `env:"SERVER_PORT,required"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"SERVER_SHUTDOWN_TIMEOUT"`
}

type LoggerConfig struct {
	Level string `yaml:"level" env:"LOGGER_LEVEL"`
}

type WebConfig struct {
	Path string `yaml:"path" env:"WEB_PATH"`
}

type CacheConfig struct {
	Capacity int `yaml:"capacity" env:"CACHE_CAPACITY"`
}

type KafkaConfig struct {
	Topic   string   `yaml:"topic" env:"KAFKA_TOPIC"`
	Brokers []string `env:"KAFKA_BROKERS,required" envSeparator:","`
}

func NewConfig() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		return nil, err
	}
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
