package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	portPtr := flag.String("p", "8080", "The web server port.")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", redirect)
	r.Get("/swagger", handleServeDoc)

	docsLoc := "./docs"
	fs := http.FileServer(http.Dir(docsLoc))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})

	log.Printf("Listening on :%s...", *portPtr)
	err := http.ListenAndServe(":"+*portPtr, r)
	if err != nil {
		log.Fatal(err)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/swagger", http.StatusPermanentRedirect)
}

type Model struct {
	Docs        []string
	SelectedDoc string
}

func handleServeDoc(w http.ResponseWriter, r *http.Request) {
	doc := r.URL.Query().Get("doc")

	files, err := ioutil.ReadDir("./docs")
	if err != nil {
		http.Error(w, "Error: could not find /docs folder",
			http.StatusInternalServerError)
		return
	}

	data := Model{}
	for _, file := range files {
		fileName := file.Name()
		if strings.Contains(fileName, "yml") ||
			strings.Contains(fileName, "yaml") {

			data.Docs = append(data.Docs,
				strings.Split(fileName, ".")[0])
		}
	}

	if len(data.Docs) == 0 {
		http.Error(w, "Error: no swagger yaml files to serve",
			http.StatusInternalServerError)
		return
	}

	data.SelectedDoc = data.Docs[0]
	if doc != "" {
		data.SelectedDoc = doc
	}

	tmpl, err := template.New("").Parse(tmpl)
	if err != nil {
		http.Error(w, "Error: couldn't find template to serve",
			http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("ERROR", err)
	}
}

const tmpl = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Swagger UI</title>
    <link
      rel="stylesheet"
      type="text/css"
      href="https://cdn.bootcdn.net/ajax/libs/swagger-ui/3.26.0/swagger-ui.css"
    />
    <link
      rel="icon"
      type="image/png"
      href="https://cdn.bootcdn.net/ajax/libs/swagger-ui/3.26.0/favicon-32x32.png"
      sizes="32x32"
    />
    <link
      rel="icon"
      type="image/png"
      href="https://cdn.bootcdn.net/ajax/libs/swagger-ui/3.26.0/favicon-16x16.png"
      sizes="16x16"
    />
    <style>
      html {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
      }
      *,
      *:before,
      *:after {
        box-sizing: inherit;
      }
      body {
        margin: 0;
        background: #fafafa;
      }
    </style>
  </head>

  <body>
    <!-- main container -->
    <div id="swagger-ui"></div>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.26.0/swagger-ui-bundle.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.26.0/swagger-ui-standalone-preset.js"></script>
    <script>
      window.onload = function() {
        onUrlSelected = function(e) {
          console.log(e.currentTarget.value);
          window.location = e.currentTarget.value
        }

        const NavPlugin = function() {
          return {
            wrapComponents: {
              // adds components above title
              InfoContainer: (Original, { React }) => (props) => {
                return React.createElement("div", {style: { marginTop: "20px"}},
                  React.createElement("select", {name: "selDocs", onChange: this.onUrlSelected},
                    {{ range $doc := .Docs }}
                        React.createElement(
                          "option", 
                          {
                            value: "/swagger?doc={{ $doc }}", 
                            name: {{ $doc }},
                            {{ if eq $doc $.SelectedDoc }}
                            selected: "selected"
                            {{ end }}
                          }, 
                          "{{ $doc }}"),
                    {{ end }}
                  ),
                  React.createElement(Original, props)
                )
              }
            }
          }
        }

        // Begin Swagger UI call region
		
		const ui = SwaggerUIBundle({
          url: "./{{ .SelectedDoc }}.yml", 
          dom_id: "#swagger-ui",
          deepLinking: true,
          presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
          plugins: [SwaggerUIBundle.plugins.DownloadUrl, NavPlugin],
          layout: "StandaloneLayout"
		});
		
        // console.log(ui)
        // End Swagger UI call region

        window.ui = ui;
      };
    </script>
  </body>
</html>`
