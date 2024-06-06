package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	Name  string
	Email string
}

func newContact(name, email string) Contact {
	return Contact{Name: name, Email: email}
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

	_ = router.Run(":8080") // запускаем сервер на порту 8080
}
