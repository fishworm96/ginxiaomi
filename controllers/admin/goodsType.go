package admin

import (
	"ginxiaomi/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type GoodsTypeController struct {
	BaseController
}

func (con GoodsTypeController) Index(c *gin.Context) {
	goodsTypeList := []models.GoodsType{}
	models.DB.Find(&goodsTypeList)
	c.HTML(http.StatusOK, "admin/goodsType/index.html", gin.H{
		"goodsTypeList": goodsTypeList,
	})
}

func (con GoodsTypeController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goodsType/add.html", gin.H{})
}

func (con GoodsTypeController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	status, err1 := models.Int(c.PostForm("status"))
	if err1 != nil {
		con.Error(c, "status类型错误", "/admin/goodsType")
		return
	}

	if title == "" {
		con.Error(c, "标题不能为空", "/admin/goodsType/add")
		return
	}

	goodsType := models.GoodsType{}
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status
	goodsType.AddTime = int(models.GetUnix())

	err := models.DB.Create(&goodsType).Error
	if err != nil {
		con.Error(c, "创建角色失败请重试", "/admin/goodsType/add")
	} else {
		con.Success(c, "创建角色成功", "/admin/goodsType/add")
	}
}

func (con GoodsTypeController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据类型错误", "/admin/goodsType")
	} else {
		goodsType := models.GoodsType{Id: id}
		models.DB.Find(&goodsType)
		c.HTML(http.StatusOK, "admin/goodsType/edit.html", gin.H{
			"goodsType": goodsType,
		})
	}
}

func (con GoodsTypeController) DoEdit(c *gin.Context) {
	id, idErr := models.Int(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")

	if idErr != nil {
		con.Error(c, "传入数据类型错误", "/admin/goodsType")
		return
	}
	goodsType := models.GoodsType{Id: id}
	models.DB.Find(&goodsType)
	goodsType.Title = title
	goodsType.Description = description
	goodsTypeErr := models.DB.Save(&goodsType).Error
	if goodsTypeErr != nil {
		con.Error(c, "修改数据失败", "/admin/goodsType/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/goodsType/edit?id="+models.String(id))
	}
}

func (con GoodsTypeController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goodsType")
		return
	}
	goodsTypeAttribute := []models.GoodsTypeAttribute{}
	models.DB.Where("cate_id = ?", id).Find(&goodsTypeAttribute)
	if len(goodsTypeAttribute) > 0 {
		con.Error(c, "请先删除子模块", "/admin/goodsType")
	} else {
		goodsType := models.GoodsType{Id: id}
		models.DB.Delete(&goodsType)
		con.Success(c, "删除数据成功", "/admin/goodsType")
	}
}
