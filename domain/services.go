package domain

import (
	db_users "github.com/bitcoin-sv/spv-wallet-web-backend/data/users"
	"github.com/bitcoin-sv/spv-wallet-web-backend/domain/config"
	"github.com/bitcoin-sv/spv-wallet-web-backend/domain/contacts"
	"github.com/bitcoin-sv/spv-wallet-web-backend/domain/rates"
	"github.com/bitcoin-sv/spv-wallet-web-backend/domain/transactions"
	"github.com/bitcoin-sv/spv-wallet-web-backend/domain/users"
	"github.com/bitcoin-sv/spv-wallet-web-backend/transports/spvwallet"
	"github.com/rs/zerolog"
)

// Services is a struct that contains all services.
type Services struct {
	UsersService        *users.UserService
	TransactionsService *transactions.TransactionService
	ContactsService     *contacts.ContactsService
	WalletClientFactory users.WalletClientFactory
	ConfigService       *config.ConfigService
	RatesService        *rates.RatesService
}

// NewServices creates services instance.
func NewServices(usersRepo *db_users.UsersRepository, log *zerolog.Logger) (*Services, error) {
	walletClientFactory := spvwallet.NewWalletClientFactory(log)
	adminWalletClient, err := walletClientFactory.CreateAdminClient()
	if err != nil {
		return nil, err
	}

	rService := rates.NewRatesService(log)
	uService := users.NewUserService(usersRepo, adminWalletClient, walletClientFactory, rService, log)

	return &Services{
		RatesService:        rService,
		UsersService:        uService,
		WalletClientFactory: walletClientFactory,
		TransactionsService: transactions.NewTransactionService(adminWalletClient, walletClientFactory, log),
		ContactsService:     contacts.NewContactsService(adminWalletClient, walletClientFactory, log),
		ConfigService:       config.NewConfigService(adminWalletClient, log),
	}, nil
}
