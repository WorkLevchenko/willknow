package apiserver

/*
Для связи с файлом config.toml.
Указывает адрес порта для запуска сервера
*/
type Config struct {
	BindAddr    string `toml:"bind_addr"` //Порт запуска
	LogLevel    string `toml:"log_level"` //Уровень логирования
	DatabaseURL string `toml:"database_url"`
}

// Для возврата инициализрованного конфига с параметрами
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
