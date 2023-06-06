package auth

import (
	"bux-wallet/domain"
	"bux-wallet/domain/users"
	"errors"
	"net/http"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware middleware that is checking the variables set in session.
type AuthMiddleware struct {
	adminBuxClient users.AdmBuxClient
	services       *domain.Services
}

// NewAuthMiddleware create middleware that is checking the variables in session.
func NewAuthMiddleware(s *domain.Services) *AuthMiddleware {
	adminBuxClient, err := s.BuxClientFactory.CreateAdminBuxClient()
	if err != nil {
		panic(fmt.Errorf("error during creating admin bux client: %w", err))
	}
	return &AuthMiddleware{
		adminBuxClient: adminBuxClient,
		services:       s,
	}
}

// ApplyToApi is a middleware which checks if the validity of variables in session.
func (h *AuthMiddleware) ApplyToApi(c *gin.Context) {
	session := sessions.Default(c)

	// Try to retrieve session token.
	token := session.Get(SessionToken)
	if token == nil || token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	// Try to retrieve session user id.
	userId := session.Get(SessionUserId)
	if userId == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	err := h.checkAccessKey(token.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(SessionToken, token)
	c.Set(SessionUserId, userId)
}

// checkAccessKey checks if access key is valid by getting it from BUX.
func (h *AuthMiddleware) checkAccessKey(token string) error {
	// TODO: access token validation
	//
	// err := h.adminBuxClient.GetAccessKey(token)
	// if err != nil {
	// 	return fmt.Errorf("error during checking access key in BUX: %w", err)
	// 	return nil
	// }
	return nil
}
