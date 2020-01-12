package models

import (
	"../structs"
	"fmt"
)

func CreatedSos(sos structs.CreateSos, notif structs.CreateSosNotif) (structs.CreateSos, error) {
	var err error
	var t = structs.Component{}
	tx := idb.DB.Begin()

	if err = tx.Error; err != nil {
		return sos, err
	}
	sos.CreatedAt = t.GetTimeNow()
	if err = tx.Table("sos_jamaah").Create(&sos).Error; err != nil {
		tx.Rollback()
		return sos, err
	}
	notif.IdSosJamaah = sos.Id
	if err = tx.Table("notif").Create(&notif).Error; err != nil {
		tx.Rollback()
		return sos, err
	}
	tx.Commit()
	return sos, err
}

func GetSosPetugas(petugasSos structs.GetSosPetugas) ([]structs.GetSosPetugas, error) {
	var data []structs.GetSosPetugas
	gets := idb.DB.Table("notif").Select("notif.id, notif.id_sos_jamaah, sos_jamaah.message, " +
		"sos_jamaah.id_users_sender, users.nama, sos_jamaah.lat, sos_jamaah.lng, " +
		"date_format(sos_jamaah.created_at, '%d-%M-%Y %H:%i') as created_at").
		Joins("JOIN sos_jamaah ON notif.id_sos_jamaah = sos_jamaah.id").
		Joins("JOIN users ON sos_jamaah.id_users_sender = users.id").
		Order("sos_jamaah.created_at desc")

	if petugasSos.Id != nil {
		gets = gets.Where("notif.id in (?)", int(*petugasSos.Id))
	}
	if petugasSos.IdUsersSender != nil {
		gets = gets.Where("sos_jamaah.id_users_sender in (?)", int(*petugasSos.IdUsersSender))
	}
	if petugasSos.Nama != nil {
		gets = gets.Where("sos_jamaah.nama  LIKE ?", "%"+string(*petugasSos.Nama)+"&")
	}
	err := gets.Find(&data).Error
	fmt.Println("get notif", data)
	return data, err
}

//
//func GetSosAdmin(adminsos structs.GetSosAdmin) ([]structs.GetSosAdmin, error) {
//	var data []structs.GetSosAdmin
//	get := idb.DB.Table("notif_admin").Select("notif_admin.id, notif_admin.id_sos_sender, " +
//		"notif_admin.id_users_admin, admin.nama as nama_admin, sos_jamaah.sos, sos_jamaah.id_users_sender, " +
//		"user_table.nama as nama_user, sos_jamaah.lat, sos_jamaah.lng, " +
//		"date_format(sos_jamaah.created_at, '%m-%a-%Y %H:%i') as created_at").
//		Joins("JOIN users AS admin ON notif_admin.id_users_admin = admin.id").
//		Joins("JOIN sos_jamaah ON notif_admin.id_sos_sender = sos_jamaah.id").
//		Joins("JOIN users AS user_table ON sos_jamaah.id_users_sender = user_table.id").
//		Order("sos_jamaah.created_at desc")
//	if adminsos.Id != nil {
//		get = get.Where("notif_admin.id in (?)", int(*adminsos.Id))
//	}
//	if adminsos.NamaUser != nil {
//		get = get.Where("sos_jamaah.nama  LIKE ?", "%"+string(*adminsos.NamaUser)+"%")
//	}
//	err := get.Find(&data).Error
//	fmt.Println("get notif", data)
//	return data, err
//}

func DetailSosPetugas(detail structs.DetailSosPetugas) (structs.DetailSosPetugas, error) {
	var data structs.DetailSosPetugas
	get := idb.DB.Table("notif").Select("notif.id, notif.id_sos_jamaah, " +
		"sos_jamaah.message, sos_jamaah.id_users_sender," +
		"date_format(sos_jamaah.created_at, '%d-%M-%Y %H:%i') as created_at, " +
		"users.nama, sos_jamaah.lat, sos_jamaah.lng, users.no_ktp, users.no_hp, " +
		"users.no_visa, users.no_passpor, tb_privileges.role ").
		Joins("JOIN sos_jamaah ON notif.id_sos_jamaah = sos_jamaah.id").
		Joins("JOIN users ON sos_jamaah.id_users_sender = users.id").
		Joins("JOIN tb_privileges ON users.id_privileges = tb_privileges.id")

	if detail.Id != nil {
		get = get.Where("notif.id in (?)", int(*detail.Id))
	}
	if detail.IdUsersSender != nil {
		get = get.Where("sos_jamaah.id_users_sender in (?)", int(*detail.IdUsersSender))
	}
	err := get.Find(&data).Error
	fmt.Println("get detail => ", data)
	return data, err
}

//func DetailSosAdmin(detail structs.DetailSosAdmin) (structs.DetailSosAdmin, error) {
//	var data structs.DetailSosAdmin
//	get := idb.DB.Table("notif_admin").Select("notif_admin.id, notif_admin.id_sos_sender, " +
//		"notif_admin.id_users_admin, admin.nama as nama_admin, sos_jamaah.sos, sos_jamaah.id_users_sender, " +
//		"user_table.nama as nama_user, sos_jamaah.lat, sos_jamaah.lng, user_table.no_ktp, user_table.no_hp, " +
//		"user_table.no_visa, user_table.no_passpor, user_table.foto, user_table.id_privileges, tb_privileges.role, " +
//		"date_format(sos_jamaah.created_at, '%m-%a-%Y %H:%i') as created_at").
//		Joins("JOIN users AS admin ON notif_admin.id_users_admin = admin.id").
//		Joins("JOIN sos_jamaah ON notif_admin.id_sos_sender = sos_jamaah.id").
//		Joins("JOIN users AS user_table ON sos_jamaah.id_users_sender = user_table.id").
//		Joins("JOIN tb_privileges ON user_table.id_privileges = tb_privileges.id")
//	if detail.Id != nil {
//		get = get.Where("notif_admin.id in (?)", int(*detail.Id))
//	}
//	err := get.Find(&data).Error
//	fmt.Println("get detail => ", data)
//	return data, err
//}
