package sql

import (
	"context"
	"database/sql"
	"sync"
)

// RunQueriesConcurrently runs multiple queries concurrently
func RunQueriesConcurrently(
	db *sql.DB,
	queries ...func(db *sql.DB) error,
) *[]error {
	// Create a wait group
	var wg sync.WaitGroup
	wg.Add(len(queries))

	// Create a channel to handle errors
	errCh := make(chan error)

	// Execute the queries concurrently
	for _, query := range queries {
		go func(query func(db *sql.DB) error) {
			defer wg.Done()
			if err := query(db); err != nil {
				// Send the error to the channel
				errCh <- err
			}
		}(query)
	}

	// Wait for all queries to complete
	wg.Wait()

	// Disconnect the error channel
	close(errCh)

	// Return the errors if any
	var errors []error
	for err := range errCh {
		errors = append(errors, err)
	}

	// Check if there are any errors
	if len(errors) == 0 {
		return nil
	}

	return &errors
}

// RunQueriesConcurrentlyWithCancel runs multiple queries concurrently with a cancel context
func RunQueriesConcurrentlyWithCancel(
	db *sql.DB,
	queries ...func(db *sql.DB, ctx context.Context) error,
) *[]error {
	// Create a context with a cancellation function
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a wait group
	var wg sync.WaitGroup
	wg.Add(len(queries))

	// Create a channel to handle errors
	errCh := make(chan error)

	// Execute the queries concurrently
	for _, query := range queries {
		go func(query func(db *sql.DB, ctx context.Context) error) {
			defer wg.Done()
			if err := query(db, ctx); err != nil {
				// Send the error to the channel
				errCh <- err

				// cancel the other queries
				cancel()
			}
		}(query)
	}

	// Wait for all queries to complete
	wg.Wait()

	// Disconnect the error channel
	close(errCh)

	// Return the errors if any
	var errors []error
	for err := range errCh {
		errors = append(errors, err)
	}

	// Check if there are any errors
	if len(errors) == 0 {
		return nil
	}

	return &errors
}
