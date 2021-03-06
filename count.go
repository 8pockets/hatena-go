package hatena

import (
	"net/url"
)

const (
	countURL = "http://api.b.st-hatena.com/entry.count"
)

// はてなブックマーク件数取得API
func Count(urlStr string) (int, error) {
	return DefaultClient.Count(urlStr)
}

func (c *Client) Count(urlStr string) (int, error) {

	v := make(url.Values)
	v.Set("url", urlStr)
	req := countURL + "?" + v.Encode()

	co := new(int)
	err := c.get(req, co, "json")

	return *co, err
}
