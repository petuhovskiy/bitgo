package bitgo

import (
	"net/http"
	"time"
)

// Session holds parameters for making api requests
type Session struct {
	AccessToken string
	BaseURL     string
	Client      *http.Client

	Debug bool
}

// Coin is a specified BitGo coin
type Coin struct {
	*Session
	Name string
}

// NewSession is Session constructor
func NewSession(accessToken string) *Session {
	return &Session{
		accessToken,
		EndpointAPI,
		&http.Client{Timeout: 20 * time.Second},
		DebugEnabled,
	}
}

// GetCoin returns coin-specified session object
func (s *Session) GetCoin(name string) *Coin {
	return &Coin{s, name}
}
