package apps

import (
	"chia-goths/internal"
	"github.com/Masterminds/sprig"
	"github.com/gorilla/csrf"
	"github.com/rs/zerolog/log"
	"github.com/unrolled/render"
	"html/template"
	"net/http"
)

const mainLayout = "layouts/main"
const htmxLayout = "layouts/htmx"

type Renderer struct {
	instance     *render.Render
	FileSystem   render.FileSystem
	ConstantData map[string]any // data that will be available via constant. persists across renders.
	Functions    template.FuncMap
	Layout       string
	HTMXLayout   string
	Directory    string
}

func (renderer *Renderer) getInstance() *render.Render {
	if renderer.instance == nil {
		if renderer.Layout == "" {
			renderer.Layout = mainLayout
		}

		if renderer.HTMXLayout == "" {
			renderer.HTMXLayout = htmxLayout
		}

		renderer.instance = render.New(render.Options{
			Directory:                   renderer.Directory,
			Layout:                      renderer.Layout,
			Extensions:                  []string{".gohtml"},
			IsDevelopment:               internal.EnvVars.DevMode,
			RequirePartials:             true,
			RenderPartialsWithoutPrefix: true,
			FileSystem:                  renderer.FileSystem,
			Funcs: []template.FuncMap{
				{
					"csrfToken": func() template.HTML {
						log.Panic().Msg("csrfToken called without request")
						return ""
					},
					"const": func(key string) any {
						if renderer.ConstantData == nil {
							return nil
						}
						return renderer.ConstantData[key]
					},
				},
				renderer.Functions,
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
		htmlOpts[0].Layout = renderer.HTMXLayout
	}

	htmlOpts = append(htmlOpts)

	return renderer.getInstance().HTML(w, http.StatusOK, templateName, data, htmlOpts...)
}
