package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 解析url获得页面内容
func Fetch(url string) ([]byte, error) {

	retest, _ := http.NewRequest("GET", url, nil)
	retest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	//retest.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	//retest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	retest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	//retest.Header.Set("Cache-Control","max-age=0")
	//retest.Header.Set("Connection","keep-alive")
	retest.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	client := &http.Client{}
	response, err := client.Do(retest)

	if err != nil {
		return nil, err
	}

	//defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code : %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
