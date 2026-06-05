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

	// List all Asset Management Companies
	fmt.Println("=== AMCs ===")
	amcs, _, err := client.ListAMCs(ctx, 10, "")
	if err != nil {
		log.Fatalf("list AMCs: %v", err)
	}
	for _, amc := range amcs {
		fmt.Printf("%s - %s\n", amc.UniqueID, amc.CompNameEN)
	}

	// Get fund profiles for an AMC
	fmt.Println("\n=== Fund Profiles ===")
	profiles, _, err := client.GetFundProfiles(ctx, sec.ProfileOptions{
		CompanyInfo: amcs[0].UniqueID,
	})
	if err != nil {
		log.Fatalf("get profiles: %v", err)
	}
	for _, p := range profiles {
		fmt.Printf("%s - %s (%s)\n", p.ProjID, p.ProjNameEN, p.FundStatus)
	}

	// Get daily NAV for the first fund
	if len(profiles) > 0 {
		fmt.Println("\n=== Daily NAV ===")
		end := time.Now()
		start := end.AddDate(0, 0, -7)
		navs, _, err := client.GetDailyNAV(ctx, sec.NAVOptions{
			ProjID:    profiles[0].ProjID,
			StartDate: start,
			EndDate:   end,
		})
		if err != nil {
			log.Fatalf("get NAV: %v", err)
		}
		for _, nav := range navs {
			fmt.Printf("%s: %.4f\n", nav.NavDate, nav.LastVal)
		}
	}
}
