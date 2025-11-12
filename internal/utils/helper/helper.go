package helper

import (
	"log"
	"os"
	"strconv"
)

type EnvType interface {
	string | int | bool
}

func GetEnvOrDefault[T EnvType](key string, defaultValue T) T {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	switch any(defaultValue).(type) {
	case string:
		return any(value).(T)

	case int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Printf("Failed: Environment variable %s (%s) is not a valid integer. Using default value.\n", key, value)
			return defaultValue
		}
		return any(intValue).(T)

	case bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			log.Printf("Failed: Environment variable %s (%s) is not a valid boolean. Using default value.\n", key, value)
			return defaultValue
		}
		return any(boolValue).(T)

	default:
		return defaultValue
	}
}
