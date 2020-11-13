# Qrcode生成器
Golang实现的二维码生成工具，支持命令行模式和服务器模式

# 安装
需要先安装golang 1.14以上

1.克隆项目
```
git clone git@github.com:boyxp/qrcode.git
```

2.编译
```
go build qrcode.go
```

# 命令行模式
二维码图片PNG输出到指定目录

```
./qrcode --content=www.google.com --size=200 --level=M --output=/tmp/qrcode.png
```

参数：
content 二维码内容，
size 像素尺寸，
level 容错级别 L/M/Q/H 四个级别，
output 输出图片路径


# 服务器模式
GET请求地址栏传参，直接输出图片到浏览器

```
./qrcode --server --port=8888 &
```

参数：
server 开启服务器模式，
port 绑定端口

请求：
```
http://localhost:8888/?content=www.google.com&size=300&level=H
```

