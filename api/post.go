package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"text/template"

	helpers "github.com/dwburke/go-tools/gorillamuxhelpers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	"github.com/dwburke/mockapi/config"
)

func MockPost(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")

	//var params = make(map[string]string)
	params := mux.Vars(r)
	var jsonObj = make(map[string]interface{})

	if contentType == "application/json" {
		if r.Body != nil {
			err := json.NewDecoder(r.Body).Decode(&jsonObj)
			if err != nil {
				helpers.RespondWithError(w, 404, err.Error())
				return
			}

			for k, v := range jsonObj {
				params[k] = cast.ToString(v)
			}
		}
	}

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

	log.Infof("path_template: %s; info.Result: %s; result: %s", path_template, info.Result, result)

	if info.ResultType == "application/json" {
		var jsonObj interface{}
		err := json.Unmarshal([]byte(result), &jsonObj)
		if err != nil {
			helpers.RespondWithError(w, 500, err.Error())
			return
		}
		helpers.RespondWithJSON(w, 200, jsonObj)
	} else {
		helpers.RespondWithString(w, 200, result)
	}
}
