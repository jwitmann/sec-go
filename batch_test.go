package sec

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestBatchGetNAVs(t *testing.T) {
	requestCount := int32(0)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)
		if r.URL.Path != "/v2/fund/daily-info/nav" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		projID := r.URL.Query().Get("proj_id")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"` + projID + `","nav_date":"2024-01-02","last_val":10.0}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	projIDs := []string{"FUND1", "FUND2", "FUND3"}
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)

	results := BatchGetNAVs(ctx, c, projIDs, start, end, BatchNAVOptions{Concurrency: 2})

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}

	for _, r := range results {
		if r.Err != nil {
			t.Errorf("unexpected error for %s: %v", r.ProjID, r.Err)
			continue
		}
		if len(r.NAVs) != 1 {
			t.Errorf("expected 1 NAV for %s, got %d", r.ProjID, len(r.NAVs))
		}
	}

	if atomic.LoadInt32(&requestCount) != 3 {
		t.Errorf("expected 3 requests, got %d", requestCount)
	}
}

func TestBatchGetNAVsProgress(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		projID := r.URL.Query().Get("proj_id")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"` + projID + `","nav_date":"2024-01-02","last_val":10.0}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	projIDs := []string{"FUND1", "FUND2"}
	var progressCalls int32

	results := BatchGetNAVs(ctx, c, projIDs, time.Time{}, time.Time{}, BatchNAVOptions{
		Concurrency: 2,
		Progress: func(completed, total int) {
			atomic.AddInt32(&progressCalls, 1)
			if total != 2 {
				t.Errorf("expected total=2, got %d", total)
			}
		},
	})

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
	if atomic.LoadInt32(&progressCalls) != 2 {
		t.Errorf("expected 2 progress calls, got %d", progressCalls)
	}
}

func TestBatchGetNAVsContextCancelled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	projIDs := []string{"FUND1", "FUND2", "FUND3", "FUND4"}
	results := BatchGetNAVs(ctx, c, projIDs, time.Time{}, time.Time{}, BatchNAVOptions{Concurrency: 2})

	hasContextErr := false
	for _, r := range results {
		if r.Err == context.Canceled {
			hasContextErr = true
			break
		}
	}
	if !hasContextErr {
		t.Error("expected at least one context canceled error")
	}
}

func TestBatchGetFundProfiles(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/fund/general-info/profiles" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		projID := r.URL.Query().Get("project_info")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"ok","items":[{"proj_id":"` + projID + `","proj_name_en":"Fund ` + projID + `"}]}`))
	}))
	defer server.Close()

	c, _ := NewClient("key", WithBaseURL(server.URL))
	ctx := context.Background()

	projIDs := []string{"FUND1", "FUND2"}
	results := BatchGetFundProfiles(ctx, c, projIDs, BatchProfileOptions{Concurrency: 2})

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("unexpected error for %s: %v", r.ProjID, r.Err)
			continue
		}
		if r.Profile == nil {
			t.Errorf("expected profile for %s", r.ProjID)
		}
	}
}
