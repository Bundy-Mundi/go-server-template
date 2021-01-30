package web

import (
	"net/http"

	"github.com/Bundy-Mundi/server-template/web/routers"
)

// StartServer - ...
func StartServer() {
	staticpath := http.Dir("./assets")

	mux := routers.NewRouters(staticpath)
	http.ListenAndServe(":3000", mux)
}
