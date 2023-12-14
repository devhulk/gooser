package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const file_uri = "https://raw.githubusercontent.com/WebBreacher/WhatsMyName/main/wmn-data.json"

func getSiteMap() (*os.File, error) {

	out, err := os.Create("wmn-data.json")
	if err != nil {
		return nil, err
	}

	defer out.Close()

	resp, err := http.Get(file_uri)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	_, err2 := io.Copy(out, resp.Body)

	if err2 != nil {
		return nil, err
	}

	return out, nil
}

func main() {

	flag.Func("username", "gooser <username> -- string representation of username.", func(v string) error {
		if v == "" {
			return errors.New("you didn't pass a value you dirty goose")
		}
		_, err := getSiteMap()
		if err != nil {
			log.Fatalln(err)
		}
		return nil
	})

	flag.Parse()

	fmt.Println(flag.Arg(0))

}
