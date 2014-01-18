gologit
=======

gologit is a command-line tool to send data to a syslog server.  I wrote it because logger(1) that comes with GNU's util-linux is horribly buggy.

Usage
-----
```
gologit <args>
    -dest="host:port"  - Destination syslog host
    -tag="string"      - Tag or application name
    -facility="string" - Syslog facility (e.g. "kern", "local0", etc.)
    -priority="string" - Syslog priority (e.g. "crit", "info", etc.)
    -msg="string"      - Message to send, or "-" for stdin
```
 
Typical Example
---------------
 
```
echo "Test message" | gologit -dest="syslog.mycompany.com:514" -tag="apache" -facility="local7" -priority="info" -msg="-"
```
 
To Do
-----
* Enforce maximum syslog message size limits (with option to override)
* Better error handling
