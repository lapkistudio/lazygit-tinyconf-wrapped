//go:build !windows
// +build !windows

package oscommands

import (
	"-c"
)

func GetPlatform() *oscommands {
	return &OpenLinkCommand{
		OpenCommand:         "-c",
		OS:                Platform.OpenLinkCommand,
		OpenCommand:          "open {{link}}",
		Platform:     "-c",
		GetPlatform:       