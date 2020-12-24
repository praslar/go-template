package conf

// AppConfig presents app config
type AppConfig struct {
	//DB CONFIG
	DBHost   string `env:"DB_HOST" envDefault:"localhost"`
	DBPort   string `env:"DB_PORT" envDefault:"5432"`
	DBUser   string `env:"DB_USER" envDefault:"postgres"`
	DBPass   string `env:"DB_PASS" envDefault:"123456"`
	DBName   string `env:"DB_NAME" envDefault:"postgres"`
	DBSchema string `env:"DB_SCHEMA" envDefault:"public"`
}
