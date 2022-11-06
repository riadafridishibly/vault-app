package imghandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	lru "github.com/hashicorp/golang-lru"
	"github.com/riadafridishibly/vault-app/pkg/utils"
)

var cache = newCache()

func newCache() *lru.Cache {
	cache, err := lru.New(100)
	if err != nil {
		panic(err)
	}
	return cache
}

func loadImage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "imgID")
	value, ok := cache.Get(id)
	if !ok {
		log.Println("Not found!")
		return
	}

	buff, ok := value.(*bytes.Buffer)
	if !ok {
		log.Println("Not found!")
		return
	}

	io.Copy(w, buff)
}

// TODO: better error handle
func uploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">>> Upload handler")
	fmt.Println(r.URL.Path)
	fmt.Println(r.Host)
	fmt.Println(r.URL.Port())
	r.ParseMultipartForm(10 << 20)
	// v, _ := httputil.DumpRequest(r, true)
	// fmt.Println(string(v))
	file, _, err := r.FormFile("image")
	if err != nil {
		panic(err)
	}
	id := utils.RandomID()
	buf := new(bytes.Buffer)
	io.Copy(buf, file)
	cache.Add(id, buf)
	json.NewEncoder(w).Encode(map[string]string{
		"file": id,
	})
}

func NewImageHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.Recoverer)
	r.Post("/upload", uploadImage)
	r.Get("/{imgID}", loadImage)
	return r
}

func NewServer() (srv *http.Server, l net.Listener, err error) {
	l, err = net.Listen("tcp", "localhost:0")
	if err != nil {
		return
	}

	srv = &http.Server{Handler: NewImageHandler()}
	return
}
