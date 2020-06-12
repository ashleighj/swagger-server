# swagger-server

## Try it out

### If you have go installed

- go (v1.12+) should be installed on your machine

1. Run `go build` in the root directory  

2. Run the `swagger-server` executable in whichever manner is appropriate for your os. If you would like to specify a port other than 8080, append `-p $(port_number)` to your terminal command, e.g. `./swagger-server -p 8088`  

3. Hit http://localhost:8080 (or with custom port) in your browser

### If you'd rather use Docker

1. Run `docker-compose up --build`

2. Hit http://localhost:8080 (or with custom port) in your browser
