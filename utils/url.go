package utils

import "net/url"

func IsValidUrl(urlStr string) bool {
	urlObj, err := url.ParseRequestURI(urlStr)
	return err == nil && (urlObj.Scheme == "http" || urlObj.Scheme == "https") && urlObj.Host != ""
}
