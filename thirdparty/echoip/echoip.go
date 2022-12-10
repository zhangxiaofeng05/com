// Package echoip: https://github.com/mpolden/echoip
package echoip

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zhangxiaofeng05/com/comhttp"
)

type EchoIp struct {
	Ip         string            `json:"ip"`
	IpDecimal  int               `json:"ip_decimal"`
	Country    string            `json:"country"`
	CountryIso string            `json:"country_iso"`
	CountryEu  bool              `json:"country_eu"`
	RegionName string            `json:"region_name"`
	RegionCode string            `json:"region_code"`
	City       string            `json:"city"`
	Latitude   float64           `json:"latitude"`
	Longitude  float64           `json:"longitude"`
	TimeZone   string            `json:"time_zone"`
	Asn        string            `json:"asn"`
	AsnOrg     string            `json:"asn_org"`
	UserAgent  map[string]string `json:"user_agent"`
}

// IfConfigJson https://ifconfig.co/json
func IfConfigJson(ctx context.Context) (*EchoIp, error) {
	url := fmt.Sprintf("%s/json", IfConfigCoUrl)
	var res EchoIp
	err := comhttp.Get(ctx, url, comhttp.DefaultHeader, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// IfConfigIp https://ifconfig.co/ip
func IfConfigIp(ctx context.Context) (string, error) {
	url := fmt.Sprintf("%s/ip", IfConfigCoUrl)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBodyBytes), nil
}
