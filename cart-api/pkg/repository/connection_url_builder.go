package repository

import (
	"fmt"

	"github.com/tohanilhan/Cart-API/vars"
)

func ConnectionURLBuilder(n string) (string, error) {
	var url string

	switch n {
	case "sentinel-postgres":
		dbHost := vars.AppConfigs.PostgresqlHost
		dbPort := vars.AppConfigs.PostgresqlPort
		dbUser := vars.AppConfigs.PostgresqlUser
		dbPassword := vars.AppConfigs.PostgresqlPass
		dbName := vars.AppConfigs.PostgresqlDb
		dbSslMode := vars.AppConfigs.PostgresqlSslMode

		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			dbHost,
			dbPort,
			dbUser,
			dbPassword,
			dbName,
			dbSslMode,
		)
	default:
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	return url, nil
}
