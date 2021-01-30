package dataThree

import (
	"fmt"
	"net/http"
)

func GetData(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Get data3")
}
func PostData(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Post data3")
}
func UpdateData(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Update data3")
}
func DeleteData(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Delete data3")
}
