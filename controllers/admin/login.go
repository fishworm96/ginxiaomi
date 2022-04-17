package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (con LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}