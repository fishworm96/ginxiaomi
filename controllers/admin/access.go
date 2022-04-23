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

func (con AccessController) DoEdit(c *gin.Context) {
	id, idErr := models.Int(c.PostForm("id"))
	if idErr != nil {
		con.Error(c, "id类型错误", "/admin/access")
		return
	}
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	moduleType, typeErr := models.Int(c.PostForm("type"))
	if typeErr != nil {
		con.Error(c, "type类型错误", "/admin/access")
		return
	}
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "admin/access/edit?id="+models.String(id))
		return
	}
	actionName := c.PostForm("action_name")
	url := c.PostForm("url")
	moduleId, moduleIdErr := models.Int(c.PostForm("module_id"))
	if moduleIdErr != nil {
		con.Error(c, "模块类型id错误", "admin/access")
		return
	}
	sort, sortErr := models.Int(c.PostForm("sort"))
	if sortErr != nil {
		con.Error(c, "sort类型错误", "/admin/access")
		return
	}
	description := c.PostForm("description")
	status, statusErr := models.Int(c.PostForm("status"))
	if statusErr != nil {
		con.Error(c, "状态类型错误", "admin/access")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)
	access.ModuleName = moduleName
	access.ActionName = actionName
	access.Type = moduleType
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Description = description
	access.Status = status
	err := models.DB.Save(&access).Error
	if err != nil {
		con.Error(c, "修改失败请重试", "/admin/access?id="+models.String(id))
		return
	}
	con.Success(c, "修改成功", "/admin/access")
}

func (con AccessController) Delete(c *gin.Context) {
	id, idErr := models.Int(c.Query("id"))
	if idErr != nil {
		con.Error(c, "id类型错误", "/admin/access")
		return
	}
	access := models.Access{Id: id}
	fmt.Println(access)
	if access.ModuleId == 0 {
		accessList := []models.Access{}
		models.DB.Where("module_id = ?", access.Id).Find(&accessList)
		if len(accessList) > 0 {
			con.Error(c, "请先删除子模块", "admin/access")
		} else {
			models.DB.Delete(&access)
		}
	} else {
		models.DB.Delete(&access)
	}
	con.Success(c, "删除成功", "/admin/access")
}