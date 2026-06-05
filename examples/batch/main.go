package main

import (
	"context"
	"fmt"
	"log"
	"time"

	sec "github.com/jwitmann/sec-go"
)

func main() {
	ctx := context.Background()

	client, err := sec.ClientFromEnv()
	if err != nil {
		log.Fatalf("create client: %v", err)
	}

	// Fund list to fetch NAVs for
	projIDs := []string{"PRINCIPALi9", "KFS1001", "TMB80/20"}

	end := time.Now()
	start := end.AddDate(0, 0, -30)

	fmt.Printf("Fetching NAV history for %d funds...\n", len(projIDs))

	results := sec.BatchGetNAVs(ctx, client, projIDs, start, end, sec.BatchNAVOptions{
		Concurrency: 4,
		Progress: func(completed, total int) {
			fmt.Printf("Progress: %d/%d\n", completed, total)
		},
	})

	for _, r := range results {
		if r.Err != nil {
			fmt.Printf("%s: error: %v\n", r.ProjID, r.Err)
			continue
		}
		fmt.Printf("%s: fetched %d NAV records\n", r.ProjID, len(r.NAVs))
		if len(r.NAVs) > 0 {
			latest := r.NAVs[len(r.NAVs)-1]
			fmt.Printf("  latest NAV on %s: %.4f\n", latest.NavDate, latest.LastVal)
		}
	}
}
