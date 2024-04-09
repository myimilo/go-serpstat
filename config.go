package serpstat

const (
	serpstatAPIURLv4 = "https://api.serpstat.com/v4"
)

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	authToken string
	BaseURL   string
}

func DefaultConfig(authToken string) ClientConfig {
	return ClientConfig{
		authToken: authToken,
		BaseURL:   serpstatAPIURLv4,
	}
}

func (ClientConfig) String() string {
	return "<Serpstat API ClientConfig>"
}
