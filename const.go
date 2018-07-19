package bitgo

import (
	"os"
)

// Useful consts
// Some of them based on env
var (
	TestEndpointAPI = "https://test.bitgo.com/api/v2"
	ProdEndpointAPI = "https://www.bitgo.com/api/v2"

	EndpointAPI string

	DebugEnabled bool
)

func init() {
	if _, prod := os.LookupEnv("prod"); prod {
		EndpointAPI = ProdEndpointAPI
	} else {
		EndpointAPI = TestEndpointAPI
	}
	_, DebugEnabled = os.LookupEnv("debug")
}
