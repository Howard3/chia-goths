package todos

import (
	"chia-goths/internal/apps"
	"embed"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

//go:embed assets/*
var embeddedAssetsFS embed.FS

//go:embed templates/*
var templatesFS embed.FS

type App struct{}

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
	return "internal/apps/todos"
}

func (receiver App) Init(config *apps.AppConfig) {
	log.Info().Msg("Initializing Todos")

	db, err := gorm.Open(sqlite.Open("internal/apps/todos/todos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}

	err = db.AutoMigrate(&Todo{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}

	config.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		var todos []Todo
		db.Find(&todos)
		config.Renderer.RenderHTML(r, w, "index", map[string]interface{}{"todos": todos})
	})

	config.Router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		title := r.Form.Get("Title")
		todo := &Todo{Title: title}
		db.Save(todo)
		config.Renderer.RenderHTML(r, w, "single", todo)
	})

	config.Router.Post("/toggle/{todoID}", func(w http.ResponseWriter, r *http.Request) {
		if todoID := chi.URLParam(r, "todoID"); todoID != "" {
			var todo Todo
			db.First(&todo, "id = ?", todoID)
			todo.Done = !todo.Done
			db.Save(&todo)
			config.Renderer.RenderHTML(r, w, "single", todo)
		} else {
			config.Renderer.RenderHTML(r, w, "404", nil)
		}
	})
}
