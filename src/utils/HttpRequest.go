package Utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"log"
)

func HttpRequest(url string) ([]byte,error) {
	res, err := http.Get(url)
	var body []byte
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		log.Println(err)
	} else {
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			log.Println(err)
		}
		return body,err
	}
	return body,err
}