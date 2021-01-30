package rooms

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoomByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Fprint(res, vars["id"])
}
func GetRooms(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "All rooms")
}
