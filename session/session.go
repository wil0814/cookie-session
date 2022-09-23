package session

import (
	"cookieAndsession/users"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var sessions = map[string]session{}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type session struct {
	username string
	expiry   time.Time
}

//determine if the session has expired
func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func Signin(c *gin.Context) {
	var creds Credentials

	err := c.Bind(&creds)
	if err != nil {
		log.Println(creds.Username)
		log.Println(creds.Password)
	}
	expcetedPassword, ok := users.Users[creds.Username]

	if !ok || expcetedPassword != creds.Password {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	sessions[sessionToken] = session{
		username: creds.Username,
		expiry:   expiresAt,
	}

	c.SetCookie("session_token", sessionToken, 120, "/", "localhost", false, true)

}

func Welcome(c *gin.Context) {
	data, err := c.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Println("data ", data)
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := data

	userSession, exists := sessions[sessionToken]
	if !exists {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	if userSession.isExpired() {
		delete(sessions, sessionToken)
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("Welcone %s!", userSession.username)))

}

func Refresh(c *gin.Context) {
	data, err := c.Cookie("session_token")

	if err != nil {
		if err == http.ErrNoCookie {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := data
	userSession, exists := sessions[sessionToken]
	if !exists {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	sessions[newSessionToken] = session{
		username: userSession.username,
		expiry:   expiresAt,
	}
	delete(sessions, sessionToken)

	c.SetCookie("session_token", newSessionToken, 120, "/", "localhost", false, true)
}

func Logout(c *gin.Context) {
	date, err := c.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := date

	delete(sessions, sessionToken)

	c.SetCookie("session_token", "", -1, "/", "localhost", false, true)

}
