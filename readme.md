## To run:
- For windows install gnu make with `choco install make` then just run `make` at the root directory

### Things I learned
- Templating with templ including passing props and dynamic rendering
- Basic file structure for go projects (createing cmd, model, view, handler/controllers folders for organizing)
- How to use Makefile's and run them in windows 
- Create and use middleware's in go
- Reading from dontenv with error handling 
- when using air with `main.go` in subdirectory you need to change `cmd = "go build -o ./tmp/main.exe ."`
to `cmd = "go build -o ./tmp/main.exe cmd/main.go"` in the `.air.toml`


#### The packages I used were:
- Godotenv for reading env file `go get github.com/joho/godotenv`
- Echo for route handling `go get github.com/labstack/echo/v4`
- First I had to install Templ (for tamplating htmx files) `go install github.com/a-h/templ/cmd/templ@latest`  [more info on temple here](https://github.com/a-h/templ) then run `go get github.com/a-h/templ` to use in my project
- Air for hot reloading
- Htmx/Tailwind scripts in base layout 


#### Resources
- I chose to follow [this walkthough](https://www.youtube.com/watch?v=wttTTFVrQiw&list=WL&index=27) to get a core understanding of how to setup a golang project using temple (formatter) and htmx
- https://htmx.org/docs/
- https://go.dev/tour
 