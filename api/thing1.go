package api

import (
	"net/http"

	helpers "github.com/dwburke/go-tools/gorillamuxhelpers"
	"github.com/gorilla/mux"
	//"github.com/spf13/cast"
)

func MockMethod(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)

	//if !helpers.CheckRequiredVar(w, params, "ip") {
	//return
	//}

	//agent := db.GetAgent(cast.ToString(params["ip"]))

	//if agent == nil {
	//helpers.RespondWithError(w, 404, "agent does not exist")
	//return
	//}

	if template, err := mux.CurrentRoute(r).GetPathTemplate(); err != nil {
		helpers.RespondWithError(w, 500, err.Error())
		return
	} else {
		helpers.RespondWithJSON(w, 200, template)
	}
	//helpers.RespondWithJSON(w, 200, r.CurrentRoute().GetPathTemplate())
}
