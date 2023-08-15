package internal

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

var EnvVars = struct {
	DevMode    bool
	CSRFKey    []byte
	ListenAddr string
}{}

// LoadEnv will load a .env file, if present, and set defaults/enforce availability of environment variables.
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		// this is only INFO because env can also be set... as env.
		log.Info().Err(err).Msg("failed to load env vars from .env file")
	}

	EnvVars.DevMode = os.Getenv("DEV_MODE") == "true"
	EnvVars.CSRFKey = []byte(os.Getenv("CSRF_KEY"))

	if EnvVars.CSRFKey == nil {
		log.Panic().Msg("failed to load CSRF_KEY")
	} else if len(EnvVars.CSRFKey) != 32 {
		log.Panic().Msg("CSRF_KEY must be 32 bytes")
	}

	EnvVars.ListenAddr = os.Getenv("LISTEN_ADDR")
	if EnvVars.ListenAddr == "" {
		EnvVars.ListenAddr = "localhost:3000"
	}
}
