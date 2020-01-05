package controllers

import (
	"../models"
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSos(c *gin.Context) {

	sos := structs.CreateSos{}
	notif := structs.CreateSosNotif{}
	admin := structs.CreateSosAdmin{}

	var co = structs.Component{}

	response := structs.JsonResponse{}

	err := c.ShouldBind(&sos)
	err2 := c.ShouldBind(&notif)
	err3 := c.ShouldBind(&admin)

	if err != nil || err2 != nil || err3 != nil {
		var me string
		if err != nil {
			me = me + err.Error()
		}
		if err2 != nil {
			me = me + err2.Error()
		}
		if err3 != nil {
			me = me + err3.Error()
		}
		response.ApiMessage = "validation " + me
		c.JSON(400, response)
	} else {
		data, errx := models.CreatedSos(sos, notif, admin)

		response.Data = data
		if errx != nil {
			response.ApiMessage = co.GetMessageErr()
			c.JSON(400, response)
		} else {
			response.ApiStatus = 1
			response.ApiMessage = co.GetMessageSucc()
			c.JSON(http.StatusOK, response)
		}
	}
}

func GetSosPetugas(c *gin.Context) {
	getsos := structs.GetSosPetugas{}
	t := structs.Component{}

	response := structs.JsonResponse{}

	err := c.BindQuery(&getsos)
	if err != nil {
		var m string
		if err != nil {
			m = m + err.Error()
		}
		response.ApiMessage = "validation " + m
		c.JSON(400, response)
	} else {
		data, errx := models.GetSosPetugas(getsos)
		response.Data = data
		if errx != nil {
			response.ApiMessage = t.GetMessageErr()
			c.JSON(400, response)
		} else {
			response.ApiStatus = 1
			response.ApiMessage = t.GetMessageErr()
			c.JSON(200, response)
		}
	}
}

func GetSosadmin(c *gin.Context) {
	getsos := structs.GetSosAdmin{}
	t := structs.Component{}

	response := structs.JsonResponse{}

	err := c.BindQuery(&getsos)
	if err != nil {
		var m string
		if err != nil {
			m = m + err.Error()
		}
		response.ApiMessage = "validation " + m
		c.JSON(400, response)
	} else {
		data, errx := models.GetSosAdmin(getsos)
		response.Data = data
		if errx != nil {
			response.ApiMessage = t.GetMessageErr()
			c.JSON(400, response)
		} else {
			response.ApiStatus = 1
			response.ApiMessage = t.GetMessageErr()
			c.JSON(200, response)
		}
	}
}