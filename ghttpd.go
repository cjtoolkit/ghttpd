package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"

	"github.com/cjtoolkit/ghttpd/config"
	"github.com/pelletier/go-toml"
)

func fatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func parseUserConfig(c *config.Config) {
	file, err := os.Open("ghttpd.toml")
	if err != nil {
		return
	}

	tree, err := toml.LoadReader(file)
	fatal(err)

	// Why? toml counterpart does not preserve input from struct, like json does.
	b, err := json.Marshal(tree.ToMap())
	fatal(err)

	err = json.Unmarshal(b, c)
	fatal(err)

	for _, m := range c.Mimes {
		mime.AddExtensionType(m.Extension, m.Type)
	}
}

func main() {
	c := config.Config{
		Http: config.Http{
			Debug:     false,
			Address:   ":8000",
			CacheTime: 3600,
		},
		Tls:   nil,
		Mimes: nil,
	}

	wd, err := os.Getwd()
	fatal(err)

	parseUserConfig(&c)

	fs := http.FileServer(http.Dir(wd))

	fn := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", c.Http.CacheTime))
		if c.Http.Debug {
			log.Println("DEBUG:", req.URL.String())
		}
		fs.ServeHTTP(res, req)
	})

	if c.Tls != nil {
		log.Println("Running Server with TLS @", c.Http.Address)
		log.Println(http.ListenAndServeTLS(c.Http.Address, c.Tls.Cert, c.Tls.Key, fn))
		return
	}
	log.Println("Running Server @", c.Http.Address)
	log.Println(http.ListenAndServe(c.Http.Address, fn))
}
