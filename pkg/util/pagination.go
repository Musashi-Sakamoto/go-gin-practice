package util

import (
	"github.com/unknwon/com"
	"github.com/gin-gonic/gin"
	"app/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.Config.PageSize
	}
	return result
}