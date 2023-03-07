package routes

import (
	"geekdemo/service/gpt"

	"github.com/gin-gonic/gin"
)

func ChatApi(v1 *gin.RouterGroup) {
	v1.GET("/openai/getModels", Wrapper(gpt.GetModels))
	v1.GET("/openai/getCompletions", Wrapper(gpt.GetCompletions))
}
