package util

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// membaca file config.json
func BindFromJSON(dest any, filename, path string) error {
	v := viper.New()

	v.SetConfigType("json")
	v.AddConfigPath(path)
	v.SetConfigName(filename)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(&dest)
	if err != nil {
		logrus.Errorf("failed to unmarshal: %v", err)
		return err
	}

	return nil
}

// Membaca konfigurasi dari .env
func ReadFromEnv() map[string]any {
	v := viper.New()
	v.AutomaticEnv() // Aktifkan pembacaan dari environment variables

	configMap := map[string]any{
		"port":                  v.GetInt("PORT"),
		"appName":               v.GetString("APP_NAME"),
		"appEnv":                v.GetString("APP_ENV"),
		"rateLimiterMaxRequest": v.GetFloat64("RATE_LIMITER_MAX_REQUEST"),
		"rateLimiterTimeSecond": v.GetInt("RATE_LIMITER_TIME_SECOND"),

		// Database
		"database.host":                  v.GetString("DB_HOST"),
		"database.port":                  v.GetInt("DB_PORT"),
		"database.name":                  v.GetString("DB_NAME"),
		"database.username":              v.GetString("DB_USERNAME"),
		"database.password":              v.GetString("DB_PASSWORD"),
		"database.maxOpenConnections":    v.GetInt("DB_MAX_OPEN_CONNECTIONS"),
		"database.maxLifeTimeConnection": v.GetInt("DB_MAX_LIFETIME_CONNECTION"),
		"database.maxIdleConnections":    v.GetInt("DB_MAX_IDLE_CONNECTIONS"),
		"database.maxIdleTime":           v.GetInt("DB_MAX_IDLE_TIME"),
	}

	return configMap
}

func BindFromEnv(dest any) error {
	v := viper.New()
	v.AutomaticEnv()

	for key, value := range ReadFromEnv() {
		v.Set(key, value)
	}

	err := v.Unmarshal(dest)
	if err != nil {
		return err
	}

	return nil
}
