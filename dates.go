package sec

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	thaiYear  = "ปี"
	thaiMonth = "เดือน"
	thaiDay   = "วัน"
	engYear   = "year"
	engYears  = "years"
	engMonth  = "month"
	engMonths = "months"
	engDay    = "day"
	engDays   = "days"
	thaiNone  = "ไม่มี"
	engNone   = "None"
)

var thaiUnitMap = map[string][2]string{
	thaiYear:  {engYear, engYears},
	thaiMonth: {engMonth, engMonths},
	thaiDay:   {engDay, engDays},
}

// TranslateThaiPeriod converts Thai period descriptions to English.
// Input examples: "3 ปี 3 เดือน", "1 เดือน 13 วัน", "6 เดือน", "15 วัน", "ไม่มี", "-"
// Output examples: "3 years 3 months", "1 month 13 days", "6 months", "15 days", "None", "-"
// Returns the original string if it does not match expected Thai patterns.
func TranslateThaiPeriod(period string) string {
	if period == "" || period == "-" {
		return period
	}

	if period == thaiNone {
		return engNone
	}

	lower := strings.ToLower(period)
	if strings.Contains(lower, "year") || strings.Contains(lower, "month") || strings.Contains(lower, "day") {
		return period
	}

	parts := strings.Fields(period)
	if len(parts) == 0 {
		return period
	}

	var result []string
	for i := 0; i < len(parts); i++ {
		part := parts[i]

		num, err := strconv.Atoi(part)
		if err != nil {
			if units, ok := thaiUnitMap[part]; ok {
				translateUnit(&result, units)
			} else {
				result = append(result, part)
			}
		} else {
			result = append(result, strconv.Itoa(num))
		}
	}

	if len(result) == 0 {
		return period
	}

	return strings.Join(result, " ")
}

func translateUnit(result *[]string, units [2]string) {
	if len(*result) == 0 {
		return
	}
	lastIdx := len(*result) - 1
	if n, err := strconv.Atoi((*result)[lastIdx]); err == nil {
		(*result)[lastIdx] = fmt.Sprintf("%d %s", n, pluralize(n, units[0], units[1]))
	}
}

// pluralize returns singular if n == 1, plural otherwise.
func pluralize(n int, singular, plural string) string {
	if n == 1 {
		return singular
	}
	return plural
}

// ConvertBuddhistDate converts a Buddhist calendar date string to Gregorian.
// Expected input format: "DD/MM/YYYY" where year is Buddhist (e.g., "03/04/2557")
// Output format: "YYYY-MM-DD" Gregorian (e.g., "2014-04-03")
// If the year is already Gregorian (< 2500), returns the date converted to ISO format.
// Returns empty string for invalid/empty input.
func ConvertBuddhistDate(dateStr string) string {
	if dateStr == "" || dateStr == "-" {
		return dateStr
	}

	if strings.Contains(dateStr, "-") && len(dateStr) == 10 {
		return dateStr
	}

	parts := strings.Split(dateStr, "/")
	if len(parts) != 3 {
		return dateStr
	}

	day, err1 := strconv.Atoi(parts[0])
	month, err2 := strconv.Atoi(parts[1])
	year, err3 := strconv.Atoi(parts[2])
	if err1 != nil || err2 != nil || err3 != nil {
		return dateStr
	}

	if year > 2500 {
		year -= 543
	}

	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}

// TranslateBenchmarkRemark translates a Thai benchmark remark to English.
// It handles both exact-match templates and variable-percentage patterns.
func TranslateBenchmarkRemark(remark string) string {
	if remark == "" {
		return remark
	}

	if trans, ok := BenchmarkRemarkTranslation[remark]; ok {
		return trans
	}

	prefix := "ปรับด้วยต้นทุนการป้องกันความเสี่ยง "
	suffix := " และปรับด้วยอัตราแลกเปลี่ยน "
	if strings.HasPrefix(remark, prefix) && strings.Contains(remark, suffix) {
		parts := strings.SplitN(strings.TrimPrefix(remark, prefix), suffix, 2)
		if len(parts) == 2 {
			return fmt.Sprintf("Adjusted with %s hedging cost and %s exchange rate", parts[0], parts[1])
		}
	}

	return remark
}

// TranslateBenchmarkName translates a Thai benchmark name to English.
// It handles exact matches and prefix patterns with variable index names.
func TranslateBenchmarkName(name string) string {
	if name == "" {
		return name
	}

	if trans, ok := BenchmarkNameTranslation[name]; ok {
		return trans
	}

	prefix := "ดัชนี "
	if strings.HasPrefix(name, prefix) {
		return strings.TrimPrefix(name, prefix) + " Index"
	}

	prefix2 := "ดัชนีผลตอบแทนรวม "
	if strings.HasPrefix(name, prefix2) {
		return strings.TrimPrefix(name, prefix2) + " Total Return Index"
	}

	prefix3 := "ผลตอบแทนรวมสุทธิของดัชนี"
	if strings.HasPrefix(name, prefix3) {
		remainder := strings.TrimSpace(strings.TrimPrefix(name, prefix3))
		return "Total Return of " + remainder + " Index"
	}

	return name
}
