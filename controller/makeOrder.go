package controller

import (
	"fmt"
	"lecture/WBABEProject-23/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (p *Controller) MakeOrder(c *gin.Context) {
	order := model.Order{}
	var jsonMap map[string]interface{}
	if err := c.BindJSON(&jsonMap); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	order.Id = primitive.NewObjectID()
	order.BusinessName = fmt.Sprintf("%v", jsonMap["businessName"])
	hexId := fmt.Sprintf("%v", jsonMap["orderer"])
	oID, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		panic(err)
	}
	order.Orderer = oID
	menu, ok := jsonMap["menu"].([]interface{})
	if !ok {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	for _, v := range menu {
		m, _ := v.(map[string]interface{})
		order.Menu = append(order.Menu, model.MenuNum{m["menu_name"].(string), int(m["number"].(float64))})
	}
	order.CreatedAt = time.Now()
	order.Status = model.Receipting
	p.md.MakeOrder(order)
	c.JSON(200, gin.H{"msg": "ok"})
}
