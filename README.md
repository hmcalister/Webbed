# Webbed
## A template for developing webapps with Go using any frontend framework you like
## Importantly, the *web* frontend is em*bed*ded into the final binary for extreme portability!

#### Running the example

Simply `git clone` this template to get started. Run `go build .` to build the project for the first time, then run the executable to show the example frontend, served embedded from the Go binary. 

#### Developing from the template

For ease of development, install [air](https://github.com/cosmtrek/air) using `go install github.com/cosmtrek/air@latest`, then run `make dev` to start the development server. Open the npm exposed dev server for hot reloading of the frontend. The backend will be hot reloaded by air as changes are made to the Go files.