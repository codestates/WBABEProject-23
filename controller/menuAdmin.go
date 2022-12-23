package controller

import (
	"fmt"
	"lecture/WBABEProject-23/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// NewMenu godoc
// @Summary call NewMenu, return ok by json.
// @새로운 메뉴 추가.
// @name NewMenu
// @Accept  json
// @Produce  json
// @Param id body model.Menu true "User input"
// @Router /menu/admin/new [POST]
// @Success 200 {object} Controller
func (p *Controller) NewMenu(c *gin.Context) {
	var menu model.Menu
	if err := c.ShouldBind(&menu); err != nil {
		c.String(http.StatusBadRequest, "Bad request: %v", err)
		return
	}
	business := c.GetHeader("Business-Id")
	p.md.CreateNewMenu(menu, business)
	c.JSON(200, gin.H{"msg": "ok"})
}

func (p *Controller) ModifyMenu(c *gin.Context) {
	var menu model.Menu
	var jsonMap map[string]interface{}
	business := c.GetHeader("Business-Id")
	if err := c.BindJSON(&jsonMap); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	toUpdate := fmt.Sprintf("%v", jsonMap["toUpdate"])
	err := mapstructure.Decode(jsonMap, &menu)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	p.md.ModifyMenu(toUpdate, business, menu)
	c.JSON(200, gin.H{"msg": "ok"})
}
