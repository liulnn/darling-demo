package plants

import (
	"darling"
	"fmt"
)

type PlantsCtrl struct {
	darling.Controller
}

func (c *PlantsCtrl) Post() {
	fmt.Println(c.Request.In)

	c.Response.StatusCode = 201
}
