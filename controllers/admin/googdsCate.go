package admin

import (
	"ginxiaomi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsCateController struct {
	BaseController
}

func (con GoodsCateController) Index(c *gin.Context) {
	goodsCateList := []models.GoodsCate{}
	models.DB.Find(&goodsCateList)
	c.HTML(http.StatusOK, "admin/goodsCate/index.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}

func (con GoodsCateController) Add(c *gin.Context) {
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid = 0").Find(&goodsCateList)
	c.HTML(http.StatusOK, "admin/goodsCate/add.html", gin.H {
		"goodsCateList": goodsCateList,
	})
}

func (con GoodsCateController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	pid, err1 := models.Int(c.PostForm("pid"))
	link := c.PostForm("link")
	template := c.PostForm("template")
	subTitle := c.PostForm("subTitle")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort, err2 := models.Int(c.PostForm("sort"))
	status, err3 := models.Int(c.PostForm("status"))

	if err1 != nil || err2 != nil {
		con.Error(c, "参数错误", "/admin/goodsCate")
		return
	}
	if err3 != nil {
		con.Error(c, "sort类型错误", "/admin/goodsCate")
	}

	cateImgDir, _ := models.UploadImg(c, "cate_img")

	goodsCate := models.GoodsCate{
		Title: title,
		Pid: pid,
		Link: link,
		Template: template,
		SubTitle: subTitle,
		Keywords: keywords,
		Description: description,
		Sort: sort,
		Status: status,
		CateImg: cateImgDir,
	}

	err := models.DB.Create(&goodsCate).Error
	if err != nil {
		con.Error(c, "添加失败", "/admin/goodsCate")
		return
	}
	con.Success(c, "添加成功", "/admin/goodsCate")
}

func (con GoodsCateController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goodsCate/edit.html", gin.H {})
}

func (con GoodsCateController) Delete(c *gin.Context) {
	con.Success(c, "删除成功" ,"/admin/goodsCate")
}