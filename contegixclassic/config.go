package contegixclassic

import (
	"fmt"
	"log"

	"github.com/randomeizer/contegix-classic"
)

type Config struct {
	AuthToken string
	CustomURL *string
}

// Client() returns a new client for accessing dnsimple.
func (c *Config) Client() (*contegixclassic.Client, error) {
	if c.CustomURL != nil {
		return c.defaultClient()
	} else {
		return c.customClient()
	}
}

func (c *Config) defaultClient() (*contegixclassic.Client, error) {
	client, err := contegixclassic.NewClient(c.AuthToken)

	if err != nil {
		return nil, fmt.Errorf("Error setting up client: %s", err)
	}

	log.Printf("[INFO] Contegix Cloud Client configured.")

	return client, nil
}

func (c *Config) customClient() (*contegixclassic.Client, error) {
	client, err := contegixclassic.NewCustomClient(c.AuthToken, c.CustomURL)

	if err != nil {
		return nil, fmt.Errorf("Error setting up client: %s", err)
	}

	log.Printf("[INFO] Contegix Cloud Client configured for '%v'.", client.URL)

	return client, nil
}
