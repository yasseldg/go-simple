package nets

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/nets/sNet"
)

func Prueba() {
	os.Setenv("SERV_Own", "ApiManagement")

	ServiceOwn, err := sNet.NewService("SERV_Own", "")
	if err != nil {
		sLog.Error("sNet.NewService(): %s", err)
		return
	}
	ServiceOwn.Log()

	request := sNet.NewRequest().MethodPost().SetEndPoint("a")

	body := "aqqbPIycNs+ToAXQIX3TInl/ArkJcqbnhmZQpneqxL6iHg+nynYVzxusUuv5Z7qfRkSVmiIXeFSTRL6FcwDiyUJBMU1r3NxarRFjvMTlTGIsbpDwwjNyqYtMfUlkDX/iEmzpnZvxri0fJ2sipA8ouQ=="
	request.SetBody(strings.NewReader(body))

	data, err := request.Call(nil, ServiceOwn, nil)
	if err != nil {
		sLog.Error("request.Call(): %s", err)
		return
	}

	sLog.Info("data: %s", data)
}

// Handler para el servidor en el puerto 8080
func handler1(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/server1" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello from server 1!")
}

// Handler para el servidor en el puerto 8081
func handler2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/server2" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello from server 2!")
}

func Main() {
	// Servidor 1 en el puerto 8080
	server1 := &http.Server{
		Addr:         ":8080",
		Handler:      http.HandlerFunc(handler1),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Servidor 2 en el puerto 8081
	server2 := &http.Server{
		Addr:         ":8081",
		Handler:      http.HandlerFunc(handler2),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Levantar ambos servidores en gorutinas separadas
	go func() {
		if err := server1.ListenAndServe(); err != nil {
			fmt.Printf("Server 1 Error: %v\n", err)
		}
	}()

	go func() {
		if err := server2.ListenAndServe(); err != nil {
			fmt.Printf("Server 2 Error: %v\n", err)
		}
	}()

	// Esperar para evitar que el programa termine inmediatamente
	select {}
}
