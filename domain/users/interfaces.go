package users

import (
	"time"

	buxmodels "github.com/BuxOrg/bux-models"
	"github.com/BuxOrg/go-buxclient/transports"
	"github.com/libsv/go-bk/bip32"
)

type (
	// AccKey is an interface that defianes access key data and methods.
	AccKey interface {
		GetAccessKey() string
		GetAccessKeyId() string
	}
	// PubKey is an interface that defines xpub key data and methods.
	PubKey interface {
		GetId() string
		GetCurrentBalance() uint64
	}

	// Transaction is an interface that defines transaction data and methods.
	Transaction interface {
		GetTransactionId() string
		GetTransactionDirection() string
		GetTransactionTotalValue() uint64
		GetTransactionFee() uint64
		GetTransactionStatus() string
		GetTransactionCreatedDate() time.Time
		GetTransactionSender() string
		GetTransactionReceiver() string
	}

	// FullTransaction is an interface that defines extended transaction data and methods.
	FullTransaction interface {
		GetTransactionId() string
		GetTransactionBlockHash() string
		GetTransactionBlockHeight() uint64
		GetTransactionTotalValue() uint64
		GetTransactionDirection() string
		GetTransactionStatus() string
		GetTransactionFee() uint64
		GetTransactionNumberOfInputs() uint32
		GetTransactionNumberOfOutputs() uint32
		GetTransactionCreatedDate() time.Time
		GetTransactionSender() string
		GetTransactionReceiver() string
	}

	// DraftTransaction is an interface that defines draft transaction data and methods.
	DraftTransaction interface {
		GetDraftTransactionHex() string
		GetDraftTransactionId() string
	}

	// UserBuxClient defines methods for bux client with user key.
	UserBuxClient interface {
		// Access Key methods
		CreateAccessKey() (AccKey, error)
		GetAccessKey(accessKeyId string) (AccKey, error)
		RevokeAccessKey(accessKeyId string) (AccKey, error)
		// XPub Key methods
		GetXPub() (PubKey, error)
		// Transaction methods
		SendToRecipients(recipients []*transports.Recipients, senderPaymail string) (Transaction, error)
		GetTransactions(queryParam transports.QueryParams, userPaymail string) ([]Transaction, error)
		GetTransaction(transactionId, userPaymail string) (FullTransaction, error)
		GetTransactionsCount() (int64, error)
		CreateAndFinalizeTransaction(recipients []*transports.Recipients, metadata *buxmodels.Metadata) (DraftTransaction, error)
		RecordTransaction(hex, draftTxId string, metadata *buxmodels.Metadata) (*buxmodels.Transaction, error)
		UnreserveUtxos(draftTxId string) error
	}

	// AdmBuxClient defines methods for bux client with admin key.
	AdmBuxClient interface {
		RegisterXpub(xpriv *bip32.ExtendedKey) (string, error)
		RegisterPaymail(alias, xpub string) (string, error)
	}

	// BuxClientFactory defines methods for bux client factory.
	BuxClientFactory interface {
		CreateWithXpriv(xpriv string) (UserBuxClient, error)
		CreateWithAccessKey(accessKey string) (UserBuxClient, error)
		CreateAdminBuxClient() (AdmBuxClient, error)
	}
)
