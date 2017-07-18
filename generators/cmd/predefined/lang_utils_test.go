package cmd

import (
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestGetLanguagePart(t *testing.T) {
	if v := getLanguagePart("ja"); v != "ja" {
		t.Fatalf(`getLanguagePart("ja") should be "ja" but "%s"`, v)
	}
	if v := getLanguagePart("ja_JP"); v != "ja" {
		t.Fatalf(`getLanguagePart("ja_JP") should be "ja" but "%s"`, v)
	}
	if v := getLanguagePart("en_US"); v != "en" {
		t.Fatalf(`getLanguagePart("en_US") should be "en" but "%s"`, v)
	}
	if v := getLanguagePart(""); v != "" {
		t.Fatalf(`getLanguagePart("") should be "" but "%s"`, v)
	}
}

func TestGetAltLang(t *testing.T) {
	if v := getAltLang("ja:en"); v != "en" {
		t.Fatalf(`getAltLang("ja:en") should be "en" but "%s"`, v)
	}
	if v := getAltLang("ja_JP:en_US"); v != "en" {
		t.Fatalf(`getAltLang("ja_JP:en_US") should be "en" but "%s"`, v)
	}
	if v := getAltLang("ja:"); v != "" {
		t.Fatalf(`getAltLang("ja:") should be "" but "%s"`, v)
	}
	if v := getAltLang("ja:x"); v != "" {
		t.Fatalf(`getAltLang("ja:x") should be "" but "%s"`, v)
	}
	if v := getAltLang("japan"); v != "" {
		t.Fatalf(`getAltLang("japan") should be "" but "%s"`, v)
	}
}

func TestGetLanguagePreference(t *testing.T) {
	if v := getLanguagePreference(languageSettings{language: "", lcAll: "", lcMessages: "", lang: ""}); v != "en" {
		t.Fatalf(`getLanguagePreferece() should return default value "en" when no environment variables specified but "%s"`, v)
	}
	if v := getLanguagePreference(languageSettings{language: "", lcAll: "", lcMessages: "", lang: "ja"}); v != "ja" {
		t.Fatalf(`getLanguagePreferece() should return "ja" which is specified in LANG environment variable but "%s"`, v)
	}
	if v := getLanguagePreference(languageSettings{language: "", lcAll: "", lcMessages: "en", lang: "ja"}); v != "en" {
		t.Fatalf(`getLanguagePreferece() should return "en" which is specified in LC_MESSAGES environment variable but "%s"`, v)
	}
	if v := getLanguagePreference(languageSettings{language: "", lcAll: "", lcMessages: "ja", lang: "de"}); v != "ja" {
		t.Fatalf(`getLanguagePreferece() should return "ja" which is specified in LC_MESSAGES environment variable but "%s"`, v)
	}
	if v := getLanguagePreference(languageSettings{language: "", lcAll: "", lcMessages: "de", lang: "ja"}); v != "ja" {
		t.Fatalf(`getLanguagePreferece() should return "ja" which is specified in LANG environment variable because "de" in LC_MESSAGE is not supported language, but "%s" returned`, v)
	}
	if v := getLanguagePreference(languageSettings{language: "", lcAll: "ja", lcMessages: "en", lang: "en"}); v != "ja" {
		t.Fatalf(`getLanguagePreferece() should return "ja" which is specified in LC_ALL environment variable, but "%s" returned`, v)
	}
	if v := getLanguagePreference(languageSettings{language: "", lcAll: "de", lcMessages: "ja", lang: "ja"}); v != "en" {
		t.Fatalf(`getLanguagePreferece() should return "en" which is specified in LC_MESSAGES environment variable because "de" in LC_ALL is not supported language, but "%s" returned`, v)
	}
	if v := getLanguagePreference(languageSettings{language: "de:ja", lcAll: "de", lcMessages: "ja", lang: "ja"}); v != "ja" {
		t.Fatalf(`getLanguagePreferece() should return "ja" which is specified in LANGUAGE as an alternative for "de" specified in LC_ALL environment variable, but "%s" returned`, v)
	}
}

type testStructForVisit struct {
	A string              `yaml:"a"`
	B string              `yaml:"b"`
	C *testStructForVisit `yaml:"c"`
}

func TestVisit(t *testing.T) {
	v := testStructForVisit{
		A: "a1",
		B: "b1",
		C: &testStructForVisit{
			A: "a2",
			B: "b2",
			C: &testStructForVisit{
				A: "a3",
				B: "b3",
			},
		},
	}

	b, err := yaml.Marshal(v)
	if err != nil {
		t.Fatalf("unable to marshal the test object to json")
	}

	var o map[interface{}]interface{}
	err = yaml.Unmarshal(b, &o)
	if err != nil {
		t.Fatalf("unable to unmarshal the test object from json")
	}

	if visit(o, "a") != "a1" {
		t.Fatalf("unable to visit to 'a'")
	}
	if visit(o, "c.b") != "b2" {
		t.Fatalf("unable to visit to 'c.b'")
	}
	if visit(o, "c.c.a") != "a3" {
		t.Fatalf("unable to visit to 'c.c.a'")
	}
	if visit(o, "a.a") != "" {
		t.Fatalf("should not be able to visit to 'a.a'")
	}
	if visit(o, "x.a") != "" {
		t.Fatalf("should not be able to visit to 'x.a'")
	}
}
