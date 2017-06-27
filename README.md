![Zenpress](https://github.com/insionng/zenpress/raw/master/public/favicon.png)

Zenpress
===
Zenpress是一个灵感来源于wordpress的cms项目，最终目标是实现一个简单和强大的Golang CMS系统网站，以内置Qlang语言的纯GO VM实现来动态解析脚本，从而实现插件支持，作小许改动后可以应用为Blog、企业站、小说站、图站等多种类型网站..


基于[MAKROSS](https://github.com/insionng/makross)框架和[GORM](https://github.com/jinzhu/gorm)，于2017年06月10日开始重写。

已实现插件机制及主题机制的支持！



## 轻量级社区版本

[YouGam社区](http://www.yougam.com/) 支持PostgreSQL、MySQL、Sqlite3、Go/Leveldb、Boltdb、Tidb等等数据库

(采用reddit系列算法而类似V2EX的轻社区,购买系统请联系作者：QQ547092001/微信xiongtuntianxia)

问题反馈：https://github.com/insionng/yougam/issues



# 通过捐款支持Zenpress项目
如果你喜欢这个项目的话， 可以通过捐款的方式， 支持作者继续更新本项目或者做出其他更多好玩好用的开源应用： 比如为本项目修补漏洞、添加更多有趣的功能， 或者发行有更多更棒特性的下一版等等。

支付宝捐款地址： 赞助我写更多更好的开源项目 insion@live.com

![Donation](https://github.com/insionng/zenpress/raw/master/public/donation.jpg)


请求
===
    Golang 1.9

安装
===
    go get -u github.com/insionng/zenpress

    cd $GOPATH/src/github.com/insionng/zenpress

    go build
    
    screen -dmS zenpress ./zenpress


## 交流联系
欢迎大家加入QQ专用交流群:245386165/作者QQ：547092001，微信账号：xiongtuntianxia

技术分享：[http://www.yougam.com](http://www.yougam.com)

凝视网络：[http://www.nilyes.com](http://www.nilyes.com)


## 授权许可
除特别声明外，本项目代码遵循[BSD 3-Clause License](<http://opensource.org/licenses/BSD-3-Clause/>)（3项条款的BSD许可协议）。
作者本人及凝视网络保留本项目著作权。