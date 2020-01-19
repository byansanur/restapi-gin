package models

import (
	"../config_db"
	"../structs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"mime/multipart"
	"strconv"
	"time"
)

// fungsi unutk membuat user baru
// parameter di inisialisasi diatas sesuai pada structnya

func CreateUsers2(users structs.CreateUsers) (structs.CreateUsers, error) {

	var err error
	var t = structs.Component{}
	tx := idb.DB.Begin()
	if err = tx.Error; err != nil {
		fmt.Println("err start tx", err.Error())
		return users, err
	}
	users.CreatedAt = t.GetTimeNow()
	if err = tx.Table("users").Create(&users).Error; err != nil {
		tx.Rollback()
		return users, err
	}
	tx.Commit()
	return users, err
}

func CreateUsers(nama string, username string, password string,
	tgl_lahir string, no_ktp string, no_hp string, no_visa string,
	no_passpor string, id_privileges string) structs.JsonResponse {

	// pembuatan variable dengan mengarahkan ke struct yang dibuat
	var (
		users   structs.CreateUsers
		CheckId structs.CheckId
		t       structs.Component
	)

	// pemanggilan json response dengan membuat variable dan panggil struct
	response := structs.JsonResponse{}

	// query mysql
	check := idb.DB.Table("users").Select("count(users.id) as count_id")

	// pemberitahuan pada db untuk variable username ada pada table users dan field username
	check = check.Where("users.username = ?", username)

	// check first id
	check = check.First(&CheckId)

	// untuk cek error
	checkx := check.Error

	// jika error itu kosong atau nil, maka tidak ada error
	if checkx == nil {
		fmt.Println("cek id ga error")

		// cek id jika id = 0 maka password di encrypt pada database dan hasil encrypt bukan hasil convert
		// melainkan hilang pada database
		if CheckId.CountId == 0 {
			encrptPassword, _ := EncryptPassword(password)
			id_privileges_conv, _ := strconv.Atoi(id_privileges)

			//url := UploadImage("user", fmt.Sprint(username), files, header)

			//if url != "" {
			fmt.Println("foto ga null")

			// pengenalan variable setiap field
			users.Nama = nama
			users.Username = username
			users.Password = encrptPassword
			users.TglLahir = tgl_lahir
			users.NoKtp = no_ktp
			users.NoHp = no_hp
			users.NoVisa = no_visa
			users.NoPasspor = no_passpor
			//users.Foto = url
			users.IdPrivileges = id_privileges_conv
			users.CreatedAt = t.GetTimeNow()

			// query untuk insert ke table users
			err := idb.DB.Table("users").Create(&users)

			// jika gagal insert ke table
			errx := err.Error

			// kondisi jika gagal buat akun
			if errx != nil {
				fmt.Println("gagal buat akun")
				response.ApiMessage = t.GetMessageErr()
			} else {
				// jika berhasil buat akun atau insert ke db maka responsenya disini
				response.ApiStatus = 1
				response.ApiMessage = t.GetMessageSucc()
				response.Data = users
			}
			//}
		} else {
			// jika user mendaftar dengan username yang sama maka akan response disini
			response.ApiMessage = "Username Already Used"
		}
	} else {
		// jika gagal cek id
		fmt.Println(checkx)
	}
	// kembalian response
	return response
}

func LoginAdmin(username string, password string) structs.JsonResponse {

	var (
		userlogin structs.CekUserLogin
		users     structs.GetUserLogin
	)

	response := structs.JsonResponse{}

	cekUsername := idb.DB.Table("users").Select("id,username,password")
	cekUsername = cekUsername.Where("users.username = ?", username)

	cekUsername = cekUsername.Scan(&userlogin)
	cekUsernames := cekUsername.RecordNotFound()

	fmt.Println(userlogin.Id)

	if cekUsernames {
		fmt.Println("belum terdaftar")
		response.ApiMessage = "belum terdaftar"
	} else {
		cekPassword, errPass := DecryptPassword(userlogin.Password)

		if errPass != nil {
			fmt.Println("username/password salah")
			response.ApiMessage = "Password salah"
		} else {
			fmt.Println("ok ada nih")
			if cekPassword == password {
				fmt.Println("Pass sama")
				err := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_" +
					"lahir, users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, " +
					"users.id_privileges, tb_privileges.role, users.created_at")
				err = err.Joins("join tb_privileges on users.id_privileges = tb_privileges.id")
				err = err.Where("users.id = ?", userlogin.Id)
				err = err.Where("users.id_privileges = 1")
				err = err.First(&users)
				errx := err.Error

				sign := jwt.New(jwt.SigningMethodHS256)
				claim := sign.Claims.(jwt.MapClaims)

				claim["authorized"] = true
				claim["client"] = users.Id
				claim["exp"] = time.Now().Add(time.Minute * 360).Unix()

				token, _ := sign.SignedString(config_db.JwtKey())
				users.Token = token

				if errx == nil {
					response.ApiStatus = 1
					response.ApiMessage = "success login"
					response.Data = users
				} else {
					response.ApiMessage = "gagal login"
				}
			} else {
				response.ApiMessage = "password salah"
			}
		}
	}
	return response
}

func LoginPetugas(username string, password string) structs.JsonResponse {

	var (
		userlogin structs.CekUserLogin
		users     structs.GetUserLogin
	)

	response := structs.JsonResponse{}

	cekUsername := idb.DB.Table("users").Select("id,username,password")
	cekUsername = cekUsername.Where("users.username = ?", username)

	cekUsername = cekUsername.Scan(&userlogin)
	cekUsernames := cekUsername.RecordNotFound()

	fmt.Println(userlogin.Id)

	if cekUsernames {
		fmt.Println("belum terdaftar")
		response.ApiMessage = "belum terdaftar"
	} else {
		cekPassword, errPass := DecryptPassword(userlogin.Password)

		if errPass != nil {
			fmt.Println("username/password salah")
			response.ApiMessage = "Password salah"
		} else {
			fmt.Println("ok ada nih")
			if cekPassword == password {
				fmt.Println("Pass sama")
				err := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir, " +
					"users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, " +
					"tb_privileges.role, users.created_at")
				err = err.Joins("join tb_privileges on users.id_privileges = tb_privileges.id")
				err = err.Where("users.id = ?", userlogin.Id)
				err = err.Where("users.id_privileges = 2")
				err = err.First(&users)
				errx := err.Error

				sign := jwt.New(jwt.SigningMethodHS256)
				claim := sign.Claims.(jwt.MapClaims)

				claim["authorized"] = true
				claim["client"] = users.Id
				claim["exp"] = time.Now().Add(time.Minute * 360).Unix()

				token, _ := sign.SignedString(config_db.JwtKey())
				users.Token = token

				if errx == nil {
					response.ApiStatus = 1
					response.ApiMessage = "success login"
					response.Data = users
				} else {
					response.ApiMessage = "gagal login"
				}
			} else {
				response.ApiMessage = "password salah"
			}
		}
	}
	return response
}

func LoginUsers(username string, password string) structs.JsonResponse {

	var (
		userlogin structs.CekUserLogin
		users     structs.GetUserLogin
	)

	response := structs.JsonResponse{}

	cekUsername := idb.DB.Table("users").Select("id,username,password")
	cekUsername = cekUsername.Where("users.username = ?", username)
	cekUsername = cekUsername.Where("users.id_privileges = 3")

	cekUsername = cekUsername.Scan(&userlogin)
	cekUsernames := cekUsername.RecordNotFound()

	fmt.Println(userlogin.Id)

	if cekUsernames {
		fmt.Println("belum terdaftar")
		response.ApiMessage = "belum terdaftar"
	} else {
		cekPassword, errPass := DecryptPassword(userlogin.Password)

		if errPass != nil {
			fmt.Println("username/password salah")
			response.ApiMessage = "Password salah"
		} else {
			fmt.Println("ok ada nih")
			if cekPassword == password {
				fmt.Println("Pass sama")
				err := idb.DB.Table("users").Select("users.id, users.nama, users.username, " +
					"users.tgl_lahir, users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, " +
					"users.foto, users.id_privileges, tb_privileges.role, users.created_at")
				err = err.Joins("join tb_privileges on users.id_privileges = tb_privileges.id")
				err = err.Where("users.id = ?", userlogin.Id)
				err = err.First(&users)
				errx := err.Error

				sign := jwt.New(jwt.SigningMethodHS256)
				claim := sign.Claims.(jwt.MapClaims)

				claim["authorized"] = true
				claim["client"] = users.Id
				claim["exp"] = time.Now().Add(time.Hour * 24).Unix()

				token, _ := sign.SignedString(config_db.JwtKey())
				users.Token = token

				if errx == nil {
					response.ApiStatus = 1
					response.ApiMessage = "success login"
					response.Data = users
				} else {
					response.ApiMessage = "gagal login"
				}
			} else {
				response.ApiMessage = "password salah"
			}
		}
	}
	return response
}

//func UserUploadFoto(upload structs.UploadFoto) (structs.UploadFoto, error) {
//	data := structs.UploadFoto{}
//	db := idb.DB.Table("users")
//}

func SearchUsers(search structs.SearchUser) ([]structs.SearchUser, error) {
	var data []structs.SearchUser
	getDb := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir," +
		" users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, " +
		"users.created_at," + "tb_privileges.role").
		Joins("JOIN tb_privileges ON users.id_privileges = tb_privileges.id").
		Where("users.id_privileges != 1")

	if search.Nama != "" {
		getDb = getDb.Where("users.nama LIKE ?", "%"+search.Nama+"%")
	}

	err := getDb.Find(&data).Error
	fmt.Println("users is  ", data)
	return data, err
}

func GetUsersAll(users structs.GetUser, limit string, offset string) ([]structs.GetUser, error) {
	var data []structs.GetUser
	get := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir," +
		" users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, " +
		"date_format(users.created_at, '%m-%a-%Y %H:%i') as created_at," + "tb_privileges.role").
		Joins("JOIN tb_privileges ON users.id_privileges = tb_privileges.id").
		Order("users.created_at desc")

	if limit != "" {
		get = get.Limit(limit)
	}
	if offset != "" {
		get = get.Offset(offset)
	}
	if users.IdPrivileges != nil {
		get = get.Where("users.id_privileges in (?)", int(*users.IdPrivileges))
	}
	if users.Id != nil {
		get = get.Where("users.id in (?)", int(*users.Id))
	}

	err := get.Find(&data).Error

	return data, err
}

func GetJamaah(created_at string, nama string, username string, tgl_lahir string, no_ktp string,
	no_hp string, no_visa string, no_passpor string, id_privileges string, offset string, limit string) structs.JsonResponse {
	var (
		getUser []structs.GetUser
		t       structs.Component
	)

	response := structs.JsonResponse{}

	err := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir," +
		" users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, " +
		"users.created_at," + "tb_privileges.role")
	err = err.Joins("JOIN tb_privileges ON users.id_privileges = tb_privileges.id")
	err = err.Where("users.id_privileges = 3")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if created_at != "" {
		err = err.Where("users.created_at LIKE ?", "%"+created_at+"%")
	}
	if nama != "" {
		err = err.Where("users.nama = ?", nama)
	}
	if username != "" {
		err = err.Where("users.username = ?", username)
	}
	if tgl_lahir != "" {
		err = err.Where("users.tgl_lahir = ?", tgl_lahir)
	}
	if no_ktp != "" {
		err = err.Where("users.no_ktp = ?", no_ktp)
	}
	if no_hp != "" {
		err = err.Where("users.no_hp = ?", no_hp)
	}
	if no_visa != "" {
		err = err.Where("users.no_visa = ?", no_visa)
	}
	if no_passpor != "" {
		err = err.Where("users.no_passpor = ?", no_passpor)
	}
	if id_privileges != "" {
		err = err.Where("users.id_privileges = ?", id_privileges)
	}

	err = err.Order("users.nama asc")

	err = err.Find(&getUser)
	errx := err.Error

	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = getUser
	}
	return response
}

func GetUsersFoto(id string, foto string, id_privileges string) structs.JsonResponse {
	var (
		getfoto []structs.GetFoto
		t       structs.Component
	)

	response := structs.JsonResponse{}
	err := idb.DB.Table("users").Select("users.id, users.foto, users.id_privileges, tb_privileges.role")
	err = err.Joins("JOIN tb_privileges ON users.id_privileges = tb_privileges.id")
	err = err.Find(&getfoto)
	errx := err.Error

	if id != "" {
		err = err.Where("users.id = ?", id)
	}
	if foto != "" {
		err = err.Where("users.foto = ?", foto)
	}
	if id_privileges != "" {
		err = err.Where("users.id_privileges = ?", id_privileges)
	}

	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = getfoto
	}
	return response
}

func GetPetugas(created_at string, nama string, username string, tgl_lahir string, no_ktp string,
	no_hp string, no_visa string, no_passpor string, id_privileges string, offset string, limit string) structs.JsonResponse {
	var (
		getUser []structs.GetUser
		t       structs.Component
	)

	response := structs.JsonResponse{}

	err := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir, " +
		"users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, " +
		"users.created_at," + "tb_privileges.role")
	err = err.Joins("JOIN tb_privileges ON users.id_privileges = tb_privileges.id")
	err = err.Where("users.id_privileges = 2")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if created_at != "" {
		err = err.Where("users.created_at LIKE ?", "%"+created_at+"%")
	}
	if nama != "" {
		err = err.Where("users.nama = ?", nama)
	}
	if username != "" {
		err = err.Where("users.username = ?", username)
	}
	if tgl_lahir != "" {
		err = err.Where("users.tgl_lahir = ?", tgl_lahir)
	}
	if no_ktp != "" {
		err = err.Where("users.no_ktp = ?", no_ktp)
	}
	if no_hp != "" {
		err = err.Where("users.no_hp = ?", no_hp)
	}
	if no_visa != "" {
		err = err.Where("users.no_visa = ?", no_visa)
	}
	if no_passpor != "" {
		err = err.Where("users.no_passpor = ?", no_passpor)
	}
	if id_privileges != "" {
		err = err.Where("users.id_privileges = ?", id_privileges)
	}

	err = err.Order("users.nama asc")

	err = err.Find(&getUser)
	errx := err.Error

	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = getUser
	}
	return response
}

func GetUsers(nama string, username string, tgl_lahir string, no_ktp string,
	no_hp string, no_visa string, no_passpor string, id_privileges string, offset string, limit string) structs.JsonResponse {
	var (
		getUser []structs.GetUser
		t       structs.Component
	)

	response := structs.JsonResponse{}

	err := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir, " +
		"users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, " +
		"users.created_at," + "tb_privileges.role")
	err = err.Joins("JOIN tb_privileges ON users.id_privileges = tb_privileges.id")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if nama != "" {
		err = err.Where("users.nama = ?", nama)
	}
	if username != "" {
		err = err.Where("users.username = ?", username)
	}
	if tgl_lahir != "" {
		err = err.Where("users.tgl_lahir = ?", tgl_lahir)
	}
	if no_ktp != "" {
		err = err.Where("users.no_ktp = ?", no_ktp)
	}
	if no_hp != "" {
		err = err.Where("users.no_hp = ?", no_hp)
	}
	if no_visa != "" {
		err = err.Where("users.no_visa = ?", no_visa)
	}
	if no_passpor != "" {
		err = err.Where("users.no_passpor = ?", no_passpor)
	}
	if id_privileges != "" {
		err = err.Where("users.id_privileges = ?", id_privileges)
	}

	err = err.Order("users.nama asc")

	err = err.Find(&getUser)
	errx := err.Error

	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = getUser
	}
	return response
}

func FetchUsersProfile(users structs.GetUser) (structs.GetUser, error) {
	data := structs.GetUser{}
	get := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir, " +
		"users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, " +
		"users.created_at," + "tb_privileges.role").
		Joins("JOIN tb_privileges ON users.id_privileges = tb_privileges.id")

	if users.Id != nil {
		get = get.Where("users.id = ?", users.Id)
	}
	if users.IdPrivileges != nil {
		get = get.Where("users.id_privileges = ?", users.IdPrivileges)
	}
	err := get.Find(&data).Error
	fmt.Println("getUser", data)
	return data, err
}

func FetchAllUsers(usersAll structs.GetUser, limit string, offset string) ([]structs.GetUser, error) {
	var data []structs.GetUser
	get := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir, " +
		"users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, " +
		"users.created_at," + "tb_privileges.role").
		Joins("JOIN tb_privileges ON users.id_privileges = tb_privileges.id")

	if usersAll.IdPrivileges != nil {
		get = get.Where("users.id_privileges = ?", usersAll.IdPrivileges)
	}
	if limit != "" {
		get = get.Limit(limit)
	}
	if offset != "" {
		get = get.Offset(offset)
	}
	err := get.Find(&data).Error
	fmt.Println("all users", data)
	return data, err
}

func GetUserDetail(id string) structs.JsonResponse {

	var (
		users structs.GetUser
		t     structs.Component
	)

	response := structs.JsonResponse{}
	err := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir, " +
		"users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, " +
		"users.created_at," + "tb_privileges.role")
	err = err.Joins("JOIN tb_privileges ON users.id_privileges = tb_privileges.id")
	err = err.Where("users.id = ?", id)
	err = err.First(&users)
	errx := err.Error

	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = users
	}
	return response
}

func UpdateUser(id string, nama string, username string, password string,
	no_hp string, no_visa string, no_passpor string, files multipart.File, header *multipart.FileHeader) structs.JsonResponse {
	var (
		userUpdate structs.UpdateUsers
		t          structs.Component
	)
	response := structs.JsonResponse{}
	encryptPassword, _ := EncryptPassword(password)
	id_conv, _ := strconv.Atoi(id)
	no_hp_conv, _ := strconv.Atoi(no_hp)

	url := UploadImage("user", fmt.Sprint(username), files, header)

	if password != "" {
		userUpdate.Password = encryptPassword
	}

	userUpdate.Nama = nama
	userUpdate.Username = username
	userUpdate.NoHp = no_hp_conv
	userUpdate.NoVisa = no_visa
	userUpdate.NoPasspor = no_passpor
	userUpdate.Foto = url

	err := idb.DB.Table("users").Where("users.id = ?", id_conv).Update(&userUpdate)
	errx := err.Error

	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = userUpdate
	}
	return response
}

func UpdateLocationUsers(id string, lat string, lng string) structs.JsonResponse {
	var (
		locUpdate structs.UpdateLocationUsers
		t         structs.Component
	)
	response := structs.JsonResponse{}
	id_conv, _ := strconv.Atoi(id)
	locUpdate.Id = id_conv
	locUpdate.Lat = lat
	locUpdate.Lng = lng
	err := idb.DB.Table("users").Where("users.id = ?", id_conv).Update(&locUpdate)
	err2 := err.Error
	if err2 != nil {
		response.ApiStatus = 0
		response.ApiMessage = err2.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = locUpdate
	}
	return response
}

func GetLocationUsers(loc structs.GetLocationUsers) ([]structs.GetLocationUsers, error) {
	var data []structs.GetLocationUsers
	get := idb.DB.Table("users").
		Select("users.id, users.nama, users.lat, users.lng, users.id_privileges, tb_privileges.role").
		Joins("Join tb_privileges on users.id_privileges = tb_privileges.id")

	if loc.IdPrivileges != nil {
		get = get.Where("users.id_privileges = ?", loc.IdPrivileges)
	}

	err := get.Find(&data).Error
	fmt.Println("loc users", data)
	return data, err
}
