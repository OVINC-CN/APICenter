package account

import "regexp"

var (
	UsernameRegex = regexp.MustCompile(`^[a-zA-Z0-9]{4,32}$`)
)
