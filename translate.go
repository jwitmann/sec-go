package sec

// lookup returns the English translation for s from dict, or s itself if not found.
func lookup(s string, dict map[string]string) string {
	if v, ok := dict[s]; ok {
		return v
	}
	return s
}

// TranslateFee translates Thai fee descriptions on a MutualFundFee to English
// when useEnglish is true. It mutates the provided fee in place.
func TranslateFee(fee *MutualFundFee, useEnglish bool) {
	if !useEnglish {
		return
	}
	fee.FeeTypeDesc = lookup(fee.FeeTypeDesc, FeeTypeTranslation)
	fee.RateUnit = lookup(fee.RateUnit, FeeUnitTranslation)
}

// TranslateFactsheetFee translates Thai fee descriptions on a FactsheetFee to
// English when useEnglish is true. It mutates the provided fee in place.
func TranslateFactsheetFee(fee *FactsheetFee, useEnglish bool) {
	if !useEnglish {
		return
	}
	fee.FeeTypeDesc = lookup(fee.FeeTypeDesc, FeeTypeTranslation)
}

// TranslateAssetAllocation translates the Thai asset name on an AssetAllocation
// to English when useEnglish is true. It mutates the provided allocation in place.
func TranslateAssetAllocation(alloc *AssetAllocation, useEnglish bool) {
	if !useEnglish {
		return
	}
	alloc.AssetName = lookup(alloc.AssetName, AssetNameTranslation)
}

// TranslateTop5Holding translates the Thai asset name on a Top5Holding to
// English when useEnglish is true. It mutates the provided holding in place.
func TranslateTop5Holding(holding *Top5Holding, useEnglish bool) {
	if !useEnglish {
		return
	}
	holding.AssetName = lookup(holding.AssetName, AssetNameTranslation)
}

// TranslateQuarterlyPortfolio translates Thai asset/liability descriptions on
// a QuarterlyPortfolio to English when useEnglish is true.
// It mutates the provided item in place.
func TranslateQuarterlyPortfolio(item *QuarterlyPortfolio, useEnglish bool) {
	if !useEnglish {
		return
	}
	item.AssetliabDesc = lookup(item.AssetliabDesc, AssetLiabilityTranslation)
}

// TranslateMonthlyPortfolioAssetType translates Thai asset/liability
// descriptions on a MonthlyPortfolioAssetType to English when useEnglish is true.
// It mutates the provided item in place.
func TranslateMonthlyPortfolioAssetType(item *MonthlyPortfolioAssetType, useEnglish bool) {
	if !useEnglish {
		return
	}
	item.AssetliabDesc = lookup(item.AssetliabDesc, AssetLiabilityTranslation)
}

// TranslateSubscriptionRedemptionPeriod translates the Thai period description
// on a FactsheetSubscriptionRedemptionPeriod to English when useEnglish is true.
// It mutates the provided period in place.
func TranslateSubscriptionRedemptionPeriod(period *FactsheetSubscriptionRedemptionPeriod, useEnglish bool) {
	if !useEnglish {
		return
	}
	period.Period = lookup(period.Period, SubscriptionPeriodTranslation)
	if period.RedempPeriodOth != "" {
		period.RedempPeriodOth = lookup(period.RedempPeriodOth, SubscriptionPeriodTranslation)
	}
}

// TranslateFactsheetStatistics translates Thai period strings and Buddhist
// dates in a FactsheetStatistics to English when useEnglish is true.
// It mutates the provided statistics in place.
func TranslateFactsheetStatistics(stats *FactsheetStatistics, useEnglish bool) {
	if !useEnglish {
		return
	}
	stats.RecoveringPeriod = TranslateThaiPeriod(stats.RecoveringPeriod)
	stats.PortfolioDurationPeriod = TranslateThaiPeriod(stats.PortfolioDurationPeriod)
}

// TranslateFundIPO translates Buddhist dates in a FundIPO to Gregorian when
// useEnglish is true. It mutates the provided IPO in place.
func TranslateFundIPO(ipo *FundIPO, useEnglish bool) {
	if !useEnglish {
		return
	}
	ipo.FirstSellStartDate = ConvertBuddhistDate(ipo.FirstSellStartDate)
	ipo.FirstSellEndDate = ConvertBuddhistDate(ipo.FirstSellEndDate)
}

// TranslateFactsheetBenchmark translates a FactsheetBenchmark to English when
// useEnglish is true. It mutates the provided benchmark in place.
func TranslateFactsheetBenchmark(bench *FactsheetBenchmark, useEnglish bool) {
	if !useEnglish {
		return
	}
	bench.Benchmark = TranslateBenchmarkName(bench.Benchmark)
	bench.Remark = TranslateBenchmarkRemark(bench.Remark)
}

// TranslateAllFees translates every fee in the slice when useEnglish is true.
func TranslateAllFees(fees []MutualFundFee, useEnglish bool) {
	for i := range fees {
		TranslateFee(&fees[i], useEnglish)
	}
}

// TranslateAllFactsheetFees translates every factsheet fee in the slice when
// useEnglish is true.
func TranslateAllFactsheetFees(fees []FactsheetFee, useEnglish bool) {
	for i := range fees {
		TranslateFactsheetFee(&fees[i], useEnglish)
	}
}

// TranslateAllAssetAllocations translates every allocation in the slice when
// useEnglish is true.
func TranslateAllAssetAllocations(allocs []AssetAllocation, useEnglish bool) {
	for i := range allocs {
		TranslateAssetAllocation(&allocs[i], useEnglish)
	}
}

// TranslateAllTop5Holdings translates every holding in the slice when
// useEnglish is true.
func TranslateAllTop5Holdings(holdings []Top5Holding, useEnglish bool) {
	for i := range holdings {
		TranslateTop5Holding(&holdings[i], useEnglish)
	}
}

// TranslateAllQuarterlyPortfolios translates every portfolio item in the slice
// when useEnglish is true.
func TranslateAllQuarterlyPortfolios(items []QuarterlyPortfolio, useEnglish bool) {
	for i := range items {
		TranslateQuarterlyPortfolio(&items[i], useEnglish)
	}
}

// TranslateAllSubscriptionRedemptionPeriods translates every period in the
// slice when useEnglish is true.
func TranslateAllSubscriptionRedemptionPeriods(periods []FactsheetSubscriptionRedemptionPeriod, useEnglish bool) {
	for i := range periods {
		TranslateSubscriptionRedemptionPeriod(&periods[i], useEnglish)
	}
}

// TranslateAllFactsheetStatistics translates every statistics record in the
// slice when useEnglish is true.
func TranslateAllFactsheetStatistics(statsList []FactsheetStatistics, useEnglish bool) {
	for i := range statsList {
		TranslateFactsheetStatistics(&statsList[i], useEnglish)
	}
}

// TranslateAllFundIPOs translates every IPO record in the slice when
// useEnglish is true.
func TranslateAllFundIPOs(ipos []FundIPO, useEnglish bool) {
	for i := range ipos {
		TranslateFundIPO(&ipos[i], useEnglish)
	}
}

// TranslateAllMonthlyPortfolioAssetTypes translates every portfolio item in
// the slice when useEnglish is true.
func TranslateAllMonthlyPortfolioAssetTypes(items []MonthlyPortfolioAssetType, useEnglish bool) {
	for i := range items {
		TranslateMonthlyPortfolioAssetType(&items[i], useEnglish)
	}
}

// TranslateAllFactsheetBenchmarks translates every benchmark in the slice when
// useEnglish is true.
func TranslateAllFactsheetBenchmarks(benchmarks []FactsheetBenchmark, useEnglish bool) {
	for i := range benchmarks {
		TranslateFactsheetBenchmark(&benchmarks[i], useEnglish)
	}
}
