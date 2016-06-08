package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type languageResourceMap map[string]interface{}

var defaultLang = "en"
var supportedLanguages = map[string]bool{"en": true, "ja": true, "zh": true}
var selectedLang = ""
var languageRegexp = regexp.MustCompile(".*:(..)_?")
var languageResources map[string]languageResourceMap

func initIfRequired() {
	if languageResources == nil {
		loadLanguageResources()
		selectedLang = getLanguagePreference(loadLanguageSettings())
	}
}

func getSelectedLanguage() string {
	return selectedLang
}

func loadLanguageResources() {
	languageResources = make(map[string]languageResourceMap)
	for lang := range supportedLanguages {
		b, err := Asset("../generators/assets/i18n/soracom-api.text." + lang + ".json")
		if err != nil {
			fmt.Printf("warning: unable to load language resource '%s'\n", lang)
			continue
		}

		var data map[string]interface{}
		err = json.Unmarshal(b, &data)
		if err != nil {
			fmt.Printf("warning: unable to parse language resource '%s'\n", lang)
			fmt.Println(err)
			continue
		}
		languageResources[lang] = data
	}
}

type languageSettings struct {
	lang       string
	lcAll      string
	lcMessages string
	language   string
}

func loadLanguageSettings() languageSettings {
	return languageSettings{
		lang:       getLanguagePart(os.Getenv("LANG")),
		lcAll:      getLanguagePart(os.Getenv("LC_ALL")),
		lcMessages: getLanguagePart(os.Getenv("LC_MESSAGES")),
		language:   os.Getenv("LANGUAGE"),
	}
}

// Gets user preference language
// Implements the same behavior with gettext:
// https://www.gnu.org/software/gettext/manual/html_node/Locale-Environment-Variables.html
// https://www.gnu.org/software/gettext/manual/html_node/The-LANGUAGE-variable.html#The-LANGUAGE-variable
func getLanguagePreference(ls languageSettings) string {
	if ls.lang == "C" || ls.lcAll == "C" {
		return defaultLang
	}

	if ls.lcAll != "" {
		if supportedLanguages[ls.lcAll] {
			return ls.lcAll
		}
		altLang := getAltLang(ls.language)
		if altLang != "" && supportedLanguages[altLang] {
			return altLang
		}
		return defaultLang
	}

	if ls.lcMessages != "" {
		if supportedLanguages[ls.lcMessages] {
			return ls.lcMessages
		}
	}

	if ls.lang != "" {
		if supportedLanguages[ls.lang] {
			return ls.lang
		}
		altLang := getAltLang(ls.language)
		if altLang != "" && supportedLanguages[altLang] {
			return altLang
		}
		return defaultLang
	}

	return defaultLang
}

func getLanguagePart(langCountry string) string {
	if len(langCountry) < 2 {
		return langCountry
	}
	return langCountry[0:2]
}

// the parameter ll should be "ja:en" or "ja_JP:en_US"
// this func returns "en" for both the arguments
func getAltLang(ll string) string {
	s := languageRegexp.FindStringSubmatch(ll)
	if len(s) > 1 && s[1] != "" {
		return s[1]
	}
	return ""
}

// TR returns translated text specified by the resourceID
func TR(resourceID string) string {
	initIfRequired()
	r := languageResources[selectedLang]
	s := visit(r, resourceID)
	if s == "" {
		r = languageResources[defaultLang]
		s = visit(r, resourceID)
		if s == "" {
			return resourceID
		}
	}
	return s
}

func visit(data map[string]interface{}, path string) string {
	if data == nil {
		return ""
	}
	i := strings.Index(path, ".")
	if i == -1 {
		obj := data[path]
		switch v := obj.(type) {
		case string:
			return v
		default:
			return ""
		}
	} else {
		name := path[0:i]
		obj := data[name]
		switch v := obj.(type) {
		case map[string]interface{}:
			return visit(v, path[i+1:])
		default:
			return ""
		}
	}
}
