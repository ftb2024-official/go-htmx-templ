package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Templates struct {
// 	templates *template.Template
// }

// func (t *Templates) Render(wr io.Writer, name string, data interface{}, ctx *gin.Context) error {
// 	return t.templates.ExecuteTemplate(wr, name, data)
// }

// func newTemplate() *Templates {
// 	return &Templates{
// 		templates: template.Must(template.ParseGlob("views/*.html")),
// 	}
// }

type Count struct {
	Count int
}

func main() {
	router := gin.Default()             // создаем роутер Gin
	router.LoadHTMLGlob("views/*.html") // указываем, где искать HTML файлы

	count := Count{Count: 0}

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", count)
	})

	router.POST("/count", func(ctx *gin.Context) {
		count.Count++
		ctx.HTML(http.StatusOK, "count", count)
	})

	_ = router.Run(":8080") // запускаем сервер на порту 8080
}
