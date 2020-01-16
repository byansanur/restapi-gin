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
	//admin := structs.CreateSosAdmin{}

	var co = structs.Component{}

	response := structs.JsonResponse{}

	err := c.ShouldBind(&sos)
	err2 := c.ShouldBind(&notif)
	//err3 := c.ShouldBind(&admin)

	if err != nil || err2 != nil {
		var me string
		if err != nil {
			me = me + err.Error()
		}
		if err2 != nil {
			me = me + err2.Error()
		}
		response.ApiMessage = "validation " + me
		c.JSON(400, response)
	} else {
		data, errx := models.CreatedSos(sos, notif)

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

	limit := c.Query("limit")
	offset := c.Query("offset")

	err := c.BindQuery(&getsos)
	if err != nil {
		var m string
		if err != nil {
			m = m + err.Error()
		}
		response.ApiMessage = "validation " + m
		c.JSON(400, response)
	} else {
		data, errx := models.GetSosPetugas(getsos, limit, offset)
		response.Data = data
		if errx != nil {
			response.ApiMessage = t.GetMessageErr()
			c.JSON(400, response)
		} else {
			response.ApiStatus = 1
			response.ApiMessage = t.GetMessageSucc()
			c.JSON(200, response)
		}
	}
}

//func GetSosadmin(c *gin.Context) {
//	getsos := structs.GetSosAdmin{}
//	t := structs.Component{}
//
//	response := structs.JsonResponse{}
//
//	err := c.BindQuery(&getsos)
//	if err != nil {
//		var m string
//		if err != nil {
//			m = m + err.Error()
//		}
//		response.ApiMessage = "validation " + m
//		c.JSON(400, response)
//	} else {
//		data, errx := models.GetSosAdmin(getsos)
//		response.Data = data
//		if errx != nil {
//			response.ApiMessage = t.GetMessageErr()
//			c.JSON(400, response)
//		} else {
//			response.ApiStatus = 1
//			response.ApiMessage = t.GetMessageSucc()
//			c.JSON(200, response)
//		}
//	}
//}

func DetailSosPetugas(c *gin.Context) {
	detailSOS := structs.DetailSosPetugas{}
	t := structs.Component{}
	response := structs.JsonResponse{}
	err := c.BindQuery(&detailSOS)
	if err != nil {
		var a string
		if err != nil {
			a = a + err.Error()
		}
		response.ApiMessage = "validation " + a
		c.JSON(400, response)
	} else {
		data, err2 := models.DetailSosPetugas(detailSOS)
		response.Data = data
		if err2 != nil {
			response.ApiMessage = t.GetMessageErr()
			c.JSON(400, response)
		} else {
			response.ApiStatus = 1
			response.ApiMessage = t.GetMessageSucc()
			c.JSON(200, response)
		}
	}
}

//func DetailSosAdmin(c *gin.Context) {
//	sos := structs.DetailSosAdmin{}
//	t := structs.Component{}
//	response := structs.JsonResponse{}
//	err := c.BindQuery(&sos)
//	if err != nil {
//		var a string
//		if err != nil {
//			a = a + err.Error()
//		}
//		response.ApiMessage = "validation " + a
//		c.JSON(400, response)
//	} else {
//		data, err2 := models.DetailSosAdmin(sos)
//		response.Data = data
//		if err2 != nil {
//			response.ApiMessage = t.GetMessageErr()
//			c.JSON(400, response)
//		} else {
//			response.ApiStatus = 1
//			response.ApiMessage = t.GetMessageSucc()
//			c.JSON(200, response)
//		}
//	}
//}
