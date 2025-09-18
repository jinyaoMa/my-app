package router

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"majinyao.cn/my-app/backend/pkg/cflog"
)

type IRouter interface {
	http.Handler
	GetAPIPath() string
	GetDocsPath() string
	GetDocsTitle() string
	GetDocsVersion() string
}

func MustNew(apiPath string, options Options, setupHandlers ...func(humaapi huma.API) error) IRouter {
	api, err := New(apiPath, options, setupHandlers...)
	if err != nil {
		panic(err)
	}
	return api
}

func New(apiPath string, options Options, setupHandlers ...func(humaapi huma.API) error) (IRouter, error) {
	return new(router).init(apiPath, options, setupHandlers...)
}

type router struct {
	mux         *chi.Mux
	apiPath     string
	docsPath    string
	docsTitle   string
	docsVersion string
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *router) GetAPIPath() string {
	return r.apiPath
}

func (r *router) GetDocsPath() string {
	return r.docsPath
}

func (r *router) GetDocsTitle() string {
	return r.docsTitle
}

func (r *router) GetDocsVersion() string {
	return r.docsVersion
}

func (r *router) init(apiPath string, options Options, setupHandlers ...func(humaapi huma.API) error) (*router, error) {
	r.docsTitle = options.DocsTitle
	r.docsVersion = options.DocsVersion
	r.docsPath = options.DocsPath
	r.apiPath = apiPath

	// setup router
	r.mux = chi.NewRouter()

	logger, err := cflog.New(options.Cflog)
	if err != nil {
		return nil, err
	}

	// A good base middleware stack
	r.mux.Use(middleware.RequestID)
	r.mux.Use(middleware.RealIP)
	r.mux.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger, NoColor: true}))
	r.mux.Use(middleware.Recoverer)

	if options.EnableTimeout {
		// Set a timeout value on the request context (ctx), that will signal
		// through ctx.Done() that the request has timed out and further
		// processing should be stopped.
		r.mux.Use(middleware.Timeout(options.Timeout * time.Second))
	}

	if options.EnableCors {
		r.mux.Use(cors.Handler(cors.Options{
			AllowedOrigins:   options.AllowedOrigins,
			AllowedMethods:   options.AllowedMethods,
			AllowedHeaders:   options.AllowedHeaders,
			AllowCredentials: true,
		}))
	}

	if options.EnableHttpRate {
		if options.LimitByIp {
			r.mux.Use(httprate.LimitByIP(options.RateLimit, 1*time.Minute))
		} else {
			r.mux.Use(httprate.LimitAll(options.RateLimit, 1*time.Minute))
		}
	}

	// setup statics
	if fi, errStat := os.Lstat(options.StaticsDirectory); errStat != nil {
		if os.IsNotExist(errStat) {
			if errMkdir := os.MkdirAll(options.StaticsDirectory, os.ModeDir); errMkdir != nil {
				return nil, errors.Join(errors.New("failed to make statics directory"), errMkdir)
			}
		} else {
			return nil, errors.Join(errors.New("failed to stat statics directory"), errStat)
		}
	} else if !fi.IsDir() {
		return nil, fmt.Errorf("statics directory `%s` is not a directory", options.StaticsDirectory)
	}

	fileServer := http.FileServer(http.Dir(options.StaticsDirectory))
	r.mux.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, ".wasm.gz") {
			w.Header().Set("Content-Encoding", "gzip")
			w.Header().Set("Content-Type", "application/wasm")
		}
		fileServer.ServeHTTP(w, r)
	}))

	// setup huma
	cfg := DocsConfig(r.docsTitle, r.docsVersion, r.docsPath, r.apiPath)
	humaapi := humachi.New(r.mux, cfg)
	docsRoute(r.mux, cfg)

	humagrp := huma.NewGroup(humaapi, r.apiPath)
	for _, setupHandler := range setupHandlers {
		err := setupHandler(humagrp)
		if err != nil {
			return nil, errors.Join(errors.New("failed to setup handler"), err)
		}
	}
	return r, nil
}
