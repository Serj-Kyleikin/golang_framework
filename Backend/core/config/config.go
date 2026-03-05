package config

import (
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	HTTP struct {
		Addr string `mapstructure:"addr"`
	} `mapstructure:"http"`

	DB struct {
		URL      string `mapstructure:"url"`
		MaxConns int32  `mapstructure:"max_conns"`
	} `mapstructure:"db"`

	Log struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"log"`
}

func Load() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	_ = v.ReadInConfig()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetDefault("http.addr", ":8080")
	v.SetDefault("db.max_conns", 10)
	v.SetDefault("log.level", "info")

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if cfg.DB.URL == "" {
		if env := os.Getenv("DB_URL"); env != "" {
			cfg.DB.URL = env
		}
	}

	if cfg.DB.URL == "" {
		return nil, &ConfigError{"db.url (or DB_URL) is required"}
	}
	return &cfg, nil
}

type ConfigError struct{ msg string }

func (e *ConfigError) Error() string { return e.msg }

func NewLogger(level string) *slog.Logger {
	var lvl slog.Level
	switch strings.ToLower(level) {
	case "debug":
		lvl = slog.LevelDebug
	case "warn", "warning":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: lvl})
	return slog.New(h)
}
