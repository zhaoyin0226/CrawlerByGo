package Fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimit = time.Tick(time.Millisecond * 30)

func Fetch(targetUrl string) ([]byte, error) {

	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, targetUrl, nil)

	request.Header.Add("Referer", targetUrl)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	request.Header.Add("Host", "movie.douban.com")
	request.Header.Add("Cookie", "ll=\"118172\"; bid=xy9mDAWHMn0; __utmz=30149280.1617803880.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utmz=223695111.1617803880.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __yadk_uid=o7Y7pneJBaZ1W9KnAA8B7c7yjeyxZdv9; _vwo_uuid_v2=DDFC5A71F7696DFC413142AC5EC244747|16daf2795f100afedcf7178c15cdc3f2; __gads=ID=c98f76575eaf8314-22ec43fb5ac70021:T=1618536159:RT=1618536159:S=ALNI_MZH_c5EaAfzjcHNkgdCI30aFQiRVQ; _pk_ses.100001.4cf6=*; ap_v=0,6.0; __utma=30149280.1816436803.1617803880.1618536143.1618721958.3; __utmb=30149280.0.10.1618721958; __utmc=30149280; __utma=223695111.1007643682.1617803880.1618536143.1618721958.3; __utmb=223695111.0.10.1618721958; __utmc=223695111; ct=y; _pk_id.100001.4cf6=470e8b3d46ff9ff7.1617803880.3.1618722068.1618536717.")
	//BA_HECTOR := &http.Cookie{
	//	Name:   "BA_HECTOR",
	//	Value:  "0hag0hala12g8g241a1g7nnk40r",
	//	Path:   "/",
	//	Domain: ".baidu.com",
	//}
	//delPer := &http.Cookie{
	//	Name:   "delPer",
	//	Value:  "0",
	//	Path:   "/",
	//	Domain: ".baidu.com",
	//}
	//H_BDCLCKID_SF_BFESS := &http.Cookie{
	//	Name:   "H_BDCLCKID_SF_BFESS",
	//	Value:  "tb48oDDMtCP3jtJdhPTqq4_QK2cO2C62aJ3G2CbvWJ5TMCo1WfJ60qD8bGOp-tv7K2QL0tP5Jh7kShPC-tn8bto0QUnPbl3bJ2nlWbvd3l02V-OIe-t2yU_VeHrR5tRMW23i0h7mWpTTsxA45J7cM4IseboJLfT-0bc4KKJxbnLWeIJIjjC5D5QyjHtfJTnfb5kXLn58a-o5jbj4bDTjh6PLh-v9BtQmJJrf2Mn1ttoqeC-93P71-UIDbqJJ3tT9Qg-q3RA2yq7zMlTvXnjpqDufXR330x-jLIOOVn0MWhjDfpR_04nJyUPUbPnnBUcm3H8HL4nv2JcJbM5m3x6qLTKkQN3T-PKO5bRh_CFbtIK5MD06D5Rb5nbH5MoX2bje2I6H06rJaDkKjPbOy4oTj6j332FHLl3-5JnJKJndKhcVf-5_WqQC3MvB-fnn5MTnKa6dXPJHb4nJjR0RQft20M4IeMtjB53ab2JZbJ7jWhvvep72yMrTQlRX5q79atTMfNTJ-qcH0KQpsIJM5-DWbT8EjH62btt_JRCtVI5P",
	//	Path:   "/",
	//	Domain: ".baidu.com",
	//}
	//refreshToken := &http.Cookie{
	//	Name:  "refreshToken",
	//	Value: "1616397597.1618818060187.1a8adff306ba4143d0deb638fedfacb7",
	//	Path: "/",
	//	Domain: ".baidu.com",
	//}
	//login_health := &http.Cookie{
	//	Name:  "login_health",
	//	Value: "05cb159d646f940be6077d26941066f8531597b3dee397907512452d95227c00522ca91ceda182b871917ceee4de0580a680fd742e864c2a3584ec0f0f8a307d",
	//	Path: "/",
	//	Domain: ".baidu.com",
	//}
	//token := &http.Cookie{
	//	Name:  "token",
	//	Value: "1616397597.1618731660187.13e4adfb703ad8928bcf3dfd3f331fe6",
	//	Path: "/",
	//	Domain: ".baidu.com",
	//}
	//request.AddCookie(BA_HECTOR)
	//request.AddCookie(delPer)
	//request.AddCookie(H_BDCLCKID_SF_BFESS)
	//request.AddCookie(refreshToken)
	//request.AddCookie(login_health)
	//request.AddCookie(token)
	<-rateLimit
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("wrong status code:", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
