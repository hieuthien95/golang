*Create image and run container*
```
$ docker build -t golang_health_image .
$ docker run -d --name golang_health_container -p 9092:9092 golang_health_image
```

# SSH

```
hieut@DESKTOP-M6CBJL7 MINGW64 ~
$ ssh hieut@35.247.146.103
Last login: Thu Mar 28 01:39:55 2019 from 14.161.15.48
```

```
[hieut@instance-1 ~]$ ifconfig
docker0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.17.0.1  netmask 255.255.0.0  broadcast 0.0.0.0
        inet6 fe80::42:42ff:fe96:9739  prefixlen 64  scopeid 0x20<link>
        ether 02:42:42:96:97:39  txqueuelen 0  (Ethernet)
        RX packets 380  bytes 30989 (30.2 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 389  bytes 44512 (43.4 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1460
        inet 10.148.0.2  netmask 255.255.255.255  broadcast 10.148.0.2
        inet6 fe80::4001:aff:fe94:2  prefixlen 64  scopeid 0x20<link>
        ether 42:01:0a:94:00:02  txqueuelen 1000  (Ethernet)
        RX packets 129807  bytes 1173264897 (1.0 GiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 99983  bytes 9387467 (8.9 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        inet6 ::1  prefixlen 128  scopeid 0x10<host>
        loop  txqueuelen 1000  (Local Loopback)
        RX packets 634  bytes 63325 (61.8 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 634  bytes 63325 (61.8 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

vethaa2979c: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet6 fe80::cc82:1cff:fef5:cd24  prefixlen 64  scopeid 0x20<link>
        ether ce:82:1c:f5:cd:24  txqueuelen 0  (Ethernet)
        RX packets 89  bytes 7909 (7.7 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 97  bytes 11132 (10.8 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
```

```
[hieut@instance-1 ~]$ exit
logout
Connection to 35.247.146.103 closed.
```

```
hieut@DESKTOP-M6CBJL7 MINGW64 ~
$ ssh hieut@35.247.146.103
Last login: Thu Mar 28 08:01:32 2019 from 14.161.15.48
[hieut@instance-1 ~]$ ssh-keygen
Generating public/private rsa key pair.
Enter file in which to save the key (/home/hieut/.ssh/id_rsa):
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /home/hieut/.ssh/id_rsa.
Your public key has been saved in /home/hieut/.ssh/id_rsa.pub.
The key fingerprint is:
SHA256:18rzFmAjRWY+Ox+Dsd/1u88umsrSOHn68PWKEzajhRc hieut@instance-1
The key's randomart image is:
+---[RSA 2048]----+
|         .+      |
|         +.      |
|         .+      |
|        . EB     |
|        S+*++   .|
|        .oB=.+ ..|
|        .O+++.. .|
|        *++=.o...|
|        .**+=o.=*|
+----[SHA256]-----+
```

```
[hieut@instance-1 ~]$ cat /home/hieut/.ssh/id_rsa.pub
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDuvVmO9M1b+ttHo56dAyimEJ0a25d6hZTvTCDRYrk9lUF3R7VgClBXkUeg0Oo58DyqsQ+slMW1otFuSxzxBGr8jCF2WNswq4UbHfYonQsfKR1XuXrTxCuzTkeAQ+hSKP6Ht3U5ssHZTBYu3zp3NOntGzzkJBrr3z+DKji1+oLsORtIEDmAnWJ7YEE4Vp2IYJRMZ+2aEBBn/VJQBs053hYJ6kJcc0posHmXDGZ4t9MLthTailmeRtljzlkIVvRVNjXGbZkmIafNCWYkVKY6fGSDmriwPaEwgleqLp6NUrK8ewvnLhXHnPu8B/IL+Qs2NMfwzIGo1U6xTC4CM6DC+XFF hieut@instance-1
```

```
[hieut@instance-1 ~]$ ls

[hieut@instance-1 ~]$ pwd
/home/hieut

[hieut@instance-1 ~]$ mkdir go

[hieut@instance-1 ~]$ cd go

[hieut@instance-1 go]$ git clone git@gitlab.com:lak8s/thienbh-lession2.git
Cloning into 'thienbh-lession2'...
The authenticity of host 'gitlab.com (35.231.145.151)' can't be established.
ECDSA key fingerprint is SHA256:HbW3g8zUjNSksFbqTiUWPWg2Bq1x8xdGUrliXFzSnUw.
ECDSA key fingerprint is MD5:f1:d0:fb:46:73:7a:70:92:5a:ab:5d:ef:43:e2:1c:35.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added 'gitlab.com,35.231.145.151' (ECDSA) to the list of known hosts.
remote: Enumerating objects: 396, done.
remote: Counting objects: 100% (396/396), done.
remote: Compressing objects: 100% (247/247), done.
remote: Total 396 (delta 147), reused 387 (delta 145)
Receiving objects: 100% (396/396), 1.06 MiB | 266.00 KiB/s, done.
Resolving deltas: 100% (147/147), done.
```
```
[hieut@instance-1 go]$ pwd
/home/hieut/go

[hieut@instance-1 go]$ ls
thienbh-lession2

[hieut@instance-1 go]$ cd thienbh-lession2

[hieut@instance-1 thienbh-lession2]$ ls
batch.sh  Dockerfile  go.mod  go.sum  main.go  README.md  vendor
```

```
[hieut@instance-1 thienbh-lession2]$ chmod +x batch.sh

[hieut@instance-1 thienbh-lession2]$ ./batch.sh
Already up-to-date.
Sending build context to Docker daemon  8.836MB
Step 1/5 : FROM golang:1.11.4-alpine
1.11.4-alpine: Pulling from library/golang
Digest: sha256:198cb8c94b9ee6941ce6d58f29aadb855f64600918ce602cdeacb018ad77d647
Status: Downloaded newer image for golang:1.11.4-alpine
 ---> f56365ec0638
Step 2/5 : WORKDIR /go/src/go-module
 ---> Running in a79c1978c79b
Removing intermediate container a79c1978c79b
 ---> 37942a74e9f2
Step 3/5 : COPY . /go/src/go-module/
 ---> 1e89ec43c82a
Step 4/5 : RUN go build -o main .
 ---> Running in 71552e959d16
Removing intermediate container 71552e959d16
 ---> 03ed76eda55e
Step 5/5 : CMD ["./main"]
 ---> Running in fa3afa40f06a
Removing intermediate container fa3afa40f06a
 ---> 29e2bc8c39f9
Successfully built 29e2bc8c39f9
Successfully tagged golang_health_image:latest
...
...
```

```
[hieut@instance-1 thienbh-lession2]$ curl "localhost:9092/healthcheck"
{"Status":"OK","StatusCode":200,"Msg":"BHT Service Runing"}
```

# fix permition ssh
```
[hieut@instance-1 thienbh-lession2]$ sudo usermod -aG docker hieut

[hieut@instance-1 thienbh-lession2]$ sudo usermod -aG root hieut

[hieut@instance-1 thienbh-lession2]$ sudo chmod 777 /var/run/docker.sock
```

```
/etc/nginx/nginx.conf
```

