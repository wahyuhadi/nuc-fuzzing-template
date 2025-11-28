### Improper Error Handling (baderror-fuzzing-938ceb626eda69150b5141d95641546aaab5a7a2d1fd7a3ed87f847fcb45a1dc) found on api-v3.career.support

----
**Details**: **baderror-fuzzing-938ceb626eda69150b5141d95641546aaab5a7a2d1fd7a3ed87f847fcb45a1dc** matched at api-v3.career.support

**Protocol**: HTTP

**Full URL**: https://api-v3.career.support/school/api/v1/landing/get/events?page=1&target='111'11'111111-'&subdomain=smkbisahebat

**Timestamp**: Sun Jun 15 21:35:12 +0800 WITA 2025

**Template Information**

| Key | Value |
| --- | --- |
| Name | Improper Error Handling |
| Authors | wahyuhadi |
| Tags | error-handling, info-leak, dast |
| Severity | medium |
| Description | Improper Error Handling occurs when an application discloses internal information through unhandled or verbose error messages. <br>These messages may expose sensitive details such as database structure, file paths, server configuration, or even full stack traces, <br>which can be leveraged by attackers to further compromise the system.<br> |
| severity | medium (5.3) |
| cvss-3.1 | CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:N/A:N |
| cvss-3.1-link | https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:N/A:N |
| tech | http |

**Request**
```http
GET /school/api/v1/landing/get/events?page=1&target='111'11'111111-'&subdomain=smkbisahebat HTTP/1.1
Host: api-v3.career.support
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0
Accept: application/json, text/plain, */*
Accept-Encoding: gzip, deflate, br, zstd
Accept-Language: en-US,en;q=0.5
Connection: keep-alive
Origin: https://smkbisahebat.career.support
Referer: https://smkbisahebat.career.support/
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: same-site


```

**Response**
```http
HTTP/1.1 500 Internal Server Error
Connection: close
Content-Length: 382
Access-Control-Allow-Origin: *
Alt-Svc: h3=":443"; ma=86400
Cache-Control: no-cache, private
Cf-Cache-Status: DYNAMIC
Cf-Ray: 95027308caeaf91c-SIN
Content-Type: application/json
Date: Sun, 15 Jun 2025 13:35:12 GMT
Nel: {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}
Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/report/v4?s=Fw2DrmtjPgzmln18GxUF9hnxX%2BIi3FG2rxhP%2FhJnWf71MTmnt42eJl4kEMzG71it38sCHRtkcd0hQUsjKO0y5du9fX6MoC3PGM1CuzizlUIhQw4lJOql%2BfbTdkoBgChR2A%3D%3D"}]}
Server: cloudflare
X-Ratelimit-Limit: 60
X-Ratelimit-Remaining: 59

{"data":null,"message":"SQLSTATE[22P02]: Invalid text representation: 7 ERROR:  invalid input syntax for type bigint: \\"'111'11'111111-'\\"\\nCONTEXT:  unnamed portal parameter $2 = '...' (Connection: pgsql, SQL: select count(*) as aggregate from \\"school_events\\" where \\"school_id\\" = 23 and \\"target\\" = '111'11'111111-' and \\"school_events\\".\\"deleted_at\\" is null)","status":500}
```
**Metadata:**

- path: '111'11'111111-'

**Fixing**

- Always handle errors gracefully and return generic error messages to users.
- Disable detailed error output in production environments.
- Log errors internally without exposing them in HTTP responses.
- Avoid displaying stack traces, file paths, or database queries in user-facing messages.




**CURL command**
```sh
curl -X 'GET' -d '' -H 'Accept: application/json, text/plain, */*' -H 'Accept-Encoding: gzip, deflate, br, zstd' -H 'Accept-Language: en-US,en;q=0.5' -H 'Connection: keep-alive' -H 'Host: api-v3.career.support' -H 'Origin: https://smkbisahebat.career.support' -H 'Referer: https://smkbisahebat.career.support/' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-site' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0' 'https://api-v3.career.support/school/api/v1/landing/get/events?page=1&target='\\''111'\\''11'\\''111111-'\\''&subdomain=smkbisahebat'
```


----

Eagle Pulsar Team :  > From Home With Love