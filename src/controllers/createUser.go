package controllers

import (
	"github.com/TonyGuerra122/meu-primeiro-crud-go/src/configurations/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rota de forma errada")
	ThrowApiError(c, err)
}
