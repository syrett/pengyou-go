package controllers

import (
	"fmt"
	"github.com/robfig/revel"
	//	"net/http"
	"pengyou/app/models"
)

type Active struct {
	*revel.Controller
}

func (ac Active) Index() revel.Result {
	return ac.Render()
}

func (ac Active) Create() revel.Result {
	//	var active *active.Active = nil

	//	for d := range args {
	//		fmt.Println(d)
	//	}
	return ac.Render()
}

func (ac Active) PostCreate(active *models.Active) revel.Result {
	active.Owner = 1 //todo get it from session
	active.Total = 0
	fmt.Println(active)
	active.Save()
	//	for d := range args {
	//		fmt.Println(d)
	//	}

	//return ac.Redirect("Index")
	return ac.Render()
}

func (ac Active) AddActive() revel.Result {
	return ac.Render()
}
