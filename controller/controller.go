package controller

import "lecture/WBABEProject-23/model"

type Controller struct {
	md *model.Model
}

func NewController(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	return r, nil
}
