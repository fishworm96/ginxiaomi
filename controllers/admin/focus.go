package admin

import (
	"fmt"
	"ginxiaomi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FocusController struct{
	BaseController
}

func (con FocusController) Index(c *gin.Context) {
	focusList := []models.Focus{}
	models.DB.Find(&focusList)
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}

func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

func (con FocusController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	focusType, err1 := models.Int(c.PostForm("focus_type"))
	sort, err2 := models.Int(c.PostForm("sort"))
	status, err3 := models.Int(c.PostForm("status"))
	if err1 != nil || err2 != nil || err3 != nil {
		con.Error(c, "参数错误", "/admin/focus/add")
		return
	}
	link := c.PostForm("link")
	focusImg, err4 := models.UploadImg(c, "focus_img")
	if err4 != nil {
		fmt.Println(err4)
	}
	focusList := models.Focus{
		Title: title,
		FocusType: focusType,
		FocusImg: focusImg,
		Link: link,
		Sort: sort,
		Status: status,
		AddTime: int(models.GetUnix()),
	}

	err5 := models.DB.Create(&focusList).Error
	if err5 != nil {
		con.Error(c, "增加轮播图失败", "/admin/focus/add")
	} else {
		con.Success(c, "增加轮播图成功", "/admin/focus")
	}
}

func (con FocusController) Edit(c *gin.Context) {
	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "参数错误", "/admin/focus")
		return
	}
	focus := models.Focus{}
	models.DB.Where("id = ?", id).Find(&focus)
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}

func (con FocusController) DoEdit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	title := c.PostForm("title")
	focusType, err1 := models.Int(c.PostForm("focus_type"))
	sort, err2 := models.Int(c.PostForm("sort"))
	status, err3 := models.Int(c.PostForm("status"))
	if err !=nil || err1 != nil || err2 != nil || err3 != nil {
		con.Error(c, "参数错误", "/admin/focus/add")
		return
	}
	link := c.PostForm("link")
	focusImg, err4 := models.UploadImg(c, "focus_img")
	if err4 != nil {
		fmt.Println(err4)
	}
	focus := models.Focus{Id: id}
	focus.Title = title
	focus.FocusType = focusType
	focus.FocusImg = focusImg
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	err5 := models.DB.Save(&focus).Error
	if err5 != nil {
		con.Error(c, "修改失败", "/admin/focus?id="+models.String(id))
	} else {
		con.Success(c, "修改成功", "/admin/focus")
	}
}

func (con FocusController) Delete(c *gin.Context) {
	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "参数错误", "/admin/focus")
		return
	}
	focus := models.Focus{Id: id}
	err2 := models.DB.Delete(&focus).Error
	if err2 != nil {
		con.Error(c, "删除失败", "/admin/focus")
	} else {
		con.Success(c, "删除成功", "/admin/focus")
	}
	
}
