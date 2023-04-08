package config

import "fmt"

type AppConfig struct {
	MySQLHost     string `yaml:"mysql_host" env:"MYSQL_HOST"`
	MySQLPort     string `yaml:"mysql_port" env:"MYSQL_PORT"`
	MySQLUserName string `yaml:"mysql_user" env:"MYSQL_USER"`
	MySQLPassword string `yaml:"mysql_pass" env:"MYSQL_PASS"`
	MySQLDatabase string `yaml:"mysql_db" env:"MYSQL_DB"`
	MySQLMaxConn  int    `yaml:"mysql_max_connection" env:"MYSQL_MAX_CONNECTION"`

	HTTPHost string `yaml:"http_host" env:"HTTP_HOST"`
	HTTPPort int    `yaml:"http_port" env:"HTTP_PORT"`

	JWTSecretKey string `yaml:"jwt_secret_key" env:"JWT_SECRET_KEY"`
	JWTDuration  int    `yaml:"jwt_duration" env:"JWT_DURATION"`
}

func (c AppConfig) ListenAddr() string {
	return fmt.Sprintf("%s:%d", c.HTTPHost, c.HTTPPort)
}
