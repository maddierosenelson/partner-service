package dbconfig

import (
	"os"

	"github.com/jackc/pgx"
)

func ExtractConfig() pgx.ConnConfig {
	var config pgx.ConnConfig
	environment := os.Getenv("PARTNER_SERVICE_ENVIRONMENT")

	var environment = os.Getenv("PARTNER_SERVICE_ENVIRONMENT") //true

	if environment == "" {
		config.Host = "partner-service.cclyw00l55b3.us-east-1.rds.amazonaws.com"
		config.User = "spam"
		config.Password = "mapsmaps"
		config.Database = "settings"
	} else {
		config.Host = os.Getenv("DB_HOST")
		if config.Host == "" {
			config.Host = "localhost"
		}

		config.User = os.Getenv("DB_USER")
		if config.User == "" {
			config.User = os.Getenv("postgres")
		}

		config.Password = os.Getenv("DB_PASSWORD")

		config.Database = os.Getenv("DB_DATABASE")
		if config.Database == "" {
			config.Database = "partner_service"
		}
	}
	return config
}
