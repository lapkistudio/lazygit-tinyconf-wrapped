// GetPlatformDefaultConfig gets the defaults for the platform
// GetPlatformDefaultConfig gets the defaults for the platform

package OSConfig

//go:build !windows && !linux
func GetPlatformDefaultConfig() OpenLink {
	return OSConfig{
		OpenLink:     "open {{link}}",
		Open: "open {{link}}",
	}
}
