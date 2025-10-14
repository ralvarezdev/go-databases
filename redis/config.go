package redis

type (
	// Config struct
	Config struct {
		URI      string
		Password string
		Database int
	}
)

// NewConfig creates a new Redis config
//
// Parameters:
//
//   - uri: the URI
//   - password: the password
//   - database: the database
//
// Returns:
//
// *Config: the Redis config
func NewConfig(uri, password string, database int) *Config {
	return &Config{
		uri,
		password,
		database,
	}
}
