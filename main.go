package main

import (
	"chia-goths/internal"
	"embed"
	"fmt"
	"github.com/gorilla/csrf"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/unrolled/render"
	"io/fs"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "embed"
)

//go:embed assets/*
var embeddedAssetsFS embed.FS

//go:embed templates/*
var templatesFS embed.FS

func assetFS() fs.FS {
	if internal.EnvVars.DevMode {
		return os.DirFS("assets")
	}

	sub, err := fs.Sub(embeddedAssetsFS, "assets")
	if err != nil {
		panic(fmt.Errorf("failed to get sub FS: %w", err))
	}

	return sub
}

func main() {
	internal.LoadEnv()

	configLogger()

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Use(middleware.Compress(5))

	renderer := internal.Renderer{}
	if !internal.EnvVars.DevMode {
		renderer.FileSystem = &render.EmbedFileSystem{FS: templatesFS}
	}

	// todo: this assets delivery works but has indexes, best to not list dir contents
	c.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(assetFS()))))

	c.Get("/", func(w http.ResponseWriter, r *http.Request) {
		renderer.RenderHTML(r, w, "index", nil)
	})
	c.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		renderer.RenderHTML(r, w, "status", nil)
	})
	c.Get("/technologies", func(w http.ResponseWriter, r *http.Request) {
		renderer.RenderHTML(r, w, "technologies", nil)
	})
	c.Get("/csrf-testing", func(w http.ResponseWriter, r *http.Request) {
		renderer.RenderHTML(r, w, "csrf-testing", nil)
	})
	c.Post("/csrf-testing", func(w http.ResponseWriter, r *http.Request) {
		renderer.RenderHTML(r, w, "csrf-testing-post", r.PostForm)
	})

	log.Info().Str("listenAddr", internal.EnvVars.ListenAddr).Msg("starting server")
	if err := http.ListenAndServe(internal.EnvVars.ListenAddr, csrf.Protect(internal.EnvVars.CSRFKey)(c)); err != nil {
		panic(fmt.Errorf("failed to listen and serve: %w", err))
	}
}

func configLogger() {
	if internal.EnvVars.DevMode {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		log.Info().Msg("dev mode enabled")
	}

	// set chi middleware logger to zerolog
	middleware.DefaultLogger = middleware.RequestLogger(
		&middleware.DefaultLogFormatter{
			Logger:  &log.Logger,
			NoColor: !internal.EnvVars.DevMode,
		})
}
