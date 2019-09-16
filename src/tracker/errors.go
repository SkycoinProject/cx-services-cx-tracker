package tracker

import "errors"

var (
	errCannotFindApplication  = errors.New("tracker service: cannot find application by hash")
	errUnableToRead           = errors.New("tracker service: unable to query persisted data")
	errMissingMandatoryFields = errors.New("tracker service: missing some mandatory fields")
	errUnableToParseConfig    = errors.New("tracker service: unable to parse provided cx application configuration data")
	errExistingGenesisHash    = errors.New("tracker service: there's already configuration with same genesis hash")
)
