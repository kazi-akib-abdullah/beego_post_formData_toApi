package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	// "golang.org/x/text/message"
)


type User struct {
	Firstname   string `form:"first_name"`
	Lastname    string `form:"last_name"`
	Email       string `form:"email"`
	Phone       string `form:"phone"`
	Password    string `form:"password"`
	DateOfBirth string `form:"dob"`
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.TplName = "index.tpl"
	flash:=beego.NewFlash()

	first_name := c.GetString("first_name")
	last_name := c.GetString("last_name")
	email := c.GetString("email")
	phone := c.GetString("phone")
	password := c.GetString("password")
	dob := c.GetString("dob")
	red := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

	valid := validation.Validation{}
	valid.Email(email, "")
	valid.Phone(phone, "")
	valid.MinSize(password, 6,"Password")
	valid.Match(dob, red ,"")

	if valid.HasErrors() {

		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			if strings.Contains(err.Message, "email"){
				c.Data["email"] = err.Message
				flash.Error(err.Message)
				// log.Println(err.Key, err.Message)
				flash.Store(&c.Controller)
			}
			if strings.Contains(err.Message, "phone"){
				c.Data["phone"] = err.Message
				flash.Error(err.Message)
				// log.Println(err.Key, err.Message)
				flash.Store(&c.Controller)
			}
			if strings.Contains(err.Message, "Minimum"){
				c.Data["password"] = err.Message
				flash.Error(err.Message)
				// log.Println(err.Key, err.Message)
				flash.Store(&c.Controller)
			}
			if strings.Contains(err.Message, "match"){
				c.Data["dob"] = err.Message
				flash.Error(err.Message)
				// log.Println(err.Key, err.Message)
				flash.Store(&c.Controller)
			}

			// flash.Error(err.Key, err.Message)
        	// flash.Store(&c.Controller)
	        return
		}
	}




	postBody, _ := json.Marshal(map[string]string{
		"FirstName":   first_name,
		"LastName":    last_name,
		"Phone":       phone,
		"Email":       email,
		"Password":    password,
		"DoB": dob,
	})
	responseBody := bytes.NewBuffer(postBody)
	fmt.Println(responseBody)
	resp, err := http.Post("http://127.0.0.1:8080/v1/user", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	flash.Error(sb)
	flash.Store(&c.Controller)
	fmt.Println(sb)
}
