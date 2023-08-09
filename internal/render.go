package internal

import (
	"github.com/Masterminds/sprig"
	"github.com/gorilla/csrf"
	"github.com/rs/zerolog/log"
	"github.com/unrolled/render"
	"html/template"
	"net/http"
)

type Renderer struct {
	instance   *render.Render
	FileSystem render.FileSystem
}

func (renderer *Renderer) getInstance() *render.Render {
	if renderer.instance == nil {
		renderer.instance = render.New(render.Options{
			Directory:                   "templates",
			Layout:                      "layouts/main",
			Extensions:                  []string{".gohtml"},
			IsDevelopment:               EnvVars.DevMode,
			RequirePartials:             true,
			RenderPartialsWithoutPrefix: true,
			FileSystem:                  renderer.FileSystem,
			Funcs: []template.FuncMap{
				{
					"csrfToken": func() template.HTML {
						log.Panic().Msg("csrfToken called without request")
						return ""
					},
				},
				sprig.FuncMap(),
			},
		})
	}

	return renderer.instance
}

func (renderer *Renderer) RenderHTML(r *http.Request, w http.ResponseWriter, templateName string, data interface{}) error {
	htmlOpts := []render.HTMLOptions{
		{
			Funcs: map[string]any{
				"csrfToken": func() template.HTML {
					return csrf.TemplateField(r)
				},
			},
		},
	}

	if r.Header.Get("HX-Request") == "true" {
		htmlOpts[0].Layout = "layouts/htmx"
	}

	htmlOpts = append(htmlOpts)

	return renderer.getInstance().HTML(w, http.StatusOK, templateName, data, htmlOpts...)
}
