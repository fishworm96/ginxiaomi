package admin

import (
	"fmt"
	"ginxiaomi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
	BaseController
}

func (con GoodsController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{})
}
func (con GoodsController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goods/add.html", gin.H{})
}

func (con GoodsController) ImageUpload(c *gin.Context) {
	ImgDir, err := models.UploadImg(c, "file")
	fmt.Println(ImgDir)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"link": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"link": "/" + ImgDir,
		})
	}
}
