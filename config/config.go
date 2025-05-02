
package config

import (
    "os"
)

func Init() {
    // nothing yet; could load .env if needed
}

func GetEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}
