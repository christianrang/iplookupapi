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

type DomainResponse struct {
	Domain     string                 `json:"domain"`
	Virustotal map[string]interface{} `json:virustotal`
}

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

func APICall(url string, out *map[string]interface{}) *http.Response {
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

	if err := json.Unmarshal(body, &out); err != nil {
		log.Println(err)
	}

	return res
}

// This function is simply used for readability purposes. It can be removed to cut lines of code at the cost of readability.
// To remove this line simply copy the everything after the return keyword and paste it whereever the function is called.
func VtIpApiCall(ip string, out *map[string]interface{}) *http.Response {
	return APICall(fmt.Sprintf("%s/ip-address/report?apikey=%s&ip=%s", vtUrl, vtApiKey, ip), out)
}

// This function is simply used for readability purposes. It can be removed to cut lines of code at the cost of readability.
// To remove this line simply copy the everything after the return keyword and paste it whereever the function is called.
func IPInfoAPICall(ip string, out *map[string]interface{}) *http.Response {
	// Returning the response here is for review of the http.Response.StatusCode
	return APICall(fmt.Sprintf("https://ipinfo.io/%s", ip), out)
}

func VtDomainApiCall(domain string, out *map[string]interface{}) *http.Response {
	// Returning the response here is for review of the http.Response.StatusCode
	// TODO: Consider only returning the http.Response.StatusCode if nothing else in http.Response if of use
	return APICall(fmt.Sprintf("%s/domain/report?apikey=%s&domain=%s", vtUrl, vtApiKey, domain), out)
}

func SearchIP(context *gin.Context) {
	ipResp := IPResponse{
		Ip: context.Params.ByName("ip"),
	}

	// TODO: Implement handling for http.Response.StatusCode
	VtIpApiCall(ipResp.Ip, &ipResp.Virustotal)
	IPInfoAPICall(ipResp.Ip, &ipResp.IPInfo)

	context.JSON(http.StatusOK, ipResp)
}

func SearchDomain(context *gin.Context) {
	dResp := DomainResponse{
		Domain: context.Params.ByName("domain"),
	}

	// TODO: Implement handling for http.Response.StatusCode
	VtDomainApiCall(dResp.Domain, &dResp.Virustotal)

	context.JSON(http.StatusOK, dResp)
}
