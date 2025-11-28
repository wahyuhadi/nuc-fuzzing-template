### Improper Error Handling (baderror-fuzzing-314866838572124b3e7d1ff2d9a7ffef0b3c1cd8f6e6e9630cabc41ee6d22adc) found on api-v3.career.support

----
**Details**: **baderror-fuzzing-314866838572124b3e7d1ff2d9a7ffef0b3c1cd8f6e6e9630cabc41ee6d22adc** matched at api-v3.career.support

**Protocol**: HTTP

**Full URL**: https://api-v3.career.support/school/api/v1/partnership/get/related?page=1&name&industry=../../~`

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
GET /school/api/v1/partnership/get/related?page=1&name&industry=../../~\` HTTP/1.1
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
Content-Length: 1529
Access-Control-Allow-Origin: *
Alt-Svc: h3=":443"; ma=86400
Cache-Control: no-cache, private
Cf-Cache-Status: DYNAMIC
Cf-Ray: 95027308bd1bfd67-SIN
Content-Type: application/json
Date: Sun, 15 Jun 2025 13:35:12 GMT
Nel: {"report_to":"cf-nel","success_fraction":0.0,"max_age":604800}
Report-To: {"group":"cf-nel","max_age":604800,"endpoints":[{"url":"https://a.nel.cloudflare.com/report/v4?s=r8tcAYWcsF2s1PRKFBjd%2FrwkuRMS%2F47YeU5%2F7zTmbQx8E1p192zQMB4tP8CDZL4yR80nghZUv23CKVd5YtCGdwmsD0%2BMmwZRfhhLwNNeb2qMllKL%2FEj%2BBtpGcRWQX4RmfQ%3D%3D"}]}
Server: cloudflare
X-Ratelimit-Limit: 60
X-Ratelimit-Remaining: 59

{"data":null,"message":"SQLSTATE[22P02]: Invalid text representation: 7 ERROR:  invalid input syntax for type bigint: \\"..\\/..\\/~\`\\"\\nCONTEXT:  unnamed portal parameter $1 = '...' (Connection: pgsql, SQL: select count(*) as aggregate from \\"company_school_partnerships\\" left join (select max(detail_mx.id) as detail_id, detail_mx.company_school_partnership_id from \\"company_school_partnership_histories\\" as \\"detail_mx\\" group by \\"detail_mx\\".\\"company_school_partnership_id\\") as \\"detail_max\\" on \\"company_school_partnerships\\".\\"id\\" = \\"detail_max\\".\\"company_school_partnership_id\\" left join (select detail_data.id, detail_data.status from \\"company_school_partnership_histories\\" as \\"detail_data\\") as \\"detail\\" on \\"detail\\".\\"id\\" = \\"detail_max\\".\\"detail_id\\" where company_school_partnerships.id in (select max(company_school_partnerships.id) from company_school_partnerships group by company_school_partnerships.company_id) and company_school_partnerships.id in (select max(company_school_partnerships.id) from company_school_partnerships group by company_school_partnerships.company_temporary_id) and company_school_partnerships.school_id = 23 and detail.status != 1 and company_school_partnerships.id in (1,2,3,4,9,10,14,15,5,6,7,8,11,12,13) and exists (select * from \\"companies\\" where \\"company_school_partnerships\\".\\"company_id\\" = \\"companies\\".\\"id\\" and \\"industry_id\\" = ..\\/..\\/~\` and \\"companies\\".\\"deleted_at\\" is null) and \\"company_school_partnerships\\".\\"deleted_at\\" is null)","status":500}
```
**Metadata:**

- path: ../../~`

**Fixing**

- Always handle errors gracefully and return generic error messages to users.
- Disable detailed error output in production environments.
- Log errors internally without exposing them in HTTP responses.
- Avoid displaying stack traces, file paths, or database queries in user-facing messages.




**CURL command**
```sh
curl -X 'GET' -d '' -H 'Accept: application/json, text/plain, */*' -H 'Accept-Encoding: gzip, deflate, br, zstd' -H 'Accept-Language: en-US,en;q=0.5' -H 'Authorization: Bearer 978|dT0wUvhvUDskmmrYYpjhc5H2C3Io3vCCbFc1furo' -H 'Connection: keep-alive' -H 'Host: api-v3.career.support' -H 'Origin: https://school-v3.career.support' -H 'Priority: u=0' -H 'Referer: https://school-v3.career.support/' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-site' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0' 'https://api-v3.career.support/school/api/v1/partnership/get/related?page=1&name&industry=../../~\`'
```


----

Eagle Pulsar Team :  > From Home With Love