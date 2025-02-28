package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	config struct {
		App      App
		Db       Db
		Jwt      Jwt
		Kafka    Kafka
		Grpc     Grpc
		Paginate Paginate
	}

	App struct {
		Name  string
		Url   string
		Stage string
	}

	Db struct {
		Url string
	}

	Jwt struct {
		AccessSecretKey  string
		RefreshSecretKey string
		ApiSecretKey     string
		AccessDurration  int64
		RefreshDurration int64
	}

	Kafka struct {
		Url    string
		ApiKey string
		Screat string
	}
	Grpc struct {
		AuthUrl     string
		PlayerUrl   string
		ItemUrl     string
		InenteryUrl string
		PaymentUrl  string
	}
	Paginate struct {
		ItemNextPageBaseUrl  string
		InventeryNextPageUrl string
	}
)

func LoadConfig(path string) config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file")
	}
	return config{
		App: App{
			Name:  os.Getenv("APP_NAME"),
			Url:   os.Getenv("APP_URL"),
			Stage: os.Getenv("APP_STAGE"),
		},
		Db: Db{
			Url: os.Getenv(""),
		},
		Jwt: Jwt{
			AccessSecretKey: os.Getenv("JWT_ACCESS_SECRET_KEY"),
			AccessDurration: func() int64 {
				result, err := strconv.ParseInt(os.Getenv("JWT_ACCESS_DURATION"), 10, 64)
				if err != nil {
					log.Fatal("Error loading Access Durration failed")
				}
				return result
			}(),
			RefreshSecretKey: os.Getenv("JWT_REFRESH_SECRET_KEY"),

			RefreshDurration: func() int64 {
				result, err := strconv.ParseInt(os.Getenv("JWT_REFRESH_DURATION"), 10, 64)
				if err != nil {
					log.Fatal("Error loading Refresh Durration failed")
				}
				return result
			}(),
			ApiSecretKey: os.Getenv("JWT_API_SECRET_KEY"),
		},
		Kafka: Kafka{
			Url:    os.Getenv("KAFKA_URL"),
			ApiKey: os.Getenv("KAFKA_API_KEY"),
			Screat: os.Getenv("KAFKA_API_SECRET"),
		},
		Grpc: Grpc{
			AuthUrl:     os.Getenv("GRPC_AUTH_URL"),
			PlayerUrl:   os.Getenv("GRPC_ITEM_URL"),
			ItemUrl:     os.Getenv("GRPC_PLAYER_URL"),
			InenteryUrl: os.Getenv("GRPC_INVENTORY_URL"),
			PaymentUrl:  os.Getenv("GRPC_PAYMENT_URL"),
		},
		Paginate: Paginate{
			ItemNextPageBaseUrl:  os.Getenv("PAGINATE_ITEM_NEXT_PAGE_BASED_URL"),
			InventeryNextPageUrl: os.Getenv("PAGINATE_INVENTORY_NEXT_PAGE_BASED_URL"),
		},
	}
}
