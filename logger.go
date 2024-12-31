package go_databases

import (
	gologger "github.com/ralvarezdev/go-logger"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
)

// Logger is the logger for the database connection
type Logger struct {
	logger gologger.Logger
}

// NewLogger is the logger for the database connection
func NewLogger(logger gologger.Logger) (*Logger, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, gologger.ErrNilLogger
	}

	return &Logger{logger: logger}, nil
}

// ConnectedToDatabase logs a success message when the server connects to the database
func (l *Logger) ConnectedToDatabase() {
	l.logger.LogMessage(
		gologger.NewLogMessage(
			"connected to database",
			gologgerstatus.StatusDebug,
			nil,
		),
	)
}

// DisconnectedFromDatabase logs a success message when the server disconnects from the database
func (l *Logger) DisconnectedFromDatabase() {
	l.logger.LogMessage(
		gologger.NewLogMessage(
			"disconnected from database",
			gologgerstatus.StatusDebug,
			nil,
		),
	)
}
