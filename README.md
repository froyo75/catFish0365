# catFish0365

<img src="images/catfish.png" width="300px">

***`catFish0365` is a simple program to automate device code phishing attacks.***

## Attack Flow

 * Step 1: The program generates a user code by requesting the [Azure AD devicecode endpoint](https://login.microsoftonline.com/common/oauth2/devicecode?api-version=1.0)
 * Step 2: The user code is sent to the victim (e.g. phishing email sent to the victim)
 * Step 3: The victim opens phishing email and inputs the user code by visiting the [devicelogin url](https://microsoft.com/devicelogin)
 * Step 4: the victim logs into his account
 * Step 5: The program polls the Azure AD for the authentication status until it gets his refresh and access tokens (successfull login) by requesting the [OAuth endpoint](https://login.microsoftonline.com/Common/oauth2/token?api-version=1.0)

***[By default, the access_tokens are valid for 60 days and refresh_tokens are valid for a year](https://learn.microsoft.com/en-us/linkedin/shared/authentication/programmatic-refresh-tokens)***

## Usage
```shell
Usage: catFish0365 <command>

A tool to automate device code phishing attacks by @froyo75

Flags:
  -h, --help                                               Show context-sensitive help.
  -c, --clientid="d3590ed6-52b3-4102-aeff-aad2292ab01c"    The Application (client) ID of an Azure App to target (default: d3590ed6-52b3-4102-aeff-aad2292ab01c).
  -r, --resource="https://graph.windows.net"               The resource to request access tokens.
  -o, --out-path=STRING                                    Logging output to a specific file.
  -v, --version                                            Print version information and quit.

Commands:
  exploitdcauth
    Launch device code phishing attack.

  getaccesstoken <refreshtoken>
    Get a new access token for the given clientid and resource using the given refresh token.

Run "catFish0365 <command> --help" for more information on a command.

catFish0365: error: expected one of "exploitdcauth",  "getaccesstoken"
```

## References

Microsoft OAuth Device Code Phishing:
 * [Microsoft 365 OAuth Device Code Flow and Phishing](https://www.optiv.com/insights/source-zero/blog/microsoft-365-oauth-device-code-flow-and-phishing)
 * [The Art of the Device Code Phish](https://0xboku.com/2021/07/12/ArtOfDeviceCodePhish.html)

Framework to abuse Azure JSON Web Token (JWT):
 * [TokenTactics](https://github.com/rvrsh3ll/TokenTactics)
 * [AADInternals](https://github.com/Gerenios/AADInternals)

Mitigation & Detection methodology:
 * [Mitigation of Microsoft OAuth Device Code Phishing](https://www.optiv.com/insights/source-zero/blog/microsoft-365-oauth-device-code-flow-and-phishing)
 * [Detection Methodology](https://www.inversecos.com/2022/12/how-to-detect-malicious-oauth-device.html)
