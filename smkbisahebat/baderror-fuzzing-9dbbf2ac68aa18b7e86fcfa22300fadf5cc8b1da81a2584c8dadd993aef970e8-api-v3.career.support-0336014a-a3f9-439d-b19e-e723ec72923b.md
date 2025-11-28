### Improper Error Handling (baderror-fuzzing-9dbbf2ac68aa18b7e86fcfa22300fadf5cc8b1da81a2584c8dadd993aef970e8) found on api-v3.career.support

----
**Details**: **baderror-fuzzing-9dbbf2ac68aa18b7e86fcfa22300fadf5cc8b1da81a2584c8dadd993aef970e8** matched at api-v3.career.support

**Protocol**: HTTP

**Full URL**: https://api-v3.career.support/admin/api/v1/general/degrees?school=1'

**Timestamp**: Sun Jun 15 21:35:13 +0800 WITA 2025

**Template Information**

| Key | Value |
| --- | --- |
| Name | Improper Error Handling |
| Authors | wahyuhadi |
| Tags | error-handling, info-leak, dast |
| Severity | medium |
| Description | Improper Error Handling occurs when an application discloses internal information through unhandled or verbose error messages. <br>These messages may expose sensitive details such as database structure, file paths, server configuration, or even full stack traces, <br>which can be leveraged by attackers to further compromise the system.<br> |
| cvss-3.1 | CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:N/A:N |
| cvss-3.1-link | https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:N/A:N |
| tech | http |
| severity | medium (5.3) |

**Request**
```http
GET /admin/api/v1/general/degrees?school=1' HTTP/1.1
Host: api-v3.career.support
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0
Accept: application/json, text/plain, */*
Accept-Encoding: gzip, deflate, br, zstd
Accept-Language: en-US,en;q=0.5
Authorization: Bearer 978|dT0wUvhvUDskmmrYYpjhc5H2C3Io3vCCbFc1furo
Connection: keep-alive
Origin: https://school-v3.career.support
Priority: u=0
Referer: https://school-v3.career.support/
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: same-site


```

**Response**
```http
HTTP/1.1 500 Internal Server Error
Connection: close
Content-Length: 324
Access-Control-Allow-Origin: *
Alt-Svc: h3=":443"; ma=86400
Cache-Control: no-cache, private
Cf-Cache-Status: DYNAMIC
Cf-Ray: 95027308cf79a07d-SIN
Content-Type: application/json
Date: Sun, 15 Jun 2025 13:35:13 GMT
Nel: {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}
Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/report/v4?s=dKleXeHVgSxeu4EyYFtjqtFlpw0A%2BaipfXYqreeVvAA92j88U6XKoz8mmVW%2B8NohH8J0%2FWgenmHnACZ24rg5O8knUpUDZqciEfaxpS62bir4%2FD4kvoMG831Ht%2B8acrJZ2g%3D%3D"}]}
Server: cloudflare
X-Ratelimit-Limit: 60
X-Ratelimit-Remaining: 59

{"data":null,"message":"SQLSTATE[22P02]: Invalid text representation: 7 ERROR:  invalid input syntax for type bigint: \\"1'\\"\\nCONTEXT:  unnamed portal parameter $1 = '...' (Connection: pgsql, SQL: select * from \\"degrees\\" where \\"school_id\\" = 1' and \\"degrees\\".\\"deleted_at\\" is null order by \\"name\\" asc)","status":500}
```
**Metadata:**

- path: 1'

**Fixing**

- Always handle errors gracefully and return generic error messages to users.
- Disable detailed error output in production environments.
- Log errors internally without exposing them in HTTP responses.
- Avoid displaying stack traces, file paths, or database queries in user-facing messages.




**CURL command**
```sh
curl -X 'GET' -d '' -H 'Accept: application/json, text/plain, */*' -H 'Accept-Encoding: gzip, deflate, br, zstd' -H 'Accept-Language: en-US,en;q=0.5' -H 'Authorization: Bearer 978|dT0wUvhvUDskmmrYYpjhc5H2C3Io3vCCbFc1furo' -H 'Connection: keep-alive' -H 'Host: api-v3.career.support' -H 'Origin: https://school-v3.career.support' -H 'Priority: u=0' -H 'Referer: https://school-v3.career.support/' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-site' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0' 'https://api-v3.career.support/admin/api/v1/general/degrees?school=1'\\'''
```


----

Eagle Pulsar Team :  > From Home With Love