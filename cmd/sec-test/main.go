package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jwitmann/sec-go"
	"github.com/jwitmann/sec-go/internal/testutil"
)

func main() {
	keys, err := testutil.LoadTestKeys()
	if err != nil {
		log.Fatalf("Failed to load keys: %v", err)
	}

	client, err := sec.NewClient(keys.Primary, sec.WithSecondaryKey(keys.Secondary))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	fmt.Println("=== Phase 2: V2 Schema Discovery ===")
	fmt.Println()

	// 1. AMCs
	probeAndSave(ctx, client, "amcs", "/v2/fund/general-info/amcs?page_size=5")

	// 2. Profiles
	probeAndSave(ctx, client, "profiles", "/v2/fund/general-info/profiles?page_size=5")

	// 3. Daily NAV (recent)
	probeAndSave(ctx, client, "daily-nav-recent", "/v2/fund/daily-info/nav?page_size=5")

	// 4. Daily NAV with date range
	end := time.Now().Format("2006-01-02")
	start := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	probeAndSave(ctx, client, "daily-nav-range",
		fmt.Sprintf("/v2/fund/daily-info/nav?start_nav_date=%s&end_nav_date=%s&page_size=5", start, end))

	fmt.Println("=== Discovery Complete ===")
	fmt.Println("Saved responses to docs/v2-schemas/")
}

func probeAndSave(ctx context.Context, client *sec.Client, name, path string) {
	fmt.Printf("--- %s ---\n", name)
	fmt.Printf("GET %s\n", path)

	data, err := client.Get(ctx, path)
	if err != nil {
		fmt.Printf("ERROR: %v\n\n", err)
		return
	}

	fmt.Printf("Success! %d bytes\n", len(data))

	filename := filepath.Join("docs/v2-schemas", fmt.Sprintf("%s.json", name))
	if err := os.WriteFile(filename, data, 0o644); err != nil {
		fmt.Printf("Failed to save: %v\n", err)
	} else {
		fmt.Printf("Saved: %s\n", filename)
	}
	fmt.Println()
}
