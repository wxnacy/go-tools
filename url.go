package tools

import (
	"fmt"
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

type URL struct {
	*url.URL
	Homepage string
	Dir      string
	FullName string
	Name     string
}

func URLParse(rawURL string) (*URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	res := &URL{URL: u}
	res.Homepage = fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	res.FullName = filepath.Base(u.String())
	res.Name = strings.Replace(res.FullName, path.Ext(res.FullName), "", 1)
	dir := filepath.Dir(rawURL)
	dir = strings.Replace(dir, u.Scheme+":/", u.Scheme+"://", 1)
	res.Dir = dir

	return res, nil
}
