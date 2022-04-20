package admin

import (
	"fmt"
	"ginxiaomi/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(c *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
}

func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

func (con RoleController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")

	if title == "" {
		con.Error(c, "标题不能为空", "/admin/role/add")
		return
	}

	role := models.Role{}
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(models.GetUnix())

	err := models.DB.Create(&role).Error
	if err != nil {
		con.Error(c, "创建角色失败请重试", "/admin/role/add")
	} else {
		con.Success(c, "创建角色成功", "/admin/role/add")
	}
}

func (con RoleController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据类型错误", "/admin/role")
	} else {
		role := models.Role{Id: id}
		models.DB.Find(&role)
		fmt.Println(role)
		c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": role,
		})
	}
}

func (con RoleController) DoEdit(c *gin.Context) {
	id, idErr := models.Int(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")

	if idErr != nil {
		con.Error(c, "传入数据类型错误", "/admin/role")
		return
	}
	role := models.Role{Id: id}
	models.DB.Find(&role)
	role.Title = title
	role.Description = description
	roleErr := models.DB.Save(&role).Error
	if roleErr != nil {
		con.Error(c, "修改数据失败", "/admin/role/edit?id=" + models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/role/edit?id=" + models.String(id))
	}
}

func (con RoleController) Delete(c *gin.Context) {
	id, idErr := models.Int(c.Query("id"))
	if idErr != nil {
		con.Error(c, "删除失败", "/admin/role")
	} else {
		role := models.Role{Id: id}
		models.DB.Delete(&role)
			con.Error(c, "删除成功", "/admin/role")
	}
}