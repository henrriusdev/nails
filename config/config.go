package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type EnvVar struct {
	Env                  string `env:"ENV" envDefault:"local"`
	BaseURL              string `env:"BASE_URL" envDefault:"http://localhost"`
	Port                 string `env:"PORT" envDefault:"8080"`
	DBName               string `env:"DB_NAME" envDefault:"db"`
	DBPort               string `env:"DB_PORT" envDefault:"5432"`
	DBUser               string `env:"DB_USERNAME" envDefault:"sera"`
	DBPass               string `env:"DB_PASSWORD" envDefault:""`
	DBHost               string `env:"DB_HOST" envDefault:"db"`
	DBSSLMode            string `env:"SSL_MODE" envDefault:"disable"`
	RedisAuthToken       string `env:"REDIS_AUTH_TOKEN" envDefault:""`
	RedisHost            string `env:"REDIS_HOST" envDefault:"redis"`
	RedisPort            string `env:"REDIS_PORT" envDefault:"6379"`
	JWTSecret            string `env:"JWT_SECRET" envDefault:"secret"`
	Sendgrid             string `env:"SENDGRID_API_KEY" envDefault:""`
	PlaidClientID        string `env:"PLAID_CLIENT_ID" envDefault:""`
	PlaidSecret          string `env:"PLAID_SECRET" envDefault:""`
	SenderEmail          string `env:"SENDER_EMAIL" envDefault:""`
	SenderPhoneNumber    string `env:"SENDER_PHONE_NUMBER" envDefault:""`
	TelnexProfileID      string `env:"TELNYX_PROFILE_ID" envDefault:""`
	TelnexAPIKey         string `env:"TELNYX_API_KEY" envDefault:""`
	MethodAPIKey         string `env:"METHOD_API_KEY" envDefault:""`
	MethodAuthHeader     string `env:"METHOD_AUTH_HEADER" envDefault:""`
	ByPassUser           string `env:"BYPASS_USERID" envDefault:""`
	PersonaWebhookSecret string `env:"PERSONA_WEBHOOK_SECRET" envDefault:""`
	PersonaAPIKey        string `env:"PERSONA_API_KEY" envDefault:""`
	StripeDefaultPlanID  string `env:"STRIPE_DEFAULT_PLAN_ID" envDefault:""`
	StripeSecretKey      string `env:"STRIPE_SECRET_KEY" envDefault:""`
	OpenAIAPIKey         string
}

var Env EnvVar

func getEnvPath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	return filepath.Join(basepath, "..", ".env")
}

func LoadEnv() {
	// Load the .env file, if it exists, this is for local development only
	if err := godotenv.Load(getEnvPath()); err != nil {
		log.Println("No .env file found")
	}

	if err := env.Parse(&Env); err != nil {
		log.Fatalf("Failed to parse env: %v", err)
	}
}
