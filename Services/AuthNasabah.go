package Services

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/zenazn/goji/web"

	P_Database "api-nasabah/Database"
	P_Model "api-nasabah/Model"
)

func AuthNasabah(c web.C, w http.ResponseWriter, r *http.Request) {

	nasabah := P_Model.Nasabah{}
	arrayNasabah := []P_Model.Nasabah{}
	response := P_Model.Response{}
	db := P_Database.Connection()

	decodeByte, _ := base64.StdEncoding.DecodeString(c.URLParams["data"])

	json.Unmarshal(decodeByte, &nasabah)

	nasabah.Status = 1

	query, err := db.Query("SELECT * FROM nasabah WHERE status=? and email=? and password=? ORDER BY id ASC LIMIT 1", nasabah.Status, nasabah.Email, nasabah.Password)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	for query.Next() {
		err = query.Scan(
			&nasabah.Id,
			&nasabah.Version,
			&nasabah.IdCard,
			&nasabah.Email,
			&nasabah.Password,
			&nasabah.Name,
			&nasabah.BirthPlace,
			&nasabah.BirthDate,
			&nasabah.Gender,
			&nasabah.Address,
			&nasabah.FamilyStatus,
			&nasabah.FamilyPosition,
			&nasabah.CreateDate,
			&nasabah.CreateBy,
			&nasabah.UpdateDate,
			&nasabah.UpdateBy,
			&nasabah.Status)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		} else {
			arrayNasabah = append(arrayNasabah, nasabah)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrayNasabah

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	log.Println(string(jsonResponse))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	return

}
