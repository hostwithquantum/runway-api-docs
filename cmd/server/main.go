// main provides a server to serve the docs
// this is the only file that should be edited
// in this repository, the rest is autogenerated
package main

import (
	"fmt"
	log "log/slog"
	"net/http"
	"os"

	"github.com/hostwithquantum/runway-api-docs/docs"
	"github.com/hostwithquantum/runway-api-docs/static"

	"github.com/gorilla/mux"
)

var listen string
var version = "dev"

func init() {
	// set port
	if os.Getenv("PORT") != "" {
		listen = ":" + os.Getenv("PORT")
	} else {
		listen = ":8484"
	}

	log.SetDefault(log.New(log.NewTextHandler(os.Stdout, &log.HandlerOptions{
		Level: log.LevelDebug,
	})).With(log.String("app", "runway-api-docs"), log.String("version", version)))
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/swagger").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	})
	r.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(docs.SwaggerJSON)
	})
	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://ui.runway.horse/favicon.ico", http.StatusPermanentRedirect)
	})
	r.HandleFunc("/rapidoc.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write(static.RapidocJS)
	})
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `<!DOCTYPE html>
<html>
<head>
	<title>Runway API Documentation</title>
	<script defer data-api="https://www.runway.horse/api/event" data-domain="runway.horse" src="https://www.runway.horse/js/script.js"></script>
</head>
<body>
	<rapi-doc
		allow-authentication ='false'
		allow-server-selection='false'
		server-url='https://api.runway.horse'
		schema-style='table'
		show-header='false'
		spec-url='/docs/swagger.json'
		theme='light'>
		<img
			slot="nav-logo"
			src="https://www.runway.horse/img/runway-logo-silverphoenix.svg"
		/>
	</rapi-doc>
	<script src="/rapidoc.js"></script>
</body>
</html>`)
	})

	log.Info("Running on " + listen)
	http.ListenAndServe(listen, r)
}
