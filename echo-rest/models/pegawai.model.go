package models

import (
	"net/http"

	"restgolang.com/restGolang/echo-rest/db"
)

type Pegawai struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

func FetchPegawai() (Response, error) {
	var obj Pegawai
	var arrObj []Pegawai
	var res Response

	conn := db.CreateCon()

	sqlStament	:= "SELECT * FROM PEGAWAI"

	rows, err := conn.Query(sqlStament)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)
		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Data berhasil di ambil"
	res.Data = arrObj

	return res, nil
}

func StorePegawai(nama string, alamat string, telepon string)(Response, error) {

	var res Response

	conn := db.CreateCon()

	sqlStament := "INSERT PEGAWAI (nama, alamat, telepon) VALUES (?, ?, ?)"

	stmt, err := conn.Prepare(sqlStament)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Data Berhasil disimpan"
	res.Data = map[string]int64{
		"lastInsertedId" : lastInsertedId,
	}

	return res, nil
}

func UpdatePegawai(id int, nama string, alamat string, telepon string) (Response, error) {

	var res Response

	conn := db.CreateCon()

	sqlStatement := "UPDATE PEGAWAI SET NAMA = ?, ALAMAT = ?, TELEPON = ? WHERE ID = ?"

	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon, id)
	if err != nil {
		return res, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Data berhasil diubah"
	res.Data = map[string]int64{
		"rowAffected" : rowAffected,
	}

	return res, nil
}