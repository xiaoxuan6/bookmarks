package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"net/http"
	"text/template"
)

var (
	config Config
	//go:embed index.html
	indexFile embed.FS
	//go:embed static
	dirStatic embed.FS
	//go:embed data/bookmarks.json
	data []byte
)

type Config struct {
	Port        string `default:"8080"`
	Title       string `default:"xiaoxuan6、s Bookmarks"`
	Author      string `default:"xiaoxuan6"`
	Description string `default:"xiaoxuan6、s Bookmarks"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("error loading .env file: %w", err))
	}
	err = envconfig.Process("", &config)
	if err != nil {
		panic(fmt.Errorf("error loading config from env: %w", err))
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/favicon.ico", favicon)
	http.HandleFunc("/api/bookmarks", bookmarks)
	http.Handle("/static/", http.FileServer(http.FS(dirStatic)))
	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		panic(fmt.Errorf("error starting server: %w", err))
	}
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon.ico")
}

func index(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFS(indexFile, "index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("error parsing template"))
		return
	}

	err = tem.Execute(w, config)
	if err != nil {
		fmt.Println("error executing template: %w", err)
		_, _ = w.Write([]byte("error executing template"))
		return
	}
}

type (
	Item struct {
		Name string `json:"name"`
		URL  string `json:"url,omitempty"`
	}

	Bookmark struct {
		Item
		Children []Bookmark `json:"children,omitempty"`
	}
)

func bookmarks(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	var bookmark Bookmark
	err := json.Unmarshal(data, &bookmark)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(fail("error parsing bookmarks"))
		fmt.Println("error parsing bookmarks: %w", err)
		return
	}

	flattenData(bookmark.Children)
	_, _ = w.Write(success(d))
}

type Data struct {
	Item []Item
}

var d Data

func flattenData(bookmark []Bookmark) {
	for _, child := range bookmark {
		if len(child.Children) > 0 {
			flattenData(child.Children)
		} else {
			d.Item = append(d.Item, Item{
				Name: child.Name,
				URL:  child.URL,
			})
		}
	}
}

func success(data interface{}) []byte {
	b, _ := json.Marshal(struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}{
		200,
		data,
		"ok",
	})
	return b
}

func fail(msg string) []byte {
	b, _ := json.Marshal(struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}{
		500,
		"",
		msg,
	})
	return b
}
