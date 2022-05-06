package admin

import (
	"fmt"
	"ginxiaomi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsCateAttributeController struct {
	BaseController
}

func (con GoodsCateAttributeController) Index(c *gin.Context) {
	cateId, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入的参数不正确", "/admin/goodsType")
		return
	}
	//获取商品类型属性
	goodsTypeAttributeList := []models.GoodsTypeAttribute{}
	models.DB.Where("cate_id=?", cateId).Find(&goodsTypeAttributeList)
	//获取商品类型属性对应的类型

	goodsType := models.GoodsType{}
	models.DB.Where("id=?", cateId).Find(&goodsType)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/index.html", gin.H{
		"cateId":                 cateId,
		"goodsTypeAttributeList": goodsTypeAttributeList,
		"goodsType":              goodsType,
	})

}

func (con GoodsCateAttributeController) Add(c *gin.Context) {
	cateId, err := models.Int(c.Query("cate_id"))
	if err != nil {
		con.Error(c, "id类型错误", "/admin/goodsType")
		return
	}
	goodsTypeList := []models.GoodsType{}
	models.DB.Where("id = ?", cateId).Find(&goodsTypeList)
	fmt.Println(goodsTypeList)
	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/add.html", gin.H{
		"goodsTypeList": goodsTypeList,
	})
}

func (con GoodsCateAttributeController) DoAdd(c *gin.Context) {
	id, err1 := models.Int(c.PostForm("cate_id"))
	if err1 != nil {
		con.Error(c, "id类型错误",  "/admin/goodsTypeAttribute/add?cate_id="+ models.String(id))
		return
	}
	title := c.PostForm("title")
	attrType, err2 := models.Int(c.PostForm("attr_type"))
	if err2 != nil {
		con.Error(c, "attr_type类型错误",  "/admin/goodsTypeAttribute/add?cate_id="+ models.String(id))
		return
	}
	attrValue := c.PostForm("attr_value")
	sort, err3 := models.Int(c.PostForm("sort"))
	if err3 != nil {
		con.Error(c, "sort类型错误", "/admin/goodsTypeAttribute/add?cate_id="+ models.String(id))
		return
	}
	goodsTypeAttribute := models.GoodsTypeAttribute{
		CateId: id,
		Title: title,
		AttrType: attrType,
		AttrValue: attrValue,
		Sort: sort,
		AddTime: int(models.GetUnix()),
	}
	err4 := models.DB.Create(&goodsTypeAttribute).Error
	if err4 != nil {
		con.Error(c, "添加失败", "/admin/goodsTypeAttribute/add?cate_id="+ models.String(id))
	} else {
		con.Success(c, "创建成功", "/admin/goodsTypeAttribute?id="+ models.String(id))
	}
}