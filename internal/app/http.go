package app

import (
	"fmt"
	"net/http"

	"github.com/aygumov-g/service-url-shortener-go/internal/config"
	"github.com/aygumov-g/service-url-shortener-go/internal/infrastructure/db"
	link_db "github.com/aygumov-g/service-url-shortener-go/internal/repository/link"
	create_link_handler "github.com/aygumov-g/service-url-shortener-go/internal/transport/http/handlers/create_link"
	get_link_handler "github.com/aygumov-g/service-url-shortener-go/internal/transport/http/handlers/get_link"
	index_handler "github.com/aygumov-g/service-url-shortener-go/internal/transport/http/handlers/index"
	redirect_handler "github.com/aygumov-g/service-url-shortener-go/internal/transport/http/handlers/redirect"
	"github.com/aygumov-g/service-url-shortener-go/internal/transport/http/router"
	"github.com/aygumov-g/service-url-shortener-go/internal/transport/http/server"
	create_link_uc "github.com/aygumov-g/service-url-shortener-go/internal/usecase/create_link"
	get_link_uc "github.com/aygumov-g/service-url-shortener-go/internal/usecase/get_link"
	update_link_uc "github.com/aygumov-g/service-url-shortener-go/internal/usecase/update_link"
	"github.com/aygumov-g/service-url-shortener-go/pkg/clock"
	"github.com/aygumov-g/service-url-shortener-go/pkg/shortcode"
	"github.com/aygumov-g/service-url-shortener-go/web/embed"
)

func buildHTTP(cfg *config.Config, db *db.Storage) (*server.Server, error) {
	gen, err := shortcode.NewEncoder(cfg.SCC.Alphabet, cfg.SCC.Secret)
	if err != nil {
		return nil, fmt.Errorf("encoder failed: %w", err)
	}

	clk := clock.NewSystemClock()

	linkRepo := link_db.NewRepository(db.Get())

	get_linkUsecase := get_link_uc.NewGetLink(linkRepo, gen)
	update_linkUsecase := update_link_uc.NewUpdateLink(linkRepo, clk)
	create_linkUsecase := create_link_uc.NewCreateLink(linkRepo, gen, clk, cfg.App.Domain)

	indexHandler := index_handler.NewHandler()
	get_linkHandler := get_link_handler.NewHandler(get_linkUsecase, cfg.App.Domain)
	create_linkHandler := create_link_handler.NewHandler(create_linkUsecase, cfg.App.Domain)
	redirectHandler := redirect_handler.NewHandler(get_linkUsecase, update_linkUsecase)

	r := router.NewRouter()
	r.Handle("/stt/*",
		http.StripPrefix("/stt/",
			http.FileServer(http.FS(embed.Public)),
		),
	)
	r.Get("/", indexHandler.Execute)
	r.Post("/api/links", create_linkHandler.Execute)
	r.Get("/api/links/{code}", get_linkHandler.Execute)
	r.Get("/{code}", redirectHandler.Execute)

	return server.NewServer(cfg.App.Port, r.Mux), nil
}
