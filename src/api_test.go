package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVTAPICall(t *testing.T) {
	ipResp := IPResponse{
		Ip: "8.8.8.8",
	}

	res := ipResp.VTAPICall(fmt.Sprintf("%s/ip-address/report?apikey=%s&ip=%s", vtUrl, vtApiKey, ipResp.Ip))
	assert.Equal(t, 200, res.StatusCode)
}

func TestIPInfoApiCall(t *testing.T) {
	ipResp := IPResponse{
		Ip: "8.8.8.8",
	}

	res := ipResp.IPInfoAPICall(fmt.Sprintf("%s%s", "https://ipinfo.io/", ipResp.Ip))
	assert.Equal(t, 200, res.StatusCode)
}
