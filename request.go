package bitgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	URL "net/url"
	"strconv"
)

func appendValues(to *URL.Values, from map[string]string) {
	for key, value := range from {
		to.Add(key, value)
	}
}

func (s *Session) setAuthorization(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+s.AccessToken) // TODO: Use their AuthV2
}

func (s *Session) executeRequest(req *http.Request) (response []byte, err error) {
	resp, err := s.Client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if s.Debug {
		log.Printf("API Response %s :: Body = %s", resp.Status, string(response))
	}

	// Check REST error
	if resp.StatusCode != http.StatusOK {
		response, err = nil, errors.New(string(response))
	}
	return
}

// DoPOST executes POST request to API endpoint
func (s *Session) DoPOST(url string, body []byte) (response []byte, err error) {
	if s.Debug {
		log.Printf("API Request POST :: %s :: Data = %s\n", url, string(body))
	}

	req, err := http.NewRequest("POST", s.BaseURL+url, bytes.NewBuffer(body))
	if err != nil {
		return
	}
	s.setAuthorization(req)
	req.Header.Set("Content-type", "application/json")
	return s.executeRequest(req)
}

// DoGET executes GET request to API endpoint
func (s *Session) DoGET(url string, args map[string]string) (response []byte, err error) {
	if s.Debug {
		log.Printf("API Request GET :: %s :: Data = %#v\n", url, args)
	}
	req, err := http.NewRequest("GET", s.BaseURL+url, nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	appendValues(&q, args)
	req.URL.RawQuery = q.Encode() // Apply modified query
	s.setAuthorization(req)
	return s.executeRequest(req)
}

// GetSessionInfo retrieves information about the current session access token
// https://www.bitgo.com/api/v2/?shell#session-information
func (s *Session) GetSessionInfo() (info *SessionInfo, err error) {
	body, err := s.DoGET("/user/session", nil)
	if err != nil {
		return
	}

	var temp struct {
		Session *SessionInfo
	}
	err = json.Unmarshal(body, &temp)
	if err != nil {
		return
	}

	info = temp.Session
	return
}

// GetWalletsPage lists user's wallets of given coin.
// https://www.bitgo.com/api/v2/?shell#list-wallets
func (c *Coin) GetWalletsPage(limit int, prevID string, allTokens bool) (wallets *WalletsPage, err error) {
	req := make(map[string]string)
	if limit != 0 {
		req["limit"] = strconv.Itoa(limit)
	}
	if prevID != "" {
		req["prevId"] = prevID
	}
	if allTokens != false {
		req["allTokens"] = strconv.FormatBool(allTokens)
	}

	url := fmt.Sprintf("/%s/wallet", c.Name)
	body, err := c.DoGET(url, req)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &wallets)
	return
}
