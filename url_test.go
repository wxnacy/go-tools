package tools

import "testing"

func TestURLParse(t *testing.T) {
	url := "https://wxnacy.com/name=wxnacy&token=%2Fwxnacy%2F/index.html"
	u, err := URLParse(url)
	if err != nil {
		t.Errorf("URLParse err: %v", err)
	}
	if u.Homepage != "https://wxnacy.com" {
		t.Errorf("%s != https://wxnacy.com", u.Homepage)
	}
	if u.FullName != "index.html" {
		t.Errorf("%s != index.html", u.FullName)
	}
	if u.Name != "index" {
		t.Errorf("%s != index", u.Name)
	}
	if u.Dir != "https://wxnacy.com/name=wxnacy&token=%2Fwxnacy%2F" {
		t.Errorf("%s  != https://wxnacy.com/name=wxnacy&token=%%2Fwxnacy%%2F", u.Dir)
	}

	u, err = URLParse("http://localhost:8005/m3u8/23293/23293.m3u8")
	if err != nil {
		t.Errorf("URLParse err: %v", err)
	}
	if u.FullName != "23293.m3u8" {
		t.Errorf("%s != 23293.m3u8", u.FullName)
	}
	if u.Name != "23293" {
		t.Errorf("%s != 23293", u.Name)
	}
}
