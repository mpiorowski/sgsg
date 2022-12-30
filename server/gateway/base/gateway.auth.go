package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	session, err := Authorization(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, session)
}
