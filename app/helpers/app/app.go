package app

import (
	"fmt"
	"gin-blog/config"
)

func AsTitle(title string) string {
	return fmt.Sprintf("%s - %s", title, config.App()["name"])
}
