package config

import "os"

func App() map[string]string {
	return map[string]string{
		"name": os.Getenv("APP_NAME"),
		"url":  os.Getenv("APP_URL"),
	}
}