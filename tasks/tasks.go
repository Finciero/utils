package tasks

import (
	"strings"

	"github.com/Finciero/errors"
)

// SubAccounts ...
// Current ...
// Historic ...
const (
	Unknown     = "unknown"
	SubAccounts = "sub-accounts"
	Current     = "current-transactions"
	Historic    = "historic-transactions"
)

var (
	keywordToID = map[string]string{
		"accountlist": SubAccounts,
		"current":     Current,
		"historic":    Historic,
	}
)

// FromScraping ...
func FromScraping(str string) (string, error) {
	received := strings.ToLower(str)
	for keyword, task := range keywordToID {
		if strings.Contains(received, keyword) {
			return task, nil
		}
	}

	err := errors.InternalServer(
		"invalid task id",
		errors.SetMeta(errors.Meta{
			"task_id": received,
		}),
	)

	return Unknown, err
}
