package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Get() Response {
	resp, err := http.Get("https://bookmeter.com/users/1348458/books/reading.json")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := string(body)
	metadata := Response{}

	json.Unmarshal([]byte(data), &metadata)
	fmt.Println(metadata)
	return metadata
	// return BookAPIResponse{Books: }
	// return BookAPIResponse{Author: prettyData.Resources[0].Book.Author.Name, Title: prettyData.Resources[0].Book.Title, ImageURL: prettyData.Resources[0].Book.ImageURL}
}

type Response struct {
	Metadata struct {
		Sort   string `json:"sort"`
		Order  string `json:"order"`
		Offset int    `json:"offset"`
		Limit  int    `json:"limit"`
		Count  int    `json:"count"`
	} `json:"metadata"`
	Resources []struct {
		Path       string `json:"path"`
		ID         int    `json:"id"`
		Priority   any    `json:"priority"`
		CreatedAt  string `json:"created_at"`
		Page       any    `json:"page"`
		AuthorName any    `json:"author_name"`
		Book       struct {
			ID         int    `json:"id"`
			Path       string `json:"path"`
			AmazonUrls struct {
				Outline      string `json:"outline"`
				Registration string `json:"registration"`
				WishBook     string `json:"wish_book"`
			} `json:"amazon_urls"`
			Title             string `json:"title"`
			ImageURL          string `json:"image_url"`
			RegistrationCount int    `json:"registration_count"`
			Page              int    `json:"page"`
			Original          bool   `json:"original"`
			IsAdvertisable    bool   `json:"is_advertisable"`
			Author            struct {
				Name string `json:"name"`
				Path string `json:"path"`
			} `json:"author"`
			BookRegistrationStatusPath string `json:"book_registration_status_path"`
		} `json:"book"`
		User struct {
			ID    int    `json:"id"`
			Path  string `json:"path"`
			Name  string `json:"name"`
			Image string `json:"image"`
		} `json:"user"`
		Contents struct {
			Book struct {
				ID         int    `json:"id"`
				Path       string `json:"path"`
				AmazonUrls struct {
					Outline      string `json:"outline"`
					Registration string `json:"registration"`
					WishBook     string `json:"wish_book"`
				} `json:"amazon_urls"`
				Title             string `json:"title"`
				ImageURL          string `json:"image_url"`
				RegistrationCount int    `json:"registration_count"`
				Page              int    `json:"page"`
				Original          bool   `json:"original"`
				IsAdvertisable    bool   `json:"is_advertisable"`
				Author            struct {
					Name string `json:"name"`
					Path string `json:"path"`
				} `json:"author"`
				BookRegistrationStatusPath string `json:"book_registration_status_path"`
			} `json:"book"`
		} `json:"contents"`
	} `json:"resources"`
}
