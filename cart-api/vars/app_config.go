package vars

var AppConfigs AppConfig

type AppConfig struct {
	ApiLabel                         string  `env:"API_LABEL,notEmpty"`
	ApiVersion                       string  `env:"API_VERSION,notEmpty"`
	ApiPort                          string  `env:"API_PORT,notEmpty"`
	ApiReadTimeout                   int     `env:"API_READ_TIMEOUT,notEmpty"`
	ApiStageStatus                   string  `env:"API_STAGE_STATUS,notEmpty"`
	PostgresqlHost                   string  `env:"DB_HOST,notEmpty"`
	PostgresqlUser                   string  `env:"DB_USER,notEmpty"`
	PostgresqlPass                   string  `env:"DB_PASS,notEmpty"`
	PostgresqlDb                     string  `env:"DB_DB,notEmpty"`
	PostgresqlPort                   string  `env:"DB_PORT,notEmpty"`
	PostgresqlMaxConnections         int     `env:"DB_MAX_CONNECTIONS,notEmpty"`
	PostgresqlMaxIdleConnections     int     `env:"DB_MAX_IDLE_CONNECTIONS,notEmpty"`
	PostgresqlMaxLifetimeConnections int     `env:"DB_MAX_LIFETIME_CONNECTIONS,notEmpty"`
	PostgresqlSslMode                string  `env:"DB_SSL_MODE,notEmpty"`
	GivenAmount                      float64 `env:"GIVEN_AMOUNT,notEmpty"`
}
