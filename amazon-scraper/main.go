package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("https://www.amazon.com/s/ref=nb_sb_noss_1?url=search-alias=aps&field-keywords=%s&ref=nb_sb_noss_1&crid=1FQ6XJD3UL1G6&sprefix=%s,aps,101",
			escape("iron pills")...),
		nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Cookie", "skin=noskin; session-id=142-0488354-2343223; session-id-time=2082787201l; i18n-prefs=USD; ubid-main=132-9538144-3343145; session-token=6aOpVLbjT5BTM190MUjKDUB5PuhW5whzadCSzCcrE8c8tDtsnFfVqj5MGLTEgHYl7sz8x5mbQhpLmHxAl0w+ck1U0h6If2f1gwBs9Sl+ZyTguNtJUHuOyE7e8J6C5DhejvOrmXBpxwZQJ7OCCSNtbpAEGEXJI5KZYz/MLxlzLxfl31aGOJSGKKCLxAfNtku0; csm-hit=tb:s-BJYAFVYHKF7FSR3HCQSH|1650943751193&t:1650943751773&adb:adblk_no")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	f, _ := os.CreateTemp(".", "iron.html")
	io.Copy(f, resp.Body)
	defer f.Close()
}

func escape(s string) []any {
	return []any{url.QueryEscape(s), url.QueryEscape(s)}
}
