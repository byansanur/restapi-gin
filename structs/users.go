package structs

import "mime/multipart"

type CreateUsers struct {
	Id        int    `json:"id" form:"id"`
	Nama      string `json:"nama" form:"nama"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	TglLahir  string `json:"tgl_lahir" form:"tgl_lahir"`
	NoKtp     string `json:"no_ktp" form:"no_ktp" binding:"required"`
	NoHp      string `json:"no_hp" form:"no_hp" binding:"required"`
	NoVisa    string `json:"no_visa" form:"no_visa" binding:"required"`
	NoPasspor string `json:"no_passpor" form:"no_passpor" binding:"required"`
	//Foto         string `json:"foto" form:"foto"`
	IdPrivileges int    `json:"id_privileges" form:"-" binding:"-"`
	CreatedAt    string `json:"created_at" form:"created_at"`
}

type UploadFoto struct {
	Id           int                   `json:"id" form:"id"`
	Foto         *multipart.FileHeader `json:"foto" form:"foto"`
	IdPrivileges int                   `json:"id_privileges" form:"id_privileges"`
}

type GetUser struct {
	Id        *int    `json:"id" form:"id"`
	Nama      *string `json:"nama" form:"nama"`
	Username  *string `json:"username" form:"username"`
	TglLahir  *string `json:"tgl_lahir" form:"tgl_lahir"`
	NoKtp     *int    `json:"no_ktp" form:"no_ktp"`
	NoHp      *int    `json:"no_hp" form:"no_hp"`
	NoVisa    *string `json:"no_visa" form:"no_visa"`
	NoPasspor *string `json:"no_passpor" form:"no_passpor"`
	//Foto         *string `json:"foto" form:"foto"`
	IdPrivileges *int    `json:"id_privileges" form:"id_privileges" binding:"required"`
	Role         *string `json:"role" form:"role"`
	CreatedAt    *string `json:"created_at" form:"created_at"`
}

type SearchUser struct {
	Id           int64  `json:"id" form:"id"`
	Nama         string `json:"nama" form:"nama"`
	Username     string `json:"username" form:"username"`
	TglLahir     string `json:"tgl_lahir" form:"tgl_lahir"`
	NoKtp        int    `json:"no_ktp" form:"no_ktp"`
	NoHp         int    `json:"no_hp" form:"no_hp"`
	NoVisa       string `json:"no_visa" form:"no_visa"`
	NoPasspor    string `json:"no_passpor" form:"no_passpor"`
	Foto         string `json:"foto" form:"foto"`
	IdPrivileges int    `json:"id_privileges" form:"id_privileges" binding:"-"`
	Role         string `json:"role" form:"role"`
	CreatedAt    string `json:"created_at" form:"created_at"`
}

type GetUserLogin struct {
	Id        int     `json:"id"`
	Nama      *string `json:"nama"`
	Username  *string `json:"username"`
	TglLahir  *string `json:"tgl_lahir"`
	NoKtp     *string `json:"no_ktp"`
	NoHp      *string `json:"no_hp"`
	NoVisa    *string `json:"no_visa"`
	NoPasspor *string `json:"no_passpor"`
	//Foto         *string `json:"foto"`
	IdPrivileges *int    `json:"id_privileges"`
	Role         *string `json:"role"`
	CreatedAt    *string `json:"created_at"`
	Token        string  `json:"token"`
}

type GetFoto struct {
	Id           int    `json:"id"`
	Foto         string `json:"foto"`
	IdPrivileges int    `json:"id_privileges"`
	Role         string `json:"role"`
}

type CekUserLogin struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CheckId struct {
	CountId int `json:"count_id"`
}

type UpdateUsers struct {
	Id        int    `json:"id"`
	Nama      string `json:"nama"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	NoHp      int    `json:"no_hp"`
	NoVisa    string `json:"no_visa"`
	NoPasspor string `json:"no_passpor"`
	Foto      string `json:"foto"`
}
