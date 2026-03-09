package link

import "errors"

var (
	ErrCannotShortenLink = errors.New("cannot shorten link")
	ErrLinkNotFound      = errors.New("link code not found")
	ErrUrlToLong         = errors.New("url to long")
)
