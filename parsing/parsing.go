package parsing

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	api "heid9/downtime/api/parsing"

	"golang.org/x/net/html"
)

const HOST = "advmusic.com"

type Parser struct{}

func NewParser() api.Parser {
	return &Parser{}
}

func (p *Parser) LoadFromUrl(url string) (io.Reader, error) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(data), nil
	}
	return nil, errors.New("HTTP :" + resp.Status)
}

func (p *Parser) LoadFromFile(path string) (io.Reader, error) {
	data, err := os.ReadFile(path)
	return bytes.NewReader(data), err
}

func (p *Parser) Match(r io.Reader) bool {
	node, err := html.Parse(r)
	if err != nil {
		return false
	}
	return dfs(node)
}

func findByKey(key string, attrs []html.Attribute) (string, bool) {
	for _, attr := range attrs {
		if key == attr.Key {
			return attr.Val, true
		}
	}
	return "", false
}

func dfs(node *html.Node) (res bool) {
	if node.Type == html.ElementNode && node.Data == "script" {
		src, ok := findByKey("src", node.Attr)
		if ok {
			return strings.Contains(src, HOST)
		}
	}
	for next := node.FirstChild; next != nil; next = next.NextSibling {
		if !res {
			res = dfs(next)
		} else {
			break
		}
	}
	return res
}
