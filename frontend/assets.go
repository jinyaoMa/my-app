package frontend

import "embed"

//go:embed all:desktop/dist/*
var Assets embed.FS
