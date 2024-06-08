package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var id = 0

type Contact struct {
	Name  string
	Email string
	Id    int
}

func newContact(name, email string) Contact {
	id++
	return Contact{Name: name, Email: email, Id: id}
}

type Contacts []Contact

type Data struct {
	Contacts Contacts
}

func (d *Data) hasEmail(email string) bool {
	for _, c := range d.Contacts {
		if c.Email == email {
			return true
		}
	}
	return false
}

func (d *Data) indexOf(id int) int {
	for i, contact := range d.Contacts {
		if contact.Id == id {
			return i
		}
	}

	return -1
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("John", "john@gmail.com"),
			newContact("Adam", "adam@gmail.com"),
		},
	}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data     Data
	FormData FormData
}

func newPage() Page {
	return Page{
		Data:     newData(),
		FormData: newFormData(),
	}
}

func main() {
	router := gin.Default()             // создаем роутер Gin
	router.LoadHTMLGlob("views/*.html") // указываем, где искать HTML файлы

	router.Static("/css", "css")
	router.Static("/images", "images")

	page := newPage()

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", page)
	})

	router.POST("/contacts", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		email := ctx.PostForm("email")

		if page.Data.hasEmail(email) {
			formData := newFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Errors["email"] = "Email already exists"

			ctx.HTML(422, "form", formData)
			return
		}

		contact := newContact(name, email)
		page.Data.Contacts = append(page.Data.Contacts, contact)

		ctx.HTML(http.StatusOK, "form", newFormData())
		ctx.HTML(http.StatusOK, "oob-contact", contact)
	})

	router.DELETE("/contacts/:id", func(ctx *gin.Context) {
		time.Sleep(time.Second * 3)

		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.String(400, "Invalid ID")
		}

		index := page.Data.indexOf(id)
		if index == -1 {
			ctx.String(400, "Contact not found")
		}

		page.Data.Contacts = append(page.Data.Contacts[:index], page.Data.Contacts[index+1:]...)
		ctx.Status(204)
	})

	_ = router.Run(":3000") // запускаем сервер на порту 3000
}
