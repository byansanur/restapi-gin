package structs

type CreateSos struct {
	Id            int    `json:"id" form:"id"`
	Sos           bool   `json:"sos" form:"sos" binding:"required"`
	IdUsersSender int    `json:"id_users_sender" form:"id_users_sender" binding:"required"`
	Lat           string `json:"lat" form:"lat" binding:"required"`
	Lng           string `json:"lng" form:"lng" binding:"-"`
	CreatedAt     string `json:"created_at" form:"created_at"`
}

type CreateSosNotif struct {
	Id              int `json:"id" form:"id"`
	IdSosJamaah     int `json:"id_sos_jamaah" form:"id_sos_jamaah"`
	IdUsersPenerima int `json:"id_users_penerima" form:"id_users_penerima"`
}

type CreateSosAdmin struct {
	Id           int `json:"id" form:"id"`
	IdSosSender  int `json:"id_sos_sender" form:"id_sos_sender"`
	IdUsersAdmin int `json:"id_users_admin" form:"id_users_admin"`
}

type GetSosPetugas struct {
	Id              *int    `json:"id" form:"id"`
	IdSosJamaah     *int    `json:"id_sos_jamaah" form:"id_sos_jamaah"`
	IdUsersPenerima *int    `json:"id_users_penerima" form:"id_users_penerima"`
	NamaPetugas     *string `json:"nama_petugas" form:"nama_petugas"`
	Sos             *bool   `json:"sos" form:"sos"`
	IdUsersSender   *int    `json:"id_users_sender" form:"id_users_sender"`
	NamaPengirim    *string `json:"nama_pengirim" form:"nama_pengirim"`
	Lat             *string `json:"lat" form:"lat"`
	Lng             *string `json:"lng" form:"lng"`
	CreatedAt       *string `json:"created_at" form:"created_at"`
}

type GetSosAdmin struct {
	Id            *int    `json:"id" form:"id"`
	IdSosSender   *int    `json:"id_sos_sender" form:"id_sos_sender"`
	IdUsersAdmin  *int    `json:"id_users_admin" form:"id_users_admin"`
	NamaAdmin     *string `json:"nama_admin" form:"nama_admin"`
	Sos           *bool   `json:"sos" form:"sos"`
	IdUsersSender *int    `json:"id_users_sender" form:"id_users_sender"`
	NamaUser      *string `json:"nama_user" form:"nama_user"`
	Lat           *string `json:"lat" form:"lat"`
	Lng           *string `json:"lng" form:"lng"`
	CreatedAt     *string `json:"created_at" form:"created_at"`
}

type DetailSosPetugas struct {
	Id              *int    `json:"id" form:"id" binding:"required"`
	IdSosJamaah     *int    `json:"id_sos_jamaah" form:"id_sos_jamaah"`
	IdUsersPenerima *int    `json:"id_users_penerima" form:"id_users_penerima"`
	NamaPetugas     *string `json:"nama_petugas" form:"nama_petugas"`
	Sos             *bool   `json:"sos" form:"sos"`
	IdUsersSender   *int    `json:"id_users_sender" form:"id_users_sender"`
	NamaPengirim    *string `json:"nama_pengirim" form:"nama_pengirim"`
	Lat             *string `json:"lat" form:"lat"`
	Lng             *string `json:"lng" form:"lng"`
	NoKtp           *int    `json:"no_ktp" form:"no_ktp"`
	NoHp            *int    `json:"no_hp" form:"no_hp"`
	NoVisa          *string `json:"no_visa" form:"no_visa"`
	NoPasspor       *string `json:"no_passpor" form:"no_passpor"`
	Foto            *string `json:"foto" form:"foto"`
	IdPrivileges    *int    `json:"id_privileges" form:"id_privileges"`
	Role            *string `json:"role" form:"role"`
	CreatedAt       *string `json:"created_at" form:"created_at"`
}

type DetailSosAdmin struct {
	Id            *int    `json:"id" form:"id" binding:"required"`
	IdSosSender   *int    `json:"id_sos_sender" form:"id_sos_sender"`
	IdUsersAdmin  *int    `json:"id_users_admin" form:"id_users_admin"`
	NamaAdmin     *string `json:"nama_admin" form:"nama_admin"`
	Sos           *bool   `json:"sos" form:"sos"`
	IdUsersSender *int    `json:"id_users_sender" form:"id_users_sender"`
	NamaUser      *string `json:"nama_user" form:"nama_user"`
	Lat           *string `json:"lat" form:"lat"`
	Lng           *string `json:"lng" form:"lng"`
	NoKtp         *int    `json:"no_ktp" form:"no_ktp"`
	NoHp          *int    `json:"no_hp" form:"no_hp"`
	NoVisa        *string `json:"no_visa" form:"no_visa"`
	NoPasspor     *string `json:"no_passpor" form:"no_passpor"`
	Foto          *string `json:"foto" form:"foto"`
	IdPrivileges  *int    `json:"id_privileges" form:"id_privileges"`
	Role          *string `json:"role" form:"role"`
	CreatedAt     *string `json:"created_at" form:"created_at"`
}
