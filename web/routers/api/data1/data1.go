package dataOne

import (
	"fmt"
	"net/http"
)

func GetData(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Get data1")
}
func PostData(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Post data1")
}
func UpdateData(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Update data1")
}
func DeleteData(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Delete data1")
}
