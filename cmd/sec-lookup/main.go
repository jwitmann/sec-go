package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jwitmann/sec-go"
	"github.com/jwitmann/sec-go/internal/testutil"
)

func main() {
	keys, err := testutil.LoadTestKeys()
	if err != nil {
		log.Fatalf("Failed to load test keys: %v", err)
	}

	client, err := sec.NewClient(keys.Primary, sec.WithSecondaryKey(keys.Secondary))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	fmt.Println("=== Search: TRAREEARTH ===")
	funds, err := client.SearchFunds(ctx, "TRAREEARTH")
	if err != nil {
		log.Fatalf("SearchFunds failed: %v", err)
	}
	fmt.Printf("Found %d funds\n", len(funds))
	for _, f := range funds {
		fmt.Printf("- %s | %s | %s | status=%s | AMC=%s\n", f.ProjID, f.ProjNameEN, f.ProjNameTH, f.FundStatus, f.CompanyName(sec.LanguageEnglish))
	}
	if len(funds) == 0 {
		fmt.Println("No TRAREEARTH found. Trying TISCO company search...")
		os.Exit(0)
	}

	projID := funds[0].ProjID
	fmt.Printf("\n=== Using proj_id: %s ===\n", projID)

	fmt.Println("\n--- Profile ---")
	profile, err := client.GetFundProfile(ctx, projID)
	if err != nil {
		fmt.Printf("Profile error: %v\n", err)
	} else {
		printJSON(profile)
	}

	fmt.Println("\n--- Latest NAV ---")
	nav, err := client.GetFundLatestNAV(ctx, projID)
	if err != nil {
		fmt.Printf("NAV error: %v\n", err)
	} else {
		printJSON(nav)
	}

	fmt.Println("\n--- Risk Spectrum ---")
	risk, err := client.GetFundRiskSpectrum(ctx, projID)
	if err != nil {
		fmt.Printf("Risk error: %v\n", err)
	} else {
		printJSON(risk)
	}

	fmt.Println("\n--- Factsheet Fees ---")
	fees, err := client.GetFundFactsheetFees(ctx, projID)
	if err != nil {
		fmt.Printf("Fees error: %v\n", err)
	} else {
		printJSON(fees)
	}

	fmt.Println("\n--- Asset Allocation ---")
	alloc, err := client.GetFundAssetAllocation(ctx, projID)
	if err != nil {
		fmt.Printf("Allocation error: %v\n", err)
	} else {
		printJSON(alloc)
	}

	fmt.Println("\n--- Top 5 Holdings ---")
	top5, err := client.GetFundTop5Holdings(ctx, projID)
	if err != nil {
		fmt.Printf("Top5 error: %v\n", err)
	} else {
		printJSON(top5)
	}

	fmt.Println("\n--- Portfolio View ---")
	portfolio, err := client.GetFundPortfolio(ctx, projID)
	if err != nil {
		fmt.Printf("Portfolio error: %v\n", err)
	} else {
		printJSON(portfolio)
	}
}

func printJSON(v any) {
	data, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(data))
}
