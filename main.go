package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("general usage:\nreq METHOD URL [-ba USERNAME PASSWORD] [-ta BEARER_TOKEN] [-c COOKIE_NAME COOKIE_VALUE] [-h HEADER_NAME HEADER_VALUE]\nexample usage:\nreq GET http://example.com -ba testuser testpwd -c myCookie 123abc -h headerName headerValue")
		os.Exit(0)
	}
	req, e := http.NewRequest(os.Args[1], os.Args[2], nil)
	printBody := true
	printHeader := true
	for i := range os.Args {
		if os.Args[i] == "-nb" {
			printBody = false
		} else if os.Args[i] == "-nh" {
			printHeader = false
		} else if os.Args[i] == "-ta" { // bearer token auth
			req.Header.Add("Authorization", "Bearer " + os.Args[i+1])
		} else if os.Args[i] == "-ba" { // basic auth
			req.SetBasicAuth(os.Args[i+1], os.Args[i+2])
		} else if os.Args[i] == "-h" { // add header
			req.Header.Add(os.Args[i+1], os.Args[i+2])
		} else if os.Args[i] == "-c" { // add cookie
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
	if printHeader {
		fmt.Println("Headers: ", resp.Header)
	}
	if printBody {
		fmt.Println("Body: ", string(body))
	}
}

func logAndExitIf(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}