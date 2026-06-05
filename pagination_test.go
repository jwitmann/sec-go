package sec

import (
	"context"
	"errors"
	"testing"
)

func TestFetchAllPages(t *testing.T) {
	ctx := context.Background()

	calls := 0
	fetch := func(ctx context.Context, cursor string) ([]int, string, error) {
		calls++
		switch cursor {
		case "":
			return []int{1, 2, 3}, "page2", nil
		case "page2":
			return []int{4, 5, 6}, "page3", nil
		case "page3":
			return []int{7, 8, 9}, "", nil
		default:
			return nil, "", errors.New("unexpected cursor")
		}
	}

	result, err := FetchAllPages(ctx, fetch)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if calls != 3 {
		t.Errorf("expected 3 fetch calls, got %d", calls)
	}

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if len(result) != len(want) {
		t.Fatalf("expected %d items, got %d", len(want), len(result))
	}
	for i, v := range want {
		if result[i] != v {
			t.Errorf("expected result[%d] = %d, got %d", i, v, result[i])
		}
	}
}

func TestFetchAllPagesEmpty(t *testing.T) {
	ctx := context.Background()

	fetch := func(ctx context.Context, cursor string) ([]int, string, error) {
		return []int{}, "", nil
	}

	result, err := FetchAllPages(ctx, fetch)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("expected empty result, got %d items", len(result))
	}
}

func TestFetchAllPagesError(t *testing.T) {
	ctx := context.Background()

	fetch := func(ctx context.Context, cursor string) ([]int, string, error) {
		return nil, "", errors.New("fetch failed")
	}

	_, err := FetchAllPages(ctx, fetch)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
