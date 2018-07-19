package bitgo

import "time"

// Timestamp respresents RFC3399 timestamp
type Timestamp string

// Parse parses timestamp
func (t Timestamp) Parse() (time.Time, error) {
	return time.Parse(time.RFC3339, string(t))
}

// SessionInfo holds session info
// https://www.bitgo.com/api/v2/?shell#session-information
type SessionInfo struct {
	Client  string
	User    string
	Scope   []string
	Expires Timestamp
	Origin  string
	Unlock  *UnlockInfo
}

// UnlockInfo describes session unlock status
// Used in SessionInfo
type UnlockInfo struct {
	Time    Timestamp
	Expires Timestamp
	TxCount int
	TxValue int
}

// WalletsPage response as described in api
// https://www.bitgo.com/api/v2/?shell#list-wallets
type WalletsPage struct {
	Wallets         []*Wallet
	NextBatchPrevID string
}

// Wallet model object
// https://www.bitgo.com/api/v2/?shell#get-wallet
type Wallet struct {
	ID                     string
	Label                  string
	Coin                   string
	Users                  []*WalletUser
	Keys                   []string
	Tags                   []string
	Balance                int64 // May not be set
	SpendableBalance       int64 // May not be set
	ConfirmedBalance       int64 // May not be set
	BalanceString          string
	SpendableBalanceString string
	ConfirmedBalanceString string
	StartDate              Timestamp
}

// WalletUser contains id and permissions array
// Used in Wallet
type WalletUser struct {
	User        string
	Permissions []string
}
