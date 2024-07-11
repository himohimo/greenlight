package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	// js := `{"status":"available", "environment": %q, "version": %q}`
	// js = fmt.Sprintf(js, app.config.env, version)

	data := map[string]string{
		"status":  "available",
		"env":     app.config.env,
		"version": version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "ERR, could not process request", http.StatusInternalServerError)
		return
	}

}
