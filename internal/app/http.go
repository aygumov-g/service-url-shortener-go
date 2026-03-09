package app

import (
	"fmt"
	"net/http"

	"github.com/aygumov-g/service-url-shortener-go/internal/config"
	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"
	"github.com/aygumov-g/service-url-shortener-go/internal/infrastructure/db"
	link_db "github.com/aygumov-g/service-url-shortener-go/internal/repository/link"
	create_link_handler "github.com/aygumov-g/service-url-shortener-go/internal/transport/http/handlers/create_link"
	index_handler "github.com/aygumov-g/service-url-shortener-go/internal/transport/http/handlers/index"
	redirect_handler "github.com/aygumov-g/service-url-shortener-go/internal/transport/http/handlers/redirect"
	"github.com/aygumov-g/service-url-shortener-go/internal/transport/http/router"
	"github.com/aygumov-g/service-url-shortener-go/internal/transport/http/server"
	create_link_uc "github.com/aygumov-g/service-url-shortener-go/internal/usecase/create_link"
	redirect_uc "github.com/aygumov-g/service-url-shortener-go/internal/usecase/redirect"
	"github.com/aygumov-g/service-url-shortener-go/pkg/clock"
	"github.com/aygumov-g/service-url-shortener-go/pkg/shortcode"
	"github.com/aygumov-g/service-url-shortener-go/web/embed"
)

func buildHTTP(
	cfg *config.Config,
	db *db.Storage,
) (*server.Server, error) {
	if err := db.Get().AutoMigrate(&link_d.Link{}); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	gen, err := shortcode.NewEncoder(cfg.SCC.Alphabet, cfg.SCC.Secret)
	if err != nil {
		return nil, fmt.Errorf("encoder failed: %w", err)
	}

	clk := clock.NewSystemClock()

	linkRepo := link_db.NewRepository(db.Get())

	rootUsecase := redirect_uc.NewRedirect(linkRepo, gen, clk)
	create_linkUsecase := create_link_uc.NewCreateLink(linkRepo, gen, clk, cfg.App.Domain)

	indexHandler := index_handler.NewHandler()
	redirectHandler := redirect_handler.NewHandler(rootUsecase)
	create_linkHandler := create_link_handler.NewHandler(create_linkUsecase, cfg.App.Domain)

	r := router.NewRouter()
	r.Handle("/stt/*",
		http.StripPrefix("/stt/",
			http.FileServer(http.FS(embed.Public)),
		),
	)
	r.Get("/", indexHandler.Execute)
	r.Post("/api/links", create_linkHandler.Execute)
	r.Get("/{code}", redirectHandler.Execute)

	return server.NewServer(cfg.App.Port, r.Mux), nil
}
