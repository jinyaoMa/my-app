package router

import (
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

func DocsConfig(title string, version string, docsPath string, apiPath string) huma.Config {
	if title == "" {
		title = "(" + apiPath + ")"
	} else {
		title += " (" + apiPath + ")"
	}

	if docsPath == "" || !strings.HasPrefix(docsPath, "/") {
		docsPath = "/docs"
	}

	openAPIPath := docsPath + "/openapi"

	schemaPrefix := "#/components/schemas/"
	schemasPath := docsPath + "/schemas"

	registry := huma.NewMapRegistry(schemaPrefix, huma.DefaultSchemaNamer)

	return huma.Config{
		OpenAPI: &huma.OpenAPI{
			OpenAPI: "3.1.0",
			Info: &huma.Info{
				Title:   title,
				Version: version,
			},
			Components: &huma.Components{
				Schemas: registry,
			},
		},
		OpenAPIPath:   openAPIPath,
		DocsPath:      docsPath,
		SchemasPath:   schemasPath,
		Formats:       huma.DefaultFormats,
		DefaultFormat: "application/json",
		CreateHooks: []func(huma.Config) huma.Config{
			func(c huma.Config) huma.Config {
				// Add a link transformer to the API. This adds `Link` headers and
				// puts `$schema` fields in the response body which point to the JSON
				// Schema that describes the response structure.
				// This is a create hook so we get the latest schema path setting.
				linkTransformer := huma.NewSchemaLinkTransformer(schemaPrefix, c.SchemasPath)
				c.OnAddOperation = append(c.OnAddOperation, linkTransformer.OnAddOperation)
				c.Transformers = append(c.Transformers, linkTransformer.Transform)
				return c
			},
		},
	}
}
