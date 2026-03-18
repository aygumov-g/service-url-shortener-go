package link

import "errors"

var (
	ErrCustomCodeAlreadyExists = errors.New("custom code is already exists")
	ErrCannotShortenLink       = errors.New("cannot shorten link")
	ErrLinkNotFound            = errors.New("link code not found")
	ErrUrlToLong               = errors.New("url to long")
)
