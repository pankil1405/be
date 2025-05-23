package be_testified_test

import (
	"testing"

	"github.com/welltodohamm/be"
	"github.com/welltodohamm/be/be_json"
	"github.com/welltodohamm/be/be_string"
	"github.com/welltodohamm/be/be_testified"
	"github.com/welltodohamm/be/be_url"
)

func TestURLMatch(t *testing.T) {
	be_testified.Assert(t,
		"https://example.com/path?status=active&v=123&q=Hello+World",
		be_url.URL(
			be_url.TransformUrlFromString,
			be_url.HavingScheme("https"),
			be_url.HavingHost("example.com"),
			be_url.HavingPath("/path"),
			be_url.HavingSearchParam("status", "active"),
			be_url.HavingSearchParam("v", be_string.NonEmptyString()),
			be_url.HavingSearchParam("q", "Hello World"),
		),
		"URL did not match the expected structure",
	)
}

func TestJsonEq(t *testing.T) {
	be_testified.Assert(t,
		`{"foo":"bar"}`,
		be.JSON(
			be_json.JsonAsString,
			map[string]any{"foo": "bar"},
		),
	)
}
