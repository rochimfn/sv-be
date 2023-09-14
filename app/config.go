package app

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_SCHEMA   string
}

func InitConfig() Config {
	return Config{
		DB_HOST:     "127.0.0.1",
		DB_PORT:     "3306",
		DB_USERNAME: "root",
		DB_PASSWORD: "password",
		DB_SCHEMA:   "article",
	}
}
