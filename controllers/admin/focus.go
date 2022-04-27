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
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{})
}
func (con FocusController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "-add--文章-")
}
