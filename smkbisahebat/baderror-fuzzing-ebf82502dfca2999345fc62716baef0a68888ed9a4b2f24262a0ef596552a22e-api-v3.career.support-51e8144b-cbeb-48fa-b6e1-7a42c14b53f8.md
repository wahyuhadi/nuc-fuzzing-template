### Improper Error Handling (baderror-fuzzing-ebf82502dfca2999345fc62716baef0a68888ed9a4b2f24262a0ef596552a22e) found on api-v3.career.support

----
**Details**: **baderror-fuzzing-ebf82502dfca2999345fc62716baef0a68888ed9a4b2f24262a0ef596552a22e** matched at api-v3.career.support

**Protocol**: HTTP

**Full URL**: https://api-v3.career.support/school/api/v1/candidate/get/chart?student_number&name&start_year&end_year=1'

**Timestamp**: Sun Jun 15 21:35:13 +0800 WITA 2025

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
GET /school/api/v1/candidate/get/chart?student_number&name&start_year&end_year=1' HTTP/1.1
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
Content-Length: 497
Access-Control-Allow-Origin: *
Alt-Svc: h3=":443"; ma=86400
Cache-Control: no-cache, private
Cf-Cache-Status: DYNAMIC
Cf-Ray: 95027308dd7cfe8a-SIN
Content-Type: application/json
Date: Sun, 15 Jun 2025 13:35:13 GMT
Nel: {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}
Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/report/v4?s=7hbwCDolY7x7L9aLElKznZL9hT91fM6QD67u693cgiY9i0HdH8PpCYb9AVLUVt5ssvyE1N5Piafm1mKCockM1lC%2Fz0QndYRP8kRC5mdAUigdjLL%2BdxaL9XV6LkUESJcelA%3D%3D"}]}
Server: cloudflare
X-Ratelimit-Limit: 60
X-Ratelimit-Remaining: 59

{"data":null,"message":"SQLSTATE[22P02]: Invalid text representation: 7 ERROR:  invalid input syntax for type integer: \\"1'\\"\\nCONTEXT:  unnamed portal parameter $10 = '...' (Connection: pgsql, SQL: select start_year, count(start_year) as total from \\"candidate_educations\\" where \\"school_id\\" = 23 and \\"is_verified\\" = 1 and \\"candidate_id\\" in (361, 381, 377, 378, 379, 362, 380) and \\"end_year\\" = 1' and \\"candidate_educations\\".\\"deleted_at\\" is null group by \\"start_year\\")","status":500}
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
curl -X 'GET' -d '' -H 'Accept: application/json, text/plain, */*' -H 'Accept-Encoding: gzip, deflate, br, zstd' -H 'Accept-Language: en-US,en;q=0.5' -H 'Authorization: Bearer 978|dT0wUvhvUDskmmrYYpjhc5H2C3Io3vCCbFc1furo' -H 'Connection: keep-alive' -H 'Host: api-v3.career.support' -H 'Origin: https://school-v3.career.support' -H 'Priority: u=0' -H 'Referer: https://school-v3.career.support/' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-site' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0' 'https://api-v3.career.support/school/api/v1/candidate/get/chart?student_number&name&start_year&end_year=1'\\'''
```


----

Eagle Pulsar Team :  > From Home With Love