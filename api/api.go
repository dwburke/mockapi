package api

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/dwburke/mockapi/config"
)

var ShutdownCh chan bool

var server *http.Server

func SetupRoutes(r *mux.Router) {
	for route, info := range config.Config.Routes {
		log.Info("Adding route: ", info.Method, " ", route)
		if info.Method == "GET" {
			r.HandleFunc(route, MockGet).Methods(info.Method)
		} else if info.Method == "POST" {
			r.HandleFunc(route, MockPost).Methods(info.Method)
		}
	}

}

func init() {
	viper.SetDefault("api.server.address", "127.0.0.1")
	viper.SetDefault("api.server.port", 9000)
	viper.SetDefault("api.server.ssl-key", "key.pem")
	viper.SetDefault("api.server.ssl-cert", "cert.pem")
	viper.SetDefault("api.server.https", false)
}

func Run() {
	var listen string = fmt.Sprintf("%s:%d", viper.GetString("api.server.address"), viper.GetInt("api.server.port"))

	if apiRunning() {
		log.Info("api: already running on ", listen)
		return
	}

	log.Info("api: not running on ", listen, "; starting")

	ShutdownCh = make(chan bool)

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	SetupRoutes(r)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	server = &http.Server{
		Addr:    listen,
		Handler: handlers.CORS(originsOk, headersOk, methodsOk)(r),
	}

	go func() {
		if viper.GetBool("api.server.https") == true {
			if err := server.ListenAndServeTLS(
				viper.GetString("api.server.ssl-cert"),
				viper.GetString("api.server.ssl-key"),
			); err != http.ErrServerClosed {
				log.Fatalf("api: %s", err)
			}
		} else {
			if err := server.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatalf("api: %s", err)
			}
		}
	}()

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func apiRunning() bool {
	var address string = fmt.Sprintf("%s:%d", viper.GetString("api.server.address"), viper.GetInt("api.server.port"))

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return false
	}

	conn.Close()

	return true
}

func Shutdown() {
	if server != nil {
		log.Info("api: [shutdown] shutting down")
		if err := server.Shutdown(context.TODO()); err != nil {
			log.Panic("api: ", err)
		}

		server = nil
	} else {
		log.Info("api: [shutdown] server == nil in this process")
	}

	if ShutdownCh != nil {
		log.Info("api: [shutdown] signaling shutdown channel")
		ShutdownCh <- true
		close(ShutdownCh)
		ShutdownCh = nil
	} else {
		log.Info("api: [shutdown] ShutdownCh == nil in this process")
	}
}
