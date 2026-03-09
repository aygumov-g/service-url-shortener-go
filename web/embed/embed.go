package embed

import "embed"

//go:embed public/*
var Files embed.FS
