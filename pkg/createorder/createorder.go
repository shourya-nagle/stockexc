package createorder

import (
	"net/http"
	"stockexchange/pkg/engine"
	"stockexchange/pkg/stocklist"

	"github.com/gin-gonic/gin"
)

func Getorder(context *gin.Context) {
	var neworder engine.Order
	if err := context.BindJSON(&neworder); err != nil {
		return
	}
	stocklist.ProcessOrder(neworder)
	context.IndentedJSON(http.StatusCreated, neworder)
}
