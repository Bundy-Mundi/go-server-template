package users

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Fprint(res, vars["id"])
}
func GetUsers(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "All users")
}
