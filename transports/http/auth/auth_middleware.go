package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bitcoin-sv/spv-wallet-web-backend/domain"
	"github.com/bitcoin-sv/spv-wallet-web-backend/domain/users"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ErrorUnauthorized is thrown if authorization failed.
var ErrorUnauthorized = errors.New("unauthorized")

// AuthMiddleware middleware that is checking the variables set in session.
type AuthMiddleware struct {
	adminWalletClient   users.AdminWalletClient
	walletClientFactory users.WalletClientFactory
	services            *domain.Services
}

// NewAuthMiddleware create middleware that is checking the variables in session.
func NewAuthMiddleware(s *domain.Services) *AuthMiddleware {
	adminWalletClient, err := s.WalletClientFactory.CreateAdminClient()
	if err != nil {
		panic(fmt.Errorf("error during creating adminWalletClient: %w", err))
	}
	return &AuthMiddleware{
		adminWalletClient:   adminWalletClient,
		walletClientFactory: s.WalletClientFactory,
		services:            s,
	}
}

// ApplyToApi is a middleware which checks if the validity of variables in session.
func (h *AuthMiddleware) ApplyToApi(c *gin.Context) {
	session := sessions.Default(c)

	accessKeyId, accessKey, userId, paymail, xPriv, err := h.authorizeSession(session)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(SessionAccessKeyId, accessKeyId)
	c.Set(SessionAccessKey, accessKey)
	c.Set(SessionUserId, userId)
	c.Set(SessionUserPaymail, paymail)
	c.Set(SessionXPriv, xPriv)
}

func (h *AuthMiddleware) authorizeSession(s sessions.Session) (accessKeyId, accessKey, userId, paymail, xPriv interface{}, err error) {
	accessKeyId = s.Get(SessionAccessKeyId)
	accessKey = s.Get(SessionAccessKey)
	userId = s.Get(SessionUserId)
	paymail = s.Get(SessionUserPaymail)
	xPriv = s.Get(SessionXPriv)

	if isNilOrEmpty(accessKeyId) ||
		isNilOrEmpty(accessKey) ||
		userId == nil ||
		paymail == nil {
		return nil, nil, nil, nil, nil, ErrorUnauthorized
	}

	err = h.checkAccessKey(accessKey.(string), accessKeyId.(string))
	if err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("%w: %w", ErrorUnauthorized, err)
	}

	return
}

func isNilOrEmpty(s interface{}) bool {
	return s == nil || s == ""
}

func (h *AuthMiddleware) checkAccessKey(accessKey, accessKeyId string) error {
	userWalletClient, err := h.walletClientFactory.CreateWithAccessKey(accessKey)
	if err != nil {
		return err
	}

	_, err = userWalletClient.GetAccessKey(accessKeyId)
	return err
}
