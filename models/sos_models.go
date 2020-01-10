package models

import (
	"../structs"
	"fmt"
)

func CreatedSos(sos structs.CreateSos, notif structs.CreateSosNotif, admin structs.CreateSosAdmin) (structs.CreateSos, error) {
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
	if err = tx.Table("notif_petugas").Create(&notif).Error; err != nil {
		tx.Rollback()
		return sos, err
	}
	admin.IdSosSender = sos.Id
	if err = tx.Table("notif_admin").Create(&admin).Error; err != nil {
		tx.Rollback()
		return sos, err
	}
	tx.Commit()
	return sos, err
}

func GetSosPetugas(petugasSos structs.GetSosPetugas) ([]structs.GetSosPetugas, error) {
	var data []structs.GetSosPetugas
	gets := idb.DB.Table("notif_petugas").Select("notif_petugas.id, notif_petugas.id_sos_jamaah, " +
		"notif_petugas.id_users_penerima, petugas.nama AS nama_petugas, sos_jamaah.sos, sos_jamaah.id_users_sender, " +
		"user_sender.nama AS nama_pengirim, sos_jamaah.lat, sos_jamaah.lng, " +
		"date_format(sos_jamaah.created_at, '%m-%a-%Y %H:%i') as created_at").
		Joins("JOIN sos_jamaah ON notif_petugas.id_sos_jamaah = sos_jamaah.id").
		Joins("JOIN users AS petugas ON notif_petugas.id_users_penerima = petugas.id").
		Joins("JOIN users AS user_sender ON sos_jamaah.id_users_sender = user_sender.id").
		Order("sos_jamaah.created_at desc")

	if petugasSos.Id != nil {
		gets = gets.Where("notif_petugas.id in (?)", int(*petugasSos.Id))
	}
	if petugasSos.NamaPengirim != nil {
		gets = gets.Where("sos_jamaah.nama  LIKE ?", "%"+string(*petugasSos.NamaPengirim)+"&")
	}
	err := gets.Find(&data).Error
	fmt.Println("get notif", data)
	return data, err
}

func GetSosAdmin(adminsos structs.GetSosAdmin) ([]structs.GetSosAdmin, error) {
	var data []structs.GetSosAdmin
	get := idb.DB.Table("notif_admin").Select("notif_admin.id, notif_admin.id_sos_sender, " +
		"notif_admin.id_users_admin, admin.nama as nama_admin, sos_jamaah.sos, sos_jamaah.id_users_sender, " +
		"user_table.nama as nama_user, sos_jamaah.lat, sos_jamaah.lng, " +
		"date_format(sos_jamaah.created_at, '%m-%a-%Y %H:%i') as created_at").
		Joins("JOIN users AS admin ON notif_admin.id_users_admin = admin.id").
		Joins("JOIN sos_jamaah ON notif_admin.id_sos_sender = sos_jamaah.id").
		Joins("JOIN users AS user_table ON sos_jamaah.id_users_sender = user_table.id").
		Order("sos_jamaah.created_at desc")
	if adminsos.Id != nil {
		get = get.Where("notif_admin.id in (?)", int(*adminsos.Id))
	}
	if adminsos.NamaUser != nil {
		get = get.Where("sos_jamaah.nama  LIKE ?", "%"+string(*adminsos.NamaUser)+"%")
	}
	err := get.Find(&data).Error
	fmt.Println("get notif", data)
	return data, err
}

func DetailSosPetugas(detail structs.DetailSosPetugas) (structs.DetailSosPetugas, error) {
	var data structs.DetailSosPetugas
	get := idb.DB.Table("notif_petugas").Select("notif_petugas.id, notif_petugas.id_sos_jamaah, " +
		"notif_petugas.id_users_penerima, petugas.nama AS nama_petugas, sos_jamaah.sos, sos_jamaah.id_users_sender," +
		" user_sender.nama AS nama_pengirim, sos_jamaah.lat, sos_jamaah.lng, user_sender.no_ktp, user_sender.no_hp, " +
		"user_sender.no_visa, user_sender.no_passpor, user_sender.foto, user_sender.id_privileges, tb_privileges.role," +
		"date_format(sos_jamaah.created_at, '%m-%a-%Y %H:%i') as created_at ").
		Joins("JOIN sos_jamaah ON notif_petugas.id_sos_jamaah = sos_jamaah.id").
		Joins("JOIN users AS petugas ON notif_petugas.id_users_penerima = petugas.id").
		Joins("JOIN users AS user_sender ON sos_jamaah.id_users_sender = user_sender.id").
		Joins("JOIN tb_privileges ON user_sender.id_privileges = tb_privileges.id")

	if detail.Id != nil {
		get = get.Where("notif_petugas.id in (?)", int(*detail.Id))
	}
	err := get.Find(&data).Error
	fmt.Println("get detail => ", data)
	return data, err
}

func DetailSosAdmin(detail structs.DetailSosAdmin) (structs.DetailSosAdmin, error) {
	var data structs.DetailSosAdmin
	get := idb.DB.Table("notif_admin").Select("notif_admin.id, notif_admin.id_sos_sender, " +
		"notif_admin.id_users_admin, admin.nama as nama_admin, sos_jamaah.sos, sos_jamaah.id_users_sender, " +
		"user_table.nama as nama_user, sos_jamaah.lat, sos_jamaah.lng, user_table.no_ktp, user_table.no_hp, " +
		"user_table.no_visa, user_table.no_passpor, user_table.foto, user_table.id_privileges, tb_privileges.role, " +
		"date_format(sos_jamaah.created_at, '%m-%a-%Y %H:%i') as created_at").
		Joins("JOIN users AS admin ON notif_admin.id_users_admin = admin.id").
		Joins("JOIN sos_jamaah ON notif_admin.id_sos_sender = sos_jamaah.id").
		Joins("JOIN users AS user_table ON sos_jamaah.id_users_sender = user_table.id").
		Joins("JOIN tb_privileges ON user_table.id_privileges = tb_privileges.id")
	if detail.Id != nil {
		get = get.Where("notif_admin.id in (?)", int(*detail.Id))
	}
	err := get.Find(&data).Error
	fmt.Println("get detail => ", data)
	return data, err
}
