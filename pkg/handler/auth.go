package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	online_diler "github.com/cora23tt/online-diler"
	"github.com/gin-gonic/gin"
)

type ErrorPageData struct {
	ErrorMessage string
}

type signInInput struct {
	Username string `json:"username" binging:"required"`
	Password string `json:"password" binging:"required"`
}

func (h *Handler) createAccount(c *gin.Context) {
	// save entered credentials into db
	emailCookie, _ := c.Request.Cookie("email")

	// input value OTP from html form
	if err := c.Request.ParseForm(); err != nil {
		fmt.Fprintf(c.Writer, "ParseForm() err: %v", err)
		return
	}
	full_name := c.Request.FormValue("FullName")
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	user := online_diler.User{
		Fullname: full_name,
		Email:    emailCookie.Value,
		Username: username,
		Password: password,
	}
	_, err := h.services.Authorisation.CreateUser(user)
	if err != nil {
		renderSignup3(c)
	}

	// return store main page
	h.indexPage(c)
}

func (h *Handler) verifyEmail(c *gin.Context) {

	// email from params
	email := c.Param("email")

	// input value OTP from html form
	if err := c.Request.ParseForm(); err != nil {
		fmt.Fprintf(c.Writer, "ParseForm() err: %v", err)
		return
	}
	OTP := c.Request.FormValue("OTP")

	// search email for matches from OTP basket
	eAuth, err := h.services.EAuthBasket.Find(email)
	if err != nil {
		h.signUp(c)
	}

	// checking inputed OTP
	if strconv.Itoa(eAuth.OTP) != OTP {
		data := ErrorPageData{ErrorMessage: "wrong one time password"}
		h.renderErrorPage(c, data)
		return
	}

	// try to save email in verified emails array
	if err := h.services.VerifyedEmailsService.Save(email); err != nil {
		// already exist
		data := ErrorPageData{ErrorMessage: err.Error()}
		h.renderErrorPage(c, data)
	}
	renderSignup3(c)
}

func (h *Handler) sendOTP(c *gin.Context) {

	// getting input email
	if err := c.Request.ParseForm(); err != nil {
		fmt.Fprintf(c.Writer, "ParseForm() err: %v", err)
		return
	}
	email := c.Request.FormValue("email")

	// sending OTP to email
	if err := h.services.SendOTP(email); err != nil {
		data := ErrorPageData{ErrorMessage: err.Error()}
		h.renderErrorPage(c, data)
		return
	}

	// rendering "OTP input" page
	t, err := template.ParseFiles("/home/cora/Documents/projects/golang-projects/online-diler/templates/sign_base.html",
		"/home/cora/Documents/projects/golang-projects/online-diler/templates/signup2.html")
	if err != nil {
		log.Println("Error parsing templates: ", err)
	}
	if err = t.ExecuteTemplate(c.Writer, "sign_base", email); err != nil {
		log.Println("Error executing template: ", err)
	}
}

func (h *Handler) signUp(c *gin.Context) {

	t, err := template.ParseFiles("/home/cora/Documents/projects/golang-projects/online-diler/templates/sign_base.html",
		"/home/cora/Documents/projects/golang-projects/online-diler/templates/signup.html")
	if err != nil {
		log.Println("Error parsing templates: ", err)
	}

	err = t.ExecuteTemplate(c.Writer, "sign_base", nil)
	if err != nil {
		log.Println("Error executing template: ", err)
	}
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorisation.GenerateToken(input.Username, input.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) renderErrorPage(c *gin.Context, data ErrorPageData) {
	t, err := template.ParseFiles("/home/cora/Documents/projects/golang-projects/online-diler/templates/sign_base.html",
		"/home/cora/Documents/projects/golang-projects/online-diler/templates/error_page.html")
	if err != nil {
		log.Println("Error parsing templates: ", err)
	}
	if err = t.ExecuteTemplate(c.Writer, "sign_base", data); err != nil {
		log.Println("Error executing template: ", err)
	}
}

func renderSignup3(c *gin.Context) {
	t, err := template.ParseFiles("/home/cora/Documents/projects/golang-projects/online-diler/templates/sign_base.html",
		"/home/cora/Documents/projects/golang-projects/online-diler/templates/signup3.html")
	if err != nil {
		log.Println("Error parsing templates: ", err)
	}
	if err = t.ExecuteTemplate(c.Writer, "sign_base", nil); err != nil {
		log.Println("Error executing template: ", err)
	}
}
