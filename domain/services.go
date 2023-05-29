package domain

import (
	db_users "bux-wallet/data/users"
	"bux-wallet/domain/users"
	buxclient "bux-wallet/transports/bux/client"
)

// Services is a struct that contains all services.
type Services struct {
	UsersService *users.UserService
}

// NewServices creates services instance.
func NewServices(usersRepo *db_users.UsersRepository, adminBuxClient *buxclient.AdminBuxClient) *Services {
	return &Services{
		UsersService: users.NewUserService(usersRepo, adminBuxClient),
	}
}
