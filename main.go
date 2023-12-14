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

func getSiteMap() (WhatsMyName, error) {

	out, err := os.Create(file_name)
	if err != nil {
		return WhatsMyName{}, err
	}

	defer out.Close()

	resp, err := http.Get(file_uri)
	if err != nil {
		return WhatsMyName{}, err
	}

	defer resp.Body.Close()

	_, err2 := io.Copy(out, resp.Body)
	if err2 != nil {
		return WhatsMyName{}, err
	}

	out.Close()

	v, err := os.Open(file_name)
	if err != nil {
		return WhatsMyName{}, err
	}
	defer v.Close()

	byteValue, _ := ioutil.ReadAll(v)

	var result WhatsMyName

	err3 := json.Unmarshal(byteValue, &result)
	if err3 != nil {
		log.Fatalln(err3)
		return WhatsMyName{}, err
	}

	return result, nil

}

func checkSites(u WhatsMyName) {

	for _, v := range u.Sites {
		fmt.Println(v.URICheck)
	}

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

		checkSites(result)

		return nil
	})

	flag.Parse()

	fmt.Println(flag.Arg(0))

}
