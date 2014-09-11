Toropress
===
Toropress是一个简单和强大的Golang CMS系统网站，作小许改动后可以应用为Blog、企业站、小说站、图站等多种类型网站..

第一个版本基于[Torgo](https://github.com/insionng/torgo)([Beego](https://github.com/astaxie/beego)山寨版)与[Qbs](https://github.com/coocood/qbs)开发.

第二版于2013年10月1日更新,由于beego最新版本几乎已经把[Torgo](https://github.com/insionng/torgo)的所有修改优化吸收完全了,从而决定了Toropress源码再次回归[Beego](https://github.com/astaxie/beego),并采用了[Xorm](https://github.com/lunny/xorm),目前已经将[Qbs](https://github.com/coocood/qbs)代码完全迁移到[Xorm](https://github.com/lunny/xorm),这将获得更好的使用体验,感谢那些捐助本项目的用户,如需定制请联系我.


##轻量级论坛版本(类似V2EX,购买请联系作者我)

[YouGam社区](http://www.yougam.com/)


##线上开源技术分享平台版本

[SudoChina社区](http://www.sudochina.com/)


## 企业型演示网站

[访问艾美](<http://www.ibeautys.com/>)

[访问因特拉](<http://www.interla.net/>)


## 实验性社区型演示网站

非常时刻 (Veryhour) this site has been down.


## 图片分享社区演示网站

梨花树社区 this site has been down.


# 通过捐款支持Toropress项目
如果你喜欢这个项目的话， 可以通过捐款的方式， 支持作者继续更新本项目或者做出其他更多好玩好用的开源应用： 比如为本项目修补漏洞、添加更多有趣的功能， 或者发行有更多更棒特性的下一版等等。

支付宝捐款地址： 赞助我写更多更好的开源项目 insion@live.com


## 感谢捐助或定制本项目的用户 排名按时间顺序
    右建华 5000,定制版
    李新友 5000,定制版 
    孙 300,捐助者要求不公开全名

##第二版安装请先更新安装beego和xorm

    go get -u github.com/astaxie/beego
    go get -u github.com/lunny/xorm

##第一版安装请先更新torgo

    go get -u github.com/insionng/torgo

##Sqlite3安装(默认,当选择其他数据库时可忽略此项)
	如果你默认使用Sqlite3安装,请确保已经安装sqlite3驱动，譬如64位的win7：
	SQLITE3驱动编译环境是TDM版MINGW64(http://tdm-gcc.tdragon.net/)，
    安装好后请把TDM版MINGW64的bin目录路径加入到你的win7环境变量path里面。

	
	go get -u github.com/mattn/go-sqlite3

    下载toropress源码后解压并cd切换目录到 toropress目录下，然后执行go build app.go，编译好后，运行./app即可。
    默认用户:root,默认密码:rootpass


安装
===
由于开发项目并不在作者名下,而是直接处于gopath的src路径下面,所以你不能直接使用go get来安装本项目!

请直接下载zip源码并解压到你的gopath/src/toropress下面


或者:

Linux/Unix/Osx
~~~
cd gopath

git clone https://github.com/insionng/toropress.git

cd toropress
chmod +x build.sh
./build.sh
./app
~~~

Windows
~~~

双击运行build.cmd进行编译,当然,前提是你的go编译环境健在...

~~~

## 交流
欢迎大家加入QQ专用交流群:231956113/作者QQ：547092001

技术分享：[http://www.sudochina.com](http://www.sudochina.com)


## 授权许可
除特别声明外，本项目代码遵循[BSD 3-Clause License](<http://opensource.org/licenses/BSD-3-Clause/>)（3项条款的BSD许可协议）。


veryhour fork by toropress
==========================

A very special time website

这是未来社区网站的新模式探索项目~
基于[Beego](https://github.com/astaxie/beego)与[xorm](https://github.com/lunny/xorm)开发的GOLANG微博社区系统网站，从toropess源码衍生而来!


mzr fork by toropress
==========================

梨花树系统开发代号[mzr](https://github.com/insionng/mzr)

基于[Beego](https://github.com/astaxie/beego)与[xorm](https://github.com/lunny/xorm)开发的一个分享图片为主的社区网站系统，从toropess源码衍生而来!



Sdc fork by toropress
==========================

问答系统开发代号[sdc](https://github.com/insionng/sdc)

基于[Beego](https://github.com/astaxie/beego)与[xorm](https://github.com/lunny/xorm)开发的这是一个问答社区网站系统，从toropess源码衍生而来!


## 交流
欢迎大家加入QQ专用交流群:231956113/作者QQ：547092001

技术分享：[http://www.sudochina.com](http://www.sudochina.com)


## 授权许可
除特别声明外，本项目代码遵循[BSD 3-Clause License](<http://opensource.org/licenses/BSD-3-Clause/>)（3项条款的BSD许可协议）。

