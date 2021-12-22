package ch1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Fetch() {
	gurl := "localhost:8089/clob?url_parameter=13&url_parameter=22"
	if !strings.HasPrefix(gurl, "http://") {
		gurl = "http://" + gurl
	}
	resp, err := http.Get(gurl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", b)
	fmt.Printf("%s\n", resp.Status)
}