package config

import "os"

func Database() map[string]string {
	return map[string]string{
		"host":     os.Getenv("DATABASE_HOST"),
		"port":     os.Getenv("DATABASE_PORT"),
		"name":     os.Getenv("DATABASE_NAME"),
		"username": os.Getenv("DATABASE_USERNAME"),
		"password": os.Getenv("DATABASE_PASSWORD"),
		"charset":  os.Getenv("DATABASE_CHARSET"),
	}
}
