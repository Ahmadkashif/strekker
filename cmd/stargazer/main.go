package main

import (
	"fmt"
	"net/http"

	httpserver "github.com/Ahmadkashif/stargazer/httpserver"
)

func main() {

	mux := http.NewServeMux()
	server := httpserver.NewServer(mux, ":8080")

	logText := `
	███████╗████████╗ █████╗ ██████╗  ██████╗  █████╗ ███████╗███████╗██████╗ 
	██╔════╝╚══██╔══╝██╔══██╗██╔══██╗██╔════╝ ██╔══██╗╚══███╔╝██╔════╝██╔══██╗
	███████╗   ██║   ███████║██████╔╝██║  ███╗███████║  ███╔╝ █████╗  ██████╔╝
	╚════██║   ██║   ██╔══██║██╔══██╗██║   ██║██╔══██║ ███╔╝  ██╔══╝  ██╔══██╗
	███████║   ██║   ██║  ██║██║  ██║╚██████╔╝██║  ██║███████╗███████╗██║  ██║
	╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝╚═╝  ╚═

	Server is running and will bring you stars you stare at for the rest of your life ...
`

	fmt.Println(logText)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintln(w, "GET request working")
		case http.MethodPost:
			fmt.Fprintln(w, "POST request")
		case http.MethodPut:
			fmt.Fprintln(w, "PUT request")
		case http.MethodDelete:
			fmt.Fprintln(w, "DELETE request")
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	server.ListenAndServe()
}
