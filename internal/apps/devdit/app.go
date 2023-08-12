package devdit

import (
	"chia-goths/internal/apps"
	"embed"
	"github.com/rs/zerolog/log"
	"math/rand"
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
	return "internal/apps/devdit"
}

func (receiver App) Init(config *apps.AppConfig) {
	log.Info().Msg("Initializing Devdit")

	footerMessages := []string{
		"It’s not a bug – it’s an undocumented feature!",
		"Last commit: \"Fixed those pesky bugs... probably.\"",
		"Uptime: 99.99% (It's that 0.01% that keeps me up at night!)",
		"This footer is a footer.",
		"Spotted a bug? Congrats! You're today's unofficial QA tester. Let us know!",
		"Fun Fact: The first computer bug was an actual bug – a moth found in a relay.",
		"Remember to stand up and stretch every once in a while!",
	}

	config.Renderer.Functions = map[string]interface{}{
		"randomFooterMessage": func() string {
			return footerMessages[rand.Intn(len(footerMessages))]
		},
	}

	config.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		config.Renderer.RenderHTML(r, w, "index", nil)
	})
}
