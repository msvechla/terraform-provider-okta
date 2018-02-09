package okta

import (
	"fmt"
	"log"

	"github.com/articulate/oktasdk-go/okta"
)

// Config is a struct containing our provider schema values
// plus the okta client object
type Config struct {
	orgName  string
	domain   string
	apiToken string

	oktaClient *okta.Client
}

func (c *Config) loadAndValidate() error {

	client, err := okta.NewClientWithDomain(nil, c.orgName, c.domain, c.apiToken)
	if err != nil {
		return err
	}

	// quick test of our credentials by listing our authorization server(s)
	url := fmt.Sprintf("authorizationServers")
	req, err := client.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error initializing test connect to Okta: %v", err)
		return fmt.Errorf("Error initializing test connect to Okta: %v", err)
	}
	_, err = client.Do(req, nil)
	if err != nil {
		log.Printf("[ERROR] Error testing connection to Okta: %v", err)
		return fmt.Errorf("Error testing connection to Okta. Please verify your credentials: %v", err)
	}

	// add our client object to Config
	c.oktaClient = client
	return nil
}