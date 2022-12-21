package route

import (
	ctl "lecture/WBABEProject-23/controller"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct: ctl}

	return r, nil
}
