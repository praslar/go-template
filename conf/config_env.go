package conf

// AppConfig presents app config
type AppConfig struct {
	//DB CONFIG
	DBHost   string `env:"DB_HOST" envDefault:"localhost"`
	DBPort   string `env:"DB_PORT" envDefault:"5433"`
	DBUser   string `env:"DB_USER" envDefault:"hieuminh"`
	DBPass   string `env:"DB_PASS" envDefault:"hieuminh"`
	DBName   string `env:"DB_NAME" envDefault:"hieuminh"`
	DBSchema string `env:"DB_SCHEMA" envDefault:"public"`
}
