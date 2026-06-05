package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

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

	fmt.Println("=== Discovering Additional Endpoints ===")
	fmt.Println()

	probeAndSave(ctx, client, "dividend-history", "/v2/fund/daily-info/dividend-history?page_size=5")
	probeAndSave(ctx, client, "asset-allocation", "/v2/fund/factsheet/asset-allocation?page_size=5")

	fmt.Println("=== Done ===")
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
