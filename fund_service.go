package sec

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

func fetchPaginated[T any](ctx context.Context, c *Client, path string, op string) ([]T, string, error) {
	data, err := c.Get(ctx, path)
	if err != nil {
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}

	var response struct {
		PaginatedResponse
		Items []T `json:"items"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, "", fmt.Errorf("unmarshal %s: %w", op, err)
	}

	return response.Items, response.NextCursor, nil
}

func setPagination(params url.Values, pageSize int, cursor string) {
	if pageSize > 0 {
		params.Set("page_size", strconv.Itoa(pageSize))
	}
	if cursor != "" {
		params.Set("next_cursor", cursor)
	}
}

func setDateRange(params url.Values, startKey, endKey string, startDate, endDate time.Time) {
	if !startDate.IsZero() {
		params.Set(startKey, startDate.Format("2006-01-02"))
	}
	if !endDate.IsZero() {
		params.Set(endKey, endDate.Format("2006-01-02"))
	}
}

func buildPath(base string, params url.Values) string {
	if len(params) == 0 {
		return base
	}
	return base + "?" + params.Encode()
}
