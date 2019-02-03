package api_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/dwburke/mockapi/api"
	"github.com/dwburke/mockapi/config"
)

func mockPostSetup() (r *mux.Router) {
	r = mux.NewRouter()

	config.Config = &config.ConfigType{
		Routes: map[string]*config.Route{
			"/post/thing/{id}": &config.Route{
				Method:     "POST",
				Result:     "{\"thing-id\":{{- .Params.id -}}}",
				ResultType: "application/json",
			},
			"/post/str/{uuid}": &config.Route{
				Method:     "POST",
				Result:     "thing-id: {{ .Params.uuid }}",
				ResultType: "text/html",
			},
		},
	}

	api.SetupRoutes(r)

	return
}

func TestMockPostString(t *testing.T) {

	r := mockPostSetup()

	req, _ := http.NewRequest("POST", "/post/str/e78a1e65-a2f4-43e5-aaf9-654ee11d68ae", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	body, err := ioutil.ReadAll(w.Body)
	expect(t, err, nil, "")

	expect(t, w.Code, http.StatusOK, string(body))
	expect(t, string(body), string("thing-id: e78a1e65-a2f4-43e5-aaf9-654ee11d68ae"), string(body))
}

func TestMockPostJson(t *testing.T) {

	r := mockPostSetup()

	req, _ := http.NewRequest("POST", "/post/thing/668", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	body, err := ioutil.ReadAll(w.Body)
	expect(t, err, nil, "")

	expect(t, w.Code, http.StatusOK, string(body))
	expect(t, string(body), string("{\"thing-id\":668}"), string(body))

	type JsonType struct {
		Id int `json:"thing-id"`
	}

	var thing JsonType
	err = json.Unmarshal(body, &thing)
	expect(t, err, nil, "")
	expect(t, thing.Id, 668, "")

}
