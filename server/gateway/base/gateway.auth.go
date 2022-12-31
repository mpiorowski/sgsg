package base

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	session, err := Authorization(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, session)
}

type IdToken struct {
	IdToken string `json:"idToken"`
}

func Login(c *gin.Context) {

	var idToken IdToken
	err := c.BindJSON(&idToken)
	if err != nil {
		log.Printf("c.BindJSON: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid idToken"})
		return
	}

	client, err := ConnectToFirebase()
	if err != nil {
		log.Printf("ConnectToFirebase: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Set session expiration to 5 days.
	expiresIn := time.Hour * 24 * 5
	cookie, err := client.SessionCookie(c, idToken.IdToken, expiresIn)
	if err != nil {
		log.Printf("client.SessionCookie: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid idToken"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cookie": cookie})
}
