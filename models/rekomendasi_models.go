package models

import (
	"../structs"
	"fmt"
	"mime/multipart"
	"strconv"
)

func CreateRekomendasi(nama string, alamat string, files multipart.File,
	header *multipart.FileHeader, lat string, lng string, rating string, id_type string) structs.JsonResponse {

	var (
		rekomCreate structs.CreateRekomendasi
		cekIdRekom  structs.CheckIdRekom
		t           structs.Component
	)

	response := structs.JsonResponse{}
	check := idb.DB.Table("rekomendasi").Select("count(rekomendasi.id) as count_id")
	check = check.Where("rekomendasi.nama = ?", nama)

	check = check.First(&cekIdRekom)

	checkx := check.Error

	if checkx == nil {
		fmt.Println("cek id ga error")

		if cekIdRekom.CountId == 0 {
			id_type_conv, _ := strconv.Atoi(id_type)

			url := UploadImage("rekom", fmt.Sprint(nama), files, header)

			if url != "" {
				fmt.Println("foto tidak kosong")

				rekomCreate.Nama = nama
				rekomCreate.Alamat = alamat
				rekomCreate.Foto = url
				rekomCreate.Lat = lat
				rekomCreate.Lng = lng
				rekomCreate.Rating = rating
				rekomCreate.IdType = id_type_conv
				rekomCreate.CreatedAt = t.GetTimeNow()

				err := idb.DB.Table("rekomendasi").Create(&rekomCreate)

				errx := err.Error

				if errx != nil {
					fmt.Println("gagal buat rekomendasi")
					response.ApiMessage = t.GetMessageErr()
				} else {
					// jika berhasil buat akun atau insert ke db maka responsenya disini
					response.ApiStatus = 1
					response.ApiMessage = t.GetMessageSucc()
					response.Data = rekomCreate
				}
			}
		} else {
			// jika user mendaftar dengan username yang sama maka akan response disini
			response.ApiMessage = "Nama is Already Used"
		}
	} else {
		// jika gagal cek id
		fmt.Println(checkx)
	}
	return response
}

func GetRekomends(nama string, alamat string, foto string, rating string, lat string, lng string, id_type string, offset string, limit string) structs.JsonResponse {
	var (
		getRekom []structs.GetRekomendasi
		t        structs.Component
	)

	response := structs.JsonResponse{}

	err := idb.DB.Table("rekomendasi").Select("rekomendasi.id, rekomendasi.nama, rekomendasi.alamat, " +
		"rekomendasi.foto, rekomendasi.rating, rekomendasi.lat, rekomendasi.lng, rekomendasi.id_type," +
		"typerekom.type_rekom")
	err = err.Joins("join typerekom on rekomendasi.id_type = typerekom.id")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if nama != "" {
		err = err.Where("rekomendasi.nama like ?", "%"+nama+"%")
	}
	if alamat != "" {
		err = err.Where("rekomendasi.alamat like ?", "%"+alamat+"%")
	}
	if rating != "" {
		err = err.Where("rekomendasi.rating like ?", "%"+rating+"%")
	}
	if lat != "" {
		err = err.Where("rekomendasi.lat")
	}
	if lng != "" {
		err = err.Where("rekomendasi.lng")
	}
	if id_type != "" {
		err = err.Where("rekomendasi.id_type = ?", id_type)
	}
	if foto != "" {
		err = err.Where("rekomendasi.foto = ?", foto)
	}

	//err = err.Order("rekomendasi.rating desc")

	err = err.Find(&getRekom)
	errx := err.Error

	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = getRekom
	}
	return response
}

func GetRekomFoto(foto structs.RekomFoto) ([]structs.RekomFoto, error) {
	var data []structs.RekomFoto
	getDb := idb.DB.Table("rekomendasi").Select("rekomendasi.id, rekomendasi.foto, rekomendasi.id_type")
	if foto.Id != nil {
		getDb = getDb.Where("rekomendasi.id in (?)", int64(*foto.Id))
	}
	if foto.Foto != nil {
		getDb = getDb.Where("rekomendasi.foto in (?)", *foto.Foto)
	}
	if foto.IdType != nil {
		getDb = getDb.Where("rekomendasi.id_type in (?)", int64(*foto.IdType))
	}

	err := getDb.Find(&data).Error
	fmt.Println("foto rekomendasi", data)
	return data, err
}

func GetRekomendsDetail(id string) structs.JsonResponse {
	var (
		rekom structs.GetRekomendasi
		t     structs.Component
	)
	response := structs.JsonResponse{}
	err := idb.DB.Table("rekomendasi").Select("rekomendasi.id, rekomendasi.nama, rekomendasi.alamat, rekomendasi.foto, " +
		"rekomendasi.rating, rekomendasi.lat, rekomendasi.lng, rekomendasi.id_type, date_format(rekomendasi.created_at, '%m-%a-%Y %H:%i') as created_at," + "typerekom.type_rekom")
	err = err.Joins("join typerekom on rekomendasi.id_type = typerekom.id")
	err = err.Where("rekomendasi.id = ?", id)
	err = err.First(&rekom)
	errx := err.Error
	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = rekom
	}
	return response
}

func UpdateRekomendasi(id string, nama string, alamat string,
	files multipart.File, header *multipart.FileHeader, rating string, lat string, lng string) structs.JsonResponse {
	var (
		updateRekom structs.UpdateRekomendasi
		t           structs.Component
	)

	response := structs.JsonResponse{}

	id_conv, _ := strconv.Atoi(id)
	url := UploadImage("nama", fmt.Sprint(nama), files, header)

	updateRekom.Id = id_conv
	updateRekom.Nama = nama
	updateRekom.Alamat = alamat
	updateRekom.Foto = url
	updateRekom.Rating = rating
	updateRekom.Lat = lat
	updateRekom.Lng = lng

	err := idb.DB.Table("rekomendasi").Where("rekomendasi.id = ?", id_conv).Update(&updateRekom)
	errx := err.Error

	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = updateRekom
	}
	return response
}
