package data

import (
	"database/sql"
	"errors"

	"github.com/go-jet/jet/v2/qrm"
)

var (
	ErrNotFound = errors.New("not found")
)

func isNotFoundError(err error) bool {
	return errors.Is(err, qrm.ErrNoRows) || errors.Is(err, sql.ErrNoRows)
}
