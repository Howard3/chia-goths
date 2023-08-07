package internal

import (
	"github.com/Masterminds/sprig"
	"github.com/gorilla/csrf"
	"github.com/rs/zerolog/log"
	"github.com/unrolled/render"
	"html/template"
	"net/http"
)

var renderer *render.Render

func getRenderer() *render.Render {
	if renderer == nil {
		renderer = render.New(render.Options{
			Directory:                   "templates",
			Layout:                      "layouts/main",
			Extensions:                  []string{".gohtml"},
			IsDevelopment:               EnvVars.DevMode,
			RequirePartials:             true,
			RenderPartialsWithoutPrefix: true,
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

	return renderer
}

func RenderHTML(r *http.Request, w http.ResponseWriter, templateName string, data interface{}) error {
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

	return getRenderer().HTML(w, http.StatusOK, templateName, data, htmlOpts...)
}
