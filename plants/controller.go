package plants

import (
	"darling"
	"fmt"
)

type PlantsCtrl struct {
	darling.Controller
}

func (c *PlantsCtrl) Post() {
	req := c.Request.In
	req.ParseForm()
	name := req.PostFormValue("name")
	fmt.Println("name:", name)

	c.Response.StatusCode = 201
}

func (c *PlantsCtrl) Get() {

	c.Response.StatusCode = 200
}

type PlantCtrl struct {
	darling.Controller
}

func (c *PlantCtrl) Get() {
	plantId := c.PathParams[0]
	fmt.Println("plantId:", plantId)
	c.Response.StatusCode = 200
}
