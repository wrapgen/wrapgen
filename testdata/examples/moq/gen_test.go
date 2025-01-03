// Code generated by WrapGen. DO NOT EDIT.
package moq

import (
	"context"
	"sync"
)


// RepoMock is a mock implementation of Repo.
type RepoMock struct {
	// ResolveIDsFunc mocks the ResolveIDs method.
	ResolveIDsFunc func(ctx context.Context, queryID string, ids []string) (map[string]bool, error)

	// calls tracks calls to the methods.
	calls struct {
		// ResolveIDs holds details about calls to the ResolveIDs method.
		ResolveIDs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// QueryID is the queryID argument value.
			QueryID string
			// Ids is the ids argument value.
			Ids []string
		}
	}
	lockResolveIDs sync.RWMutex
}


// ResolveIDs calls ResolveIDsFunc.
func (mock *RepoMock) ResolveIDs(ctx context.Context, queryID string, ids []string) (map[string]bool, error) {
	if mock.ResolveIDsFunc == nil {
		panic("RepoMock.ResolveIDsFunc: method is nil but Repo.ResolveIDs was just called")
	}
	callInfo := struct {
		Ctx context.Context
		QueryID string
		Ids []string
	}{
		Ctx: ctx,
		QueryID: queryID,
		Ids: ids,
	}
	mock.lockResolveIDs.Lock()
	mock.calls.ResolveIDs = append(mock.calls.ResolveIDs, callInfo)
	mock.lockResolveIDs.Unlock()
	return mock.ResolveIDsFunc(ctx, queryID, ids)
}

// ResolveIDs gets all the calls that were made to ResolveIDs.
// Check the length with:
//     len(RepoMock.ResolveIDsCalls())
func (mock *RepoMock) ResolveIDsCalls() []struct{
    // Ctx is the ctx argument value.
    Ctx context.Context
    // QueryID is the queryID argument value.
    QueryID string
    // Ids is the ids argument value.
    Ids []string
} {
    mock.lockResolveIDs.RLock()
    defer mock.lockResolveIDs.RUnlock()
    return mock.calls.ResolveIDs
}

