package middlewares

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"idist-core/app/collections"
	"idist-core/app/controllers/auth"
	"idist-core/app/providers/configProvider"
	"idist-core/app/providers/loggerProvider"
	"log"
	"time"
	// "fmt"
)

var identityKey = "id"

func AuthMiddleware() *jwt.GinJWTMiddleware {
	config := configProvider.GetConfig()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           config.GetString("app.name"),
		Key:             []byte(config.GetString("app.secret.jwt")),
		Timeout:         24 * time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     identityKey,
		TokenLookup:     "cookie: access_token, header: Authorization",
		TokenHeadName:   "Bearer",
		CookieName:      "access_token",
		TimeFunc:        time.Now,
		SecureCookie:    true,
		CookieHTTPOnly:  false,
		SendCookie:      true,
		PayloadFunc:     auth.PayloadFunc,
		IdentityHandler: auth.IdentityHandler,
		Authenticator:   auth.PostLogin,
		LoginResponse:   auth.ResponseLogin,
		Authorizator:    auth.Authorizator,
		Unauthorized:    auth.Unauthorized,
		RefreshResponse: auth.RefreshResponse,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}

func Gate(Subject, Action string) gin.HandlerFunc {
	user := collections.User{}
	return func(c *gin.Context) {
		if Subject == "" || Action == "" {
			c.Next()
			return
		}
		loggerProvider.GetLogger().Info("Subject:" + Subject + " Action:" + Action)
		if _, exist := c.Get("user"); exist == true {
			user = c.MustGet("user").(collections.User)
		}

		if user.HasPermission() {
			c.Next()
			return
		}
		c.AbortWithStatus(403)
		return
	}
}
