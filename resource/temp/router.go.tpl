package router

import (
    "github.com/gin-gonic/gin"

    "{{.ModName}}/{{.Abbr}}/apis/{{.PackageName}}"
)

func register{{.SupStructName}}(router *gin.RouterGroup) {
	router.POST("/{{.RouterName}}s", {{.PackageName}}.Create{{.SupStructName}})
	router.PUT("/{{.RouterName}}s/:id", {{.PackageName}}.Update{{.SupStructName}})
	router.GET("/{{.RouterName}}s/:id", {{.PackageName}}.Get{{.SupStructName}})
	router.GET("/{{.RouterName}}s", {{.PackageName}}.List{{.SupStructName}})
	router.DELETE("/{{.RouterName}}s/:id", {{.PackageName}}.Delete{{.SupStructName}})
}
