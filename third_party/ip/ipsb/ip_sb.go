// Package ipsb get ip information from https://ip.sb/
package ipsb

import (
	"context"
	"net/http"

	"github.com/zhangxiaofeng05/com/com_http"
)

type JsonIpRes struct {
	Ip string `json:"ip"`
}

// JsonIp https://api.ip.sb/jsonip
func JsonIp(ctx context.Context, client *http.Client) (string, error) {
	var ip JsonIpRes
	// why: https://ip.sb/api/
	header := map[string]string{
		"user-agent": "Mozilla",
	}
	err := com_http.Get(ctx, client, IpSbApiUrlJsonIp, header, &ip)
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
func GeoIp(ctx context.Context, client *http.Client) (*GeoIpRes, error) {
	var geoIp GeoIpRes
	header := map[string]string{
		"user-agent": "Mozilla",
	}
	err := com_http.Get(ctx, client, IpSbApiUrlGeoIp, header, &geoIp)
	if err != nil {
		return nil, err
	}
	return &geoIp, nil
}
