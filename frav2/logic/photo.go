package logic

import (
	"Toupiao/frav2/model"
	"Toupiao/frav2/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpLoad() gin.HandlerFunc {
	return func(context *gin.Context) {
		file, _ := context.FormFile("photofile")
		//a为从jwt中取出的1
		//a := int64(2)
		context.Set("userID", int64(12))
		jid, _ := context.Get("userID")
		id := jid.(int64)
		path := "/user/" + strconv.FormatInt(id, 10) + "pic.png"
		context.SaveUploadedFile(file, "./view"+path)
		//图片地址保存
		g := path
		fmt.Println(g)
		context.JSON(http.StatusOK, tools.Code{})
	}
}
func UpLoadUserPhoto() gin.HandlerFunc {
	return func(context *gin.Context) {
		var MyEntry model.Entry
		var MyRess []model.Res
		jid, _ := context.Get("userID")
		id := jid.(int64)
		//获取数据
		MyEntry.Key.Name = "翠绿之影"
		MyEntry.Key.Main = "风元素伤害加成"
		MyEntry.Key.Asse = "空之杯"
		MyEntry.Value.CN = 33
		MyEntry.Value.DGj = 5
		MyEntry.Value.GJ = 17
		MyEntry.Value.DSM = 5
		role := tools.FindRole(MyEntry.Key.Name)
		if role[0].RoleName == "过渡使用" {
			fmt.Println("不推荐")
		}
		var resRoles []model.Role
		switch MyEntry.Key.Asse {
		case "时之沙":
			for _, r := range role {
				for _, hourglass := range r.RoleHourglass {
					if MyEntry.Key.Main == hourglass {
						resRoles = append(resRoles, r)
						break
					}
				}
			}
		case "空之杯":
			for _, r := range role {
				for _, cup := range r.RoleCup {
					if MyEntry.Key.Main == cup {
						resRoles = append(resRoles, r)
						break
					}
				}
			}
		case "理之冠":
			for _, r := range role {
				for _, pieum := range r.RolePileum {
					if MyEntry.Key.Main == pieum {
						resRoles = append(resRoles, r)
						break
					}
				}
			}
		default:
			for _, r := range role {
				resRoles = append(resRoles, r)
				return
			}
		}
		a := tools.FindDate()
		for _, resRole := range resRoles {
			MyRess = append(MyRess, model.Res{
				Name:   resRole.RoleName,
				Source: strconv.FormatFloat(tools.Count(MyEntry.Value, a), 'f', 1, 64),
			})
		}
		path := "/view/res/" + strconv.FormatInt(id, 10) + "pic.png"
		fmt.Println(MyRess)
		tools.SavePhoto(MyRess, path)
		context.JSON(http.StatusOK, tools.Code{Data: path})
	}
}
