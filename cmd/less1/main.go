package main

import (
	"context"
	"net/http"

	// "github.com/AleksandrMac/contentman/config"
	// accountHTTP "github.com/AleksandrMac/contentman/pkg/account/http"

	"github.com/AleksandrMac/goback2less1/config"
	"github.com/AleksandrMac/goback2less1/pkg/community"
	"github.com/AleksandrMac/goback2less1/pkg/environment"
	envUser "github.com/AleksandrMac/goback2less1/pkg/environment-user"
	"github.com/AleksandrMac/goback2less1/pkg/group"
	"github.com/AleksandrMac/goback2less1/pkg/organization"
	"github.com/AleksandrMac/goback2less1/pkg/project"
	"github.com/AleksandrMac/goback2less1/pkg/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	logger := httplog.NewLogger("contentman", httplog.Options{
		LogLevel: cfg.LogLevel,
		JSON:     true})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbpool, err := initDBPool(ctx, cfg.DBURL)
	if err != nil {
		panic(err)
	}
	defer dbpool.Close()

	_ = &http.Server{
		Addr: ":" + cfg.ServerPort,
		Handler: InitRoutes(
			dbpool,
			&logger,
		),
	}

	// go watchSignals(cancel, &logger)

	// go func() {
	// 	logger.Info().Msg("server starting")
	// 	logger.Fatal().Err(srv.ListenAndServe())
	// }()

	// ListenChan(ctx, &logger, cfg.ServerGraceFull, srv)
}

func initDBPool(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.Connect(ctx, dbURL)
	if err != nil {
		return nil, err
	}

	if err := dbpool.Ping(ctx); err != nil {
		return nil, err
	}
	return dbpool, nil
}

func InitRoutes(dbpool *pgxpool.Pool, log *zerolog.Logger) http.Handler {
	userHTTP := user.NewHTTP(log, user.NewService(user.NewStorage(dbpool)))
	envHTTP := environment.NewHTTP(log, environment.NewService(environment.NewStorage(dbpool)))
	environmentUser := envUser.NewHTTP(log, envUser.NewService(envUser.NewStorage(dbpool)))
	projectHTTP := project.NewHTTP(log, project.NewService(project.NewStorage(dbpool)))
	organizationHTTP := organization.NewHTTP(log, organization.NewService(organization.NewStorage(dbpool)))
	groupHTTP := group.NewHTTP(log, group.NewService(group.NewStorage(dbpool)))
	communityHTTP := community.NewHTTP(log, community.NewService(community.NewStorage(dbpool)))

	// ctrl := &controller.Controller{
	// 	Context:    ctx,
	// 	Logger:     logger,
	// 	ActiveUser: map[usr.SecretKey]usr.User{},
	// 	User: userService.New(
	// 		userStorage.NewWithPGX(ctx, dbpool)),
	// 	Product: productService.New(
	// 		productStorage.NewWithPGX(ctx, dbpool)),
	// 	ShoppingCart: shoppingCartService.New(
	// 		shoppingCartStorage.NewWithPGX(ctx, dbpool)),
	// }

	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		// r.Post("/auth/signIn", ctrl.SignIn)

		// r.Group(func(r chi.Router) {
		// 	r.Use(ctrl.AuthCheck)
		// 	r.Get("/cart/{user_id}", ctrl.CartGet)
		// 	r.Group(func(r chi.Router) {
		// 		r.Use(ctrl.IsManager)
		// 		r.Post("/product", ctrl.ProductCreate)
		// 		r.Post("/product/{operation}", ctrl.ProductPushPop)
		// 	})
		// 	r.Group(func(r chi.Router) {
		// 		r.Use(ctrl.IsUser)
		// 		r.Post("/cart/{user_id}/{operation}", ctrl.CartPushPop)
		// 	})
		// })
	})
	return r //CORS(r)
}
