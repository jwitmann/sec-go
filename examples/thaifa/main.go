package main

import (
	"context"
	"fmt"
	"log"
	"time"

	sec "github.com/jwitmann/sec-go"
)

// Example: How THAIFA might use sec-go as a fallback data source.
// This demonstrates the integration pattern — not a working THAIFA component.

func main() {
	ctx := context.Background()

	secClient, err := sec.ClientFromEnv()
	if err != nil {
		log.Fatalf("create SEC client: %v", err)
	}

	// THAIFA's internal fund IDs would map to SEC proj_ids
	thaifaToSEC := map[string]string{
		"finnomena-123": "PRINCIPALi9",
		"finnomena-456": "KFS1001",
	}

	from := time.Now().AddDate(0, 0, -90)
	to := time.Now()

	for finnomenaID, secProjID := range thaifaToSEC {
		navs, _, err := secClient.GetDailyNAV(ctx, sec.NAVOptions{
			ProjID:    secProjID,
			StartDate: from,
			EndDate:   to,
		})
		if err != nil {
			fmt.Printf("%s (SEC:%s): fallback failed: %v\n", finnomenaID, secProjID, err)
			continue
		}
		fmt.Printf("%s (SEC:%s): %d bars from SEC fallback\n", finnomenaID, secProjID, len(navs))
	}
}
