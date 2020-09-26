# GoForum
goforum为go语言写的一个Go Web应用

goforum 项目地址： https://github.com/xChen16/forum


## 使用

​本项目使用gomod进行版本管理

使用forum.sql中的建表语句在新建的数据库中新建数据表，参照config.example.json文件配置新建config.json

直接运行
```
go run main.go
```

应用主页地址默认为：http://localhost:8080/

使用goi18n把消息本地化

```
/forum$ cd locales/
/forum/locales$ goi18n merge active.en.json active.zh.json
```

## 协议

[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
