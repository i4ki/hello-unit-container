# Unit Container Example

```
make
make push IMAGE=<image tag>
```

To setup Unit, edit the files `app.json` and `listener.json`, then run:

```
$ curl --unix-socket <unit-controller-sock> \
    http://127.0.0.1:8080/config/applications/helloworld -XPUT \
    --data-binary @app.json
{
	"success": "Reconfiguration done."
}
$ curl --unix-socket <unit-controller-sock> -XPUT \
    http://127.0.0.1:8080/config/listeners/*:8300 \
    --data-binary @listener.json
{
	"success": "Reconfiguration done."
}
```

The program output is:

```
$ curl localhost:8300
Unit container example!
My PID: 1
rootfs:
.dockerenv .old bin boot dev etc home lib lib64 media mnt opt proc root run sbin srv sys tmp unit-example usr var 
UID: 65534
GID: 65534

Environment variables: 
SSH_CONNECTION=10.0.2.2 50938 10.0.2.15 22
LANG=C.UTF-8
XDG_SESSION_ID=34
USER=vagrant
PWD=/vagrant
HOME=/home/vagrant
SSH_CLIENT=10.0.2.2 50938 22
XDG_DATA_DIRS=/usr/local/share:/usr/share:/var/lib/snapd/desktop
SSH_TTY=/dev/pts/3
MAIL=/var/mail/vagrant
TERM=xterm-256color
SHELL=/bin/bash
SHLVL=1
LOGNAME=vagrant
XDG_RUNTIME_DIR=/run/user/1000
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/usr/local/go/bin
_=./build/unitd
OLDPWD=/vagrant/unit-example
NXT_UNIT_INIT=1.10.0;1;28987,0,9;29027,0,10;2,
NXT_UNIT_CNT_PID=29027
```
