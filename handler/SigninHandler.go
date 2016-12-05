package handler

import (
	"github.com/insionng/macross"
	"github.com/insionng/macross/jwt"
	"github.com/insionng/zenpress/helper"
	"github.com/insionng/zenpress/models"
	"time"
)

func SigninGetHandler(self *macross.Context) error {

	claims := jwt.GetMapClaims(self)
	var uid int64
	if jwtUserId, okay := claims["UserId"].(float64); okay {
		uid = int64(jwtUserId)
	}

	if uid > 0 {
		return self.Redirect("/", 302)
	} else {
		return self.Render("signin")
	}

}

func SigninPostHandler(self *macross.Context) error {
	username := self.Args("username").String()
	password := self.Args("password").String()

	if username != "" && password != "" {

		if usr := models.GetUserByNickname(username); usr.Password != "" {

			if helper.ValidateHash(usr.Password, password) {
				claims := jwt.NewMapClaims()
				claims["IsRoot"] = (usr.Role == -1000)
				claims["UserId"] = usr.Id
				claims["Username"] = usr.Nickname
				claims["exp"] = time.Now().Add(jwt.DefaultJWTConfig.Expires).Unix()

				var secret string
				if signingKey, okay := jwt.DefaultJWTConfig.SigningKey.(string); okay {
					secret = signingKey
				}
				tokenString, _ := jwt.NewTokenString(secret, "HS256", claims)

				self.Response.Header.Set(macross.HeaderAccessControlExposeHeaders, "Authorization")
				self.Response.Header.Set("Authorization", jwt.Bearer+" "+tokenString)

				return self.Redirect("/", 302)

			}
		}
	}
	return self.Redirect("/signin/", 302)
}
