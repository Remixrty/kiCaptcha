package handler

import (
	"errors"
	"net/http"

	// "time"
	"github.com/gin-gonic/gin"
	"github.com/romanSeleznev/kiCaptcha/backend/pkg/model"
)

//struct for sing in ad sing up
//------------------------
type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//Struct for delete
//------------------
type deleteInfo struct {
	Id       int    `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//Post Method singIn
//Use when sing in account
//--------------------------
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, map[string]interface{}{
			"error": "Incorrect request",
		})
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusOK, map[string]interface{}{
			"error": "Incorrect email or password",
		})
		return
	}
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Expose-Headers", "X-Custom")
	// c.JSON(http.StatusOK, map[string]interface{}{
	// 	"token": token,
	// })

	c.SetCookie("jwt", token, 60*60*12, "http://localhost:8080/sign-check", "", true, true)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Correct log in",
	})
}

//Get Method singCheck
//Only for debuging
//------------------------
func (h *Handler) signCheck(c *gin.Context) {
	// var user model.User

	cookieName := &http.Cookie{
		Name:   "name",
		Value:  "Andrew",
		MaxAge: 300,
	}
	cookieLogin := &http.Cookie{
		Name:   "email",
		Value:  "remixrty@gmail.com",
		MaxAge: 300,
	}
	cookieID := &http.Cookie{
		Name:   "ID",
		Value:  "1",
		MaxAge: 300,
	}
	http.SetCookie(c.Writer, cookieID)
	http.SetCookie(c.Writer, cookieLogin)
	http.SetCookie(c.Writer, cookieName)
	err := errors.New("")
	c.JSON(http.StatusOK, err)

}

//Unused
//-------------------------
func (h *Handler) signOut(c *gin.Context) {

}

//Post Method signUp
//Add user when registration
//--------------------------
func (h *Handler) signUp(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, map[string]interface{}{
			"error": "Incorrect request",
		})
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})

}

// Post method DeleteUser
//----------------------
func (h *Handler) deleteUser(c *gin.Context) {
	var input deleteInfo

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, map[string]interface{}{
			"error": "Incorrect request",
		})
		return
	}
	err := h.services.Authorization.DeleteUsers(input.Id, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"info": "User was deleted",
	})
}
