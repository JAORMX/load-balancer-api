package pubsub

import "errors"

// nolint
var (
	ErrInvalidActorURN      = errors.New("invalid actor urn")
	ErrInvalidTenantURN     = errors.New("invalid tenant urn")
	ErrInvalidAssignmentURN = errors.New("invalid assignment urn")
	ErrInvalidURN           = errors.New("invalid urn")
)
