package store

import (
	"log"

	"github.com/henrriusdev/nails/config"
	"github.com/supabase-community/postgrest-go"
)

func NewConnection(cfg config.EnvVar) (*postgrest.Client, error) {
	// Create a new Supabase client
	client := postgrest.NewClient(cfg.SupabaseURL, "", nil)
	if client.ClientError != nil {
		log.Fatalf("Failed to create a new Supabase client: %v", client.ClientError)
		return nil, client.ClientError
	}

	log.Println("Successfully connected to Supabase")
	return client, nil
}
