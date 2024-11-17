package controllers

import (
	"github.com/TonyGuerra122/meu-primeiro-crud-go/src/configurations/rest_err"
	"github.com/gin-gonic/gin"
)

func ThrowApiError(c *gin.Context, r *rest_err.RestErr) {
	c.JSON(int(r.Code), r)
}
