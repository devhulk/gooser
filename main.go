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
	"strings"

	"github.com/zchee/color/v2"
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

func checkSites(w WhatsMyName, u string) []string {

	var hits []string

	for _, v := range w.Sites {

		uri := strings.Replace(v.URICheck, "{account}", u, 1)
		resp, err := http.Get(uri)
		if err != nil {
			log.Println(err)
			continue
		}

		if resp.StatusCode == v.ECode {
			color.Green("Hit: %v\n", uri)
			hits = append(hits, uri)

		} else if resp.StatusCode == v.MCode {
			color.Red("Nope: %v\n", uri)
			continue
		}
	}

	return hits

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

		hits := checkSites(result, v)
		fmt.Println(hits)

		return nil
	})

	flag.Parse()

	fmt.Println(flag.Arg(0))

}
