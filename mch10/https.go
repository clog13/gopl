package mch10

import (
	"net/http"
)

func HttpRe() {
	client := http.Client()
	req, err := http.NewRequest("GET", "https://www.toutiao.com/", nil)
	// Check.err(err)
	// req.Header.Set()
}
