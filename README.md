# swagger-server

## Try it out

### If you have go installed

- go (v1.12+) should be installed on your machine

1. Run `go build` in the root directory  

2. Run the `swagger-server` executable in whichever manner is appropriate for your os. If you would like to specify a port other than 8080, append `-p $(port_number)` to your terminal command, e.g. `./swagger-server -p 8088`  

3. Hit http://localhost:8080 (or with custom port) in your browser

(you can also just run `go run main.go`)

Note: the executable should always find itself in the same directory as a `/docs` folder containing swagger yaml files in order to work. You don't need to worry about this if you are simply running the app from within the project root.

### If you'd rather use Docker

1. Run `docker-compose up --build`

2. Hit http://localhost:8080 (or with custom port) in your browser

** Add your own swagger yaml files to the `/docs` folder to have them served. You can do this while running the app - simply refresh the browser page to see your new file available in the dropdown at the top of the page.
