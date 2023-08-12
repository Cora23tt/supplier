package handler

import (
	"fmt"
	"html/template"
	"log"

	online_diler "github.com/cora23tt/online-diler"
	"github.com/gin-gonic/gin"
)

type ErrorPageData struct {
	ErrorMessage string
}

// type signInInput struct {
// 	Username string `json:"username" binging:"required"`
// 	Password string `json:"password" binging:"required"`
// }

// func (h *Handler) createAccount(c *gin.Context) {
// 	// save entered credentials into db
// 	emailCookie, _ := c.Request.Cookie("email")
// 	// input value OTP from html form
// 	if err := c.Request.ParseForm(); err != nil {
// 		fmt.Fprintf(c.Writer, "ParseForm() err: %v", err)
// 		return
// 	}
// 	full_name := c.Request.FormValue("FullName")
// 	username := c.Request.FormValue("username")
// 	password := c.Request.FormValue("password")
// 	user := online_diler.User{
// 		FirstName: full_name,
// 		Email:     emailCookie.Value,
// 		Username:  username,
// 		Password:  password,
// 	}
// 	_, err := h.services.Authorisation.CreateUser(user)
// 	if err != nil {
// 		renderSignup3(c)
// 	}
// 	// return store main page
// 	h.indexPage(c)
// }

func (h *Handler) verify(c *gin.Context) {
	email := c.Param("email")

	if err := c.Request.ParseForm(); err != nil {
		fmt.Fprintf(c.Writer, "ParseForm() err: %v", err)
		return
	}
	OTP := c.Request.FormValue("OTP")

	err := h.services.EAuthService.CheckOTP(email, OTP)
	if err != nil {
		h.render(c, ErrorPageData{ErrorMessage: "wrong one time password"}, "sign_base",
			"/home/cora/Documents/projects/golang-projects/online-diler/templates/sign_base.html",
			"/home/cora/Documents/projects/golang-projects/online-diler/templates/error_page.html",
		)
		return
	}

	// save user credentials in DB
	session, err := h.services.EAuthService.Get(email)
	if err != nil {
		h.render(c, ErrorPageData{ErrorMessage: "internal system error, can`t create user"}, "sign_base",
			"/home/cora/Documents/projects/golang-projects/online-diler/templates/sign_base.html",
			"/home/cora/Documents/projects/golang-projects/online-diler/templates/error_page.html",
		)
		return
	}

	user := online_diler.User{
		Email:     session.Email,
		FirstName: session.FirstName,
		LastName:  session.LastName,
		Password:  session.Password,
	}

	_, err = h.services.Authorisation.CreateUser(user)
	if err != nil {
		h.render(c, ErrorPageData{ErrorMessage: "internal system error, can`t create user"}, "sign_base",
			"/home/cora/Documents/projects/golang-projects/online-diler/templates/sign_base.html",
			"/home/cora/Documents/projects/golang-projects/online-diler/templates/error_page.html",
		)
		return
	}

	access_token, err := h.services.Authorisation.GenerateToken(session.Email, "")
	if err != nil {
		log.Println("Can`t generate access_token hash", err.Error())
		h.render(c, ErrorPageData{ErrorMessage: "internal system error, can`t create user"}, "sign_base",
			"/home/cora/Documents/projects/golang-projects/online-diler/templates/sign_base.html",
			"/home/cora/Documents/projects/golang-projects/online-diler/templates/error_page.html",
		)
		return
	}
	c.SetCookie("access_token", access_token, 43200, "/", "", false, false)
	c.Redirect(301, "/")
}

func (h *Handler) sendOTP(c *gin.Context) {

	// getting input credentials
	if err := c.Request.ParseForm(); err != nil {
		fmt.Fprintf(c.Writer, "ParseForm() err: %v", err)
		return
	}
	user := online_diler.User{
		Email:     c.Request.FormValue("email"),
		FirstName: c.Request.FormValue("firstname"),
		LastName:  c.Request.FormValue("surname"),
		Password:  c.Request.FormValue("password"),
	}

	// sending OTP to email
	if err := h.services.EAuthService.SendOTP(user); err != nil {
		h.render(
			c, ErrorPageData{ErrorMessage: err.Error()}, "sign_base",
			"/home/cora/Documents/projects/golang-projects/online-diler/templates/sign_base.html",
			"/home/cora/Documents/projects/golang-projects/online-diler/templates/error_page.html",
		)
		return
	}

	// rendering "OTP input" page
	h.render(
		c, c.Request.FormValue("email"), "sign_base",
		"/home/cora/Documents/projects/golang-projects/online-diler/templates/sign_base.html",
		"/home/cora/Documents/projects/golang-projects/online-diler/templates/signup2.html",
	)
}

func (h *Handler) signUp(c *gin.Context) {
	h.render(
		c, nil, "sign_base",
		"/home/cora/Documents/projects/golang-projects/online-diler/templates/sign_base.html",
		"/home/cora/Documents/projects/golang-projects/online-diler/templates/signup.html",
	)
}

// func (h *Handler) signIn(c *gin.Context) {
// 	var input signInInput
// 	if err := c.BindJSON(&input); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	token, err := h.services.Authorisation.GenerateToken(input.Username, input.Password)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	c.JSON(http.StatusOK, map[string]interface{}{
// 		"token": token,
// 	})
// }

func (h *Handler) render(c *gin.Context, data any, tmp_name string, pages ...string) {
	t, err := template.ParseFiles(pages...)
	if err != nil {
		log.Println("Error parsing templates: ", err)
	}
	if err = t.ExecuteTemplate(c.Writer, tmp_name, data); err != nil {
		log.Println("Error executing template: ", err)
	}
}

// func renderSignup3(c *gin.Context) {
// 	t, err := template.ParseFiles("/home/cora/Documents/projects/golang-projects/online-diler/templates/sign_base.html",
// 		"/home/cora/Documents/projects/golang-projects/online-diler/templates/signup3.html")
// 	if err != nil {
// 		log.Println("Error parsing templates: ", err)
// 	}
// 	if err = t.ExecuteTemplate(c.Writer, "sign_base", nil); err != nil {
// 		log.Println("Error executing template: ", err)
// 	}
// }
