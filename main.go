package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

type URLMap map[string]*url.URL

type server struct {
	db *cached
}

type cached struct {
	sync.RWMutex
	value         URLMap
	updatedAt     time.Time
	ttl           time.Duration
	sheetProvider *sheet
}

func GetUrls(data [][]interface{}) URLMap {

	output := make(URLMap)

	for _, v := range data {

		key, ok := v[0].(string)

		if !ok {

			continue
		}

		key = strings.ToLower(key)

		value, ok := v[1].(string)

		if !ok {

			continue
		}

		url, err := url.Parse(value)

		if err != nil {

			continue
		}

		output[key] = url

	}

	return output

}

func main() {

	port := os.Getenv("PORT")

	sheetId := os.Getenv("SHEETID")

	sheetName := os.Getenv("SHEETNAME")

	ttlValue := os.Getenv("TTL")

	ttl := time.Second * 5

	if ttlValue != "" {

		v, err := time.ParseDuration(ttlValue)

		if err != nil {

			log.Fatalf("Error parsing time duration %v", err)
		}

		ttl = v

	}

	if port == "" {

		port = "8080"
	}

	srv := &server{
		db: &cached{
			ttl: ttl,
			sheetProvider: &sheet{
				Name: sheetName,
				Id:   sheetId,
			},
		},
	}

	http.HandleFunc("/", srv.Handler)

	log.Printf("Server listening on port: %s", port)

	http.ListenAndServe(":"+port, nil)

}

func (s *server) Handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {

		s.home(w)

		return

	}

	s.redirector(w, r)

}

func (s *server) home(w http.ResponseWriter) {

	w.WriteHeader(http.StatusNotFound)

	fmt.Fprintf(w, `<!DOCTYPE html>
<html><head><title>Not found</title></head><body><h1>Not found :(</h1>
	<p>This is home page for a URL redirector service.</p>
	<p>The URL is missing the shortcut in the path.</p>
	</body></html>`)

}

func (c *cached) refreshData() error {

	c.Lock()

	defer c.Unlock()

	if time.Since(c.updatedAt) <= c.ttl {

		return nil
	}

	rows, err := c.sheetProvider.Query()

	if err != nil {

		return err

	}

	c.value = GetUrls(rows)

	c.updatedAt = time.Now()

	return nil
}

func (c *cached) Get(query string) (*url.URL, error) {

	if err := c.refreshData(); err != nil {

		return nil, err
	}

	c.RLock()

	defer c.RUnlock()

	return c.value[query], nil

}

func (s *server) redirector(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()
	}

	path := strings.TrimPrefix(r.URL.Path, "/")

	url, err := s.db.Get(path)

	if err != nil {

		http.Error(w, "Error querying data", http.StatusInternalServerError)

		return
	}

	if url == nil {

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprintf(w, `404 Not Found`)

		return
	}

	fmt.Printf("Redirecting %s to %s\n", r.URL, url)

	http.Redirect(w, r, url.String(), http.StatusFound)

}
