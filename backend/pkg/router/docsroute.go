package router

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/chi/v5"
	"majinyao.cn/my-app/backend/pkg/router/assets"
)

func docsRoute(mux *chi.Mux, cfg huma.Config) {
	mux.Get(cfg.DocsPath, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <link rel="icon" href="` + cfg.DocsPath + `/favicon.ico" />
    <meta name="referrer" content="same-origin" />
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
    <title>` + cfg.OpenAPI.Info.Title + `</title>
    <style>
      body {
        margin: 0;
        padding: 1rem 0.5rem;
      }
    </style>
    <script>
      function updateLayout() {
        var elementsApi = document.querySelector("elements-api");
        if (elementsApi.offsetWidth < 1024) {
          document.body.style.padding = "1rem 0.5rem"
          elementsApi.layout = "stacked";
        } else {
          document.body.style.padding = "0"
          elementsApi.layout = "sidebar";
        }
      }
      document.addEventListener("DOMContentLoaded", updateLayout);
      window.addEventListener("resize", updateLayout);
    </script>
    <!-- Embed elements Elements via Web Component -->
    <link href="` + cfg.DocsPath + `/styles" rel="stylesheet" />
    <script src="` + cfg.DocsPath + `/web-components"></script>
  </head>
  <body style="height: 100vh;">
    <elements-api
      apiDescriptionUrl="` + cfg.OpenAPIPath + `.yaml"
      router="hash"
      layout="stacked"
      tryItCredentialsPolicy="same-origin"
    />
  </body>
</html>`))
	})
	mux.Get(cfg.DocsPath+"/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write([]byte(assets.Favicon))
	})
	mux.Get(cfg.DocsPath+"/styles", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		w.Write([]byte(assets.Styles))
	})
	mux.Get(cfg.DocsPath+"/web-components", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write([]byte(assets.WebComponents))
	})
}
