package config

import "fmt"

type AppConfig struct {
	MySQLHost     string `yaml:"mysql_host" env:"MYSQL_HOST"`
	MySQLPort     string `yaml:"mysql_port" env:"MYSQL_PORT"`
	MySQLUserName string `yaml:"mysql_user" env:"MYSQL_USER"`
	MySQLPassword string `yaml:"mysql_pass" env:"MYSQL_PASS"`
	MySQLDatabase string `yaml:"mysql_db" env:"MYSQL_DB"`
	MySQLMaxConn  int    `yaml:"mysql_max_connection" env:"MYSQL_MAX_CONNECTION"`

	RedisHost     string `yaml:"redis_host" env:"REDIS_HOST"`
	RedisPort     string `yaml:"redis_port" env:"REDIS_PORT"`
	RedisPassword string `yaml:"redis_pass" env:"REDIS_PASS"`

	KafkaHost string `yaml:"kafka_host" env:"KAFKA_HOST"`
	KafkaPort string `yaml:"kafka_port" env:"KAFKA_PORT"`

	ZKHost     string `yaml:"zk_host" env:"ZK_HOST"`
	ZKPort     string `yaml:"zk_port" env:"ZK_PORT"`
	ZKUserName string `yaml:"zk_user" env:"ZK_USER"`
	ZKPassword string `yaml:"zk_pass" env:"ZK_PASS"`

	HTTPHost string `yaml:"http_host" env:"HTTP_HOST"`
	HTTPPort int    `yaml:"http_port" env:"HTTP_PORT"`

	JWTSecretKey string `yaml:"jwt_secret_key" env:"JWT_SECRET_KEY"`
	JWTDuration  int    `yaml:"jwt_duration" env:"JWT_DURATION"`
}

func (c AppConfig) ListenAddr() string {
	return fmt.Sprintf("%s:%d", c.HTTPHost, c.HTTPPort)
}
