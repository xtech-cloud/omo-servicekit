package account

import (
	"osk/http"
	"osk/model"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func HandleCurrent(_uri string, _group *gin.RouterGroup) {

	_group.GET(_uri, func(_context *gin.Context) {
		defer http.CatchRenderError()

		claims := jwt.ExtractClaims(_context)
		uuid := claims["id"].(string)

		dao := model.NewAccountDAO()
		account, err := dao.Find(uuid)

		http.TryRenderDatabaseError(_context, err)

		rsp := gin.H{
			"name":   account.Name,
			"avatar": account.Avatar,
			"uuid":   account.UUID,
		}
		http.RenderOK(_context, rsp)
	})
}
