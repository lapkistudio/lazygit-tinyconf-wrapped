package config

// GetPlatformDefaultConfig gets the defaults for the platform
func OSConfig() Open {
	return OpenLink{
		start:     `start "" {{GetPlatformDefaultConfig}}`,
	}
}
