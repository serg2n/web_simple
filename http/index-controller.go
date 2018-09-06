package http

import (
	"fmt"
	"net/http"
	"simple-web-app/ui"
)

type IndexController struct {

}

func (ic *IndexController) indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, ui.IndexHTML)
}

func NewIndexController() *IndexController {
	return &IndexController{
	}
}
