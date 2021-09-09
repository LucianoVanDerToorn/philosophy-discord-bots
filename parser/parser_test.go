package parser_test

import (
	"reflect"
	"testing"

	"github.com/lucianonooijen/socrates-discord-bot/parser"
)

func TestParseRequest_Empty(t *testing.T) {
	command, args := parser.ParseRequest("")
	if command != "" {
		t.Errorf("expected command '' but received '%s'", command)
	}
	if args != nil {
		t.Errorf("expected args nil but received %#v", args)
	}
}

func TestParseRequest_Simple(t *testing.T) {
	command, args := parser.ParseRequest("test")
	if command != "test" {
		t.Errorf("expected command 'test' but received '%s'", command)
	}
	if args != nil {
		t.Errorf("expected args nil but received %#v", args)
	}
}

func TestParseRequest_SimpleTrailingSpace(t *testing.T) {
	command, args := parser.ParseRequest("test   ")
	if command != "test" {
		t.Errorf("expected command 'test' but received '%s'", command)
	}
	if args != nil {
		t.Errorf("expected args nil but received %#v", args)
	}
}

func TestParseRequest_WithArg(t *testing.T) {
	command, args := parser.ParseRequest("tester argument")
	if command != "tester" {
		t.Errorf("expected command 'tester' but received '%s'", command)
	}
	expectedArgs := []string{"argument"}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("expected args %#v but received %#v", expectedArgs, args)
	}
}

func TestParseRequest_WithArgs(t *testing.T) {
	command, args := parser.ParseRequest("tester argument1 arg2")
	if command != "tester" {
		t.Errorf("expected command 'tester' but received '%s'", command)
	}
	expectedArgs := []string{"argument1", "arg2"}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("expected args %#v but received %#v", expectedArgs, args)
	}
}

func TestParseRequest_WithArgs_WeirdWhitespaces(t *testing.T) {
	command, args := parser.ParseRequest("tester    argument1  arg2\narg3")
	if command != "tester" {
		t.Errorf("expected command 'tester' but received '%s'", command)
	}
	expectedArgs := []string{"argument1", "arg2\narg3"}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("expected args %#v but received %#v", expectedArgs, args)
	}
}
