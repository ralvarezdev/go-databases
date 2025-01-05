package go_databases

import (
	gologger "github.com/ralvarezdev/go-logger"
	gologgermode "github.com/ralvarezdev/go-logger/mode"
	gologgermodenamed "github.com/ralvarezdev/go-logger/mode/named"
)

// Logger is the logger for the database connection
type Logger struct {
	logger gologgermodenamed.Logger
}

// NewLogger is the logger for the database connection
func NewLogger(header string, modeLogger gologgermode.Logger) (*Logger, error) {
	// Check if the logger is nil
	if modeLogger == nil {
		return nil, gologger.ErrNilLogger
	}

	// Initialize the mode named logger
	namedLogger, _ := gologgermodenamed.NewDefaultLogger(header, modeLogger)

	return &Logger{logger: namedLogger}, nil
}

// ConnectedToDatabase logs a success message when the server connects to the database
func (l *Logger) ConnectedToDatabase() {
	l.logger.Debug(
		"connected to database",
	)
}

// DisconnectedFromDatabase logs a success message when the server disconnects from the database
func (l *Logger) DisconnectedFromDatabase() {
	l.logger.Debug(
		"disconnected from database",
	)
}
