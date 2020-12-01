package Week02

import (
	"database/sql"

	"github.com/pkg/errors"
)

func QueryRows() error {

	return errors.Wrap(sql.ErrNoRows, "query rows failed")
}

func IsErrNoRows(err error) bool {
	if errors.Cause(err) == sql.ErrNoRows {
		return true
	}
	return false
}
