package controller

import (
	"lecture/WBABEProject-23/model"
	"lecture/WBABEProject-23/protocol"

	"github.com/gin-gonic/gin/binding"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

// CreateMenu godoc
// @Summary call CreateMenu, return ok by json.
// @새로운 메뉴 추가.
// @name CreateMenu
// @Accept  json
// @Produce  json
// @Param id body CreateMenuInput true "메뉴 입력"
// @Router /menu [POST]
// @Success 200 {object} Controller

type CreateMenuInput struct {
	Name       string `bson:"name" binding:"required"`
	Price      int    `bson:"price" binding:"required,gte=0"`
	Origin     string `bson:"origin" binding:"required"`
	Category   string `bson:"category" binding:"required"`
	BusinessID string `bson:"business_id,omitempty" binding:"required"`
}

func (p *Controller) CreateMenu(c *gin.Context) {
	var body CreateMenuInput
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		protocol.Fail(err, protocol.BadRequest).Response(c)
		return
	}
	menu, res := p.createMenuInputValidate(body)
	if res != nil {
		res.Response(c)
		return
	}
	result := p.md.CreateMenu(*menu)
	result.Response(c)
}

func (p *Controller) createMenuInputValidate(body CreateMenuInput) (res *model.Menu, errorRes *protocol.ApiResponse[any]) {
	res = new(model.Menu)
	var err error
	bID, err := primitive.ObjectIDFromHex(body.BusinessID)
	if err != nil {
		return nil, protocol.Fail(err, protocol.BadRequest)
	}
	if r, e := p.md.CheckBusinessID(bID); !r {
		return nil, protocol.FailCustomMessage(e, "No document was found with the business id", protocol.BadRequest)
	}
	res.BusinessID = bson.M{"$ref": "business", "$id": bID}
	res.Name = body.Name
	res.Price = body.Price
	res.Origin = body.Category
	res.Category = body.Category
	res.State = 1
	res.IsDeleted = false
	return res, nil
}

// UpdateMenu godoc
// @Summary call UpdateMenu, return ok by json.
// @메뉴 수정/삭제.
// @name UpdateMenu
// @Accept  json
// @Produce  json
// @Param id body UpdateMenuInput true "User input 바꿀 메뉴 이름 toUpdate로 추가, 바꿀내용만 작성"
// @Router /menu [PATCH]
// @Success 200 {object} Controller
func (p *Controller) UpdateMenu(c *gin.Context) {
	var body UpdateMenuInput
	if err := c.ShouldBindJSON(&body); err != nil {
		protocol.Fail(err, protocol.BadRequest).Response(c)
		return
	}
	menu, res := p.updateMenuInputValidate(body)
	if res != nil {
		res.Response(c)
	}
	res = p.md.UpdateMenu(menu)
	res.Response(c)
}

func (p *Controller) updateMenuInputValidate(body UpdateMenuInput) (*model.Menu, *protocol.ApiResponse[any]) {
	toUpdate, err := primitive.ObjectIDFromHex(body.ID)
	if err != nil {
		return nil, protocol.Fail(err, protocol.BadRequest)
	}
	menu, err := p.md.GetMenuById(toUpdate)
	if err != nil {
		return nil, protocol.Fail(err, protocol.BadRequest)
	}

	if body.Name != "" {
		menu.Name = body.Name
	}
	if body.State != 0 {
		menu.State = body.State
	}
	if body.Price != 0 {
		menu.Price = body.Price
	}
	if body.Origin != "" {
		menu.Origin = body.Origin
	}
	if body.Category != "" {
		menu.Category = body.Category
	}
	if (menu.Price < 0) || (menu.State != 1 && menu.State != 2) {
		return nil, protocol.FailCustomMessage(err, "Invalid Input", protocol.BadRequest)
	}

	return menu, nil
}

type UpdateMenuInput struct {
	ID        string `bson:"id" binding:"required"`
	Name      string `bson:"name,omitempty"`
	State     int    `bson:"state,omitempty"`
	Price     int    `bson:"price,omitempty"`
	Origin    string `bson:"origin,omitempty"`
	Category  string `bson:"category,omitempty"`
	IsDeleted bool   `bson:"is_deleted,omitempty"`
}
