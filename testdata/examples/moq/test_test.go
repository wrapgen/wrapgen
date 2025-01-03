package moq

import (
	"context"
	"errors"
	"testing"
)

func TestDoStuff(t *testing.T) {
	for _, tc := range []struct {
		name         string
		mock         func(t *testing.T) *RepoMock
		expectError  string
		expectResult int
		verify       func(t *testing.T, r *RepoMock)
	}{
		{
			name: "empty result",
			mock: func(t *testing.T) *RepoMock {
				return &RepoMock{
					ResolveIDsFunc: func(ctx context.Context, queryID string, ids []string) (map[string]bool, error) {
						return nil, nil
					},
				}
			},
			verify: func(t *testing.T, r *RepoMock) {
				if calls := len(r.ResolveIDsCalls()); calls != 1 {
					t.Errorf("expect 1 call got %d", calls)
				}
			},
		},
		{
			name: "fail with error",
			mock: func(t *testing.T) *RepoMock {
				return &RepoMock{
					ResolveIDsFunc: func(ctx context.Context, queryID string, ids []string) (map[string]bool, error) {
						return nil, errors.New("some error")
					},
				}
			},
			expectError: "failed to resolve ids: some error",
		},
		{
			name: "counting",
			mock: func(t *testing.T) *RepoMock {
				return &RepoMock{
					ResolveIDsFunc: func(ctx context.Context, queryID string, ids []string) (map[string]bool, error) {
						return map[string]bool{
							"a": true, "b": false, "c": true, "d": false,
						}, nil
					},
				}
			},
			expectResult: 2,
			verify: func(t *testing.T, r *RepoMock) {
				if calls := len(r.ResolveIDsCalls()); calls != 1 {
					t.Errorf("expect 1 call got %d", calls)
				}
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := tc.mock(t)
			result, err := DoStuff(context.Background(), r)
			if tc.expectError != "" {
				if err.Error() != tc.expectError {
					t.Errorf("expected error %q got %q", tc.expectError, err.Error())
				}
			} else {
				if tc.expectResult != result {
					t.Errorf("expected result %d got %d", tc.expectResult, result)
				}
			}
			if tc.verify != nil {
				tc.verify(t, r)
			}
		})
	}
}
