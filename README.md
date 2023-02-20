# Best Kubernetes Repos (GitHub-API-Search)

This is a simple project built with Go (using Gin HTTP web framework) that demonstrates how to use the GitHub API to search for repositories that match certain criteria, and how to display the search results on a web page with pagination.

![Image screenshot](https://github.com/sv222/k8s-best-repos/raw/main/screenshot.png)

## Installation

1. Clone the repository:

```shell
git clone https://github.com/sv222/k8s-best-repos.git
```

2. Install the dependencies:

```shell
go mod tidy
```

3. Create a `config.json` file in the root directory with your GitHub access token in the following format:

```json
   {
       "access_token": "YOUR_ACCESS_TOKEN"
   }
```

Make sure your access token has the necessary permissions to search for repositories (see GitHub documentation for details).
4. Start the server: 

```shell
go run main.go
```

5. Open your web browser and go to <http://localhost:8080>

## Usage

In the current example the web page displays a table of the top 10 GitHub repositories that match the following criteria:

- The repository description includes the word "kubernetes"
- The repository has more than 100 stars
- The repositories are sorted in descending order of the number of stars.
- The table includes the repository name, the number of stars, and a link to the repository.

## Contributing

Feel free to contribute to this project by submitting pull requests or reporting issues.

## License

This project is licensed under the MIT License.