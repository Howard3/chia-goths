package apps

import (
	"chia-goths/internal"
	"embed"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
	"io/fs"
	"net/http"
	"os"
	"path"
)

func assetFS(app App) fs.FS {
	fsDef := app.GetAssetsFS()

	relativePath := fsDef.RelativePath
	if relativePath == "" {
		relativePath = "assets"
	}

	if internal.EnvVars.DevMode {
		osFSPath := path.Join(app.GetAppPath(), relativePath)
		return os.DirFS(osFSPath)
	}

	sub, err := fs.Sub(fsDef.EmbeddedFS, relativePath)
	if err != nil {
		panic(fmt.Errorf("failed to get sub FS: %w", err))
	}

	return sub
}

// AppConfig allows an app to be initialized with the necessary building blocks of the chia-goths stack.
// It ideally should not be created directly and should instead be called with NewAppConfig.
type AppConfig struct {
	Renderer Renderer
	Router   chi.Router
	SubPath  string
}

func (config *AppConfig) InitApp(app App) {
	// todo: this assets delivery works but has indexes, best to not list dir contents
	appAssetPath := path.Join(config.SubPath, "assets")
	config.Router.Handle("/assets/*", http.StripPrefix(appAssetPath, http.FileServer(http.FS(assetFS(app)))))

	templateFS := app.GetTemplatesEmbedFS()
	osFSPath := path.Join(app.GetAppPath(), templateFS.RelativePath)
	config.Renderer = Renderer{
		ConstantData: map[string]any{
			"AppPath": config.SubPath,
		},
		Directory: osFSPath,
	}

	if !internal.EnvVars.DevMode {
		config.Renderer.FileSystem = &render.EmbedFileSystem{FS: templateFS.EmbeddedFS}
		config.Renderer.Directory = templateFS.RelativePath
	}

	app.Init(config)
}

func NewAppConfig(router *chi.Mux, subPath string) *AppConfig {
	config := AppConfig{}

	router.Route(subPath, func(router chi.Router) {
		config.Router = router
	})
	config.SubPath = subPath

	return &config
}

type App interface {
	Init(config *AppConfig)
	GetAppPath() string
	GetAssetsFS() AssetsFS
	GetTemplatesEmbedFS() TemplatesFS
}

type AssetsFS struct {
	EmbeddedFS   fs.FS
	RelativePath string
}

type TemplatesFS struct {
	EmbeddedFS   embed.FS
	RelativePath string
}
