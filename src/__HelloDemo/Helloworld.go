package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main(){

	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet,"http://www.zhenai.com/zhenghun",nil)
	request.Header.Add("Referer","https://www.zhenai.com/")
	request.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36")
	request.Header.Add("X-Requested-With","XMLHttpRequest")
	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		fmt.Println(" error status is --",resp.StatusCode)
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n",all)
	printCityListAll(all)
}

func printCityListAll(contents []byte) {

	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`)
	subMatches := compile.FindAllSubmatch(contents, -1)
	//fmt.Printf("subMatches is ----- %s\n",subMatches)
	for _, subMatche := range subMatches {
		//fmt.Printf("match is ----- %s\n",subMatche)
		for _, bytes := range subMatche {
			fmt.Printf(" byte is ---%s\n",bytes)
		}
	}
	fmt.Printf("Matches found : %d\n", len(subMatches))
}
