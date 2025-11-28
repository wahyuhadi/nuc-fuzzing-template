### Improper Error Handling (baderror-fuzzing-c46c12a98c83fbabd76d82d3fe1c9fce995a6ecf042bb69ba00bc0c099363d60) found on api-v3.career.support

----
**Details**: **baderror-fuzzing-c46c12a98c83fbabd76d82d3fe1c9fce995a6ecf042bb69ba00bc0c099363d60** matched at api-v3.career.support

**Protocol**: HTTP

**Full URL**: https://api-v3.career.support/school/api/v1/candidate/get/verified?page=1&student_number&name&start_year&end_year=1'

**Timestamp**: Sun Jun 15 21:35:13 +0800 WITA 2025

**Template Information**

| Key | Value |
| --- | --- |
| Name | Improper Error Handling |
| Authors | wahyuhadi |
| Tags | error-handling, info-leak, dast |
| Severity | medium |
| Description | Improper Error Handling occurs when an application discloses internal information through unhandled or verbose error messages. <br>These messages may expose sensitive details such as database structure, file paths, server configuration, or even full stack traces, <br>which can be leveraged by attackers to further compromise the system.<br> |
| cvss-3.1-link | https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:N/A:N |
| tech | http |
| severity | medium (5.3) |
| cvss-3.1 | CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:N/A:N |

**Request**
```http
GET /school/api/v1/candidate/get/verified?page=1&student_number&name&start_year&end_year=1' HTTP/1.1
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
Content-Length: 1119
Access-Control-Allow-Origin: *
Alt-Svc: h3=":443"; ma=86400
Cache-Control: no-cache, private
Cf-Cache-Status: DYNAMIC
Cf-Ray: 95027308c8f6401c-SIN
Content-Type: application/json
Date: Sun, 15 Jun 2025 13:35:13 GMT
Nel: {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}
Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/report/v4?s=nAr10eP%2B%2FlRKbbn5Y%2FdsePjpCd3ri0%2B5D5Ejt0LPcgf4bvL3tKqaa0Pm0UlMtl434BVm2PUojLyQs4NGN2wEUbLbZITZgQtkaEtUE%2BSNCZPm5EmnMC1oCktCkG0zVXL2dA%3D%3D"}]}
Server: cloudflare
X-Ratelimit-Limit: 60
X-Ratelimit-Remaining: 58

{"data":null,"message":"SQLSTATE[22P02]: Invalid text representation: 7 ERROR:  invalid input syntax for type integer: \\"1'\\"\\nCONTEXT:  unnamed portal parameter $2 = '...' (Connection: pgsql, SQL: select count(*) as aggregate from \\"candidates\\" left join (select max(detail_mx.id) as detail_id, detail_mx.candidate_id from \\"candidate_statuses\\" as \\"detail_mx\\" group by \\"detail_mx\\".\\"candidate_id\\") as \\"detail_max\\" on \\"candidates\\".\\"id\\" = \\"detail_max\\".\\"candidate_id\\" left join (select detail_data.id, detail_data.status from \\"candidate_statuses\\" as \\"detail_data\\") as \\"detail\\" on \\"detail\\".\\"id\\" = \\"detail_max\\".\\"detail_id\\" where detail.status = 3 and exists (select * from \\"school_candidates\\" where \\"candidates\\".\\"id\\" = \\"school_candidates\\".\\"candidate_id\\" and \\"school_id\\" = 23 and \\"school_candidates\\".\\"deleted_at\\" is null) and exists (select * from \\"candidate_educations\\" where \\"candidates\\".\\"id\\" = \\"candidate_educations\\".\\"candidate_id\\" and \\"end_year\\" = 1' and \\"candidate_educations\\".\\"deleted_at\\" is null) and \\"candidates\\".\\"deleted_at\\" is null)","status":500}
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
curl -X 'GET' -d '' -H 'Accept: application/json, text/plain, */*' -H 'Accept-Encoding: gzip, deflate, br, zstd' -H 'Accept-Language: en-US,en;q=0.5' -H 'Authorization: Bearer 978|dT0wUvhvUDskmmrYYpjhc5H2C3Io3vCCbFc1furo' -H 'Connection: keep-alive' -H 'Host: api-v3.career.support' -H 'Origin: https://school-v3.career.support' -H 'Priority: u=0' -H 'Referer: https://school-v3.career.support/' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-site' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0' 'https://api-v3.career.support/school/api/v1/candidate/get/verified?page=1&student_number&name&start_year&end_year=1'\\'''
```


----

Eagle Pulsar Team :  > From Home With Love