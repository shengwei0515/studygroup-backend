package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func SaveAuthSession(c *gin.Context, id string) error {
	session := sessions.Default(c)
	session.Set("userId", id)
	err := session.Save()
	return err
}

func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if session == nil {
		log.Printf("In HasSession with session is nil")
		return false
	}
	if sessionValue := session.Get("userId"); sessionValue == nil {
		return false
	}
	return true
}

func GetSessionUserId(c *gin.Context) string {
	session := sessions.Default(c)
	sessionValue := session.Get("userId")
	if sessionValue == nil {
		return ""
	}
	return sessionValue.(string)
}
