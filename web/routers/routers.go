package routers

import (
	"fmt"
	"net/http"

	"github.com/Bundy-Mundi/server-template/web/routers/rooms"
	"github.com/Bundy-Mundi/server-template/web/routers/users"
	"github.com/gorilla/mux"
)

// NewRouters - Handle all routers from here
func NewRouters(dir http.Dir) http.Handler {

	mux := mux.NewRouter()

	/* Home */
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Home")
	})

	/* Room Router */
	s1 := mux.PathPrefix("/rooms").Subrouter()
	s1.HandleFunc("/", rooms.GetRooms)
	s1.HandleFunc("/{id:[0-9]+}", rooms.GetRoomByID)

	/* User Router */
	s2 := mux.PathPrefix("/users").Subrouter()
	s2.HandleFunc("/", users.GetUsers)
	s2.HandleFunc("/{id:[0-9]+}", users.GetUserByID)

	/* Static File Server */
	/* ABOUT FILE PATH:  https://stackoverflow.com/questions/52141282/http-fileserverhttp-dir-not-working-in-separate-package */
	// fs := http.StripPrefix("/static/", http.FileServer(dir))
	// mux.Handle("/static/", fs)
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	return mux
}
