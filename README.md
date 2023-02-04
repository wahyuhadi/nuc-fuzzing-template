# Upstream proxy service to help create nuclei template for fuzzing attack

## Template design  in template.txt
---
- you can modify this template

```yaml
id: {{.Name}}
info:
  name: {{ .TempName }}
  author: Rahmat Wahyu Hadi
  severity: info

requests:
  - raw:
      - |
        {{ .Raw }}

    payloads:
      path: {{ .Payload }}
    matchers-condition: or
    matchers:
      - type: status
        part: header
        status:
          - 500
          - 503
      - type: regex
        regex:
          # MySQL
          - "SQL syntax.*?MySQL"
          - "Warning.*?\\Wmysqli?_"
          - "MySQLSyntaxErrorException"
          ...
```

## Default fuzzing payload in payload/default.txt
----

## Video 

[![Toturial](https://i9.ytimg.com/vi_webp/pJEwuUrUN2c/mq2.webp?sqp=CNiG-p4G-oaymwEmCMACELQB8quKqQMa8AEB-AH-CYAC0AWKAgwIABABGGUgZShlMA8=&rs=AOn4CLAitrzCASMwSj9MASrEMVsUxeDg4g)](https://www.youtube.com/watch?v=pJEwuUrUN2c)
