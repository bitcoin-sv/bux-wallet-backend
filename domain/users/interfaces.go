package users

import "github.com/libsv/go-bk/bip32"

type (
	// AccKey is an interface that defianes access key data and methods.
	AccKey interface {
		GetAccessKey() string
		GetAccessKeyId() string
	}
	// PubKey is an interface that defines xpub key data and methods.
	PubKey interface {
		GetId() string
		GetXPub() string
		GetCurrentBalance() uint64
	}
	// UserBuxClient defines methods for bux client with user key.
	UserBuxClient interface {
		// Access Key methods
		CreateAccessKey() (AccKey, error)
		GetAccessKey(accessKeyId string) (AccKey, error)
		RevokeAccessKey(accessKeyId string) (AccKey, error)
		// XPub Key methods
		GetXPub() (PubKey, error)
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
