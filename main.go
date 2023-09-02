package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func fetchStatus(number string) (status string, err error) {
	data := url.Values{}
	data.Set("number01", number)

	resp, err := http.PostForm("https://toi.kuronekoyamato.co.jp/cgi-bin/tneko", data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	var f func(*html.Node) bool
	f = func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "class" && strings.Contains(a.Val, "js-tracking-detail") {
					status = extractText(n)
					return true
				}
			}
		}
	
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if f(c) {
				return true
			}
		}
		return false
	}

	f(doc)
	return status, nil
}

func main() {
	inputData, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	trackingNumbers := strings.Split(string(inputData), "\n")
	outputData := ""

	for _, number := range trackingNumbers {
		if number == "" {
			outputData += "\n"
			continue
		}
		status, err := fetchStatus(number)
		if err != nil {
			log.Fatal(err)
		}
		outputData += fmt.Sprintf("%s\t%s\n", number, strings.TrimSpace(status))
	}

	err = os.WriteFile("output.txt", []byte(outputData), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func extractText(n *html.Node) string {
	text := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text += c.Data
		}
	}
	return text
}
