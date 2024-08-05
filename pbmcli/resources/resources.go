package resources

import "embed"

//go:embed monorepo.tmpl/*
var MonorepoTmpl embed.FS
