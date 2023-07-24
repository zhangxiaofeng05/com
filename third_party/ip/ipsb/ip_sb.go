// Package ipsb get ip information from https://ip.sb/
package ipsb

import (
	"context"
	"fmt"

	"github.com/zhangxiaofeng05/com/com_http"
)

type JsonIpRes struct {
	Ip string `json:"ip"`
}

// JsonIp https://api.ip.sb/jsonip
func JsonIp(ctx context.Context) (string, error) {
	url := fmt.Sprintf("%s/jsonip", IpSbApiUrl)
	var ip JsonIpRes
	// why: https://ip.sb/api/
	header := map[string]string{
		"user-agent": "Mozilla",
	}
	err := com_http.Get(ctx, url, header, &ip)
	if err != nil {
		return "", err
	}
	return ip.Ip, nil
}

type GeoIpRes struct {
	Organization    string  `json:"organization"`
	Longitude       float64 `json:"longitude"`
	City            string  `json:"city"`
	Timezone        string  `json:"timezone"`
	Isp             string  `json:"isp"`
	Offset          int64   `json:"offset"`
	Region          string  `json:"region"`
	Asn             int64   `json:"asn"`
	AsnOrganization string  `json:"asn_organization"`
	Country         string  `json:"country"`
	Ip              string  `json:"ip"`
	Latitude        float64 `json:"latitude"`
	ContinentCode   string  `json:"continent_code"`
	CountryCode     string  `json:"country_code"`
	RegionCode      string  `json:"region_code"`
}

// GeoIp https://api.ip.sb/geoip
func GeoIp(ctx context.Context) (*GeoIpRes, error) {
	url := fmt.Sprintf("%s/geoip", IpSbApiUrl)
	var geoIp GeoIpRes
	header := map[string]string{
		"user-agent": "Mozilla",
	}
	err := com_http.Get(ctx, url, header, &geoIp)
	if err != nil {
		return nil, err
	}
	return &geoIp, nil
}
