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

func (c *PlantsCtrl) Get() {

	c.Response.StatusCode = 200
}

type PlantCtrl struct {
	darling.Controller
}

func (c *PlantCtrl) Get() {
	fmt.Println(c.PathParams[0])
	c.Response.StatusCode = 200
}
