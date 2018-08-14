package config

type Config struct {
	Port   string `conf:"PORT" default:"8080"`
	DBHost string `conf:"DB_HOST" default:"localhost"`
	DBName string `conf:"DB_NAME" default:"my_db"`
	DBUser string `conf:"DB_USER" default:"root"`
	DBPass string `conf:"DB_PASS" default:"123456"`
}
