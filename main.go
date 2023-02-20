package main

import (
	"context"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

type Config struct {
	AccessToken string `json:"access_token"`
}

func main() {
	router := gin.Default()

	// Load the access token from the config file
	configData, err := os.ReadFile("config.json")
	if err != nil {
		panic("Error loading config file")
	}
	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		panic("Error loading config file")
	}

	// Set up the GitHub API client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.AccessToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	router.GET("/", func(c *gin.Context) {
		// Parse the page parameter
		pageStr := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}

		// Search for repositories using the GitHub API
		query := "kubernetes in:description stars:>100"
		opts := &github.SearchOptions{Sort: "stars", Order: "desc", ListOptions: github.ListOptions{Page: page, PerPage: 10}}
		result, _, err := client.Search.Repositories(ctx, query, opts)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error searching for repositories")
			return
		}

		// Convert the slice of pointers to a slice of values
		repos := make([]github.Repository, len(result.Repositories))
		for i, r := range result.Repositories {
			repos[i] = *r
		}

		// Render the HTML template with the search results
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Error rendering template")
			return
		}
		err = tmpl.Execute(c.Writer, struct {
			Repos []github.Repository
			Page  int
			Prev  int
			Next  int
		}{
			Repos: repos,
			Page:  page,
			Prev:  page - 1,
			Next:  page + 1,
		})
		if err != nil {
			c.String(http.StatusInternalServerError, "Error rendering template")
			return
		}
	})

	router.Run(":8080")
}
