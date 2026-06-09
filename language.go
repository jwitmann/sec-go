package sec

// Language represents the preferred display language for bilingual fields.
type Language string

const (
	LanguageThai    Language = "th"
	LanguageEnglish Language = "en"
)

func normalizeLanguage(lang Language) Language {
	switch lang {
	case LanguageThai, "TH", "tha":
		return LanguageThai
	case LanguageEnglish, "EN", "eng":
		return LanguageEnglish
	default:
		return LanguageEnglish
	}
}

// Name returns the company name in the client's preferred language.
// Falls back to the other language if the preferred one is empty.
func (a AMC) Name(lang Language) string {
	return pickString(lang, a.CompNameTH, a.CompNameEN)
}

// Name returns the fund name in the client's preferred language.
// Falls back to the other language if the preferred one is empty.
func (p FundProfile) Name(lang Language) string {
	return pickString(lang, p.ProjNameTH, p.ProjNameEN)
}

// CompanyName returns the AMC name in the client's preferred language.
func (p FundProfile) CompanyName(lang Language) string {
	return pickString(lang, p.CompNameTH, p.CompNameEN)
}

// EntityName returns the entity name in the client's preferred language.
func (f FundInvolveParty) EntityName(lang Language) string {
	return pickString(lang, f.EntityNameTH, f.EntityNameEN)
}

func pickString(lang Language, thai, english string) string {
	lang = normalizeLanguage(lang)
	if lang == LanguageThai {
		if thai != "" {
			return thai
		}
		return english
	}
	if english != "" {
		return english
	}
	return thai
}
