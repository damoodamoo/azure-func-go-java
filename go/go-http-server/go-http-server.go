package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	//"flag"
)

var shoppingList = NewShoppingList()

func add(w http.ResponseWriter, r *http.Request) {

	// where are we called from?
	invocationid := r.Header.Get("X-Azure-Functions-InvocationId")
	if invocationid != "" {
		fmt.Println("Called from Azure Function")
	}

	// get the json and unmarshall to a struct
	b, _ := ioutil.ReadAll(r.Body)
	var item ShoppingItem
	json.Unmarshal(b, &item)
	fmt.Println(fmt.Sprintf("Adding item: %v", item))

	// add to repo
	shoppingList.Add(item)

	// respond
	w.Write([]byte("Hello World from go worker!"))
}

func get(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	id := 0
	for k, v := range queryParams {
		if k == "id" {
			id, _ = strconv.Atoi(v[0])
		}
	}

	if id > 0 {
		item := shoppingList.Get(id)
		fmt.Println(fmt.Sprintf("Returning item: %v", item))
		js, _ := json.Marshal(item)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	} else {
		http.Error(w, "No item found", http.StatusNotFound)
	}
}

func list(w http.ResponseWriter, r *http.Request) {

	items := shoppingList.List()
	js, _ := json.Marshal(items)
	fmt.Println(fmt.Sprintf("Returning items: %v", items))
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func sendItems(w http.ResponseWriter, r *http.Request) {

	items := shoppingList.List()

	outputs := make(map[string]interface{})
	outputs["output1"] = items

	invokeResponse := InvokeResponse{outputs, []string{"Sent items - somewhere!"}, http.StatusCreated}

	js, _ := json.Marshal(invokeResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func processItems(w http.ResponseWriter, r *http.Request) {

	// item from queue comes in InvokeRequest model
	var invokeReq InvokeRequest
	d := json.NewDecoder(r.Body)
	_ = d.Decode(&invokeReq)

	// ... imagine we do something useful here ...

	// push... somewhere else
	outputs := make(map[string]interface{})
	outputs["output1"] = invokeReq.Data
	outputs["output2"] = invokeReq.Data
	invokeResponse := InvokeResponse{outputs, []string{"Processed + sent items"}, http.StatusCreated}

	js, _ := json.Marshal(invokeResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {

	// bootstrap the HTTP server + serve
	httpInvokerPort, exists := os.LookupEnv("FUNCTIONS_HTTPWORKER_PORT")
	if exists {
		fmt.Println("FUNCTIONS_HTTPWORKER_PORT: " + httpInvokerPort)
	}
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/add", add)
	mux.HandleFunc("/get", get)
	mux.HandleFunc("/list", list)

	mux.HandleFunc("/send-items", sendItems)
	mux.HandleFunc("/process-items", processItems)

	log.Println("Go server Listening...on httpInvokerPort:", httpInvokerPort)
	log.Fatal(http.ListenAndServe(":"+httpInvokerPort, mux))
}
