package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func MakeRequest(URL string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("can't requet to url", err)
	}
	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)
	return string(resBody)
}
