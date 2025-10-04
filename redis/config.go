package redis

type (
	// Config interface
	Config interface {
		URI() string
		Password() string
		Database() int
	}

	// ConnConfig struct
	ConnConfig struct {
		uri      string
		password string
		database int
	}
)

// NewConnConfig creates a new Redis config
//
// Parameters:
//
//   - uri: the URI
//   - password: the password
//   - database: the database
//
// Returns:
//
// *ConnConfig: the Redis config
func NewConnConfig(uri, password string, database int) *ConnConfig {
	return &ConnConfig{
		uri,
		password,
		database,
	}
}

// URI returns the URI
//
// Returns:
//
// string: the URI
func (c ConnConfig) URI() string {
	return c.uri
}

// Password returns the password
//
// Returns:
//
// string: the password
func (c ConnConfig) Password() string {
	return c.password
}

// Database returns the database
//
// Returns:
//
// int: the database
func (c ConnConfig) Database() int {
	return c.database
}
