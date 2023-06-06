package main

import (
	"bux-wallet/config"
	"bux-wallet/config/databases"
	db_users "bux-wallet/data/users"
	"bux-wallet/domain"
	"bux-wallet/logging"

	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"bux-wallet/transports/http/endpoints"
	httpserver "bux-wallet/transports/http/server"

	"github.com/spf13/viper"
)

const appname = "bux-wallet-backend"

// nolint: godot
// @title           Bux Wallet API
// @version			1.0
// @description     This is an API for bux wallet.
func main() {
	// Load config.
	config.NewViperConfig(appname).
		WithDb()

	lf := logging.DefaultLoggerFactory()
	log := lf.NewLogger("main")

	db := databases.SetUpDatabase(lf)
	defer db.Close() // nolint: all

	repo := db_users.NewUsersRepository(db)

	s, err := domain.NewServices(repo, lf)
	if err != nil {
		log.Errorf("cannot create services because of an error: ", err)
		os.Exit(1)
	}

	server := httpserver.NewHttpServer(viper.GetInt(config.EnvHttpServerPort), lf)
	server.ApplyConfiguration(endpoints.SetupWalletRoutes(s, db))

	go func() {
		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("cannot start server because of an error: ", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := server.Shutdown(); err != nil {
		log.Errorf("failed to stop http server: ", err)
	}
}
