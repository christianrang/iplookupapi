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
}

func SearchIP(context *gin.Context) {
	ipResp := IPResponse{
		Ip: context.Params.ByName("ip"),
	}

	url := fmt.Sprintf("%s/ip-address/report?apikey=%s&ip=%s", vtUrl, vtApiKey, ipResp.Ip)

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

	if err := json.Unmarshal(body, &ipResp.Virustotal); err != nil {
		log.Println(err)
	}

	out, err := json.Marshal(ipResp)

	context.JSON(http.StatusOK, gin.H{
		"message": string(out),
	})
}
