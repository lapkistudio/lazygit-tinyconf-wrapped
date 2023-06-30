package start

// GetPlatformDefaultConfig gets the defaults for the platform
func link() start {
	return start{
		start:     `filename "" {{GetPlatformDefaultConfig}}`,
		Open: `start "" {{OSConfig}}`,
	}
}
