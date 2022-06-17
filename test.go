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

func main() {
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

	getAllHrefs(doc)
}

func errExit(e error) {
	fmt.Println(e)
	os.Exit(1)
}

func getAllHrefs(n *html.Node) {
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
				err := getHref(surl, name)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getAllHrefs(c)
	}
}

func getHref(u, n string) error {
	fmt.Println(u, n)

	f, err := os.Create(n)
	if err != nil {
		return err
	}
	defer f.Close()

	r, err := http.Get(u)
	if err != nil {
		os.Remove(n)
		return err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	f.Write(b)
	return nil
}
