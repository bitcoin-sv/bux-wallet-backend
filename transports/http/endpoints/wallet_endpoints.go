package endpoints

import (
	"database/sql"
	"errors"

	"github.com/bitcoin-sv/spv-wallet-web-backend/domain"
	"github.com/bitcoin-sv/spv-wallet-web-backend/transports/http/endpoints/status"
	"github.com/bitcoin-sv/spv-wallet-web-backend/transports/http/endpoints/swagger"
	"github.com/bitcoin-sv/spv-wallet-web-backend/transports/websocket"

	"github.com/rs/zerolog"

	"github.com/bitcoin-sv/spv-wallet-web-backend/transports/http/auth"
	"github.com/bitcoin-sv/spv-wallet-web-backend/transports/http/endpoints/api/access"
	"github.com/bitcoin-sv/spv-wallet-web-backend/transports/http/endpoints/api/config"
	"github.com/bitcoin-sv/spv-wallet-web-backend/transports/http/endpoints/api/contacts"
	"github.com/bitcoin-sv/spv-wallet-web-backend/transports/http/endpoints/api/transactions"
	"github.com/bitcoin-sv/spv-wallet-web-backend/transports/http/endpoints/api/users"
	router "github.com/bitcoin-sv/spv-wallet-web-backend/transports/http/endpoints/routes"
	httpserver "github.com/bitcoin-sv/spv-wallet-web-backend/transports/http/server"

	"github.com/gin-gonic/gin"
)

// SetupWalletRoutes main point where we're registering endpoints registrars (handlers that will register endpoints in gin engine)
//
//	and middlewares. It's returning function that can be used to setup engine of httpserver.HttpServer
func SetupWalletRoutes(s *domain.Services, db *sql.DB, log *zerolog.Logger, ws websocket.Server) httpserver.GinEngineOpt {
	accessRootEndpoints, accessApiEndpoints := access.NewHandler(s, log)
	usersRootEndpoints, usersApiEndpoints := users.NewHandler(s, log)

	routes := []interface{}{
		swagger.NewHandler(),
		status.NewHandler(),
		config.NewHandler(s),
		usersRootEndpoints,
		usersApiEndpoints,
		accessRootEndpoints,
		accessApiEndpoints,
		transactions.NewHandler(s, log, ws),
		contacts.NewHandler(s, log),
	}

	return func(engine *gin.Engine) {
		apiMiddlewares := router.ToHandlers(
			auth.NewSessionMiddleware(db, engine),
			auth.NewAuthMiddleware(s),
		)

		rootRouter := engine.Group("")
		apiRouter := engine.Group("/api/v1", apiMiddlewares...)
		for _, r := range routes {
			switch r := r.(type) {
			case router.RootEndpoints:
				r.RegisterEndpoints(rootRouter)
			case router.ApiEndpoints:
				r.RegisterApiEndpoints(apiRouter)
			default:
				panic(errors.New("unexpected router endpoints registrar"))
			}
		}
	}
}
