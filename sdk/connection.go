package sdk

import (
	"os"
)

// NewConnection create connection for server api
func NewConnection(accessKey string, secretKey string) *Conn {
	conn := &Conn{
		accessKey: accessKey,
		secretKey: secretKey,
		apiURL:    "https://ncloud.apigw.ntruss.com/",
	}

	// for other phase(dev, test, beta ...) test
	if os.Getenv("NCLOUD_API_GW") != "" {
		conn.apiURL = os.Getenv("NCLOUD_API_GW")
	}

	return conn
}
