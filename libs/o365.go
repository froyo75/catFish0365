package libs

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

const DefaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"
const DeviceCodeURL = "https://login.microsoftonline.com/common/oauth2/devicecode?api-version=1.0"
const OAuthURL = "https://login.microsoftonline.com/Common/oauth2/token?api-version=1.0"
const DeviceCodeGrantType = "urn:ietf:params:oauth:grant-type:device_code"

func HttpRequest(method string, url string, data url.Values) *http.Response {
	cli := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(data.Encode()))
	CheckErrors(err)

	req.Header.Set("User-Agent", DefaultUserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := cli.Do(req)
	CheckErrors(err)

	return resp
}

func GetuserCode(clientid string, resource string) map[string]interface{} {
	data := url.Values{
		"client_id": {clientid},
		"resource":  {resource},
	}

	resp := HttpRequest("POST", DeviceCodeURL, data)
	var respjson map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&respjson)

	return respjson
}

func GetAccessRefreshTokens(clientid string, resource string, respjson map[string]interface{}) map[string]interface{} {

	tokens := map[string]interface{}{
		"access_token":  "",
		"refresh_token": "",
	}

	var total int64 = 0
	devicecode := respjson["device_code"].(string)
	interval, err := strconv.ParseInt(respjson["interval"].(string), 10, 0)
	CheckErrors(err)
	expiration_time, err := strconv.ParseInt(respjson["expires_in"].(string), 10, 0)
	CheckErrors(err)

	data := url.Values{
		"client_id":  {clientid},
		"grant_type": {DeviceCodeGrantType},
		"code":       {devicecode},
		"resource":   {resource},
	}

	log.Debug().Msg("[*] Authorization Pending...")

	for total < expiration_time {
		resp := HttpRequest("POST", OAuthURL, data)
		var respjson map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&respjson)

		if resp.StatusCode == 200 {
			log.Debug().Msg("[*] Authorization Succeeded !")
			tokens["access_token"] = respjson["access_token"].(string)
			tokens["refresh_token"] = respjson["refresh_token"].(string)
			return tokens
		} else {
			if val, ok := respjson["error_description"]; ok {
				error_description := val.(string)
				if strings.Contains(error_description, "polling") {
					time.Sleep(time.Duration(interval) * time.Second)
					total += interval
				} else {
					log.Error().Msg("[!] Error: " + error_description)
					return nil
				}
			}
		}
	}
	log.Warn().Msg("[!] Timeout occurred.")
	return nil
}

func GetAccessTokenFromResfreshToken(clientid string, resource string, scope string, resfreshtoken string) map[string]interface{} {
	tokens := map[string]interface{}{
		"access_token":  "",
		"refresh_token": "",
	}

	data := url.Values{
		"client_id":     {clientid},
		"grant_type":    {"refresh_token"},
		"scope":         {scope},
		"resource":      {resource},
		"refresh_token": {resfreshtoken},
	}

	resp := HttpRequest("POST", OAuthURL, data)
	var respjson map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&respjson)
	if resp.StatusCode == 200 {
		tokens["access_token"] = respjson["access_token"].(string)
		tokens["refresh_token"] = respjson["refresh_token"].(string)
		log.Debug().Msg("[*] Successfully got a new access token !")
		return tokens
	} else {
		if val, ok := respjson["error_description"]; ok {
			error_description := val.(string)
			log.Error().Msg("[!] Error: " + error_description)
			return nil
		}
	}
	return nil
}
