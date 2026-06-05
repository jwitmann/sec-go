//go:build integration

package sec

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jwitmann/sec-go/internal/testutil"
)

func integrationClient(t *testing.T) *Client {
	t.Helper()

	keys, err := testutil.LoadTestKeys()
	if err != nil {
		key := os.Getenv("SEC_API_KEY")
		if key == "" {
			t.Skip("Skipping integration test: no API keys available (config/sec-keys.json or SEC_API_KEY env)")
		}
		client, err := NewClient(key)
		if err != nil {
			t.Fatalf("Failed to create client from env: %v", err)
		}
		return client
	}

	client, err := NewClient(keys.Primary, WithSecondaryKey(keys.Secondary))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	return client
}

func TestIntegration_ListAMCs(t *testing.T) {
	client := integrationClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	amcs, _, err := client.ListAMCs(ctx, 5, "")
	if err != nil {
		t.Fatalf("ListAMCs failed: %v", err)
	}
	if len(amcs) == 0 {
		t.Fatal("expected at least one AMC")
	}
	t.Logf("Got %d AMCs, first: %s", len(amcs), amcs[0].CompNameEN)
}

func TestIntegration_GetFundProfiles(t *testing.T) {
	client := integrationClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	profiles, _, err := client.GetFundProfiles(ctx, ProfileOptions{PageSize: 5})
	if err != nil {
		t.Fatalf("GetFundProfiles failed: %v", err)
	}
	if len(profiles) == 0 {
		t.Fatal("expected at least one profile")
	}
	t.Logf("Got %d profiles, first: %s", len(profiles), profiles[0].ProjID)
}

func TestIntegration_GetDailyNAV(t *testing.T) {
	client := integrationClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	end := time.Now().AddDate(0, 0, -1)
	start := end.AddDate(0, 0, -7)

	navs, _, err := client.GetDailyNAV(ctx, NAVOptions{
		PageSize:  5,
		StartDate: start,
		EndDate:   end,
	})
	if err != nil {
		t.Fatalf("GetDailyNAV failed: %v", err)
	}
	t.Logf("Got %d NAV records", len(navs))
}

func TestIntegration_GetFactsheetFees(t *testing.T) {
	client := integrationClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fees, _, err := client.GetFactsheetFees(ctx, FactsheetOptions{PageSize: 5, Latest: true})
	if err != nil {
		t.Fatalf("GetFactsheetFees failed: %v", err)
	}
	t.Logf("Got %d factsheet fees", len(fees))
}

func TestIntegration_GetRiskSpectrum(t *testing.T) {
	client := integrationClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	spectrums, _, err := client.GetRiskSpectrum(ctx, FactsheetOptions{PageSize: 5, Latest: true})
	if err != nil {
		t.Fatalf("GetRiskSpectrum failed: %v", err)
	}
	t.Logf("Got %d risk spectrums", len(spectrums))
}

func TestIntegration_GetFactsheetStatistics(t *testing.T) {
	client := integrationClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stats, _, err := client.GetFactsheetStatistics(ctx, FactsheetOptions{PageSize: 5, Latest: true})
	if err != nil {
		t.Fatalf("GetFactsheetStatistics failed: %v", err)
	}
	t.Logf("Got %d statistics records", len(stats))
}

func TestIntegration_GetFactsheetDividendPolicy(t *testing.T) {
	client := integrationClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	policies, _, err := client.GetFactsheetDividendPolicy(ctx, FactsheetOptions{PageSize: 5, Latest: true})
	if err != nil {
		t.Fatalf("GetFactsheetDividendPolicy failed: %v", err)
	}
	t.Logf("Got %d dividend policies", len(policies))
}

func TestIntegration_GetFactsheetBenchmarks(t *testing.T) {
	client := integrationClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	benchmarks, _, err := client.GetFactsheetBenchmarks(ctx, FactsheetOptions{PageSize: 5, Latest: true})
	if err != nil {
		t.Fatalf("GetFactsheetBenchmarks failed: %v", err)
	}
	t.Logf("Got %d benchmarks", len(benchmarks))
}

func TestIntegration_GetFundInvolveParties(t *testing.T) {
	client := integrationClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	parties, _, err := client.GetFundInvolveParties(ctx, InvolvePartyOptions{PageSize: 5})
	if err != nil {
		t.Fatalf("GetFundInvolveParties failed: %v", err)
	}
	t.Logf("Got %d involve parties", len(parties))
}

func TestIntegration_GetFundFactsheetURLs(t *testing.T) {
	client := integrationClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	urls, _, err := client.GetFundFactsheetURLs(ctx, FeeOptions{PageSize: 5})
	if err != nil {
		t.Fatalf("GetFundFactsheetURLs failed: %v", err)
	}
	t.Logf("Got %d factsheet URLs", len(urls))
}

func TestIntegration_GetQuarterlyPortfolio(t *testing.T) {
	client := integrationClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	portfolio, _, err := client.GetQuarterlyPortfolio(ctx, OutstandingOptions{PageSize: 5})
	if err != nil {
		t.Fatalf("GetQuarterlyPortfolio failed: %v", err)
	}
	t.Logf("Got %d portfolio items", len(portfolio))
}

func TestIntegration_GetDividendHistory(t *testing.T) {
	client := integrationClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	end := time.Now().AddDate(0, 0, -1)
	start := end.AddDate(0, 0, -30)

	history, _, err := client.GetDividendHistory(ctx, DividendHistoryOptions{
		PageSize:  5,
		StartDate: start,
		EndDate:   end,
	})
	if err != nil {
		t.Fatalf("GetDividendHistory failed: %v", err)
	}
	t.Logf("Got %d dividend history records", len(history))
}
