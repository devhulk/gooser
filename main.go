package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const file_uri = "https://raw.githubusercontent.com/WebBreacher/WhatsMyName/main/wmn-data.json"
const file_name = "wmn-data.json"

func getSiteMap() (map[string]interface{}, error) {

	out, err := os.Create(file_name)
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

	out.Close()

	v, err := os.Open(file_name)
	if err != nil {
		return nil, err
	}
	defer v.Close()

	byteValue, _ := ioutil.ReadAll(v)

	var result map[string]interface{}

	err3 := json.Unmarshal([]byte(byteValue), &result)
	if err3 != nil {
		return nil, err
	}

	return result, nil

}

func main() {

	flag.Func("username", "gooser <username> -- string representation of username.", func(v string) error {
		if v == "" {
			return errors.New("you didn't pass a value you dirty goose")
		}
		result, err := getSiteMap()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(result["categories"])
		fmt.Println(result["sites"])

		return nil
	})

	flag.Parse()

	fmt.Println(flag.Arg(0))

}
