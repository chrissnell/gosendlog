gosendlog
=======

gosendlog is a command-line tool to send data to a syslog server.  I wrote it because logger(1) that comes on most Linux distros is horribly buggy and has a 1000-character maximum packet size limit.  gosendlog is fast (50,000 syslog messages/sec on a single host is no problem) and supports larger syslog packets if your syslog server will.

Installation
------------
Install Go: http://golang.org/doc/install.

Then:
```
git clone git@github.com:chrissnell/gosendlog.git
cd gosendlog
go build gosendlog.go
cp gosendlog <wherever you want it>
```



Usage
-----
```
gosendlog <args>
    -dest="host:port"  - Destination syslog host
    -proto="protocol"  - Protocol (e.g. "udp" or "tcp").  Defaults to udp
    -tag="string"      - Tag or application name
    -facility="string" - Syslog facility (e.g. "kern", "local0", etc.)
    -priority="string" - Syslog priority (e.g. "crit", "info", etc.)
    -msg="string"      - Message to send.  For stdin, just exclude this flag.
```
 
Typical Examples
----------------

Logging from stdin: 
```
echo "Test message" | gosendlog -dest="syslog.mycompany.com:514" -proto=udp -tag="apache" -facility="local7" -priority="info"
```

Logging JSON-formatted access logs with Apache
----------------------------------------------
```
    LogFormat "{ \
            \"@timestamp\": \"%{%Y-%m-%dT%H:%M:%S%z}t\", \
            \"@version\": \"1\", \
            \"vip\": \"www.mycompany.com\", \
            \"message\": \"%h %l %u %t \\\"%r\\\" %>s %b\", \
            \"clientip\": \"%a\", \
            \"duration\": %D, \
            \"status\": %>s, \
            \"request\": \"%U%q\", \
            \"urlpath\": \"%U\", \
            \"urlquery\": \"%q\", \
            \"bytes\": %B, \
            \"method\": \"%m\", \
            \"referer\": \"%{Referer}i\", \
            \"useragent\": \"%{User-agent}i\" \
           }" access_log_json


    CustomLog "| gosendlog -dest=\"syslog-server.mycompany.com:514\" -tag=\"app-server-httpd\" -facility=\"local7\" -priority=\"info\" -proto=\"udp\"" access_log_json
```

Logging with command-line arguments:
```
gosendlog -dest="syslog.mycompany.com:514" -proto=udp -tag="apache" -facility="local7" -priority="info" -msg="Test message"
```
 
To Do
-----
* TLS support
* Ability to read/tail a file and deal with logrotate
* Enforce maximum syslog message size limits (with option to override)
* Better error handling
* This is my first Go language program ever so there are most certainly many ways to improve the code
