package main

import(
	"net/http"
	"fmt"
	"encoding/json"
)

func main() {
	const host = "localhost";
	const port = "8010";

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := "Hello from Golang"
		m := map[string]string{"message": message}
		enc := json.NewEncoder(w)
		enc.Encode(m)
    })

	fmt.Println("Server is running on http://"+ host + port)
	http.ListenAndServe(":" + port, nil)
	
}