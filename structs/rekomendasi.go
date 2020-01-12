package structs

type CreateRekomendasi struct {
	Id        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Foto      string `json:"foto"`
	Lat       string `json:"lat"`
	Lng       string `json:"lng"`
	Rating    string `json:"rating"`
	IdType    int    `json:"id_type"`
	CreatedAt string `json:"created_at"`
}
type GetRekomendasi struct {
	Id        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Foto      string `json:"foto"`
	Rating    string `json:"rating"`
	Lat       string `json:"lat"`
	Lng       string `json:"lng"`
	IdType    int    `json:"id_type"`
	TypeRekom string `json:"type_rekom"`
}

type RekomFoto struct {
	Id     *int64  `json:"id"`
	Foto   *string `json:"foto"`
	IdType *int64  `json:"id_type"`
}

type UpdateRekomendasi struct {
	Id     int    `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Foto   string `json:"foto"`
	Rating string `json:"rating"`
	Lat    string `json:"lat"`
	Lng    string `json:"lng"`
}

type CheckIdRekom struct {
	CountId int `json:"count_id"`
}
