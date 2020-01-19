package controllers

import (
	"../models"
	"../structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func GetUsersNew(c *gin.Context) {
	getu := structs.GetUser{}
	t := structs.Component{}

	limit := c.Query("limit")
	offset := c.Query("offset")

	response := structs.JsonResponse{}

	err := c.BindQuery(&getu)

	if err != nil {
		var m string
		if err != nil {
			m = m + err.Error()
		}
		response.ApiMessage = "validation " + m
		c.JSON(400, response)
	} else {
		data, errx := models.GetUsersAll(getu, limit, offset)
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

func CreateUsers(c *gin.Context) {
	nama := c.PostForm("nama")
	username := c.PostForm("username")
	password := c.PostForm("password")
	tgl_lahir := c.PostForm("tgl_lahir")
	no_ktp := c.PostForm("no_ktp")
	no_hp := c.PostForm("no_hp")
	no_visa := c.PostForm("no_ktp")
	no_passpor := c.PostForm("no_passpor")
	//photo, header, _ := c.Request.FormFile("photo")
	id_privileges := c.PostForm("id_privileges")

	response := structs.JsonResponse{}

	if nama == "" {
		response.ApiMessage = "isi nama"
	} else if username == "" {
		response.ApiMessage = "isi username"
	} else if password == "" {
		response.ApiMessage = "isi password"
	} else if tgl_lahir == "" {
		response.ApiMessage = "input tgl lahir"
	} else if no_ktp == "" {
		response.ApiMessage = "masukan nomor kependudukan anda"
	} else if no_visa == "" {
		response.ApiMessage = "masukan nomor visa anda"
	} else if no_passpor == "" {
		response.ApiMessage = "masukan nomor passpor anda"
	} else if no_hp == "" {
		response.ApiMessage = "masukan nomor hp anda"
	} else if id_privileges == "" {
		response.ApiMessage = "isi privileges"
	} else {
		strs := strings.Split(id_privileges, ",")
		ary := make([]int, len(strs))
		for i := range ary {
			ary[i], _ = strconv.Atoi(strs[i])
		}
		fmt.Println("arry ", ary)
		response = models.CreateUsers(nama, username, password, tgl_lahir, no_ktp, no_hp, no_visa, no_passpor, id_privileges)
	}

	if response.ApiStatus == 1 {

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func LoginAdmin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	response := structs.JsonResponse{}

	if username == "" {
		response.ApiMessage = "Required username"
	} else if password == "" {
		response.ApiMessage = "Required Password"
	} else {
		response = models.LoginAdmin(username, password)
	}

	if response.ApiStatus == 1 {

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func LoginPetugas(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	response := structs.JsonResponse{}

	if username == "" {
		response.ApiMessage = "Required username"
	} else if password == "" {
		response.ApiMessage = "Required Password"
	} else {
		response = models.LoginPetugas(username, password)
	}

	if response.ApiStatus == 1 {

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func LoginUsers(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	response := structs.JsonResponse{}

	if username == "" {
		response.ApiMessage = "Required username"
	} else if password == "" {
		response.ApiMessage = "Required Password"
	} else {
		response = models.LoginUsers(username, password)
	}

	if response.ApiStatus == 1 {

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func FetchUsers(c *gin.Context) {
	users := structs.GetUser{}
	t := structs.Component{}

	response := structs.JsonResponse{}
	err := c.BindQuery(&users)
	if err != nil {
		var m string
		if err != nil {
			m = m + err.Error()
		}
		response.ApiMessage = "validation " + m
		c.JSON(400, response)
	} else {
		data, errs := models.FetchUsersProfile(users)
		response.Data = data
		if errs != nil {
			response.ApiMessage = t.GetMessageErr()
			c.JSON(400, response)
		} else {
			response.ApiStatus = 1
			response.ApiMessage = t.GetMessageSucc()
			c.JSON(200, response)
		}
	}
}

func FetchAllUsers(c *gin.Context) {
	users := structs.GetUser{}
	t := structs.Component{}
	limit := c.Query("limit")
	offset := c.Query("offset")

	response := structs.JsonResponse{}
	err := c.BindQuery(&users)
	if err != nil {
		var m string
		if err != nil {
			m = m + err.Error()
		}
		response.ApiMessage = "validation " + m
		c.JSON(400, response)
	} else {
		data, errs := models.FetchAllUsers(users, limit, offset)
		response.Data = data
		if errs != nil {
			response.ApiMessage = t.GetMessageErr()
			c.JSON(400, response)
		} else {
			response.ApiStatus = 1
			response.ApiMessage = t.GetMessageSucc()
			c.JSON(200, response)
		}
	}
}

func FetchLocationUsers(c *gin.Context) {
	userLoc := structs.GetLocationUsers{}
	t := structs.Component{}

	response := structs.JsonResponse{}
	err := c.BindQuery(&userLoc)
	if err != nil {
		var m string
		if err != nil {
			m = m + err.Error()
		}
		response.ApiMessage = "validation " + m
		c.JSON(400, response)
	} else {
		data, errs := models.GetLocationUsers(userLoc)
		response.Data = data
		if errs != nil {
			response.ApiMessage = t.GetMessageErr()
			c.JSON(400, response)
		} else {
			response.ApiStatus = 1
			response.ApiMessage = t.GetMessageSucc()
			c.JSON(200, response)
		}
	}
}
func GetJamaah(c *gin.Context) {
	created_at := c.Query("created_at")
	nama := c.Query("nama")
	username := c.Query("username")
	limit := c.Query("limit")
	offset := c.Query("offset")
	tgl_lahir := c.Query("tgl_lahir")
	no_ktp := c.Query("no_ktp")
	no_hp := c.Query("no_hp")
	no_visa := c.Query("no_ktp")
	no_passpor := c.Query("no_passpor")
	id_privileges := c.Query("id_privileges")

	response := structs.JsonResponse{}

	response = models.GetJamaah(created_at, nama, username, tgl_lahir, no_ktp, no_hp, no_visa, no_passpor, id_privileges, offset, limit)

	if response.ApiStatus == 1 {

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func SearchUser(c *gin.Context) {
	search := structs.SearchUser{}
	t := structs.Component{}
	res := structs.JsonResponse{}
	err := c.BindQuery(&search)
	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}
		res.ApiMessage = "validate " + mess
		c.JSON(400, res)
	} else {
		data, errs := models.SearchUsers(search)
		res.Data = data
		if errs != nil {
			res.ApiMessage = t.GetMessageErr()
			c.JSON(400, res)
		} else {
			res.ApiStatus = 1
			res.ApiMessage = t.GetMessageSucc()
			c.JSON(http.StatusOK, res)
		}
	}
}

func GetFoto(c *gin.Context) {
	id := c.Query("id")
	foto := c.Query("foto")
	id_privileges := c.Query("id_privileges")

	response := structs.JsonResponse{}
	response = models.GetUsersFoto(id, foto, id_privileges)

	if response.ApiStatus == 1 {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func GetPetugas(c *gin.Context) {
	created_at := c.Query("created_at")
	nama := c.Query("nama")
	username := c.Query("username")
	limit := c.Query("limit")
	offset := c.Query("offset")
	tgl_lahir := c.Query("tgl_lahir")
	no_ktp := c.Query("no_ktp")
	no_hp := c.Query("no_hp")
	no_visa := c.Query("no_ktp")
	no_passpor := c.Query("no_passpor")
	id_privileges := c.Query("id_privileges")

	response := structs.JsonResponse{}

	response = models.GetPetugas(created_at, nama, username, tgl_lahir, no_ktp, no_hp, no_visa, no_passpor, id_privileges, offset, limit)

	if response.ApiStatus == 1 {

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func GetUsers(c *gin.Context) {
	nama := c.Query("nama")
	username := c.Query("username")
	limit := c.Query("limit")
	offset := c.Query("offset")
	tgl_lahir := c.Query("tgl_lahir")
	no_ktp := c.Query("no_ktp")
	no_hp := c.Query("no_hp")
	no_visa := c.Query("no_ktp")
	no_passpor := c.Query("no_passpor")
	id_privileges := c.Query("id_privileges")

	response := structs.JsonResponse{}

	response = models.GetUsers(nama, username, tgl_lahir, no_ktp, no_hp, no_visa, no_passpor, id_privileges, offset, limit)

	if response.ApiStatus == 1 {

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func GetUserDetail(c *gin.Context) {
	id := c.Query("id")

	responses := structs.JsonResponse{}

	if id == "" {
		responses.ApiStatus = 0
		responses.ApiMessage = "Required Id"
		responses.Data = nil

	} else {
		responses = models.GetUserDetail(id)

	}
	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(400, responses)
	}
}

func UpdateUser(c *gin.Context) {

	id := c.PostForm("id")
	nama := c.PostForm("nama")
	username := c.PostForm("username")
	password := c.PostForm("password")
	no_hp := c.PostForm("no_hp")
	no_visa := c.PostForm("no_visa")
	no_passpor := c.PostForm("no_passpor")
	foto, header, _ := c.Request.FormFile("foto")

	responses := structs.JsonResponse{}

	if id == "" {
		responses.ApiMessage = "Required Id"
	} else {

		responses = models.UpdateUser(id, nama, username, password, no_hp, no_visa, no_passpor, foto, header)
	}

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}

func UpdateLocUser(c *gin.Context) {

	id := c.PostForm("id")
	lat := c.PostForm("lat")
	lng := c.PostForm("lng")

	responses := structs.JsonResponse{}

	if id == "" {
		responses.ApiMessage = "Required Id"
	} else {

		responses = models.UpdateLocationUsers(id, lat, lng)
	}

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}
