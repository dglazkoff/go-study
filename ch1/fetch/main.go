// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"strings"

	// "io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// Modify fetch to add the prefix http:// to each argument URL if it is missing. You might want to use strings.HasPrefix.
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(resp.Status)

		//b, err := ioutil.ReadAll(resp.Body)
		//resp.Body.Close()
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		//	os.Exit(1)
		//}
		//fmt.Printf("%s", b)

		// The function call io.Copy(dst, src) reads from src and writes to dst. Use it instead of ioutil.
		//ReadAll to copy the response body to os.Stdout without requiring a buffer large enough to hold the entire stream.
		//Be sure to check the error result of io.Copy.
		io.Copy(os.Stdout, resp.Body)
	}
}

//!-
