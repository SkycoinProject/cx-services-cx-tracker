package tracker

import "errors"

var (
	errCannotFindUser      = errors.New("tracker service: cannot find application by hash ")
	errUnableToRead        = errors.New("tracker service: unable to query persisted data")
	errUnableToParseConfig = errors.New("tracker service: unable to parse provided cx application configuration data")
)
