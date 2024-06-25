package moq

import "context"

//wrapgen:generate -template moq -package moq -destination gen_test.go
type Repo interface {
	ResolveIDs(ctx context.Context, queryID string, ids []string) (map[string]bool, error)
}
