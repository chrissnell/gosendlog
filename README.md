gologit
=======

gologit is a command-line tool to send data to a syslog server.  I wrote it because logger(1) that comes with GNU's util-linux is horribly buggy.

Installation
------------
Install Go: http://golang.org/doc/install.

Then:
```
git clone git@github.com:chrissnell/gologit.git
cd gologit
go build gologit.go
cp gologit <wherever you want it>
```



Usage
-----
```
gologit <args>
    -dest="host:port"  - Destination syslog host
    -proto="protocol"  - Protocol (e.g. "udp" or "tcp").  Defaults to udp
    -tag="string"      - Tag or application name
    -facility="string" - Syslog facility (e.g. "kern", "local0", etc.)
    -priority="string" - Syslog priority (e.g. "crit", "info", etc.)
    -msg="string"      - Message to send, or "-" for stdin
```
 
Typical Example
---------------
 
```
echo "Test message" | gologit -dest="syslog.mycompany.com:514" -proto=udp -tag="apache" -facility="local7" -priority="info" -msg="-"
```
 
To Do
-----
* Enforce maximum syslog message size limits (with option to override)
* Better error handling
* This is my first Go language program ever so there is most certainly many ways to improve the code
