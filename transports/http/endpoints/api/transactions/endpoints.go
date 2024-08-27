package transactions

import (
	"net/http"
	"strconv"

	"github.com/bitcoin-sv/spv-wallet-web-backend/domain"
	"github.com/bitcoin-sv/spv-wallet-web-backend/domain/transactions"
	"github.com/bitcoin-sv/spv-wallet-web-backend/domain/users"
	"github.com/bitcoin-sv/spv-wallet-web-backend/notification"
	"github.com/bitcoin-sv/spv-wallet-web-backend/spverrors"
	"github.com/bitcoin-sv/spv-wallet-web-backend/transports/http/auth"
	router "github.com/bitcoin-sv/spv-wallet-web-backend/transports/http/endpoints/routes"
	"github.com/bitcoin-sv/spv-wallet-web-backend/transports/spvwallet"
	"github.com/bitcoin-sv/spv-wallet-web-backend/transports/websocket"
	"github.com/bitcoin-sv/spv-wallet/models/filter"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type handler struct {
	uService users.UserService
	tService transactions.TransactionService
	log      *zerolog.Logger
	ws       websocket.Server
}

// FullTransaction is used for swagger generation
type FullTransaction = spvwallet.FullTransaction

// NewHandler creates new endpoint handler.
func NewHandler(s *domain.Services, log *zerolog.Logger, ws websocket.Server) router.APIEndpoints {
	return &handler{
		uService: *s.UsersService,
		tService: *s.TransactionsService,
		log:      log,
		ws:       ws,
	}
}

// RegisterAPIEndpoints registers routes that are part of service API.
func (h *handler) RegisterAPIEndpoints(router *gin.RouterGroup) {
	user := router.Group("/transaction")
	{
		user.POST("", h.createTransaction)
		user.POST("/search", h.getTransactions)
		user.GET("/:id", h.getTransaction)
	}
}

// Get all user transactions.
//
//	@Summary Get all transactions.
//	@Tags transaction
//	@Produce json
//	@Success 200 {object} transactions.PaginatedTransactions
//	@Router /api/v1/transaction/search [post]
func (h *handler) getTransactions(c *gin.Context) {
	var req SearchTransaction
	if err := c.Bind(&req); err != nil {
		spverrors.ErrorResponse(c, spverrors.ErrCannotBindRequest, h.log)
		return
	}

	if req.QueryParams == nil {
		req.QueryParams = &filter.QueryParams{
			Page:     1,
			PageSize: 10,
		}
	}

	// Get user transactions.
	txs, err := h.tService.GetTransactions(c.GetString(auth.SessionAccessKey), c.GetString(auth.SessionUserPaymail), req.QueryParams)
	if err != nil {
		spverrors.ErrorResponse(c, err, h.log)
		return
	}

	c.JSON(http.StatusOK, txs)
}

// Get specific transactions.
//
//	@Summary Get transaction by id.
//	@Tags transaction
//	@Produce json
//	@Success 200 {object} FullTransaction
//	@Router /api/v1/transaction/{id} [get]
//	@Param id path string true "Transaction id"
func (h *handler) getTransaction(c *gin.Context) {
	transactionID := c.Param("id")

	// Get transaction by id.
	transaction, err := h.tService.GetTransaction(c.GetString(auth.SessionAccessKey), transactionID, c.GetString(auth.SessionUserPaymail))
	if err != nil {
		spverrors.ErrorResponse(c, err, h.log)
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// Create transactions.
//
//	@Summary Create transaction.
//	@Tags transaction
//	@Produce json
//	@Success 200 {object} FullTransaction
//	@Router /api/v1/transaction [post]
//	@Param data body CreateTransaction true "Create transaction data"
func (h *handler) createTransaction(c *gin.Context) {
	var reqTransaction CreateTransaction
	if err := c.Bind(&reqTransaction); err != nil {
		spverrors.ErrorResponse(c, spverrors.ErrCannotBindRequest, h.log)
		return
	}

	// Validate user.
	xpriv, err := h.uService.GetUserXpriv(c.GetInt(auth.SessionUserID), reqTransaction.Password)
	if err != nil {
		spverrors.ErrorResponse(c, err, h.log)
		return
	}

	events := make(chan notification.TransactionEvent)
	err = h.tService.CreateTransaction(c.GetString(auth.SessionUserPaymail), xpriv, reqTransaction.Recipient, reqTransaction.Satoshis, events)
	if err != nil {
		spverrors.ErrorResponse(c, err, h.log)
		return
	}
	go func() {
		transaction := <-events
		h.ws.GetSocket(strconv.Itoa(c.GetInt(auth.SessionUserID))).Notify(transaction)
	}()

	c.Status(http.StatusOK)
}
