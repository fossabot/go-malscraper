package common

import "errors"

var (
	// Err3LettersSearch will throw if search query string is less than 3 letters.
	Err3LettersSearch = errors.New("search query needs at least 3 letters")
	// ErrInvalidMainType will throw if not a valid type.
	ErrInvalidMainType = errors.New("invalid type")
	// ErrInvalidSeason will throw if value is not a valid season name.
	ErrInvalidSeason = errors.New("invalid season name")
	// ErrMissingRedis will throw if redis client is nil.
	ErrMissingRedis = errors.New("missing redis client")
)
