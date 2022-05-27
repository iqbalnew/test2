package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

///database mahasiswa
type Mahasiswa struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
	Nim  string `json:"nim"`
}

var db *gorm.DB
var err error

func main() {
	fmt.Println(1 + 1)

	dsn := "root:@tcp(127.0.0.1:3306)/17220294?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// di test ini
	if err != nil {
		fmt.Println("koneksi gagal")
	} else {
		fmt.Println("koneksi berhasil")
	}

	db.AutoMigrate(&Mahasiswa{})

	handleRequest()

}

func handleRequest() {
	fmt.Println("router runnung in port:9999")
	r := mux.NewRouter()

	r.HandleFunc("/", homePage)
	r.HandleFunc("/mahasiswa/", mahasiswaAdd).Methods("POST")
	r.HandleFunc("/mahasiswa/", getmahasiswas).Methods("GET")
	r.HandleFunc("/mahasiswa/{id}", getmahasiswa).Methods("GET")
	r.HandleFunc("/mahasiswa/{id}", mahasiswaupd).Methods("PUT")
	r.HandleFunc("/mahasiswa/{id}", delmahasiswa).Methods("DELETE")

	http.ListenAndServe(":9999", r)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")

}
func getmahasiswas(w http.ResponseWriter, r *http.Request) {
	mahasiswa := []Mahasiswa{}

	db.Find(&mahasiswa)

	res, _ := json.Marshal(mahasiswa)

	w.Header().Set("content-type", "application/type")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func getmahasiswa(w http.ResponseWriter, r *http.Request) {
	mahasiswa := []Mahasiswa{}
	mahasiswaId := mux.Vars(r)["id"]
	db.First(&mahasiswa, mahasiswaId)

	res, _ := json.Marshal(mahasiswa)

	w.Header().Set("content-type", "application/type")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func delmahasiswa(w http.ResponseWriter, r *http.Request) {

	mahasiswaId := mux.Vars(r)["id"]
	var mahasiswa Mahasiswa
	db.First(&mahasiswa, mahasiswaId)
	db.Delete(&mahasiswa)

	res, _ := json.Marshal(mahasiswa)

	w.Header().Set("content-type", "application/type")
	w.WriteHeader(http.StatusOK)

	w.Write(res)

}

func mahasiswaAdd(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Manambah data")

	payloads, _ := ioutil.ReadAll(r.Body)
	var mahasiswa Mahasiswa
	json.Unmarshal(payloads, &mahasiswa)
	db.Create(&mahasiswa)

	res, _ := json.Marshal(mahasiswa)

	w.Header().Set("content-type", "application/type")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	w.Write(payloads)

}

func mahasiswaupd(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Manambah data")

	payloads, _ := ioutil.ReadAll(r.Body)
	var mahasiswaup Mahasiswa
	json.Unmarshal(payloads, &mahasiswaup)

	mahasiswaId := mux.Vars(r)["id"]
	var mahasiswa Mahasiswa
	db.First(&mahasiswa, mahasiswaId)
	db.Model(&mahasiswa).Updates(mahasiswaup)

	res, _ := json.Marshal(mahasiswa)

	w.Header().Set("content-type", "application/type")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	w.Write(payloads)

}
