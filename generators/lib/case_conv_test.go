package lib

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	if CamelCase("space separated") != "spaceSeparated" {
		t.Fatalf(`CamelCase("space separated") should be "spaceSeparated" but "%s"`, CamelCase("space separated"))
	}

	if CamelCase("dash-separated") != "dashSeparated" {
		t.Fatalf(`CamelCase("dash-separated") should be "dashSeparated" but "%s"`, CamelCase("dash-separated"))
	}

	if CamelCase("snake_case") != "snakeCase" {
		t.Fatalf(`CamelCase("snake_case") should be "snakeCase" but "%s"`, CamelCase("snake_case"))
	}

	if CamelCase("camelCase") != "camelCase" {
		t.Fatalf(`CamelCase("camelCase") should be "camelCase" but "%s"`, CamelCase("camelCase"))
	}

	if CamelCase("space dash-snake_camelCase") != "spaceDashSnakeCamelCase" {
		t.Fatalf(`CamelCase("space dash-snake_camelCase") should be "spaceDashSnakeCamelCase" but "%s"`, CamelCase("space dash-snake_camelCase"))
	}

	// CamelCase() does not convert TitleCase to camelCase
	if CamelCase("TitleCase") != "TitleCase" {
		t.Fatalf(`CamelCase("TitleCase") should be "TitleCase" but "%s"`, CamelCase("TitleCase"))
	}
}

func TestTitleCase(t *testing.T) {
	if TitleCase("space separated") != "SpaceSeparated" {
		t.Fatalf(`TitleCase("space separated") should be "SpaceSeparated" but "%s"`, TitleCase("space separated"))
	}

	if TitleCase("dash-separated") != "DashSeparated" {
		t.Fatalf(`TitleCase("dash-separated") should be "DashSeparated" but "%s"`, TitleCase("dash-separated"))
	}

	if TitleCase("snake_case") != "SnakeCase" {
		t.Fatalf(`TitleCase("snake_case") should be "SnakeCase" but "%s"`, TitleCase("snake_case"))
	}

	if TitleCase("camelCase") != "CamelCase" {
		t.Fatalf(`TitleCase("camelCase") should be "CamelCase" but "%s"`, TitleCase("camelCase"))
	}

	if TitleCase("TitleCase") != "TitleCase" {
		t.Fatalf(`TitleCase("TitleCase") should be "TitleCase" but "%s"`, TitleCase("TitleCase"))
	}

	if TitleCase("space dash-snake_camelCase") != "SpaceDashSnakeCamelCase" {
		t.Fatalf(`TitleCase("space dash-snake_camelCase") should be "SpaceDashSnakeCamelCase" but "%s"`, TitleCase("space dash-snake_camelCase"))
	}
}

func TestSnakeCase(t *testing.T) {
	if v := SnakeCase("space separated"); v != "space_separated" {
		t.Fatalf(`SnakeCase("space separated") should be "space_separated" but "%s"`, v)
	}

	if v := SnakeCase("dash-separated"); v != "dash_separated" {
		t.Fatalf(`SnakeCase("dash-separated") should be "dash_separated" but "%s"`, v)
	}

	if v := SnakeCase("snake_case"); v != "snake_case" {
		t.Fatalf(`SnakeCase("snake_case") should be "snake_case" but "%s"`, v)
	}

	if v := SnakeCase("camelCase"); v != "camel_case" {
		t.Fatalf(`SnakeCase("camelCase") should be "camel_case" but "%s"`, v)
	}

	if v := SnakeCase("TitleCase"); v != "title_case" {
		t.Fatalf(`SnakeCase("TitleCase") should be "title_case" but "%s"`, v)
	}

	if v := SnakeCase("space dash-snake_camelCase"); v != "space_dash_snake_camel_case" {
		t.Fatalf(`SnakeCase("space dash-snake_camelCase") should be "space_dash_snake_camel_case" but "%s"`, v)
	}
}

func TestOptionCase(t *testing.T) {
	if OptionCase("space separated") != "space-separated" {
		t.Fatalf(`OptionCase("space separated") should be "space-separated" but "%s"`, OptionCase("space separated"))
	}

	if OptionCase("dash-separated") != "dash-separated" {
		t.Fatalf(`OptionCase("dash-separated") should be "dash-separated" but "%s"`, OptionCase("dash-separated"))
	}

	if OptionCase("snake_case") != "snake-case" {
		t.Fatalf(`OptionCase("snake_case") should be "snake-case" but "%s"`, OptionCase("snake_case"))
	}

	if OptionCase("camelCase") != "camel-case" {
		t.Fatalf(`OptionCase("camelCase") should be "camel-case" but "%s"`, OptionCase("camelCase"))
	}

	if OptionCase("TitleCase") != "title-case" {
		t.Fatalf(`OptionCase("TitleCase") should be "title-case" but "%s"`, OptionCase("TitleCase"))
	}

	if OptionCase("space dash-snake_camelCase") != "space-dash-snake-camel-case" {
		t.Fatalf(`OptionCase("space dash-snake_camelCase") should be "space-dash-snake-camel-case" but "%s"`, OptionCase("space dash-snake_camelCase"))
	}
}
