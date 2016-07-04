package env

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"testing"
)

func TestGetString_WithExistingValue(t *testing.T) {
	tests := []struct {
		key, existingValue, defaultValue string
	}{
		{"FOO_TEST", "foo", "bar"},
		{"BAR_TEST", "bax", "baz"},
		{"BAZ_TEST", "oof", "rab"},
	}

	for _, tt := range tests {
		tmp := os.Getenv(tt.key)
		defer func(key, defaultValue string) {
			os.Setenv(key, defaultValue)
		}(tt.key, tmp)
		os.Setenv(tt.key, tt.existingValue)

		got := GetString(tt.key, tt.defaultValue)
		equals(t, tt.existingValue, got)
	}
}

func TestGetString_DefaultValue(t *testing.T) {
	tests := []struct {
		key, existingValue, defaultValue string
	}{
		{"FOO_TEST", "foo", ""},
		{"BAR_TEST", "bax", ""},
		{"BAZ_TEST", "oof", ""},
	}

	for _, tt := range tests {
		got := GetString(tt.key, tt.defaultValue)
		equals(t, tt.defaultValue, got)
	}
}

func TestGetInt_WithExistingValue(t *testing.T) {
	tests := []struct {
		key, existingValue string
		defaultValue       int
	}{
		{"FOO_TEST", "1", 2},
		{"BAR_TEST", "2", 3},
		{"BAZ_TEST", "3", 4},
	}

	for _, tt := range tests {
		tmp := os.Getenv(tt.key)
		defer func(key, oldValue string) {
			os.Setenv(key, oldValue)
		}(tt.key, tmp)
		os.Setenv(tt.key, tt.existingValue)

		got := GetInt(tt.key, tt.defaultValue)
		expected, err := strconv.Atoi(tt.existingValue)
		ok(t, err)
		equals(t, expected, got)
	}
}

func TestGetInt_DefaultValue(t *testing.T) {
	tests := []struct {
		key, existingValue string
		defaultValue       int
	}{
		{"FOO_TEST", "", 2},
		{"BAR_TEST", "", 3},
		{"BAZ_TEST", "", 4},
	}

	for _, tt := range tests {
		got := GetInt(tt.key, tt.defaultValue)
		equals(t, tt.defaultValue, got)
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}
