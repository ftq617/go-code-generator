package {{.PackageName}}

import (
    "net/http"
    "strconv"

	"github.com/gin-gonic/gin"

    "{{.ModName}}/{{.Abbr}}/request"
	"{{.ModName}}/{{.Abbr}}/service"
	"{{.ModName}}/pkg/resp"
)


func Create{{.SupStructName}}(c *gin.Context) {
    ctx := c.Request.Context()
	var req request.Create{{.SupStructName}}Req

    if err := c.ShouldBindJSON(&req);err != nil {
		resp.ErrorParam(c, err)
		return
	}
	s := service.NewService()
    id, err := s.{{.SupStructName}}().Create(ctx, &req)
    if err != nil {
    	resp.Error(c, err)
    	return
    }
   	c.JSON(http.StatusOK, gin.H{"id":id})
}


func Delete{{.SupStructName}}(c *gin.Context) {
    ctx := c.Request.Context()
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err !=nil{
        resp.ErrorParam(c, err)
        return
    }

	s := service.NewService()
	if err := s.{{.SupStructName}}().Deleted(ctx, id); err != nil {
		resp.Error(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}



func Update{{.SupStructName}}(c *gin.Context) {
    ctx := c.Request.Context()
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err !=nil{
        resp.ErrorParam(c, err)
        return
    }

    var req request.Update{{.SupStructName}}Req

    if err := c.ShouldBindJSON(&req);err != nil {
		resp.ErrorParam(c, err)
		return
	}
	s := service.NewService()
    err = s.{{.SupStructName}}().Update(ctx, id, &req)
    if err != nil {
       resp.Error(c, err)
       return
    }
    c.Status(http.StatusNoContent)
}


func Get{{.SupStructName}}(c *gin.Context) {
    ctx := c.Request.Context()
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err !=nil{
        resp.ErrorParam(c, err)
        return
    }

    s := service.NewService()
	result, err := s.{{.SupStructName}}().Get(ctx, id)
	if err != nil {
		resp.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}


func List{{.SupStructName}}(c *gin.Context) {
    ctx := c.Request.Context()
    var req request.Query{{.SupStructName}}Req

    if err := c.ShouldBindQuery(&req);err != nil {
		resp.ErrorParam(c, err)
		return
	}
	s := service.NewService()
    result, err := s.{{.SupStructName}}().List(ctx, &req)
    if err != nil {
       resp.Error(c, err)
       return
    }
    c.JSON(http.StatusOK, result)
}
