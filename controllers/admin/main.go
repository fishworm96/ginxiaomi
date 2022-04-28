package admin

import (
	"encoding/json"
	"fmt"
	"ginxiaomi/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MainController struct{}

func (con MainController) Index(c *gin.Context) {
	// 获取userinfo 对应的session
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	// 类型断言来判断userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)
	if ok {
		// 获取用户信息
		var userinfoStruct []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

		// 获取所有的权限
		accessList := []models.Access{}
		models.DB.Where("module_id=?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort DESC")
		}).Order("sort DESC").Find(&accessList)

		// 获取当前角色拥有的权限，并把权限id放在一个map对象里面
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}

		// 循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中，如果是的话给当前数据加入checked属性
		for i := 0; i < len(accessList); i++ {
			if _, ok := roleAccessMap[accessList[i].Id]; ok {
				accessList[i].Checked = true
				for j := 0; j < len(accessList[i].AccessItem); j++ {
					if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
						accessList[i].AccessItem[j].Checked = true
					}
				}
			}
		}

		fmt.Printf("%#v", accessList)
		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username":   userinfoStruct[0].Username,
			"accessList": accessList,
			"isSuper":    userinfoStruct[0].IsSuper,
		})
	} else {
		c.Redirect(302, "/admin/login")
	}

}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

func (con MainController) ChangeStatus(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "非法请求",
		})
	}
	table := c.Query("table")
	field := c.Query("field")
	err1 := models.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id = ?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "修改失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "修改成功",
		})
	}
}
