package req

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/HandsFree/teacherui-backend/api"
	"github.com/HandsFree/teacherui-backend/cfg"
	"github.com/HandsFree/teacherui-backend/util"
)

func GetLogOutRequest() gin.HandlerFunc {
	return func(s *gin.Context) {
		session := sessions.Default(s)
		session.Clear()

		if err := session.Save(); err != nil {
			util.Error("LogOutRequest", err.Error())
			return
		}

		logoutLink := fmt.Sprintf("https://core.beaconing.eu/auth/logout?client_id=%s&redirect_uri=%s",
			cfg.Beaconing.Auth.ID,
			api.GetLink())

		// fmt.Println(logoutLink)

		s.Redirect(http.StatusTemporaryRedirect, logoutLink)
	}
}
