package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var vtApiKey = os.Getenv("VIRUSTOTAL_API_KEY")
var vtUrl = "https://www.virustotal.com/vtapi/v2"

type IPResponse struct {
	Ip         string                 `json:"ip"`
	Virustotal map[string]interface{} `json:"virustotal"`
	IPInfo     map[string]interface{} `json:"ipinfo"`
}

// type DomainResponse struct {
// 	Domain     string `json:"domain"`
// 	Virustotal string `json:virustotal`
// }

// type IPInfo struct {
// 	Ip       string `json:"ip"`
// 	Hostname string `json:"hostname"`
// 	City     string `json:"city"`
// 	Region   string `json:"region"`
// 	Country  string `json:"country"`
// 	Loc      string `json:"loc"`
// 	Org      string `json:"org"`
// 	Postal   string `json:"postal"`
// 	Timezone string `json:"timezone"`
// 	Readme   string `json:"readme"`
// }

func APICall(url string) (*http.Response, []byte) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	return res, body
}

func (ipResp *IPResponse) VTAPICall(url string) *http.Response {
	res, body := APICall(url)

	if err := json.Unmarshal(body, &ipResp.Virustotal); err != nil {
		log.Println(err)
	}

	return res
}

func (ipResp *IPResponse) IPInfoAPICall(url string) *http.Response {
	res, body := APICall(url)

	if err := json.Unmarshal(body, &ipResp.IPInfo); err != nil {
		log.Println(err)
	}

	return res
}

func SearchIP(context *gin.Context) {
	ipResp := IPResponse{
		Ip: context.Params.ByName("ip"),
	}

	ipResp.VTAPICall(fmt.Sprintf("%s/ip-address/report?apikey=%s&ip=%s", vtUrl, vtApiKey, ipResp.Ip))
	ipResp.IPInfoAPICall(fmt.Sprintf("%s%s", "https://ipinfo.io/", ipResp.Ip))

	context.JSON(http.StatusOK, ipResp)
}
