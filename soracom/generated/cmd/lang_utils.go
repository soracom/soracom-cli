package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type languageResourceMap map[interface{}]interface{}

var defaultLang = "en"
var supportedLanguages = map[string]bool{"en": true, "ja": true, "zh": true}
var selectedLang = ""
var languageRegexp = regexp.MustCompile(".*:(..)_?")
var apiResources map[string]languageResourceMap
var cliResources map[string]languageResourceMap

func initIfRequired() {
	if selectedLang == "" {
		selectedLang = getLanguagePreference(loadLanguageSettings())
	}
	if apiResources == nil {
		loadAPIResources()
	}
	if cliResources == nil {
		loadCLIResources()
	}
}

func getSelectedLanguage() string {
	return selectedLang
}

func loadAPIResources() {
	apiResources = make(map[string]languageResourceMap)
	for lang := range supportedLanguages {
		b, err := Asset("../generators/assets/soracom-api." + lang + ".yaml")
		if err != nil {
			fmt.Printf("warning: unable to load API language resource '%s'\n", lang)
			continue
		}

		var data map[interface{}]interface{}
		err = yaml.Unmarshal(b, &data)
		if err != nil {
			fmt.Printf("warning: unable to parse API language resource '%s'\n", lang)
			fmt.Println(err)
			continue
		}
		apiResources[lang] = data
	}
}

func loadCLIResources() {
	cliResources = make(map[string]languageResourceMap)
	for lang := range supportedLanguages {
		b, err := Asset("../generators/assets/cli/" + lang + ".yaml")
		if err != nil {
			fmt.Printf("warning: unable to load CLI language resource '%s'\n", lang)
			continue
		}

		var data map[interface{}]interface{}
		err = yaml.Unmarshal(b, &data)
		if err != nil {
			fmt.Printf("warning: unable to parse CLI language resource '%s'\n", lang)
			fmt.Println(err)
			continue
		}
		cliResources[lang] = data
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

func TRAPI(pathAndMethodAndField string) string {
	initIfRequired()
	r := apiResources[selectedLang]
	s := getStringResource(r, pathAndMethodAndField)
	if s == "" {
		r = apiResources[defaultLang]
		s = getStringResource(r, pathAndMethodAndField)
		if s == "" {
			return pathAndMethodAndField
		}
	}
	return s
}

func getStringResource(data map[interface{}]interface{}, pathAndMethodAndField string) string {
	if data == nil || len(data) == 0 {
		return ""
	}

	pmf := strings.Split(pathAndMethodAndField, ":")
	if len(pmf) < 3 {
		return ""
	}

	paths := data["paths"].(map[interface{}]interface{})
	methods := paths[pmf[0]].(map[interface{}]interface{})
	methodInfo := methods[pmf[1]].(map[interface{}]interface{})

	str := methodInfo[pmf[2]].(string)
	return str
}

func TRCLI(resourceID string) string {
	initIfRequired()
	r := cliResources[selectedLang]
	s := visit(r, resourceID)
	if s == "" {
		r = cliResources[defaultLang]
		s = visit(r, resourceID)
		if s == "" {
			return resourceID
		}
	}
	return s
}

/*
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
*/

func visit(data map[interface{}]interface{}, path string) string {
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
		case map[interface{}]interface{}:
			return visit(v, path[i+1:])
		default:
			return ""
		}
	}
}
