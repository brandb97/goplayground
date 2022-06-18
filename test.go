package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var dir string
var ch chan error = make(chan error)

func main() {
	if len(os.Args) >= 2 {
		dir = os.Args[1]
	}

	v := url.Values{}
	v.Add("i_psw", "physica")
	r, err := http.PostForm("http://eedevice.com/course.aspx", v)
	if err != nil {
		errExit(err)
	}

	doc, err := html.Parse(r.Body)
	if err != nil {
		errExit(err)
	}

	n := getAllHrefs(doc)

	for i := 0; i < n; i++ {
		err := <-ch
		if err != nil {
			fmt.Println(err)
		}
	}
}

func errExit(e error) {
	fmt.Println(e)
	os.Exit(1)
}

func getAllHrefs(n *html.Node) int {
	var nhrefs int

	if n.Type == html.ElementNode {
		var surl string
		var name string
		for _, a := range n.Attr {
			if a.Key == "href" {
				surl = a.Val
				if !strings.HasPrefix(surl, "http") {
					surl = "http://eedevice.com" + a.Val
				}
				name = filepath.Base(a.Val)
				go getHref(surl, name)
				nhrefs++
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nhrefs += getAllHrefs(c)
	}

	return nhrefs
}

func getHref(u, n string) {
	n = dir + n
	fmt.Println(u, n)

	f, err := os.Create(n)
	if err != nil {
		ch <- err
		return
	}
	defer f.Close()

	r, err := http.Get(u)
	if err != nil {
		os.Remove(n)

		ch <- err
		return
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ch <- err
		return
	}
	f.Write(b)

	ch <- nil
}
