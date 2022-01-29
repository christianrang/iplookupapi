package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVtIpApiCall(t *testing.T) {
	ipResp := IPResponse{
		Ip: "8.8.8.8",
	}

	res := VtIpApiCall(ipResp.Ip, &ipResp.Virustotal)
	assert.Equal(t, 200, res.StatusCode)
}

func TestIPInfoApiCall(t *testing.T) {
	ipResp := IPResponse{
		Ip: "8.8.8.8",
	}

	res := IPInfoAPICall(ipResp.Ip, &ipResp.IPInfo)
	assert.Equal(t, 200, res.StatusCode)
}

func TesVtDomainApiCall(t *testing.T) {
	resp := DomainResponse{
		Domain: "google.com",
	}

	res := VtDomainApiCall(resp.Domain, &resp.Virustotal)
	assert.Equal(t, 200, res.StatusCode)
}
