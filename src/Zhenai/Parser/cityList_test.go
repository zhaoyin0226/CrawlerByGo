package Parser

import (
	"crawlerByGo/src/Fetcher"
	"testing"
)
const resultLenght = 470
func TestParseCityList(t *testing.T) {
	contenrs, err := Fetcher.Fetch("http://www.zhenai.com/zhenghun/aba")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contenrs)
	if len(result.Requests) != resultLenght {
		t.Errorf("result should have %d requests ; but got %d",resultLenght, len(result.Requests))
	}
	if len(result.Items) != resultLenght {
		t.Errorf("result should have %d Items ; but got %d",resultLenght, len(result.Items))
	}

}
