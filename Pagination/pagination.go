package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// https://jsonplaceholder.typicode.com/

func main() {
	fmt.Println("Home Page")
	http.HandleFunc("/post", paginationFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Post struct {
	UserID int64  `json:"userId"`
	Id     int64  `json:"id"`
	Tittle string `json:"tittle"`
	Body   string `json:"body"`
}

func paginationFunc(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(data, &posts)
	total := int64(len(posts))
	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 0)
	per_page, err := strconv.ParseInt(r.URL.Query().Get("per_page"), 10, 0)
	perPage := validatePerPage(per_page)
	page = validatePage(page, perPage, int64(total))
	
	skipCount := validateSkipCount(page, perPage)
	
	apiResult := make([]Post, minInt(perPage, total-int64(skipCount)))
	var response []Post
	for i := int64(0); i < int64(len(apiResult)); i++ {
		response = append(response, posts[skipCount+i])
	}
	result, _ := json.Marshal(response)
	w.Write([]byte(result))
	fmt.Println(response)
}

func validatePerPage(perPage int64) int64 {
	if perPage < 1 {
		return 5
	} else if perPage > 10 {
		return 10
	}
	return perPage
}

func validatePage(page, perPage, totalCount int64) int64 {
	if perPage < 1 {
		return 1
	}
	
	pages := ((totalCount - 1) / perPage) + 1
	if page < 1 {
		return 1
	} else if page > pages {
		page = pages
	}
	return page
}

func validateSkipCount(page, perPage int64) int64 {
	skipCount := (page - 1) * perPage
	if skipCount < 0 {
		return 0
	}
	return skipCount
}

func minInt(a, b int64) int64 {
	if (a > b) {
		return b
	}
	return a
}
