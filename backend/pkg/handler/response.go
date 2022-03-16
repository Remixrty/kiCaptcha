package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// WTF warning with error
//-------------------
// type error struct {
// 	Message string `json:"message"`
// }

//For our error response
//----------------------
func newErrorResponse(c *gin.Context, statusCode int, message map[string]interface{}) {
	logrus.Error(message)
	c.JSON(statusCode, message)
}
