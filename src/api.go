package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

var vtApiKey = os.Getenv("VIRUSTOTAL_API_KEY")
var vtUrl = "https://www.virustotal.com/vtapi/v2"

type IPResponse struct {
	Ip         string                 `json:"ip"`
	Virustotal map[string]interface{} `json:"virustotal"`
	IPInfo     map[string]interface{} `json:"ipinfo"`
	IPApi      map[string]interface{} `json:"ip-api"`
}

type DomainResponse struct {
	Domain     string                 `json:"domain"`
	Virustotal map[string]interface{} `json:"virustotal"`
	IPApi      map[string]interface{} `json:"ip-api"`
}

type FileHashResponse struct {
	FileHash   string                 `json:"file_hash"`
	Virustotal map[string]interface{} `json:"virustotal"`
}

func OutboundAPICalltoJson(url string, out *map[string]interface{}) *http.Response {
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
	return OutboundAPICalltoJson(fmt.Sprintf("%s/ip-address/report?apikey=%s&ip=%s", vtUrl, vtApiKey, ip), out)
}

// This function is simply used for readability purposes. It can be removed to cut lines of code at the cost of readability.
// To remove this line simply copy the everything after the return keyword and paste it whereever the function is called.
func IPInfoAPICall(ip string, out *map[string]interface{}) *http.Response {
	// Returning the response here is for review of the http.Response.StatusCode
	return OutboundAPICalltoJson(fmt.Sprintf("https://ipinfo.io/%s", ip), out)
}

// This function is simply used for readability purposes. It can be removed to cut lines of code at the cost of readability.
// To remove this line simply copy the everything after the return keyword and paste it whereever the function is called.
func VtDomainApiCall(domain string, out *map[string]interface{}) *http.Response {
	// Returning the response here is for review of the http.Response.StatusCode
	return OutboundAPICalltoJson(fmt.Sprintf("%s/domain/report?apikey=%s&domain=%s", vtUrl, vtApiKey, domain), out)
}

// This function is simply used for readability purposes. It can be removed to cut lines of code at the cost of readability.
// To remove this line simply copy the everything after the return keyword and paste it whereever the function is called.
func VtFileHashApiCall(hash string, out *map[string]interface{}) *http.Response {
	return OutboundAPICalltoJson(fmt.Sprintf("%s/file/report?apikey=%s&resource=%s", vtUrl, vtApiKey, hash), out)
}

// This function is simply used for readability purposes. It can be removed to cut lines of code at the cost of readability.
// To remove this line simply copy the everything after the return keyword and paste it whereever the function is called.
func IPApiApiCall(ipOrDomain string, out *map[string]interface{}) *http.Response {
	return OutboundAPICalltoJson(fmt.Sprintf("http://ip-api.com/json/%s?fields=66846719", ipOrDomain), out)
}

// Handles non 200 status codes from the virustotal api. Passing the necessary information to the end user.
func CheckVtStatusCode(code int, context *gin.Context) {
	switch code {
	case 204:
		context.JSON(http.StatusNoContent, gin.H{
			"message": "Our provider has limited our requests. Please try again later.",
		})
		break
	case 400:
		context.JSON(http.StatusNoContent, gin.H{
			"message": "Our provider was unable to process this request",
		})
		break
	default:
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Experienced an error trying to provide requested data",
		})
	}
}

func SearchIP(context *gin.Context) {
	if !validateIP(context.Params.ByName("ip")) {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid IP provided",
		})
		return
	}

	msg := IPResponse{
		Ip: context.Params.ByName("ip"),
	}

	if resp := VtIpApiCall(msg.Ip, &msg.Virustotal); resp.StatusCode != 200 {
		// CheckVtStatusCode will send a msg to the user forwarding the error encountered
		CheckVtStatusCode(resp.StatusCode, context)
		return
	}

	IPInfoAPICall(msg.Ip, &msg.IPInfo)
	IPApiApiCall(msg.Ip, &msg.IPApi)

	context.JSON(http.StatusOK, msg)
}

func SearchDomain(context *gin.Context) {
	if !govalidator.IsDNSName(context.Params.ByName("domain")) {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid domain provided",
		})
		return
	}
	msg := DomainResponse{
		Domain: context.Params.ByName("domain"),
	}

	if resp := VtDomainApiCall(msg.Domain, &msg.Virustotal); resp.StatusCode != 200 {
		// CheckVtStatusCode will send a msg to the user forwarding the error encountered
		CheckVtStatusCode(resp.StatusCode, context)
		return
	}

	IPApiApiCall(msg.Domain, &msg.IPApi)

	context.JSON(http.StatusOK, msg)
}

func SearchFileHash(context *gin.Context) {
	hash := context.Params.ByName("file_hash")

	if !govalidator.IsHash(hash, "md5") && !govalidator.IsHash(hash, "sha1") && !govalidator.IsHash(hash, "sha256") {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid file hash provided",
		})
		return
	}
	msg := FileHashResponse{
		FileHash: hash,
	}

	if resp := VtFileHashApiCall(msg.FileHash, &msg.Virustotal); resp.StatusCode != 200 {
		// CheckVtStatusCode will send a msg to the user forwarding the error encountered
		CheckVtStatusCode(resp.StatusCode, context)
		return
	}

	context.JSON(http.StatusOK, msg)
}
