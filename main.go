package main 

import (
    "log"
    "net/http"
)

// Define a home handler function which writes a byte slice containing 
// "Hello from Snippetbox" as the response body. 

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Hello from Snippetbox\n"))
}

// Add a snippetView handler function. 
func snippetView(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a specific snippet...\n"))
}

// Add a snippetCreate handler function 
func snippetCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Write a new snippet\n"))
}

func main() {
    // Use the http.NewServeMux() function to initialize a new servemux, then
    // register the home function as the handler for the "/" URL pattern 
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate)
    /* 
     * Use the http.listenAndServe() function to star a new web server. We pass in 
     * two parameters: the TCP network address to listen on (in this case ":4000") 
     * and the servemux we just created. If http.ListenAndServe() returns an error
     * we ues the log.Fatal() function to log the error message and exit. Note 
     * that any error returned by http.ListenAndServe() is always not-nil. 
     */
    log.Print("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux) 
    log.Fatal(err)
}
