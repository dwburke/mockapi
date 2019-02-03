package api

import (
	"bytes"
	"net/http"
	"text/template"

	helpers "github.com/dwburke/go-tools/gorillamuxhelpers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/dwburke/mockapi/config"
)

func MockMethod(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	path_template, err := mux.CurrentRoute(r).GetPathTemplate()
	if err != nil {
		helpers.RespondWithError(w, 404, err.Error())
		return
	}

	info, ok := config.Config.Routes[path_template]
	if !ok {
		helpers.RespondWithError(w, 404, path_template+" not configured")
		return
	}

	t := template.New("result")
	t2, err := t.Parse(info.Result)
	if err != nil {
		helpers.RespondWithError(w, 500, err.Error())
		return
	}

	args := map[string]interface{}{
		"Params": params,
		"Route":  info,
	}

	var tpl bytes.Buffer

	err = t2.Execute(&tpl, args)
	if err != nil {
		helpers.RespondWithError(w, 500, err.Error())
		return
	}

	result := tpl.String()

	log.Debugf("path_template: %s; info.Result: %s; result: %s", path_template, info.Result, result)

	helpers.RespondWithJSON(w, 200, result)
}
