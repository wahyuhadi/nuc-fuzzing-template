### Improper Error Handling (baderror-fuzzing-899407cae926abda7aacbc1c7c612bc2000550dab06bfbbdea4a23c124235825) found on api-v3.career.support

----
**Details**: **baderror-fuzzing-899407cae926abda7aacbc1c7c612bc2000550dab06bfbbdea4a23c124235825** matched at api-v3.career.support

**Protocol**: HTTP

**Full URL**: https://api-v3.career.support/school/api/v1/tracer-study/get?page=1&name=print(')

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
GET /school/api/v1/tracer-study/get?page=1&name=print(') HTTP/1.1
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
Content-Length: 951
Access-Control-Allow-Origin: *
Alt-Svc: h3=":443"; ma=86400
Cache-Control: no-cache, private
Cf-Cache-Status: DYNAMIC
Cf-Ray: 95027308bca5ce02-SIN
Content-Type: application/json
Date: Sun, 15 Jun 2025 13:35:12 GMT
Nel: {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}
Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/report/v4?s=CK4YXb5WjozHELFnvOBt7rOEi4KxVudxoFFbiUhAPEqHlbovry0G8n9YPWzoWbgzFreSX%2Bcsh4fgVdda5Ci5pDWwrjvW2pp%2FOLZAkWhf9HqIeUAfuw%3D%3D"}]}
Server: cloudflare
X-Ratelimit-Limit: 60
X-Ratelimit-Remaining: 59

{"data":null,"message":"SQLSTATE[42601]: Syntax error: 7 ERROR:  unterminated quoted string at or near \\"') and \\"tracer_studies\\".\\"deleted_at\\" is null\\"\\nLINE 1: ...3 and lower(tracer_studies.name) ~ lower('print(')') and \\"tr...\\n                                                             ^ (Connection: pgsql, SQL: select count(*) as aggregate from \\"tracer_studies\\" left join (select max(detail_mx.id) as detail_id, detail_mx.tracer_study_id from \\"tracer_studies_histories\\" as \\"detail_mx\\" group by \\"detail_mx\\".\\"tracer_study_id\\") as \\"detail_max\\" on \\"tracer_studies\\".\\"id\\" = \\"detail_max\\".\\"tracer_study_id\\" left join (select detail_data.id, detail_data.status from \\"tracer_studies_histories\\" as \\"detail_data\\") as \\"detail\\" on \\"detail\\".\\"id\\" = \\"detail_max\\".\\"detail_id\\" where tracer_studies.school_id = 23 and lower(tracer_studies.name) ~ lower('print(')') and \\"tracer_studies\\".\\"deleted_at\\" is null)","status":500}
```
**Metadata:**

- path: print(')

**Fixing**

- Always handle errors gracefully and return generic error messages to users.
- Disable detailed error output in production environments.
- Log errors internally without exposing them in HTTP responses.
- Avoid displaying stack traces, file paths, or database queries in user-facing messages.




**CURL command**
```sh
curl -X 'GET' -d '' -H 'Accept: application/json, text/plain, */*' -H 'Accept-Encoding: gzip, deflate, br, zstd' -H 'Accept-Language: en-US,en;q=0.5' -H 'Authorization: Bearer 978|dT0wUvhvUDskmmrYYpjhc5H2C3Io3vCCbFc1furo' -H 'Connection: keep-alive' -H 'Host: api-v3.career.support' -H 'Origin: https://school-v3.career.support' -H 'Priority: u=0' -H 'Referer: https://school-v3.career.support/' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-site' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0' 'https://api-v3.career.support/school/api/v1/tracer-study/get?page=1&name=print('\\'')'
```


----

Eagle Pulsar Team :  > From Home With Love