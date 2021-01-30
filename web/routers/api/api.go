package api

import (
	dataOne "github.com/Bundy-Mundi/server-template/web/routers/api/data1"
	dataTwo "github.com/Bundy-Mundi/server-template/web/routers/api/data2"
	dataThree "github.com/Bundy-Mundi/server-template/web/routers/api/data3"
	"github.com/gorilla/mux"
)

// EnrollRouter - Enroll Api
func EnrollRouter(mux *mux.Router) {

	s1 := mux.PathPrefix("/data1/").Subrouter()
	s1.HandleFunc("/", dataOne.GetData).Methods("GET")
	s1.HandleFunc("/", dataOne.GetData).Methods("POST")
	s1.HandleFunc("/", dataOne.GetData).Methods("DELETE")
	s1.HandleFunc("/", dataOne.GetData).Methods("PUT")

	s2 := mux.PathPrefix("/data2/").Subrouter()
	s2.HandleFunc("/", dataTwo.GetData).Methods("GET")
	s2.HandleFunc("/", dataTwo.GetData).Methods("POST")
	s2.HandleFunc("/", dataTwo.GetData).Methods("DELETE")
	s2.HandleFunc("/", dataTwo.GetData).Methods("PUT")

	s3 := mux.PathPrefix("/data3/").Subrouter()
	s3.HandleFunc("/", dataThree.GetData).Methods("GET")
	s3.HandleFunc("/", dataThree.GetData).Methods("POST")
	s3.HandleFunc("/", dataThree.GetData).Methods("DELETE")
	s3.HandleFunc("/", dataThree.GetData).Methods("PUT")
}
