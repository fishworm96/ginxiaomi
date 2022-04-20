package admin

import (
	"ginxiaomi/models"
	"net/http"
	// "strings"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	managerList := []models.Manager{}
	models.DB.Find(&managerList)
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})

}

func (con ManagerController) Add(c *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

func (con ManagerController) DoAdd(c *gin.Context) {
	// username := strings.Trim(c.PostForm("username"), " ")
	// passowrd := models.Md5(c.PostForm("password"))
	// mobile := c.PostForm("mobile")
	// email := c.PostForm("email")

}

func (con ManagerController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{})
}
func (con ManagerController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "-add--文章-")
}
