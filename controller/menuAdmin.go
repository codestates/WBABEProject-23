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
// @Param business_id header string true "사업체 ID"
// @Param id body NewMenuInput true "메뉴 입력"
// @Router /menu/admin/new [POST]
// @Success 200 {object} Controller
func (p *Controller) NewMenu(c *gin.Context) {
	var menu model.Menu
	if err := c.ShouldBind(&menu); err != nil {
		c.String(http.StatusBadRequest, "Bad request: %v", err)
		return
	}
	menu.State = 1
	menu.IsDeleted = false
	business := c.GetHeader("business_id")
	result := p.md.CreateNewMenu(menu, business)
	msg := [3]string{
		"No document was found with the business id",
		"Internl Error",
		"ok",
	}
	statusCode := [3]int{400, 500, 200}
	c.JSON(statusCode[result], gin.H{"msg": msg[result]})
}

// swag input 용
type NewMenuInput struct {
	Name     string `bson:"name"`
	Price    int    `bson:"price"`
	Origin   string `bson:"origin"`
	Category string `bson:"category"`
}

// ModifyMenu godoc
// @Summary call ModifyMenu, return ok by json.
// @메뉴 수정/삭제.
// @name ModifyMenu
// @Accept  json
// @Produce  json
// @Param business_id header string true "사업체 ID"
// @Param id body ModifyMenuInput true "User input 바꿀 메뉴 이름 toUpdate로 추가, 바꿀내용만 작성"
// @Router /menu/admin/modify [PATCH]
// @Success 200 {object} Controller
func (p *Controller) ModifyMenu(c *gin.Context) {
	var menu model.Menu
	var jsonMap map[string]interface{}
	business := c.GetHeader("business_id")
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

type ModifyMenuInput struct {
	State    int    `bson:"state"`
	Price    int    `bson:"price"`
	Origin   string `bson:"origin"`
	Category string `bson:"category"`
	ToUpdate string `bson:"toUpdate"`
}
