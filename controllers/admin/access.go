package admin

import (
	"fmt"
	"ginxiaomi/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccessController struct {
	BaseController
}

func (con AccessController) Index(c *gin.Context) {
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

	fmt.Printf("%#v", accessList)
	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) Add(c *gin.Context) {
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

	fmt.Printf("%#v", accessList)
	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) DoAdd(c *gin.Context) {
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	moduleType, typeErr := models.Int(c.PostForm("type"))
	actionName := strings.Trim(c.PostForm("action_name"), " ")
	url := strings.Trim(c.PostForm("url"), " ")
	moduleId, moduleIdErr := models.Int(c.PostForm("module_id"))
	sort, sortErr := models.Int(c.PostForm("sort"))
	description := c.PostForm("description")
	status, statusErr := models.Int(c.PostForm("status"))
	if typeErr != nil || moduleIdErr != nil || sortErr != nil || statusErr != nil {
		con.Error(c, "参数类型错误", "/admin/access")
		return
	}
	accessList := models.Access{
		ModuleName:  moduleName,
		ActionName:  actionName,
		Type:        moduleType,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}
	models.DB.Create(&accessList)
	con.Success(c, "创建成功", "/admin/access")
}

func (con AccessController) Edit(c *gin.Context) {
	id, IdErr := models.Int(c.Query("id"))
	if IdErr != nil {
		con.Error(c, "id类型错误", "/admin/access")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)

	accessList := []models.Access{}
	models.DB.Where("module_id=0").Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"access": access,
		"accessList": accessList,
	})
}
