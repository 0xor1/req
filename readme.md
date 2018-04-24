req
===

A Simple cmdln tool to make http requests.

```
go install github.com/0xor1/req
#general
req METHOD URL [-c COOKIE_NAME COOKIE_VALUE] [-h HEADER_NAME HEADER_VALUE]
#example
req GET http://example.com -c myCookie 123abc -h Authorization myUsr&Pwd
```