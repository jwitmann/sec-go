package sec

import "context"

func FetchAllPages[T any](ctx context.Context, fetch func(ctx context.Context, cursor string) ([]T, string, error)) ([]T, error) {
	var all []T
	cursor := ""

	for {
		items, nextCursor, err := fetch(ctx, cursor)
		if err != nil {
			return nil, err
		}

		all = append(all, items...)

		if nextCursor == "" {
			break
		}
		cursor = nextCursor
	}

	return all, nil
}
