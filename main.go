package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Go Web App</title>
</head>
<body>
    <h1>Welcome to Go Web App</h1>
    <p>This is a simple Go web application!</p>
    <p>Current path: %s</p>
</body>
</html>
`, r.URL.Path)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>About - Go Web App</title>
</head>
<body>
    <h1>About</h1>
    <p>This is a simple Go web application built with the standard library.</p>
    <a href="/">Go back home</a>
</body>
</html>
`)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}