package server

import (
	"context"
	"helios/internal/api/handlers/light"
	"helios/internal/api/handlers/sensor"
	"helios/internal/api/middleware"
	"helios/internal/api/service"
	"log"
	"net/http"
)

type Server struct {
	ctx        context.Context
	HTTPServer *http.Server
	logger     *log.Logger
}

func NewServer(ctx context.Context, sf *service.ServiceFactory, logger *log.Logger) *Server {

	mux := http.NewServeMux()
	err := setupDataHandlers(mux, sf, logger)
	if err != nil {
		logger.Fatalf("Error setting up data handlers: %v", err)
	}

	middlewares := []middleware.Middleware{
		middleware.BasicAuthenticationMiddleware,
		middleware.CommonMiddleware,
	}

	return &Server{
		ctx:    ctx,
		logger: logger,
		HTTPServer: &http.Server{
			Handler: middleware.ChainMiddleware(mux, middlewares...),
		},
	}
}

func (api *Server) Shutdown() error {
	api.logger.Println("Gracefully shutting down server...")
	return api.HTTPServer.Shutdown(api.ctx)
}

func (api *Server) ListenAndServe(addr string) error {
	api.HTTPServer.Addr = addr
	return api.HTTPServer.ListenAndServe()
}

// * REST API handlers
func setupDataHandlers(mux *http.ServeMux, sf *service.ServiceFactory, logger *log.Logger) error {

	ds, err := sf.CreateDataService(service.SQLiteDataService)
	if err != nil {
		return err
	}

	mux.HandleFunc("OPTIONS /*", func(w http.ResponseWriter, r *http.Request) {
		sensor.OptionsHandler(w, r)
	})
	mux.HandleFunc("POST /sensor", func(w http.ResponseWriter, r *http.Request) {
		sensor.PostHandler(w, r, logger, ds)
	})
	mux.HandleFunc("PUT /sensor", func(w http.ResponseWriter, r *http.Request) {
		sensor.PutHandler(w, r, logger, ds)
	})
	mux.HandleFunc("GET /sensor", func(w http.ResponseWriter, r *http.Request) {
		sensor.GetHandler(w, r, logger, ds)
	})
	mux.HandleFunc("GET /sensor/{id}", func(w http.ResponseWriter, r *http.Request) {
		sensor.GetByIDHandler(w, r, logger, ds)
	})
	mux.HandleFunc("DELETE /sensor/{id}", func(w http.ResponseWriter, r *http.Request) {
		sensor.DeleteHandler(w, r, logger, ds)
	})

	mux.HandleFunc("GET /light", func(w http.ResponseWriter, r *http.Request) {
		light.GetHandler(w, r, logger)
	})
	mux.HandleFunc("PUT /light", func(w http.ResponseWriter, r *http.Request) {
		light.PutLightHandler(w, r, logger)
	})
	return err
}
