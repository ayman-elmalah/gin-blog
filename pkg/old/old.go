package old

import (
	"github.com/gin-gonic/gin"
)

var oldList = make(map[string][]string)

func Init() {
	oldList = map[string][]string{}
}

func Set(c *gin.Context) {
	_ = c.Request.ParseForm()

	oldList = c.Request.PostForm
}

func Get() map[string][]string {
	return oldList
}
