Generate assets
===============

Environment
-------------

For example

    [vagrant@localhost origin]$ mkdir -p /go/bin
    [vagrant@localhost origin]$ cp /data/bin/go-bindata /go/bin/

CI
---

Update custome assets

    [vagrant@localhost origin]$ GOPATH=/go ./hack/vendor-console.sh 
    Making sure go-bindata binary has been built...
    No git ref specified, using current state of the repo including any unstaged changes...
    Building bindata.go files...
    Done vendoring.  To run the console, run 'make clean build' and restart your origin server.
    
    [vagrant@localhost origin]$ ls -ltr -R pkg/assets/
    pkg/assets/:
    总用量 10952
    -rw-r--r--. 1 vagrant vagrant     1341 10月 19 04:33 mime.go
    -rw-r--r--. 1 vagrant vagrant     1102 10月 19 04:33 join.go
    drwxr-xr-x. 1 vagrant vagrant      102 10月 19 04:33 java
    -rw-r--r--. 1 vagrant vagrant      120 10月 19 04:33 interfaces.go
    -rw-r--r--. 1 vagrant vagrant     6700 10月 19 04:33 handlers_test.go
    -rw-r--r--. 1 vagrant vagrant     9169 10月 19 04:33 handlers.go
    -rw-r--r--. 1 vagrant vagrant     5323 10月 19 04:33 extensions.go
    -rw-r--r--. 1 vagrant vagrant      384 10月 19 04:33 doc.go
    -rw-r--r--. 1 vagrant vagrant       96 10月 19 04:33 assets.go
    -rw-r--r--. 1 vagrant vagrant 11164962 10月 19 05:38 bindata.go
    
    pkg/assets/java:
    总用量 9968
    -rw-r--r--. 1 vagrant vagrant 10205743 10月 19 05:38 bindata.go
    
    [vagrant@localhost origin]$ date
    2016年 10月 19日 星期三 05:40:17 UTC
