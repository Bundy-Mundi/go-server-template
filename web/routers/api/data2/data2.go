package dataTwo

import (
	"fmt"
	"net/http"
)

func GetData(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Get data2")
}
func PostData(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Post data2")
}
func UpdateData(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Update data2")
}
func DeleteData(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Delete data2")
}
