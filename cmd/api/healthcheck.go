package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	// js := `{"status":"available", "environment": %q, "version": %q}`
	// js = fmt.Sprintf(js, app.config.env, version)

	envelope := envelope{
		"status": "available",
		"system_info": map[string]string{
			"env":     app.config.env,
			"version": version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, envelope, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "ERR, could not process request", http.StatusInternalServerError)
		return
	}

}
