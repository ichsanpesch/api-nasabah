package model

type Nasabah struct {
	Id             int64  `json:"id"`
	Version        string `json:"version"`
	IdCard         string `json:"id_card"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Name           string `json:"name"`
	BirthPlace     string `json:"birth_place"`
	BirthDate      string `json:"birth_date"`
	Gender         string `json:"gender"`
	Address        string `json:"address"`
	FamilyStatus   int64  `json:"family_status"`
	FamilyPosition string `json:"family_position"`
	CreateDate     string `json:"create_date"`
	CreateBy       string `json:"create_by"`
	UpdateDate     string `json:"update_date"`
	UpdateBy       string `json:"update_by"`
	Status         int64  `json:"status"`
}

type Response struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Nasabah `json:"data"`
}
