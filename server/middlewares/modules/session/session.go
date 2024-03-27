package sessionMiddleWareModule

import (
	"github.com/Xi-Yuer/cms/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Session(keyPairs string) gin.HandlerFunc {
	store := sessionConfig()
	return sessions.Sessions(keyPairs, store)
}
func sessionConfig() sessions.Store {
	sessionMaxAge := 3600
	sessionSecret := config.Config.APP.SESSIONSECRET
	var store sessions.Store
	store = cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge, //seconds
		Path:   "/",
	})
	return store
}
