package app

import (
	"fmt"
	"napoleon-email/src/pkg/logger"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func AppEnv() string { return env("APP_ENV", "development") }
func IsProd() bool { return isProduction() }
func AppPort() string { return env("SERVER_PORT", "3000") }
func AppVersion() string { return env("APP_VERSION", "1.0.0") }
func GoogleApplicationCredentials() string { return env("GOOGLE_APPLICATION_CREDENTIALS", "") }
func EmailHost() string { return env("EMAIL_HOST", "") }
func EmailPort() string { return env("EMAIL_PORT", "") }
func EmailAddress() string { return env("EMAIL_ADDRESS", "") }
func EmailPassword() string { return env("EMAIL_PASSWORD", "") }
func EmailNapoleonTo() string { return env("EMAIL_NAPOLEON_TO", "") }
func EmailGroupNapoleonTo() string { return env("EMAIL_GROUP_NAPOLEON_TO", "") }
func EmailNapoleonMineTo() string { return env("EMAIL_NAPOLEON_MINE_TO", "") }
func EmailBankGoldTo() string { return env("EMAIL_BANK_GOLD_TO", "") }

var envLoaded = false
var envMutex sync.Mutex

func ValidateEnvironment() error {
	requiredVars := []string{
		"APP_ENV",
		"SERVER_PORT",
		"APP_VERSION",
		"GOOGLE_APPLICATION_CREDENTIALS",
		"EMAIL_HOST",
		"EMAIL_PORT",
		"EMAIL_ADDRESS",
		"EMAIL_PASSWORD",
		"EMAIL_NAPOLEON_TO",
		"EMAIL_GROUP_NAPOLEON_TO",
		"EMAIL_NAPOLEON_MINE_TO",
		"EMAIL_BANK_GOLD_TO",
	}
	var missing []string
	for _, v := range requiredVars {
		val := env(v, "")
		if val == "" {
			logger.LogWarn(fmt.Sprintf("Required environment variable %s is not set", v), logger.LogStruct{Action: "env_variable_missing", User: 0, Data: v})
			missing = append(missing, v)
		}
	}
	if len(missing) > 0 {
		logger.LogError(fmt.Sprintf("missing required environment variables: %v", missing), nil, logger.LogStruct{Action: "env_variables_missing", User: 0, Data: missing})
		return fmt.Errorf("missing required environment variables: %v", missing)
	}
	return nil
}

func LoadEnv() {
	envMutex.Lock()
	defer envMutex.Unlock()

	if envLoaded {
		return
	}
	if err := godotenv.Load(); err != nil {
		logger.LogError("Error loading .env file", err, logger.LogStruct{Action: "env_load_error", User: 0})
	}
	envLoaded = true
}

func env(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	if defaultValue != "" {
		return defaultValue
	}
	logger.LogInfo(fmt.Sprintf("Environment variable %s not set, using empty string as default", key), logger.LogStruct{Action: "env_variable_missing", User: 0, Data: key})
	logger.LogWarn("Environment variable not set", logger.LogStruct{Action: "env_variable_missing", User: 0, Data: key})
	return ""
}

func isProduction() bool {
	return AppEnv() == "production"
}
