package tasks

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestFromScraping(t *testing.T) {
	tests := []struct {
		input, output string
	}{
		{"AccountList", "sub-accounts"},
		{"AccountList", SubAccounts},
		{"CheckingCurrent", "current-transactions"},
		{"CheckingCurrent", Current},
		{"CheckingHistoric", "historic-transactions"},
		{"CheckingHistoric", Historic},
		{"VistaCurrent", "current-transactions"},
		{"VistaCurrent", Current},
		{"VistaHistoric", "historic-transactions"},
		{"VistaHistoric", Historic},
		{"ClCurrent", "current-transactions"},
		{"ClCurrent", Current},
		{"ClHistoric", "historic-transactions"},
		{"ClHistoric", Historic},
		{"CcNationalCurrentTransactions", "current-transactions"},
		{"CcNationalCurrentTransactions", Current},
		{"CcNationalHistoricTransactions", "historic-transactions"},
		{"CcNationalHistoricTransactions", Historic},
		{"CcInternationalCurrentTransactions", "current-transactions"},
		{"CcInternationalCurrentTransactions", Current},
		{"CcInternationalHistoricTransactions", "historic-transactions"},
		{"CcInternationalHistoricTransactions", Historic},
	}

	for _, tt := range tests {
		got, err := FromScraping(tt.input)
		if err != nil {
			t.Fatal(err)
		}

		equals(t, tt.output, got)
	}
}

func TestTransformstringID_InvalidTaskIDS(t *testing.T) {
	tests := []struct {
		input, output string
	}{
		{"invalid-task-id", "unknown"},
		{"invalid-task-id", Unknown},
		{"foo bar", "unknown"},
		{"foo bar", Unknown},
	}

	for _, tt := range tests {
		got, err := FromScraping(tt.input)
		if err == nil {
			t.Errorf("expected to return an error, input: '%s', got: '%s'", tt.input, got)
		}

		if got != Unknown {
			t.Errorf("expected to return an empty string, input: '%s', got: '%s'", tt.input, got)
		}
		equals(t, tt.output, got)
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
