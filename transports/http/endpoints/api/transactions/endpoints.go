package transactions

import (
	"bux-wallet/domain"
	"bux-wallet/domain/transactions"
	"bux-wallet/domain/users"
	"bux-wallet/logging"
	"net/http"
	"strconv"

	"bux-wallet/transports/http/auth"
	router "bux-wallet/transports/http/endpoints/routes"

	"github.com/gin-gonic/gin"
	"github.com/mrz1836/go-datastore"
)

type handler struct {
	uService users.UserService
	tService transactions.TransactionService
	log      logging.Logger
}

// NewHandler creates new endpoint handler.
func NewHandler(s *domain.Services, lf logging.LoggerFactory) router.ApiEndpoints {
	return &handler{
		uService: *s.UsersService,
		tService: *s.TransactionsService,
		log:      lf.NewLogger("transactions-handler"),
	}
}

// RegisterApiEndpoints registers routes that are part of service API.
func (h *handler) RegisterApiEndpoints(router *gin.RouterGroup) {
	user := router.Group("/transaction")
	{
		user.GET("", h.getTransactions)
		user.POST("", h.createTransaction)
		user.GET("/:id", h.getTransaction)
	}
}

// Get all user transactions.
//
//	@Summary Get all transactions.
//	@Tags transaction
//	@Produce json
//	@Success 200 {object} []buxclient.Transaction
//	@Router /api/v1/transaction [get]
func (h *handler) getTransactions(c *gin.Context) {
	page := c.Query("page")
	pageSize := c.Query("page_size")
	orderBy := c.Query("order")
	sort := c.Query("sort")

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		pageNumber = 1
	}

	pageSizeNumber, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeNumber = 10
	}

	queryParam := datastore.QueryParams{
		Page:          pageNumber,
		PageSize:      pageSizeNumber,
		OrderByField:  orderBy,
		SortDirection: sort,
	}

	// Get user transactions.
	transactions, err := h.tService.GetTransactions(c.GetString(auth.SessionAccessKey), c.GetString(auth.SessionUserPaymail), queryParam)
	if err != nil {
		h.log.Errorf("An error occurred while trying to get a list of transactions: %s", err)
		c.JSON(http.StatusInternalServerError, "An error occurred while trying to get a list of transactions")
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// Get specific transactions.
//
//	@Summary Get transaction by id.
//	@Tags transaction
//	@Produce json
//	@Success 200 {object} buxclient.FullTransaction
//	@Router /api/v1/transaction/{id} [get]
//	@Param id path string true "Transaction id"
func (h *handler) getTransaction(c *gin.Context) {
	transactionId := c.Param("id")

	// Get transaction by id.
	transaction, err := h.tService.GetTransaction(c.GetString(auth.SessionAccessKey), transactionId, c.GetString(auth.SessionUserPaymail))
	if err != nil {
		h.log.Errorf("An error occurred while trying to get transaction details: %s", err)
		c.JSON(http.StatusInternalServerError, "An error occurred while trying to get transaction details")
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// Create transactions.
//
//	@Summary Create transaction.
//	@Tags transaction
//	@Produce json
//	@Success 200 {object} buxclient.FullTransaction
//	@Router /api/v1/transaction [post]
//	@Param data body CreateTransaction true "Create transaction data"
func (h *handler) createTransaction(c *gin.Context) {
	var reqTransaction CreateTransaction
	err := c.Bind(&reqTransaction)
	if err != nil {
		h.log.Errorf("Invalid payload: %s", err)
		c.JSON(http.StatusBadRequest, "Invalid request. Please check transaction details")
		return
	}

	// Validate user.
	xpriv, err := h.uService.GetUserXpriv(c.GetInt(auth.SessionUserId), reqTransaction.Password)
	if err != nil {
		h.log.Errorf("Invalid password: %s", err)
		c.JSON(http.StatusBadRequest, "Invalid password.")
		return
	}

	// Create transaction.
	err = h.tService.CreateTransaction(c.GetString(auth.SessionUserPaymail), xpriv, reqTransaction.Recipient, reqTransaction.Satoshis)
	if err != nil {
		h.log.Errorf("An error occurred while creating a transaction: %s", err)
		c.JSON(http.StatusBadRequest, "An error occurred while creating a transaction.")
		return
	}

	c.Status(http.StatusOK)
}
