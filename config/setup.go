package config

import (
	"context"
	"log"

	"github.com/Hazem-Ben-Abdelhafidh/golangChat/ent"
	_ "github.com/lib/pq"
)

func ConnectDB() *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5434 user=root dbname=goChat password=secret sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
