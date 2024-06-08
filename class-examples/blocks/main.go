package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Block struct {
	Id int
}

type Blocks struct {
	Start  int
	Next   int
	More   bool
	Blocks []Block
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../../views/*.html")

	router.GET("/blocks", func(ctx *gin.Context) {
		startStr := ctx.Query("start")
		start, err := strconv.Atoi(startStr)
		if err != nil {
			start = 0
		}

		blocks := []Block{}
		for i := 0; i < start+10; i++ {
			blocks = append(blocks, Block{Id: i})
		}

		template := "blocks"
		if start == 0 {
			template = "blocks-index"
		}

		ctx.HTML(http.StatusOK, template, Blocks{
			Start:  start,
			Next:   start + 10,
			More:   start+10 < 100,
			Blocks: blocks,
		})
	})

	_ = router.Run(":3000")
}
