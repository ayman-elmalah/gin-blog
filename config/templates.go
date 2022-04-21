package config

func Templates() map[string]string {
	return map[string]string{
		"pattern":   "templates/**/**/*",
		"extension": "tmpl",
	}
}
