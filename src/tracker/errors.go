package tracker

import "errors"

var (
	errCannotFindUser = errors.New("tracker service: cannot find application by hash ")
	errUnableToRead   = errors.New("user service: unable to query persisted data")
)
