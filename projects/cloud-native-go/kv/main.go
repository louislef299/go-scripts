package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	store = struct {
		sync.RWMutex
		m map[string]string
	}{m: make(map[string]string)}

	ErrorNoSuchKey = errors.New("no such key")
)

func Put(key, value string) error {
	store.Lock()
	store.m[key] = value
	store.Unlock()
	return nil
}

func Get(key string) (string, error) {
	store.RLock()
	value, ok := store.m[key]
	store.RUnlock()

	if !ok {
		return "", ErrorNoSuchKey
	}
	return value, nil
}

func Delete(key string) error {
	store.Lock()
	delete(store.m, key)
	store.Unlock()
	return nil
}

// Deletes the key and value from the store
// Returns a http status OK if successful
func keyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Gathers the value from the store and returns it to the writer
func keyValueGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := Get(key)
	if errors.Is(err, ErrorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(value))
}

// keyValuePutHandler expects to be called with a PUT request for
// the "/v1/key/{key}" resource
func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/{key}", keyValuePutHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", keyValueGetHandler).Methods("GET")
	r.HandleFunc("/v1/{key}", keyValueDeleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
