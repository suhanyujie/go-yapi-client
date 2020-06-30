package request

import (
	"testing"
)

func TestPost(t *testing.T) {
	url := "https://www.baidu.com"
	params := map[string]string{
		"a": "1",
	}
	body, err := Post(url, params, map[string]string{
		"content-type": "application/json",
	})
	if err != nil {
		t.Logf("%s", err)
	}
	t.Logf("%s", body)
}

func TestGet(t *testing.T) {
	url := "https://www.baidu.com"
	params := map[string]string{
		"wd": "ishenghuo",
	}
	body, err := Get(url, params)
	if err != nil {
		t.Logf("%s", err)
	}
	t.Logf("%s", body)
}
