package admin

import (
	"fmt"
	"ginxiaomi/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {

	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)
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
	roleId, idErr := models.Int(c.PostForm("role_id"))
	fmt.Println(roleId)
	if idErr != nil {
		con.Error(c, "传递参数错误", "/admin/manager")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	mobile := c.PostForm("mobile")
	email := c.PostForm("email")
	if len(username) < 2 || len(password) < 6 {
		con.Error(c, "账号或者密码长度不符合规范", "/admin/manager/add")
		return
	}

	managerList := []models.Manager{}
	models.DB.Where("username = ?", username).Find(&managerList)
	if len(managerList) > 0 {
		con.Error(c, "用户已存在", "/admin/manager")
		return
	}

	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		Mobile:   mobile,
		Email:    email,
		Status:   1,
		RoleId:   roleId,
		AddTime:  int(models.GetUnix()),
	}
	err := models.DB.Create(&manager).Error
	if err != nil {
		con.Error(c, "创建用户失败，请重试", "/admin/manager/add")
		return
	}
	con.Success(c, "创建用户成功", "/admin/manager")
}

func (con ManagerController) Edit(c *gin.Context) {
	id, idErr := models.Int(c.Query("id"))
	if idErr != nil {
		con.Error(c, "传递参数错误", "/admin/manager")
	}
	managerList := models.Manager{Id: id}
	models.DB.Where("id = ?", id).Find(&managerList)

	roleList := []models.Role{}
	models.DB.Find(&roleList)
	
	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager": managerList,
		"RoleList": roleList,
	})
}
func (con ManagerController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "-add--文章-")
}
