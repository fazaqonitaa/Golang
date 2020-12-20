package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type college struct {
	Idmahasiswa string `json:"id_mahasiswa"`
	Nama        string `json:"nama"`
	Alamat      struct {
		Jalan     string `json:"jalan"`
		Kelurahan string `json:"kelurahan"`
		Kecamatan string `json:"kecamatan"`
		Kabupaten string `json:"kabupaten"`
		Provinsi  string `json:"provinsi"`
	} `json:"alamat"`
	Fakultas string  `json:"fakultas"`
	Jurusan  string  `json:"jurusan"`
	Nilai    []nilai `json:"Nilai"`
}

type nilai struct {
	Idmahasiswa string  `json:"id_mahasiswa"`
	Idmatkul    string  `json:"id_matkul"`
	Mkuliah     string  `json:"m_kuliah"`
	Nilai       float32 `json:"nilai"`
	Semester    int8    `json:"semester"`
}

func main() {

	url := "http://localhost:8080/mahasiswa/1811082001"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	college := college{}
	jsonErr := json.Unmarshal(body, &college)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println("ID Mahasiswa :", college.Idmahasiswa)
	fmt.Println("Nama :", college.Nama)
	fmt.Println("Fakultas :", college.Fakultas)
	fmt.Println("Jurusan :", college.Jurusan)
	fmt.Println("Alamat :")

	fmt.Println("Jalan : ", college.Alamat.Jalan)
	fmt.Println("Kelurahan : ", college.Alamat.Kelurahan)
	fmt.Println("Kecamatan : ", college.Alamat.Kecamatan)
	fmt.Println("Kabupaten : ", college.Alamat.Kabupaten)
	fmt.Println("Provinsi : ", college.Alamat.Provinsi)

	for _, nilai := range college.Nilai {
		fmt.Println("Nama Matkul : ", nilai.Mkuliah)
		fmt.Println("Nilai : ", nilai.Nilai)
		fmt.Println("Semester : ", nilai.Semester)
	}

}
