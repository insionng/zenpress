mux = http.serveMux()
mux.handle("/404", http.notFoundHandler())
mux.handleFunc("/", fn(w, req) {
    fprintln(w, "host:", req.host, "path:", req.URL)
})

mux.handleFunc("/hello", fn(w, req) {
    fprintln(w, "hello Typhoon!!!!!" )
})
