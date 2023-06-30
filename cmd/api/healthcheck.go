package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"version": version,
			"env":     app.config.env,
		},
	}
	app.writeJSON(w, http.StatusOK, data, nil)

}
