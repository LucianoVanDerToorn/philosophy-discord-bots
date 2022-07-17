package util_test

import (
	"reflect"
	"testing"

	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/util"
)

func TestRemoveEmptyFromStringArray(t *testing.T) {
	input := []string{"a", "b", "", "c", "d", "e", "", ""}
	expected := []string{"a", "b", "c", "d", "e"}
	got := util.RemoveEmptyFromStringArray(input)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("did not get expected output:\ngot:%#v\nexp:%#v", got, expected)
	}
}
