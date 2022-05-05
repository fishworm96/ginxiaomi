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
	c.HTML(http.StatusOK, "admin/goodsCate/add.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}

func (con GoodsCateController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	pid, err1 := models.Int(c.PostForm("pid"))
	link := c.PostForm("link")
	template := c.PostForm("template")
	subTitle := c.PostForm("sub_title")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort, err2 := models.Int(c.PostForm("sort"))
	status, err3 := models.Int(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.Error(c, "参数错误", "/admin/goodsCate")
		return
	}
	if err2 != nil {
		con.Error(c, "sort类型错误", "/admin/goodsCate")
		return
	}

	cateImgDir, _ := models.UploadImg(c, "cate_img")

	goodsCate := models.GoodsCate{
		Title:       title,
		Pid:         pid,
		Link:        link,
		Template:    template,
		SubTitle:    subTitle,
		Keywords:    keywords,
		Description: description,
		Sort:        sort,
		Status:      status,
		CateImg:     cateImgDir,
		AddTime: int(models.GetUnix()),
	}

	err := models.DB.Create(&goodsCate).Error
	if err != nil {
		con.Error(c, "添加失败", "/admin/goodsCate")
		return
	}
	con.Success(c, "添加成功", "/admin/goodsCate")
}

func (con GoodsCateController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "id类型错误", "/admin/goodsCate")
		return
	}
	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)

	goodsCateList := []models.GoodsCate{}
	models.DB.Find(&goodsCateList)
	c.HTML(http.StatusOK, "admin/goodsCate/edit.html", gin.H{
		"goodsCate":     goodsCate,
		"goodsCateList": goodsCateList,
	})
}

func (con GoodsCateController) DoEdit(c *gin.Context) {
	id, err := models.Int(c.PostForm("id"))
	if err != nil {
		con.Error(c, "id类型错误", "/admin/goodsCate")
		return
	}
	title := c.PostForm("title")
	pid, err1 := models.Int(c.PostForm("pid"))
	link := c.PostForm("link")
	template := c.PostForm("template")
	subTitle := c.PostForm("sub_title")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort, err2 := models.Int(c.PostForm("sort"))
	status, err3 := models.Int(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.Error(c, "参数错误", "/admin/goodsCate")
		return
	}
	if err2 != nil {
		con.Error(c, "sort类型错误", "/admin/goodsCate")
		return
	}

	cateImgDir, _ := models.UploadImg(c, "cate_img")

	goodsCate := models.GoodsCate{Id: id}
	goodsCate.Title = title
	goodsCate.Pid = pid
	goodsCate.Link = link
	goodsCate.Template = template
	goodsCate.SubTitle = subTitle
	goodsCate.Keywords = keywords
	goodsCate.Description = description
	goodsCate.Sort = sort
	goodsCate.Status = status
	goodsCate.CateImg = cateImgDir
	err4 := models.DB.Save(&goodsCate).Error
	if err4 != nil {
		con.Error(c, "修改失败", "/admin/goodsCate/edit?id="+models.String(id))
		return
	}
	con.Success(c, "修改成功", "/admin/goodsCate")
}

func (con GoodsCateController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "id类型错误", "/admin/goodsCate")
	}
	goodsCate := models.GoodsCate{Id: id}
	if goodsCate.Pid == 0 {
		goodsCateList := []models.GoodsCate{}
		models.DB.Where("pid = ?", goodsCate.Id).Find(&goodsCateList)
		if len(goodsCateList) > 0 {
			con.Error(c, "请先删除子模块", "/admin/goodsCate")
			return
		} else {
			models.DB.Delete(&goodsCate)
		}
	} else {
		models.DB.Delete(&goodsCate)
	}
	con.Success(c, "删除成功", "/admin/goodsCate")
}
