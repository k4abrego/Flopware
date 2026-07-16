//ando practicando go, esto aún no funciona
package main

import (
	"context"
	"time"
)

func (s *Service) Back(ctx context.Context) error {
	// Simulate a long-running operation
	select {
	case <-time.After(5 * time.Second):
		// Operation completed successfully
		return nil
	case <-ctx.Done():
		// Context was canceled or timed out
		return ctx.Err()
	}
}