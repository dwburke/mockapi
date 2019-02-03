package api_test

import (
	//"bytes"
	//"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"

	"github.com/dwburke/mockapi/api"
	"github.com/dwburke/mockapi/config"
)

func setup() (r *mux.Router) {
	r = mux.NewRouter()

	config.Config = &config.ConfigType{
		Routes: map[string]*config.Route{
			"/get/thing/{id}": &config.Route{
				Method:     "GET",
				Result:     "{\"thing-id\":{{- .Params.id -}}}",
				ResultType: "application/json",
			},
			"/get/str/{uuid}": &config.Route{
				Method:     "GET",
				Result:     "thing-id: {{ .Params.uuid }}",
				ResultType: "text/html",
			},
		},
	}

	api.SetupRoutes(r)

	return
}

func TestMockGetString(t *testing.T) {

	r := setup()

	req, _ := http.NewRequest("GET", "/get/str/d78a1e65-a2f4-43e5-aaf9-654ee11d68ae", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	body, err := ioutil.ReadAll(w.Body)
	expect(t, err, nil, "")

	expect(t, w.Code, http.StatusOK, string(body))
	expect(t, string(body), string("thing-id: d78a1e65-a2f4-43e5-aaf9-654ee11d68ae"), string(body))
}

func TestMockGetJson(t *testing.T) {

	r := setup()

	req, _ := http.NewRequest("GET", "/get/thing/667", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	body, err := ioutil.ReadAll(w.Body)
	expect(t, err, nil, "")

	expect(t, w.Code, http.StatusOK, string(body))
	expect(t, string(body), string("{\"thing-id\":667}"), string(body))

	//var room types.Room
	//err = json.Unmarshal(body, &room)
	//expect(t, err, nil, "")
	//expect(t, room.Vnum, 1265, "")

}

func expect(t *testing.T, a interface{}, b interface{}, body string) {
	if a != b {
		t.Errorf("Expected [%v] (type %v) - Got [%v] (type %v) : %s", b, reflect.TypeOf(b), a, reflect.TypeOf(a), body)
	}
}
