package Services

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/zenazn/goji/web"

	P_Database "api-nasabah/Database"
	P_Model "api-nasabah/Model"
)

func ServiceGetNasabahAll(c web.C, w http.ResponseWriter, r *http.Request) {

	nasabah := P_Model.Nasabah{}
	arrayNasabah := []P_Model.Nasabah{}
	response := P_Model.Response{}
	db := P_Database.Connection()

	query, err := db.Query("SELECT * FROM nasabah WHERE status=1")
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

func ServiceGetNasabahById(c web.C, w http.ResponseWriter, r *http.Request) {

	nasabah := P_Model.Nasabah{}
	arrayNasabah := []P_Model.Nasabah{}
	response := P_Model.Response{}
	db := P_Database.Connection()

	query, err := db.Query("SELECT * FROM nasabah WHERE status=1 and id=?", c.URLParams["id"])
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

func ServiceAddNasabah(c web.C, w http.ResponseWriter, r *http.Request) {

	nasabah := P_Model.Nasabah{}
	arrayNasabah := []P_Model.Nasabah{}
	response := P_Model.Response{}
	db := P_Database.Connection()

	decodeByte, _ := base64.StdEncoding.DecodeString(c.URLParams["data"])

	json.Unmarshal(decodeByte, &nasabah)

	nasabah.Version = strconv.FormatInt(time.Now().UTC().UnixNano()/1000000, 10)
	nasabah.CreateDate = time.Now().UTC().Local().Format("2006-01-02")
	nasabah.CreateBy = "SYSTEM"
	nasabah.Status = 1

	query, err := db.Prepare("INSERT INTO nasabah(version, create_date, create_by, id_card, email, password, name) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	run, err := query.Exec(nasabah.Version, nasabah.CreateDate, nasabah.CreateBy, nasabah.IdCard, nasabah.Email, nasabah.Password, nasabah.Name)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	resultLastInsertId, err := run.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	resultRowsAffected, err := run.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	nasabah.Id = resultLastInsertId
	arrayNasabah = append(arrayNasabah, nasabah)

	response.Status = 200
	response.Message = "Success inserted " + strconv.FormatInt(resultRowsAffected, 16) + " row"
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

func ServiceUpdateNasabahById(c web.C, w http.ResponseWriter, r *http.Request) {

	nasabah := P_Model.Nasabah{}
	arrayNasabah := []P_Model.Nasabah{}
	response := P_Model.Response{}
	db := P_Database.Connection()

	decodeByte, _ := base64.StdEncoding.DecodeString(c.URLParams["data"])

	json.Unmarshal(decodeByte, &nasabah)

	nasabah.Version = strconv.FormatInt(time.Now().UTC().UnixNano()/1000000, 10)
	nasabah.UpdateDate = time.Now().UTC().Local().Format("2006-01-02")
	nasabah.UpdateBy = "SYSTEM"

	query, err := db.Prepare("UPDATE nasabah SET version = ?, id_card = ?, email = ?, password = ?, name = ?, birth_place = ?, birth_date = ?, gender = ?, address = ?, family_status = ?, family_position = ?, create_date = ?, create_by = ?, update_date = ?, update_by = ?, status = ? WHERE id = ?")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	run, err := query.Exec(nasabah.Version, nasabah.IdCard, nasabah.Email, nasabah.Password, nasabah.Name, nasabah.BirthPlace, nasabah.BirthDate, nasabah.Gender, nasabah.Address, nasabah.FamilyStatus, nasabah.FamilyPosition, nasabah.CreateDate, nasabah.CreateBy, nasabah.UpdateDate, nasabah.UpdateBy, nasabah.Status, nasabah.Id)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	resultRowsAffected, err := run.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	arrayNasabah = append(arrayNasabah, nasabah)

	response.Status = 200
	response.Message = "Success updated " + strconv.FormatInt(resultRowsAffected, 16) + " row"
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

func ServiceDeleteNasabahById(c web.C, w http.ResponseWriter, r *http.Request) {

	nasabah := P_Model.Nasabah{}
	arrayNasabah := []P_Model.Nasabah{}
	response := P_Model.Response{}
	db := P_Database.Connection()

	query, err := db.Prepare("DELETE FROM nasabah WHERE id=?")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	run, err := query.Exec(c.URLParams["id"])
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	resultRowsAffected, err := run.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	arrayNasabah = append(arrayNasabah, nasabah)

	response.Status = 200
	response.Message = "Success deleted " + strconv.FormatInt(resultRowsAffected, 16) + " row"
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
