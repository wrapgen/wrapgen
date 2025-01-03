package moq

import (
	"context"
	"fmt"
)

//wrapgen:generate -template moq -package moq -destination gen_test.go
type Repo interface {
	ResolveIDs(ctx context.Context, queryID string, ids []string) (map[string]bool, error)
}

// DoStuff does some fictional work.
func DoStuff(ctx context.Context, r Repo) (int, error) {
	result, err := r.ResolveIDs(ctx, "foobar",
		[]string{"af6799c8-0c16-4668-b9f4-ef9ac2fd354e", "5764932b-bafb-4566-826a-161942ac9c9f"})
	if err != nil {
		return 0, fmt.Errorf("failed to resolve ids: %w", err)
	}

	var count int
	for _, match := range result {
		if match {
			count++
		}
	}

	return count, nil
}
