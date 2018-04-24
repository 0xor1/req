package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("general usage:\nreq METHOD URL [-c COOKIE_NAME COOKIE_VALUE] [-h HEADER_NAME HEADER_VALUE]\nexample usage:\nreq GET http://example.com -c myCookie 123abc -h Authorization myUsr&Pwd")
		os.Exit(0)
	}
	req, e := http.NewRequest(os.Args[1], os.Args[2], nil)
	for i := range os.Args {
		if os.Args[i] == "-h" {
			fmt.Println(os.Args[i+1], os.Args[i+2])
			req.Header.Set(os.Args[i+1], os.Args[i+2])
		} else if os.Args[i] == "-c" {
			req.AddCookie(&http.Cookie{Name: os.Args[i+1], Value: os.Args[i+2]})
		}
	}
	logAndExitIf(e)
	resp, e := http.DefaultClient.Do(req)
	logAndExitIf(e)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	body, e := ioutil.ReadAll(resp.Body)
	logAndExitIf(e)
	fmt.Println("Status: ", resp.StatusCode)
	fmt.Println("Headers: ", resp.Header)
	fmt.Println("Body: ", string(body))
}

func logAndExitIf(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}