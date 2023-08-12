package about

import (
	"chia-goths/internal/apps"
	"embed"
	"net/http"
)

//go:embed assets/*
var embeddedAssetsFS embed.FS

//go:embed templates/*
var templatesFS embed.FS

type App struct{}

func (a App) Init(config *apps.AppConfig) {
	c := config.Router
	renderer := config.Renderer

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
}

func (a App) GetAssetsFS() apps.AssetsFS {
	return apps.AssetsFS{
		EmbeddedFS:   embeddedAssetsFS,
		RelativePath: "assets",
	}
}

func (a App) GetTemplatesEmbedFS() apps.TemplatesFS {
	return apps.TemplatesFS{
		EmbeddedFS:   templatesFS,
		RelativePath: "templates",
	}
}

func (a App) GetAppPath() string {
	return "internal/apps/about"
}
