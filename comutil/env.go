package comutil

import "os"

// GetEnv return default value if not get value or value is empty
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

//LookupEnv return default value if not get value
func LookupEnv(key, defaultValue string) string {
	value, exist := os.LookupEnv(key)
	if exist {
		return value
	}
	return defaultValue
}
