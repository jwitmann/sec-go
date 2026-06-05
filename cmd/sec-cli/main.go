package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jwitmann/sec-go"
	"github.com/jwitmann/sec-go/internal/testutil"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	ctx := context.Background()
	client := newClient()

	switch cmd {
	case "amcs":
		handleAMCs(ctx, client)
	case "profiles":
		handleProfiles(ctx, client, os.Args[2:])
	case "specifications":
		handleSpecifications(ctx, client, os.Args[2:])
	case "fees":
		handleFees(ctx, client, os.Args[2:])
	case "involve-parties":
		handleInvolveParties(ctx, client, os.Args[2:])
	case "nav":
		handleNAV(ctx, client, os.Args[2:])
	case "dividend-history":
		handleDividendHistory(ctx, client, os.Args[2:])
	case "factsheet-urls":
		handleFactsheetURLs(ctx, client, os.Args[2:])
	case "ipos":
		handleIPOs(ctx, client, os.Args[2:])
	case "benchmarks":
		handleBenchmarks(ctx, client, os.Args[2:])
	case "subscription-minimums":
		handleSubscriptionMinimums(ctx, client, os.Args[2:])
	case "subscription-periods":
		handleSubscriptionPeriods(ctx, client, os.Args[2:])
	case "risk-spectrum":
		handleRiskSpectrum(ctx, client, os.Args[2:])
	case "statistics":
		handleStatistics(ctx, client, os.Args[2:])
	case "dividend-policy":
		handleDividendPolicy(ctx, client, os.Args[2:])
	case "factsheet-fees":
		handleFactsheetFees(ctx, client, os.Args[2:])
	case "performance":
		handlePerformance(ctx, client, os.Args[2:])
	case "asset-allocation":
		handleAssetAllocation(ctx, client, os.Args[2:])
	case "top5":
		handleTop5(ctx, client, os.Args[2:])
	case "portfolio":
		handlePortfolio(ctx, client, os.Args[2:])
	case "portfolio-asset-type":
		handlePortfolioAssetType(ctx, client, os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: sec-cli <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  amcs                      List AMCs")
	fmt.Println("  profiles                  Get fund profiles")
	fmt.Println("  specifications            Get fund specifications")
	fmt.Println("  fees                      Get mutual fund fees")
	fmt.Println("  involve-parties           Get fund involve parties")
	fmt.Println("  nav                       Get daily NAV")
	fmt.Println("  dividend-history          Get dividend history")
	fmt.Println("  factsheet-urls            Get factsheet URLs/PDFs")
	fmt.Println("  ipos                      Get IPO offering periods")
	fmt.Println("  benchmarks                Get fund benchmarks")
	fmt.Println("  subscription-minimums     Get subscription/redemption minimums")
	fmt.Println("  subscription-periods      Get subscription/redemption periods")
	fmt.Println("  risk-spectrum             Get risk spectrum")
	fmt.Println("  statistics                Get fund statistics")
	fmt.Println("  dividend-policy           Get dividend policy")
	fmt.Println("  factsheet-fees            Get factsheet fees")
	fmt.Println("  performance               Get historical performance")
	fmt.Println("  asset-allocation          Get asset allocation")
	fmt.Println("  top5                      Get top 5 holdings")
	fmt.Println("  portfolio                 Get quarterly portfolio")
	fmt.Println("  portfolio-asset-type      Get monthly portfolio by asset type")
	fmt.Println()
	fmt.Println("Common flags:")
	fmt.Println("  -proj-id string           Project ID filter")
	fmt.Println("  -fund-class-name string   Fund class name filter")
	fmt.Println("  -start string             Start date (YYYY-MM-DD) or start period (YYYYMM)")
	fmt.Println("  -end string               End date (YYYY-MM-DD) or end period (YYYYMM)")
	fmt.Println("  -latest                   Return only latest factsheet data")
	fmt.Println("  -page-size int            Items per page (default 10)")
}

func newClient() *sec.Client {
	keys, err := testutil.LoadTestKeys()
	if err != nil {
		client, err := sec.ClientFromEnv()
		if err != nil {
			log.Fatalf("Failed to load API keys: %v", err)
		}
		return client
	}

	client, err := sec.NewClient(keys.Primary, sec.WithSecondaryKey(keys.Secondary))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

func commonFlags() (*string, *string, *string, *string, *bool, *int) {
	fs := flag.NewFlagSet("common", flag.ContinueOnError)
	projID := fs.String("proj-id", "", "Project ID")
	fundClassName := fs.String("fund-class-name", "", "Fund class name")
	start := fs.String("start", "", "Start date (YYYY-MM-DD) or period (YYYYMM)")
	end := fs.String("end", "", "End date (YYYY-MM-DD) or period (YYYYMM)")
	latest := fs.Bool("latest", false, "Return only latest factsheet data")
	pageSize := fs.Int("page-size", 10, "Items per page")
	_ = fs.Parse(os.Args[3:])
	return projID, fundClassName, start, end, latest, pageSize
}

func parseDate(s string) time.Time {
	if s == "" {
		return time.Time{}
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		log.Fatalf("Invalid date %q: %v", s, err)
	}
	return t
}

func printJSON(v any) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Println(string(data))
}

func handleAMCs(ctx context.Context, client *sec.Client) {
	items, _, err := client.ListAMCs(ctx, 10, "")
	if err != nil {
		log.Fatalf("ListAMCs failed: %v", err)
	}
	printJSON(items)
}

func handleProfiles(ctx context.Context, client *sec.Client, args []string) {
	fs := flag.NewFlagSet("profiles", flag.ExitOnError)
	companyInfo := fs.String("company-info", "", "Company info search")
	fundStatus := fs.String("fund-status", "", "Fund status filter")
	_ = fs.Parse(args)

	items, _, err := client.GetFundProfiles(ctx, sec.ProfileOptions{
		CompanyInfo: *companyInfo,
		FundStatus:  *fundStatus,
		PageSize:    10,
	})
	if err != nil {
		log.Fatalf("GetFundProfiles failed: %v", err)
	}
	printJSON(items)
}

func handleSpecifications(ctx context.Context, client *sec.Client, args []string) {
	projID, fundClassName, _, _, _, pageSize := commonFlags()
	items, _, err := client.GetFundSpecifications(ctx, sec.FeeOptions{
		ProjID:        *projID,
		FundClassName: *fundClassName,
		PageSize:      *pageSize,
	})
	if err != nil {
		log.Fatalf("GetFundSpecifications failed: %v", err)
	}
	printJSON(items)
}

func handleFees(ctx context.Context, client *sec.Client, args []string) {
	projID, fundClassName, _, _, _, pageSize := commonFlags()
	items, _, err := client.GetMutualFundFees(ctx, sec.FeeOptions{
		ProjID:        *projID,
		FundClassName: *fundClassName,
		PageSize:      *pageSize,
	})
	if err != nil {
		log.Fatalf("GetMutualFundFees failed: %v", err)
	}
	printJSON(items)
}

func handleInvolveParties(ctx context.Context, client *sec.Client, args []string) {
	fs := flag.NewFlagSet("involve-parties", flag.ExitOnError)
	projID := fs.String("proj-id", "", "Project ID")
	entityType := fs.String("entity-type", "", "Entity type code")
	pageSize := fs.Int("page-size", 10, "Items per page")
	_ = fs.Parse(args)

	items, _, err := client.GetFundInvolveParties(ctx, sec.InvolvePartyOptions{
		ProjID:     *projID,
		EntityType: *entityType,
		PageSize:   *pageSize,
	})
	if err != nil {
		log.Fatalf("GetFundInvolveParties failed: %v", err)
	}
	printJSON(items)
}

func handleNAV(ctx context.Context, client *sec.Client, args []string) {
	projID, fundClassName, start, end, _, pageSize := commonFlags()
	items, _, err := client.GetDailyNAV(ctx, sec.NAVOptions{
		ProjID:        *projID,
		FundClassName: *fundClassName,
		StartDate:     parseDate(*start),
		EndDate:       parseDate(*end),
		PageSize:      *pageSize,
	})
	if err != nil {
		log.Fatalf("GetDailyNAV failed: %v", err)
	}
	printJSON(items)
}

func handleDividendHistory(ctx context.Context, client *sec.Client, args []string) {
	fs := flag.NewFlagSet("dividend-history", flag.ExitOnError)
	projID := fs.String("proj-id", "", "Project ID")
	classAbbrName := fs.String("class-abbr-name", "", "Class abbreviation")
	start := fs.String("start", "", "Start dividend date (YYYY-MM-DD)")
	end := fs.String("end", "", "End dividend date (YYYY-MM-DD)")
	pageSize := fs.Int("page-size", 10, "Items per page")
	_ = fs.Parse(args)

	items, _, err := client.GetDividendHistory(ctx, sec.DividendHistoryOptions{
		ProjID:        *projID,
		ClassAbbrName: *classAbbrName,
		StartDate:     parseDate(*start),
		EndDate:       parseDate(*end),
		PageSize:      *pageSize,
	})
	if err != nil {
		log.Fatalf("GetDividendHistory failed: %v", err)
	}
	printJSON(items)
}

func handleFactsheetURLs(ctx context.Context, client *sec.Client, args []string) {
	projID, fundClassName, _, _, _, pageSize := commonFlags()
	items, _, err := client.GetFundFactsheetURLs(ctx, sec.FeeOptions{
		ProjID:        *projID,
		FundClassName: *fundClassName,
		PageSize:      *pageSize,
	})
	if err != nil {
		log.Fatalf("GetFundFactsheetURLs failed: %v", err)
	}
	printJSON(items)
}

func handleIPOs(ctx context.Context, client *sec.Client, args []string) {
	projID, _, start, end, latest, pageSize := commonFlags()
	items, _, err := client.GetFundIPOs(ctx, sec.FactsheetOptions{
		ProjID:    *projID,
		StartDate: parseDate(*start),
		EndDate:   parseDate(*end),
		Latest:    *latest,
		PageSize:  *pageSize,
	})
	if err != nil {
		log.Fatalf("GetFundIPOs failed: %v", err)
	}
	printJSON(items)
}

func handleBenchmarks(ctx context.Context, client *sec.Client, args []string) {
	projID, _, start, end, latest, pageSize := commonFlags()
	items, _, err := client.GetFactsheetBenchmarks(ctx, sec.FactsheetOptions{
		ProjID:    *projID,
		StartDate: parseDate(*start),
		EndDate:   parseDate(*end),
		Latest:    *latest,
		PageSize:  *pageSize,
	})
	if err != nil {
		log.Fatalf("GetFactsheetBenchmarks failed: %v", err)
	}
	printJSON(items)
}

func handleSubscriptionMinimums(ctx context.Context, client *sec.Client, args []string) {
	projID, fundClassName, start, end, latest, pageSize := commonFlags()
	items, _, err := client.GetFactsheetSubscriptionRedemptionMinimums(ctx, sec.FactsheetOptions{
		ProjID:        *projID,
		FundClassName: *fundClassName,
		StartDate:     parseDate(*start),
		EndDate:       parseDate(*end),
		Latest:        *latest,
		PageSize:      *pageSize,
	})
	if err != nil {
		log.Fatalf("GetFactsheetSubscriptionRedemptionMinimums failed: %v", err)
	}
	printJSON(items)
}

func handleSubscriptionPeriods(ctx context.Context, client *sec.Client, args []string) {
	projID, fundClassName, start, end, latest, pageSize := commonFlags()
	items, _, err := client.GetFactsheetSubscriptionRedemptionPeriods(ctx, sec.FactsheetOptions{
		ProjID:        *projID,
		FundClassName: *fundClassName,
		StartDate:     parseDate(*start),
		EndDate:       parseDate(*end),
		Latest:        *latest,
		PageSize:      *pageSize,
	})
	if err != nil {
		log.Fatalf("GetFactsheetSubscriptionRedemptionPeriods failed: %v", err)
	}
	printJSON(items)
}

func handleRiskSpectrum(ctx context.Context, client *sec.Client, args []string) {
	projID, _, start, end, latest, pageSize := commonFlags()
	items, _, err := client.GetRiskSpectrum(ctx, sec.FactsheetOptions{
		ProjID:    *projID,
		StartDate: parseDate(*start),
		EndDate:   parseDate(*end),
		Latest:    *latest,
		PageSize:  *pageSize,
	})
	if err != nil {
		log.Fatalf("GetRiskSpectrum failed: %v", err)
	}
	printJSON(items)
}

func handleStatistics(ctx context.Context, client *sec.Client, args []string) {
	projID, fundClassName, start, end, latest, pageSize := commonFlags()
	items, _, err := client.GetFactsheetStatistics(ctx, sec.FactsheetOptions{
		ProjID:        *projID,
		FundClassName: *fundClassName,
		StartDate:     parseDate(*start),
		EndDate:       parseDate(*end),
		Latest:        *latest,
		PageSize:      *pageSize,
	})
	if err != nil {
		log.Fatalf("GetFactsheetStatistics failed: %v", err)
	}
	printJSON(items)
}

func handleDividendPolicy(ctx context.Context, client *sec.Client, args []string) {
	projID, fundClassName, start, end, latest, pageSize := commonFlags()
	items, _, err := client.GetFactsheetDividendPolicy(ctx, sec.FactsheetOptions{
		ProjID:        *projID,
		FundClassName: *fundClassName,
		StartDate:     parseDate(*start),
		EndDate:       parseDate(*end),
		Latest:        *latest,
		PageSize:      *pageSize,
	})
	if err != nil {
		log.Fatalf("GetFactsheetDividendPolicy failed: %v", err)
	}
	printJSON(items)
}

func handleFactsheetFees(ctx context.Context, client *sec.Client, args []string) {
	projID, fundClassName, start, end, latest, pageSize := commonFlags()
	items, _, err := client.GetFactsheetFees(ctx, sec.FactsheetOptions{
		ProjID:        *projID,
		FundClassName: *fundClassName,
		StartDate:     parseDate(*start),
		EndDate:       parseDate(*end),
		Latest:        *latest,
		PageSize:      *pageSize,
	})
	if err != nil {
		log.Fatalf("GetFactsheetFees failed: %v", err)
	}
	printJSON(items)
}

func handlePerformance(ctx context.Context, client *sec.Client, args []string) {
	projID, fundClassName, start, end, latest, pageSize := commonFlags()
	items, _, err := client.GetFactsheetPerformance(ctx, sec.FactsheetOptions{
		ProjID:        *projID,
		FundClassName: *fundClassName,
		StartDate:     parseDate(*start),
		EndDate:       parseDate(*end),
		Latest:        *latest,
		PageSize:      *pageSize,
	})
	if err != nil {
		log.Fatalf("GetFactsheetPerformance failed: %v", err)
	}
	printJSON(items)
}

func handleAssetAllocation(ctx context.Context, client *sec.Client, args []string) {
	projID, fundClassName, start, end, latest, pageSize := commonFlags()
	items, _, err := client.GetAssetAllocation(ctx, sec.FactsheetOptions{
		ProjID:        *projID,
		FundClassName: *fundClassName,
		StartDate:     parseDate(*start),
		EndDate:       parseDate(*end),
		Latest:        *latest,
		PageSize:      *pageSize,
	})
	if err != nil {
		log.Fatalf("GetAssetAllocation failed: %v", err)
	}
	printJSON(items)
}

func handleTop5(ctx context.Context, client *sec.Client, args []string) {
	projID, _, start, end, latest, pageSize := commonFlags()
	items, _, err := client.GetTop5Holdings(ctx, sec.FactsheetOptions{
		ProjID:    *projID,
		StartDate: parseDate(*start),
		EndDate:   parseDate(*end),
		Latest:    *latest,
		PageSize:  *pageSize,
	})
	if err != nil {
		log.Fatalf("GetTop5Holdings failed: %v", err)
	}
	printJSON(items)
}

func handlePortfolio(ctx context.Context, client *sec.Client, args []string) {
	projID, _, start, end, _, pageSize := commonFlags()
	items, _, err := client.GetQuarterlyPortfolio(ctx, sec.OutstandingOptions{
		ProjID:      *projID,
		StartPeriod: *start,
		EndPeriod:   *end,
		PageSize:    *pageSize,
	})
	if err != nil {
		log.Fatalf("GetQuarterlyPortfolio failed: %v", err)
	}
	printJSON(items)
}

func handlePortfolioAssetType(ctx context.Context, client *sec.Client, args []string) {
	projID, _, start, end, _, pageSize := commonFlags()
	items, _, err := client.GetMonthlyPortfolioAssetType(ctx, sec.OutstandingOptions{
		ProjID:      *projID,
		StartPeriod: *start,
		EndPeriod:   *end,
		PageSize:    *pageSize,
	})
	if err != nil {
		log.Fatalf("GetMonthlyPortfolioAssetType failed: %v", err)
	}
	printJSON(items)
}

func init() {
	// Suppress flag package's default error handling on unknown flags.
	flag.CommandLine.Init(os.Args[0], flag.ContinueOnError)
	flag.CommandLine.SetOutput(os.Stderr)
	_ = strings.NewReader("")
}
