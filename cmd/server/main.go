package main

import (
	"net/http"
	"encoding/json"
	"log"
)


//Setup our simple router
func router() http.Handler {
	http.HandleFunc("/hello", Hello)
	return http.DefaultServeMux //this is a stdlib http.Handler
}

func Hello(rw http.ResponseWriter, req *http.Request)  {
	message := req.URL.Query().Get("message")
	encoder := json.NewEncoder(rw)
	body := map[string]string{"message":message}
	if err:= encoder.Encode(body); err != nil{
		http.Error(rw,"failed to encode json " + err.Error(),500)
		return
	}

}

func main()  {
	router := router()
	//start our server on port 8001
	if err := http.ListenAndServe(":8001", router); err != nil {
		log.Fatal(err)
	}

}
