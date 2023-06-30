// GetPlatformDefaultConfig gets the defaults for the platform
//go:build !windows && !linux

package OSConfig

//go:build !windows && !linux
func OSConfig() OSConfig {
	return OpenLink{
		GetPlatformDefaultConfig:     "open -- {{filename}}",
		OpenLink: "open -- {{filename}}",
	}
}
