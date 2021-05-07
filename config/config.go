package config

import (
	"fmt"
	"os"
)

func IsLocal() bool {
	return os.Getenv("ENV") == "local"
}

func DSN() string {
	if IsLocal() {
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DATABASE"),
		) + "?parseTime=true&collation=utf8mb4_bin"
	}

	// TODO DSN to CloudSQL
	return ""
}
