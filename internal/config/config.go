package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Redis struct {
		Host     string
		Port     int
		Password string
		DB       int
	}
	Postgres struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
	JWT struct {
		SecretKey string
	}
	Server struct {
		Port int
	}
	Tg struct {
		Token            string
		ModerationChatId int64
		HelpChatId       int64
		MainCharId       int64
	}
}

func (c *Config) GetRedisAddr() string {
	return fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port)
}

func (c *Config) GetPostgresConnStr() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC", c.Postgres.Host, c.Postgres.Port, c.Postgres.User, c.Postgres.Password, c.Postgres.DBName)
}

func LoadConfig() *Config {
	config := &Config{}

	config.Redis.Host = os.Getenv("REDIS_HOST")
	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		panic(err)
	}
	config.Redis.Port = port
	config.Redis.Password = os.Getenv("REDIS_PASSWORD")
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}
	config.Redis.DB = db

	config.Postgres.Host = os.Getenv("POSTGRES_HOST")
	port, err = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		panic(err)
	}
	config.Postgres.Port = port
	config.Postgres.User = os.Getenv("POSTGRES_USER")
	config.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	config.Postgres.DBName = os.Getenv("POSTGRES_DB_NAME")

	config.JWT.SecretKey = os.Getenv("JWT_SERCRET_KEY")

	port, err = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		panic(err)
	}
	config.Server.Port = port

	config.Tg.Token = "7803762373:AAGh8EClQr47d1sr6fNbD2BobzsgNlS4fp8"
	config.Tg.ModerationChatId = 2
	config.Tg.HelpChatId = 59
	config.Tg.MainCharId = -1002608013658

	return config
}
