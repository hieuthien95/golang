# GOLANG CHO LẬP TRÌNH BACKEND

## Week 1
### Giới thiệu về Golang

– Ngôn ngữ lập trình Go

– Một vòng tham quan trực tuyến về Go

– Cài đặt, thiết lập môi trường

– Cách tổ chức  mã nguồn Go

Thực hành: Clone Underscore

## Week 2
### Đi sâu vào Golang

– Tìm hiểu go routine, channel, defer, panic, recover

– Lập trình đa luồng đồng thời với Go

– Ghi nhật ký (Logging) và kiểm thử đơn vị (unit test)

Thực hành: Xây dựng ứng dụng Crawler với Go

## Week 3
### Xây dựng RESTful hiệu suất cao với Golang

– Làm việc với JSON

– Làm việc với Gin

– Làm việc với GORM / MySQL

– Xác thực và phân quyền

Thực hành: Sử dụng https://reqres.in/

## Week 4
### Phát triển ứng dụng có khả năng xử lý đồng thời hiệu suất cao và dễ mở rộng

– Tối ưu hóa mã hóa / giải mã JSON

– Tìm hiểu cách sử dụng kết nối Pool

– Tìm hiểu cách sử dụng bộ đệm (Cache)

– Một số thủ thuật hay

Thực hành: Xây dựng dịch vụ để tạo id tăng tự động

## Week 5
### Thực hành phát triển dự án phần mềm với Golang

– Cài đặt môi trường với Docker

– Thiết lập TDD

– API kiểm tra hiệu suất

– Triển khai & giám sát

Thực hành: Xây dựng ứng dụng Todo với GraphQL

## Week 6
### Xây dựng đồ án cuối khóa cá nhân và thuyết trình

Xây dựng phần mềm nguồn mở của riêng bạn với Golang và thuyết trình sản phẩm.

# build golang

# vendor
```
$ go mod init go-module
go: creating new go.mod: module go-module

$ go build
go: finding github.com/labstack/gommon/color latest
go: finding github.com/labstack/gommon/log latest
go: finding golang.org/x/crypto/acme/autocert latest
go: finding golang.org/x/crypto/acme latest
go: finding golang.org/x/crypto latest
go: downloading golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a
go: extracting golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a
go: finding github.com/valyala/fasttemplate v1.0.1
go: finding github.com/mattn/go-isatty v0.0.7
go: downloading github.com/mattn/go-isatty v0.0.7
go: downloading github.com/valyala/fasttemplate v1.0.1
go: extracting github.com/mattn/go-isatty v0.0.7
go: extracting github.com/valyala/fasttemplate v1.0.1

$ go mod vendor
go: downloading golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223
go: extracting golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223

$ go run main.go

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v3.3.10-dev
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:9090
```

# ISSUE
## go get ussue: permission
Mac
```
env GIT_TERMINAL_PROMPT=1 go mod vendor
env GIT_TERMINAL_PROMPT=1 go get gitlab.ghn.vn/common-projects/go-sdk v0.1.31
```
## go get ussue: 410 gone
Mac
```
$ export GO111MODULE=on
$ export GOPROXY=direct
$ export GOSUMDB=off
$ go get gitlab.ghn.vn/common-projects/go-sdk v0.1.31
```

Window
```
$env:GO111MODULE="on"
$env:GOPROXY="direct"
$env:GOSUMDB="off"
go get gitlab.ghn.vn/common-projects/go-sdk v0.1.31
```
