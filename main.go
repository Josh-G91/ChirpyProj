package main

import (
        "net/http"
	"fmt"
)


func healthHandler(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
   mux := http.NewServeMux()
   mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("./app"))))
   mux.HandleFunc("/healthz", healthHandler)
   server := &http.Server{
       Addr: ":8080",
       Handler: mux,
   }

   err := server.ListenAndServe()

   if err != nil {
        //Handle error
   }
}
