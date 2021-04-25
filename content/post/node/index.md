---
# Documentation: https://sourcethemes.com/academic/docs/managing-content/

title: "学习笔记"
subtitle: ""
summary: "平时学习中,比较重要的知识点,记下来备忘"

# 标签
tags:
- Note

#关键字
keywords:
- 学习笔记
- 笔记
- Linux

# 文章分类,在列表页显示(暂时使用tags)
categories:

# 自定义url
url: ""
# 文章排序权重
weight: 100
# 是否精选文章,显示在首页
featured: true
# 草稿开关,开启则文章不会打包
draft: false
# 是否显示目录
toc: true
# 类型,默认是目录名,所以"post",也可以写"widget_page","dosc"等
type: post
# 文章添加时间
date: 2020-10-04
# 文章更新时间,可开启 自动使用git时间
# lastmod: 2020-05-03T19:54:10+08:00
# 定时发布时间
# publishDate: 2020-05-03T00:00:00+08:00
# markdown文档
markup: "md"
# 显示"阅读时长"(params.toml也有统一的设置)
reading_time: true
# 底部显示社交链接
share: true
# 底部显示作者资料
profile: false
# 是否允许评论(总开关在params.toml,comments=>engine)
commentable: true
# 是否允许编辑(跳到github,前提是github有权限)
editable: false
# 作者
authors:
- admin

# Markdown Slides (optional).
#   Associate this talk with Markdown slides.
#   Simply enter your slide deck's filename without extension.
#   E.g. `slides = "example-slides"` references `content/slides/example-slides.md`.
#   Otherwise, set `slides = ""`.
#slides: example
#links:
#  - icon_pack: fab
#    icon: twitter
#    name: 这里显示外链显示在详情页
#    url: 'https://twitter.com/twbworld'
#  - icon_pack: fab
#    icon: twitter
#    name: 也可以链接pdf文件,对现场讲解很有用
#    url: 'demo.pdf'

# Projects (optional).
#   Associate this post with one or more of your projects.
#   Simply enter your project's folder or file name without extension.
#   E.g. `projects = ["internal-project"]` references `content/project/deep-learning/index.md`.
#   Otherwise, set `projects = []`.
# 关联项目
projects: []

# Featured image
# To use, add an image named `featured.jpg/png` to your page's folder.
# Placement options: 1 = Full column width, 2 = Out-set, 3 = Screen-width
# Focal points: Smart, Center, TopLeft, Top, TopRight, Left, Right, BottomLeft, Bottom, BottomRight.
# 控文章列表页图片,也可以在当前目录下放置featured.jpg/png文件
image:
  placement: 3
  image: "https://cdn.jsdelivr.net/gh/twbworld/hosting@master/img/20201004235821.jpg"
  caption: "图片来源: [**wikihow**](https://www.wikihow.com/)"
  focal_point: "Center"
  preview_only: true
  alt_text: "Note"

# 文章页最头部的图片,如果当前目录下没有featured.jpg/png文件,且,三方image没设置image,则列表页也会显示这图片(image如设置了image,此处无效)
#header:
#  placement: 3
#  image: ""
#  caption: ""
#  focal_point: "Center"
#  preview_only: false
#  alt_text: ""


---



## 扩展知识

### 双系统制作
> <https://blog.csdn.net/mengxiangjia_linxi/article/details/78965103>

1. 电脑已有 `Windows系统` 的前提下,用 `UltraISO` 把 `centos系统` 刻入U盘
2. 启动系统U盘(新机用 `uefi启动` ),在系统选择页面,按下 `tab健` 修改启动项,把 `LABEL=CentOS\x207\x20x86_64` 改为 `/dev/sdb4` (不一定是 `sdb4` ,可能是 `sda4` , `sdc4` 等等,具体百度)
3. 安装系统界面,选择硬盘,自动分配分区得以回收空间
4. 重做 `PE启动盘` ,启动盘用 `NTBOOT` 软件做 `双系统引导修复`






## 常用 命令&&快捷键


### Linux

| 命令 | 描述 |
| ---- | ---- |
| ..                                | 上一级目录 |
| cd -                              | 上一次目录 |
| ctrl+C                            | 强制中断 |
| ctrl+U                            | 清除输入命令 |
| !v                                | 执行上一条以'v'开头的命令 |
| ll                                | ls -l 缩写 |
| 2>&1                              | 输出错误 |
| \|grep                             | 过滤结果 |
| ctrl+a/e                          | 命令最前/后 |
| clear                             | 清屏(或ctrl+L) |
| pwd                               | 当前目录 |
| cat                               | 查看文件内容 |
| mkdir/mv/cp/rm                    | 文件名操作 |
| \cp                               | 当有同名文件,不需要输"Y"即可覆盖 |
| useradd                           | 创建新用户 |
| which                             | 查找命令 |
| chown                             | 设置文件拥有者 |
| chmod                             | 设置u/g/o权限 |
| setfacl -R -m u:apache:rwx /var   | 设置单独用户的权限(getfacl查看) |
| chattr                            | 设置文件属性,如设置成禁止删除文件 |
| ps -ef(ps aux)                    | 查看所有进程 |
| pstree -p                         | 查看进程树 |
| pkill httpd                       | 关闭某进程(或killall) |
| uptime                            | 计算机负载 |
| top                               | 实时负载(可安装htop加强版) |
| df -h                             | 磁盘占用 |
| du -h --max-depth=0 \* (du -sh \*) | 目录下各个文件大小 |
| pkill -kill -t pts/1              | 强退指定用户 |
| who                               | 在线用户(或缩写w) |
| last                              | 用户登录历史记录 |
| id                                | 查看用户所在的组 |
| lsof -i:80                        | 查看端口 |
| netstat -antup \| grep nginx       | 查看进程或端口等 |
| man                               | 命令帮助(可安装中文软件) |
| curl ifconfig.me                  | 查看公网ip |
| rpm -qa\|grep                      | 查看应用程序 |
| pkill -HUP nginx                  | 系统级别重载配置文件 |
| systemctl status nginx.service    | 查看状态(或service nginx status) |
| chkconfig iptables off            | 永久关闭防火墙 |
| chkconfig iptables on             | 永久关闭后重启 |
| ln -s                                | 软连接 |
| nohup <shell\>                     | 在后台运行shell命令 |
| tree -LNFC 2                       | tree软件常用命令 |
| tar -xaf                        | 识别压缩文件类型,进行解压(-caf) |
| unzip -O cp936                  | zip文件解压避免乱码 |
| scp -r /var root@127.0.0.1:/var/ | 目录传输 |
| find / -type f -name "*.txt" \| xargs grep "hello" | 查找文件内容 |
| cat <filename\> \| openssl dgst -sha256 -binary \| openssl base64 -A | 计算文件sha256校验值(css校验方法:`<link href="filename.css" integrity="sha256-h20CPZ0QyXlBuAw7A+KluUYx/3pK+c7lYEpqLTlxjYQ=">`) |
| sha256sum filename | 计算文件md5校验值 |
| docker container list -aq \|xargs docker container rm | docker清理容器(同样效果 `docker container prune`) |
| docker volume ls -q \|xargs docker volume rm | docker清理卷(同样效果 `docker volume prune`) |


### Git

| 命令 | 描述 |
| ---- | ---- |
| `git diff <file>`                             | 工作区版本库对比 |
| `git diff <commit-id> <commit-id> --stat`    | 比较两个版本变化了的文件 |
| `git commit --amend`                          | 修改已commit的备注 |
| `git commit –am <备注>`                       | 直接跳过add命令 |
| `git reset --hard HEAD`                       | 还原最新版 |
| `git reset --hard <commit-id>`                | 还原指定版 |
| `git reset HEAD <file>`                       | 放弃暂存区的修改 |
| `git checkout -- <file>`                      | 放弃工作区的修改 |
| `git checkout . && git clean -df`             | 清除所有新增的文件目录 |
| `git checkout –b <name>`                      | 创建并切换分支 |
| `git branch -D <name>`                        | 丢弃未合并的分支 |
| `git branch -r -d origin/<name> && git push origin :<name>`     | 删除远程分支 |
| `git mergetool`                               | 解决冲突的合并工具 |
| `git tag -a <name> -m <备注> <commit-id>`     | 创建并指定标签信息 |
| `git show <name>`                             | 查看标签信息 |
| `git stash`                                   | 保存当前工作现场 |
| `git stash pop`                               | 恢复并删除工作现场 |
| `git stash list`                              | 查看工作现场 |
| `git stash apply <stash@{0}>`                 | 恢复工作现场 |
| `git stash drop <stash@{0}>`                  | 删除工作现场 |
| `git reflog <file>`                           | 命令历史 |
| `git revert <commit-id>`                      | 撤销某个版本 |
| `git cherry-pick <commit-id>`                 | 提取某个分支的某个版本到当前分支 |
| `git check-ignore -v <file>`                  | 检查哪个规则忽略了文件 |
| `git rm -rf --cached . && git add -A`        | 清徐缓存,用于更新gitignore |
| `git submodule add <远程项目地址> <路径>`      | 添加子模块 |
| `git submodule update --init --recursive`     | 下载所有子模块 |
| `git submodule update --remote`               | 更新所有子模块 |
| `git worktree add ../folder master`           | 关联某分支到目录 |
| `git remote show origin`                      | 详细的远程分支的信息 |
| `git log --oneline --graph --decorate --all -9`   | 可视化显示版本树 |
| `git config --global alias.lg "log --color --graph --all --oneline  --decorate --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"` | 可视化显示版本树配置别名 |
| `git config --global core.filemode false`            | 关闭git对文件权限的监视 |
| `git config --global core.quotepath false`        | 正常显示中文文件名 |
| `git config --global color.ui true`               | 彩色显示 |


### Vim

| 命令 | 效果|
| ---- | ---- |
| !<shell\>             | 执行系统命令 |
| u                     | 后退 |
| Ctrl + r              | 前进 |
| i/I/a/A/o/O           | 前/行首/后/行尾/上行/下行 插入 |
| b/e/^/$/H/M/L/gg/G    | 移动光标到 跳前/跳后/行首/行尾/窗头/窗中/窗尾/头部/尾部 |
| Ctrl + b/f/d/u             | 翻一页/半页 |
| D                     | 删除到行尾 |
| dG                    | 向下删除 |
| dd                    | 删除当前行 |
| yy                    | 复制 |
| p                     | 粘贴 |
| ddp                   | 与下行互换位置 |
| yyp                   | 新增一行 |
| ggvG                  | 全选 |
| set number            | 显示行数 |
| vim +X <filename\>       | 文件加密 |
|:<number\>             | 跳到指定行 |
|/<name\>               | 搜索(按 n/N 向 下/上 查找) |


### Sublime

| 快捷键 | 描述 |
| ---- | ---- |
| Ctrl+Shift+L      | 所有选中的行,在行后后出现光标 |
| Ctrl+Shift+M      | 选中括号里的内容 |
| Ctrl+M            | 光标在括号前后跳 |
| Ctrl+Shift+[ / ]  | 隐藏/显示 内容 |
| Ctrl+J            | 多行成一行 |
| Ctrl+KK           | 删除行后的内容 |
| Ctrl+K+U / L      | 选中的内容 大写/小写 |
| Ctrl+F2           | 标记行 |
| F6                | 检错 |
| Alt+.             | 补充标签 |
| Alt+数字          | 跳转tag |
| shift+pgup        | 向下选择 |


### Vimium

| 快捷键 | 描述 |
| ---- | ---- |
| o       | 搜索 |
| j/k     | 滚动 |
| gg/G    | 头/底部 |
| d/u     | 翻页 |
| f       | 打开连接 |
| x/X     | 关闭/恢复标签页 |
| J/K     | 上/下一个标签页 |
| yy      | 复制链接 |
| H/L     | 后退/前进记录 |
| p       | 刷新 |
| t       | 新建标签页 |


### chrome

| 快捷键 | 描述 |
| ---- | ---- |
| Ctrl+0 | 恢复100% |
| Ctrl+j | 下载页 |
| Ctrl+Shift+Delete | 清除数据 |
| Ctrl+u |  html源码 |
| chrome://net-internals/#dns | 清除DNS |


### windows

| 快捷键 | 描述 |
| ---- | ---- |
| Win+←                  |  小化 |
| win+v                  |  粘贴板 |
| Win+i                  |  设置页 |
| Shift+Ctrl+Esc         |  任务管理器 |
| Alt+F4                 | 关机 |
| Alt+1/Enter/双击       |  查看文件属性 |
| Alt+↑                  |  上一级目录 |
| Alt+2(Shift+Ctrl+n)    |  新建文件夹 |
| 右击+v+d               |    隐藏图标(等等操作) |
| ctrl+win+D             |   新建桌面 |
| ctrl+win+←             |  返回桌面(老板键) |









## Linux

* 端口号<1024为系统预留端口
* uid < 500 为系统用户
* 非`root`用户不能监听<1024的端口(所以一般`root`启动`nginx`)
* nobody为Linux非登录用户,其作为nginx运行用户可保证安全
* `Linux`的共享:`NFS`协议
* `windows`与`linux`共享:`SMB`协议(`Samba`软件)


| 权限 | 符号 | 描述 |
| ---- | ---- | ---- |
| 可执行       | x | 可以进入该目录,无法读取该目录中的内容 |
| 可执行可读   | xr | 可以进入该目录并读取该目录中的内容,不能创建文件 |
| 可执行可写   | xw | 可以创建文件但是不能读取 |


* 目录描述 :

![](https://cdn.jsdelivr.net/gh/twbworld/hosting@master/img/20210404165626.jpg)


* ubuntu环境安装流程 :
    ``` txt
    .bashrc命令别名
    ufw防火墙关闭
    静态ip
    ali源
    vim安装及.vimrc解决乱码
    ssh安装及配置
    samba安装及配置;用户:root;密码:root;
    传项目文件
    python升级
    net-tools安装
    htop安装
    git安装及配置
    tig安装
    tree安装
    screen安装
    docker安装
    docker-compose安装
    golang安装
    hugo安装
    node安装
    php安装
    mysql安装
    ```





## Git

* 初始配置Git
    > 1. 设置本地用户和邮箱
    `git config --global user.name <your_name> && git config --global user.email <your_email>`
    > 2. 本地生成公钥
    `ssh-keygen -t rsa -C <your_email>`
    > 3. 找到公钥,复制到github账号(即添加过公钥的用户才被允许管理github项目)
    > 4. 测试公钥
    `ssh -T git@github.com`
    > 5. 到github创建项目
    > 6. 本地创建项目并初始化
    `git init`
    > 7. 关联远程库和本地
    `git remote add <自定义远程仓库名,如origin> <远程项目地址>`
    > 8. 可查看本地关联的所有仓库主机
    `git remote -v 或者 git remote show origin`
    > 9. 把项目添加到缓存区
    `git add . 和 git commit –m <备注>`
    > 10. 提交项目到origin仓库的master分支
    `git push origin master (添加"--tags"可同时上传tag)`

* git hook :
    > * 第三方仓库git hook
    > * 服务器部署git裸仓库
        <http://kongfangyu.com/2016/02/12/git-deploy>

* 备注规范 :
    > * feat: 新功能
    > * fix: 修复 bug
    > * perf: 提升性能的改动
    > * refactor: bug 修复和添加新功能之外的代码改动,重构大改动
    > * chore: 其他修改,构建过程或辅助工具和库（如文档生成）的更改
    > * docs: 文档变动
    > * style: 格式调整 , 对代码实际运行没有改动 , 例如添加空行、格式化等
    > * test: 添加或修正测试代码
    > * ci: 持续集成相关文件修改
    > * release: 发布新版本
    > * revert: 恢复上一次提交

* Git工作流
    > ![来源:https://nvie.com/posts/a-successful-git-branching-model/](https://cdn.jsdelivr.net/gh/twbworld/hosting@master/img/20210404170704.png)






## 编译安装
``` sh

# 配置
./configure --prefix=/usr/local/nginx
# 清除编译文件
make clean
# 编译
make
# 安装
make install
# 软连接
ln -s /usr/local/nginx/sbin/nginx /usr/bin/nginx
# 开机自启
vim /etc/init.d/nginx && chmod a+x /etc/init.d/nginx
# 加入chkconfig管理列表
chkconfig --add /etc/init.d/nginx
chkconfig nginx on
systemctl status nginx.service

```







## CentOS指定版本安装

    1. 到官网找到对应版本的rpm包,如:http://nginx.org/packages
    2. 下载rpm包:wget http://nginx.org/packages/.../xxx.rpm
    3. 安装:yum install -y xxx.rpm
    4. /etc/rc.d/rc.local是系统启动后执行的文件,可在里添加自启的命令(但不推荐用这方式)
    5. /etc/profile是系统变量的文件





## 网络

* `TCP` :需要三次握手,建立了 `TCP` 虚拟通道,之后, `TCP` 运输 `HTTP流`
* `SSL` : 位于 `TCP` 与 `HTTP` 之间,作为 `HTTP` 的安全供应商 , 全权负责 `HTTP` 的安全加密工作。
* `TLS` : 在 `SSL3.0` 版本的基础上 , 重新设计并命名了这个协议 , 其全新的名字为 `TLS` ,形成: `TCP连接时间` + `TLS 连接时间` + `HTTP交易时间`
* `HTTPS` : 通常将 `TLS` 安全保护的 `HTTP` 通信 , 称之为 `HTTPS`  , 以区别于没有 `TLS` 安全防护的 `HTTP` 明文通信。
* `HTTP/2` : 第一次页面与第二次页面都是同一个网站服务器,重用第一个页面 `TCP` 连接
* `QUIC` : `http/2` 去掉 `TCP` ,改用不需要连接的 `UDP` ,形成: `UDP / QUIC` ,第一次 `2.5RTT` ,完成 `QUIC` 交易的连接的 `Session ID` 会缓存在浏览器内存里,第二次,使用 `Session ID` ,重连 `TLS` 连接是一个 `0 RTT` 事件
* `HTTP/3` : 把 `QUIC` 与 `HTTP` 分离,形成: `UDP / QUIC / HTTP`
* 使用 `IPv4` 进行路由 , 使用 `TCP` 进行连接层面的流量控制 , 使用 `SSL/TLS` 协议实现传输安全 , 使用 `DNS` 进行域名解析 , 使用 `HTTP` 进行应用数据的传输。

* `IP` 是网络层
* `TCP` 是传输层
* 网络层 `ARP` : 保存 `IP` 与 `mac` 地址的映射列表,没有则会广播

* `TCP/IP` 连接 + `http` 传输 --> 网卡 --> 内核(通过资源包的四元组信息,信息的其中之二是访问的 `IP+端口` ) --> `socket`(它是文件系统,是资源,储存着端对端的四元组信息,被进程使用) --> 进程
* `socket` 是网络与进程的中介
* https大致流程 :
    1. 网站发送自己的 `公钥(网站)` 到 信任机构, 机构用自己的 `私钥(机构)` 对 `公钥(网站)` 进行加盐加密, 生成 `签名` , 与其他的信息组合成 `证书`, 向网站颁发 这个`证书`
    2. 客户端浏览器建立 https 链接 , 接收到网站给的 `证书` + `公钥(网站)` ; 浏览器用自带的 信任机构 的 `公钥(机构)`, 对 `证书` 内的 `签名` 进行 解密, 解密后对比 `公钥(网站)` 是否正确
    3. https验证成功, 即 `非对称加密` 结束; 接下来 , 双方约定一个 `对称加密` , 用于对接下来的数据进行加密
        > 因为 `非对称加密` 非常耗资源, 所以为了速度, 使用 `对称加密` 而不是 `非对称加密`
* 加密种类 :
    * base64 : 可逆, 可理解为 编码 而不是 真正的加密
    * 散列 : 又称hash(哈希), 不可逆 , 包含 MD5, SHA(SHA-256,SHA-512)
    * 对称加密 : 加密和解密共用秘钥
    * 非对称加密 : 一对一对存在(私钥+公钥), 公钥用于加密, 私钥用于解密




## 隧道

> <https://www.zhihu.com/question/39382183>
* 隧道是一种技术,有 `SSH隧道` 和 `http隧道` 和 `UDP隧道(p2p)`
* 隧道的中间服务器是不参与流量分析的,只用来传输流量,只是用于简单的过度
* 代理是可以参与流量的处理的,比如 `http代理`
* 但 `http代理` 不能传输 `https流量` ,因为 `https` 是端对端的加密,没有中间人
* 为了解决不能代理 `https流量` , `http隧道` 配合 `http` 的 `CONNECT报文` 解决了这个问题;
* `SOCKS5` 方式可以代替 `CONNECT报文` 的方法, `CONNECT报文` 只能用于 `http流量` ,而 `SOCKS5` 可以传输任意协议,如 `https` 或 `ftp`
* 大致原理: `本地---(1)--->代理---(2)--->服务器`; 本地与代理建立动态端口隧道,过程`(1)` 使用 `SOCKS5` 或 `http` 的 `CONNECT报文(包裹着https流量)` , 过程 `(2)` 使用 `https`;

> <https://www.zsythink.net/archives/2450>

1. 动态转发(科学上网)
    ``` sh
    # socks5代理的建立
    ssh -N -f -D 127.0.0.1:6666 root@121.199.63.39

    # 原理:本地(127.0.0.1:6666)与代理(121.199.63.39)服务器建立隧道后,本地(127.0.0.1)使用socks5协议在外层加密https协议,通过本地代理端口(6666)-(隧道)-代理(121.199.63.39)-https访问服务端(https://www.google.com)
    curl -x socks5://127.0.0.1:6666 https://www.google.com #http为例
    ssh -o ProxyCommand="nc -X 5 -x 127.0.0.1:6666 %h %p" https://www.google.com #ssh为例,参数-X默认socks5
    ```

2. 本地转发
    ``` sh
    ssh -N -f -L 192.168.2.2:6666:121.199.63.39:22 root@121.199.63.39
    ```

3. 远程转发
    ``` sh
    ssh -N -f -R 121.199.63.39:6666:127.0.0.1:22 root@121.199.63.39
    ```

3. 通过 `121.199.63.39` 的 `22` 端口,登录 `121.199.63.39localhost` 的 `6666` 端口
    ``` sh
    ssh -J root@121.199.63.39:22 -p6666 localhost
    ```



## Nftables

防火墙配置例子 :

``` sh
# my.nft
table inet filter {
    chain input {
        type filter hook input priority 0; policy drop;
        ct state established,related accept
        iif "lo" accept comment "一律接受本地环回"
        ct state invalid drop
        ip protocol icmp icmp type echo-request ct state new accept
        ip protocol udp ct state new jump UDP
        ip protocol tcp tcp flags & (fin | syn | rst | ack) == syn ct state new jump TCP
        ip protocol udp reject
        ip protocol tcp reject with tcp reset
        meta nfproto ipv4 counter packets 0 bytes 0 reject with icmp type prot-unreachable
    }

    chain forward {
        type filter hook forward priority 0; policy drop;
    }

    chain output {
        type filter hook output priority 0; policy accept;
    }

    chain TCP {
        tcp dport ssh ct state new limit rate 15/minute accept comment "避免对SSH施加暴力"
        tcp dport { http, https, mysql } accept comment "80,443,3306"
        tcp dport domain accept comment "DNS:53"
        tcp dport { netbios-ns, netbios-dgm, netbios-ssn, microsoft-ds } accept comment "Samba:137,138,139,445"
        tcp dport { xtel, xtelw } accept comment ":Hugo:1313,1314"
        ip saddr { 192.168.2.100 } drop
    }

    chain UDP {
        udp dport domain accept
    }
}

```





## Ftp

1. 关闭防火墙
2. 安装 `ftp`
3. 修改三份文件 `ftpusers` `user_list` `vsftpd.conf`
4. 重启 `ftp`
5. 防火墙配置文件允许 `21` 端口
6. 开启防火墙
7. `ftp` 限制用户访问目录思路:修改 `/etc/vsftpd/vsftpd.conf` 文件中 `chroot_local_user` , `chroot_list_enable` , `chroot_list_file` 三项; 修改用户默认进入的家目录,则修改 `/etc/passwd`

`vsftpd.conf` 文件参考配置 :

``` sh
listen_ipv6=YES
pam_service_name=vsftpd
tcp_wrappers=YES
write_enable=YES
local_umask=022             #用户ftp创建目录的权限掩码,022即为755
listen=NO                   #监听,yes可能导致无法启动ftp
anonymous_enable=NO         #拒绝匿名用户
local_root=/var/www/html    #本地用户初始目录
#登录限制
local_enable=YES
userlist_enable=YES
userlist_deny=NO
userlist_file=/etc/vsftpd/user_list
#目录限制
chroot_local_user=YES
chroot_list_enable=YES
chroot_list_file=/etc/vsftpd/chroot_list
allow_writeable_chroot=YES
```




## Nginx


* `Nginx` (web服务器)偏向于 `静态资源` ,处理高并发,可用 `反向代理` 从而实现 `负载均衡` (类似 `dns` )
* `Apache` 偏向于` PHP` 的动态资源,bug少稳定
* `负载均衡` 实现思路: `动静态资源分离` ,主服务器使用 `Nginx` ,动态资源则 `反向代理` 到本机或集群的 `Apache` 服务器(一台服务器安装两个 `Apache` 服务可防止 `宕机` )
* `linux` 有端口复用(套接字端口共享)功能( `nginx` 开启 `reuseport` );从内核层面做 `负载均衡` ,避免 `锁竞争` ( `惊群效应` ): <https://www.zhihu.com/question/51618274>
* `Nginx` 启用 `PHP` 需要配置文件
* `Nginx` 等软件自启,需要参考官方的启动代码,添加到 `/etc/init.d/` 下
* `Nginx` 支持缓存
* `Nginx` 支持 `gzip`



* Nginx实现负载均衡需要源码的同步, 使用 `rsync` (或+ `sersync` )实现集群服务器源码同步,源服务器执行的命令:
`rsync -avH --progress --delete --exclude-from=/etc/exclude.txt --password-file=/etc/pw.txt /var/www/html/ username@192.168.43.175::module1`
    >其中按顺序是:含有忽略同步列表的 `exclude.txt` 文件,含有密码的 `pw.txt` 文件,源服务器的同步目录, `username` 用户名,ip地址,对应的模块
    > <https://blog.51cto.com/chenfei123/1707746>

`负载均衡` 和 `前后端分离` 配置 :

``` sh
#后端应反向给Apache代理
upstream proxys {
    server 127.0.0.1:8080 weight=2;
    server 192.168.43.175:9090 weight=3;
    ip_hash;
}
#前端应反向给Nginx代理
upstream proxys2 {
    server 127.0.0.1:8081 weight=3;
    server 192.168.43.175:9090 weight=2;
}
server {
    listen 80;
    server_name www.domain.top domain.top;
    index index.html index.htm index.php;
    location / {
        proxy_pass http://proxys;
        proxy_http_version 1.1;
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        add_header X-Slave $upstream_addr;
    }
    location ~* \.(css|js|png|jpg|jpeg|gif|txt|ico)$|^~/static {
       proxy_pass http://proxys2;
       proxy_set_header Host $http_host;
       proxy_set_header X-Real-IP $remote_addr;
       proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
       add_header X-Slave $upstream_addr;
    }
}
#前端请求代理的配置
server {
    listen 8081;
    server_name www.domain.top domain.top;
    location / {
       root /var/www/html/tp/public;
    }
}

```


`缓存` 和 `gzip` 配置 :

``` sh
http {

  # 开启缓存
  proxy_connect_timeout 10;
  proxy_read_timeout 180;
  proxy_send_timeout 5;
  proxy_buffer_size 16k;
  proxy_buffers 4 32k;
  proxy_busy_buffers_size 96k;
  proxy_temp_file_write_size 96k;
  proxy_temp_path /tmp/temp_dir;
  proxy_cache_path /tmp/cache levels=1:2 keys_zone=cache_one:100m inactive=1d max_size=10g;

  server
  {
      listen 8081 reuseport;
      server_name localhost;
      root /usr/share/nginx/tp/public/;

      location /
      {
      }
  }


  server
  {
      #开启gzip压缩资源
      gzip  on;
      gzip_vary on;
      gzip_proxied any;
      gzip_comp_level 6;
      gzip_buffers 16 8k;
      gzip_http_version 1.0;
      gzip_min_length 100;
      gzip_types text/plain text/css application/javascript application/json application/x-javascript text/xml application/xml application/xml+rss text/javascript application/vnd.ms-fontobject application/x-font-ttf font/opentype image/svg+xml image/x-icon;

      #资源文件缓存
      location ~* \.(css|js|json|png|jpg|jpeg|gif|txt|ico|bmp|swf)$|^~/static/
      {
          proxy_pass http://127.0.0.1:8081;
          proxy_redirect off;
          proxy_set_header Host $host;
          proxy_cache cache_one; #cache_one在http设置层proxy_cache_path定义
          proxy_cache_valid 200 302 24h;
          proxy_cache_valid 301 30d;
          proxy_cache_valid any 5m;
          expires 90d;
          add_header wall  "hey!这文件进行了缓存和压缩哟!";
      }
  }

}
```




## 数据库读写分离

实现 `读写分离` 分为两大步 :
* [主从同步](https://www.jianshu.com/p/bfca0cdfb169)
* 读写分离;有两种方式 :
    * 通过程序实现,如TP框架,性能最好;
    * 使用代理工具,如 :
        * [Mycat](https://www.jianshu.com/p/cb7ec06dae05)
        * [Mysql-Proxy](https://www.jianshu.com/p/cadf337274c1)



> `mysql` 之间可相互主从同步,避免 `单点` ;当有主从数据库发生宕机,可使用 `percona-tooldit` 工具解决恢复宕机后主从数据不同步的问题
    <https://blog.51cto.com/moerjinrong/2352317>

`Mysql-Proxy` 配置文件 `mysql-proxy.cnf` 参考 :

``` sh
[mysql-proxy]
  #运行mysql-proxy用户
user=root
  #主从mysql共有的用户
admin-username=proxy
  #用户的密码
admin-password=root
  #mysql-proxy运行ip(不能是127.0.0.1)和端口,不加端口,默认4040
proxy-address=192.168.1.1:4040
  #slave从数据库Ip地址,默认端口3306
proxy-read-only-backend-addresses=192.168.1.2:3306
  #master主数据库Ip地址,默认端口3306
proxy-backend-addresses=192.168.1.3:3306
  #读写分离配置文件
proxy-lua-script=/usr/lib64/mysql-proxy/lua/rw-splitting.lua
  #管理脚本
admin-lua-script=/usr/lib64/mysql-proxy/lua/admin.lua
  #日志
log-file=/var/log/mysql-proxy.log
  #日志级别,高到低有error|warning|info|message|debug
log-level=info
  #以守护进程方式运行
daemon=true
  #崩溃时,尝试重启
keepalive=true

```



## 数据库

> `mysql8` 特性 :
> * `utf8mb4` 储存表情符号
> * 默认 `InnoDB`
> * 降序索引
> * JSON支持,直接操作json里的数据
> * 跳过锁等待
> * 用户角色管理权限
> * `innodb_dedicated_server` 自适应参数,自动配置 `mysql`

* `redis` 只做缓存使用,避免进行 `持久性` && `队列` && `事务` (总之一句话:专人专事)
* `redis` 和 `DB` 一致性方案:只读 `redis` ,当更新 `DB` 之后把受影响的缓存全部干掉(注意:是干掉,不是更新),并且设定过期时间: <https://www.zhihu.com/question/319817091>
* 出现不能连接数据库的情况,除了防火墙,也可能是数据库没有设置对外添加权限
* 出现用户不能登录的情况,可能是默认登录了 `匿名用户` , 要把用户名为 `''` 的用户清除
* `MyISAM` 存储引擎偏向于大量查询和插入事件,表锁
* `InnoDB` 存储引擎偏向于增删改事件,支持事务,支持索引行锁
* 索引才用 行锁
* `悲观锁` 常用两种 `排它锁` (FOR UPDATE) 和 `共享锁` (LOCK IN SHARE MODE)
* `乐观锁` 通过设置version字段, 用业务判断
* 事务并发会出现 `脏读` `不可重复读` `幻读` ; 数据库隔离等级可解决
* 通过建立 `联合索引` , 从而出现 `索引覆盖` 和 `索引下推`, 达到减少 `回表` 的目的  
<br/>
* 数据库类型
    <http://www.jb51.net/article/55853.htm>
* 免密码登录
    `skip-grant-tables`
* 性能测试
    `explain select * from dnname`
* 性能测试工具(mysql自带)
    `mysqlslap ...`
* 数据库导出
    `mysqldump -uroot -p dbname > dbname.sql`
* 新建数据库
    `CREATE DATABASE mydb CHARACTER SET utf8 COLLATE utf8_general_ci`
* 插入数据
    `insert into $my_name (id,name,tel) values('2','tab',null)`
* 删除数据
    `delete from my_name where id = 1;`
* 修改数据
    `update my_name set id='2',name='tab' where id=1;`
* 插入字段
    `alter table my_name add id varchar(32)`
* 删除字段
    `alter table my_name drop id`
* 去掉name重复的
    `SELECT DISTINCT name FROM my_name`
* 计算个数
    `select count(*) from my_name`
* 计算总和
    `select sum(*) from my_name`
* 计算平均
    `select sum(*)/count(*) from my_name`
* 计算平均
    `select avg(name) from my_name`
* 最小值
    `select min(name) from my_name`

* 双重查询,查询指定最小值对应的的所有字段
    `SELECT * FROM my_name WHERE name=(SELECT max(name)FROM my_name)`
* 根据指定字段分类
    `select * from my_name group by name;`
* 表的联合查询
    `SELECT * FROM my_name as a INNER JOIN my_name2 as b on a.name=b.id`
* 左查询,就是当my_name跟my_name2的的值不相对应则以my_name为主要,my_name2多出的值则省略
    `SELECT * FROM my_name as a left outer JOIN my_name2 as b on a.name=b.id`





## PHP

> php7特性
> * 用绝对路径
> * 重定向后使用exit
> * array_walk
> * &引用节省内存
> * += array()
> * isset($a{1})
> * $_SERVER['REQUEST_TIME']
> * strtr
> * 高亮PHP代码highlight_file
> * 2 <=> 1
> * echo "\u{41}",PHP_EOL;
> * use some\namespace\{ClassA, ClassB, ClassC as C};
> * function arraysSum(int ...$ints): string
> * ``` php
>    class A{static private $b = 'hello world';}
>    $bb = function (){return A::$b;};
>    echo $bb->call(new A);
>    ```

* bug
  ``` php
  [] == []; //为false

  'a' == 0;

  $a = '0.3' - '0.1';
  json_encode($a);
  0.55 * 100 == 55; //为false
  ```
* 单点登录使用 `jwt` 标准进行用户验证
* `PSR` 是php编码规范, `Restful` 是API规范
* json_encode((array)$array, JSON_UNESCAPED_UNICODE|JSON_UNESCAPED_SLASHES|JSON_PRETTY_PRINT)
* 有时候exec命令执行失败,却没有报错,请这样使用:`exec('ls 2>&1', $a, $b)`
* `fastcgi`就是执行 `cgi协议` 的,用于定义 `Nginx` 调用 `php` 时的数据格式,并减少对 `php.ini` 的读取
* `php-fpm`就是 `Nginx` 服务调用的程序,管理 `fastcgi` 进程
* `php-cli`就是 `shell` 命令行调用的程序
* `|||nochange|||`表示 `input` 输入框为空
* `htmlentities` / `htmlspecialchars` 可以过滤成 `html格式` , 防止 `XSS`
* `addslashes` 防止 `sql注入`
* 引用字符串方式: `Nowdoc` / `Heredoc`
* 中文简体格式
    `header('Content-type:text/html;charset=utf-8')`
* 解决textarea标签回车编码问题
    `str_replace(" "," ",str_replace("\n","<br/>",$data))`
* 单入口写上: `session_start()` 才可以使用 `$_SESSION`
* 跳转到#页面
    `header('Location:#')`
* php格式插入html遍历
    ``` php
    <?php foreach ($my_name as $x => $y){ ?>
        <li><?php echo $y['z'] ?></li>
    <?php } ?>
    ```
* 解决ajax跨域(CORS)
    ``` php
    header('Access-Control-Allow-Origin:demain.com'); //可以用*允许所有
    header('Content-Type:application/json; charset=utf-8');
    header('Access-Control-Allow-Methods: GET, POST, DELETE');
    header('Access-Control-Allow-Headers: Origin, X-Requested-With, Content-Type, Accept');
    ```
* 内容返回
    ``` php
    ob_start(); //开启echo缓存区
    echo "Hello"; //输出
    ob_end_flush(); //输出全部内容到浏览器,包括echo
    ```

| 函数 | 描述 |
| ---- | ---- |
| rand(x,y)                   | 在xy之间取随机数 |
| implode('xx','$my_name')        | 用xx隔开my_name中每个元素 |
| explode()                   | 把字符串打散为数组. |
| arry_merge()                | 打散混乱数组重新排列 |
| in_array($xy,$my_name)          | 数组$my_name是否含$xy |
| mb_substr($my_name,x,y,'utf-8') | 截取字符串,从第x个字符开始截取y个字符 |
| strtotime()                 | 将任何英文文本的日期或时间描述解析为 Unix 时间戳. |
| 字符串类型 |-- |
| addcslashes()               | 返回在指定的字符前添加反斜杠的字符串. |
| addslashes()                | 返回在预定义的字符前添加反斜杠的字符串. |
| chop()                      | 删除字符串右侧的空白字符或其他字符.. |
| chunk_split()               | 把字符串分割为一系列更小的部分. |
| html_entity_decode()        | 把 HTML 实体转换为字符. |
| htmlentities()              | 把字符转换为 HTML 实体. |
| htmlspecialchars_decode()   | 把一些预定义的 HTML 实体转换为字符. |
| htmlspecialchars()          | 把一些预定义的字符转换为 HTML 实体. |
| join()                      | implode() 的别名. |
| ltrim()                     | 移除字符串左侧的空白字符或其他字符. |
| md5()                       | 计算字符串的 MD5 散列. |
| number_format()             | 以千位分组来格式化数字. |
| parse_str()                 | 把查询字符串解析到变量中. |
| rtrim()                     | 移除字符串右侧的空白字符或其他字符. |
| sha1()                      | 计算字符串的 SHA-1 散列. |
| similar_text()              | 计算两个字符串的相似度. |
| sprintf()                   | 把格式化的字符串写入变量中. |
| str_pad()                   | 把字符串填充为新的长度. |
| str_repeat()                | 把字符串重复指定的次数. |
| str_replace()               | 替换字符串中的一些字符(对大小写敏感) |
| str_shuffle()               | 随机地打乱字符串中的所有字符. |
| str_split()                 | 把字符串分割到数组中. |
| strip_tags()                | 剥去字符串中的 HTML 和 PHP 标签. |
| stripcslashes()             | 删除由 addcslashes() 函数添加的反斜杠. |
| stripslashes()              | 删除由 addslashes() 函数添加的反斜杠. |
| stripos()                   | 返回字符串第一次出现的位置(大小写不敏感) |
| strpos()                    | 返回字符串第一次出现的位置(大小写敏感) |
| stristr()                   | 查找字符串第一次出现的位置(大小写不敏感) |
| strrchr()                   | 查找字符串最后一次出现. |
| strripos()                  | 查找字符串最后一次出现的位置(大小写不敏感) |
| strrpos()                   | 查找字符串最后一次出现的位置(大小写敏感) |
| strstr()                    | 查找字符串第一次出现(大小写敏感) |
| strchr()                    | 查找字符串第一次出现,strstr() 的别名 |
| strrev()                    | 反转字符串. |
| strlen()                    | 返回字符串的长度. |
| strtolower()                | 把字符串转换为小写字母. |
| strtoupper()                | 把字符串转换为大写字母. |
| strtr()                     | 转换字符串中特定的字符. |
| substr()                    | 返回字符串的一部分. |
| substr_count()              | 计算子串在字符串中出现的次数. |
| substr_replace()            | 把字符串的一部分替换为另一个字符串. |
| ucfirst()                   | 把字符串中的首字符转换为大写. |
| ucwords()                   | 把字符串中每个单词的首字符转换为大写. |
| 数组函数 | -- |
| array_change_key_case()     | 把数组中所有键更改为小写或大写. |
| array_chunk()               | 把一个数组分割为新的数组块. |
| array_column()              | 返回输入数组中某个单一列的值. |
| array_combine()             | 通过合并两个数组来创建一个新数组. |
| array_count_values()        | 用于统计数组中所有值出现的次数. |
| array_diff()                | 比较数组,返回差集(只比较键值) |
| array_fill()                | 用给定的键值填充数组. |
| array_fill_keys()           | 用指定键名的给定键值填充数组. |
| array_filter()              | 用回调函数过滤数组中的元素. |
| array_flip()                | 交换数组中的键和值. |
| array_intersect()           | 比较数组,返回交集(只比较键值) |
| array_key_exists()          | 检查指定的键名是否存在于数组中. |
| array_keys()                | 返回数组中所有的键名. |
| array_map()                 | 把数组中的每个值发送到用户自定义函数,返回新的值. |
| array_merge()               | 把一个或多个数组合并为一个数组. |
| array_multisort()           | 对多个数组或多维数组进行排序. |
| array_pad()                 | 用值将数组填补到指定长度. |
| array_pop()                 | 删除数组的最后一个元素(出栈) |
| array_push()                | 将一个或多个元素插入数组的末尾(入栈) |
| array_replace()             | 使用后面数组的值替换第一个数组的值. |
| array_reverse()             | 以相反的顺序返回数组. |
| array_search()              | 搜索数组中给定的值并返回键名. |
| array_shift()               | 删除数组中首个元素,并返回被删除元素的值. |
| array_slice()               | 返回数组中被选定的部分. |
| array_splice()              | 删除并替换数组中指定的元素. |
| array_sum()                 | 返回数组中值的和. |
| array_udiff()               | 比较数组,返回差集 |
| array_unique()              | 删除数组中的重复值.(集合的概念) |
| array_unshift()             | 在数组开头插入一个或多个元素. |
| array_values()              | 返回数组中所有的值. |
| arsort()                    | 对关联数组按照键值进行降序排序. |
| asort()                     | 对关联数组按照键值进行升序排序. |
| compact()                   | 创建包含变量名和它们的值的数组. |
| count()                     | 返回数组中元素的数目. |
| current()                   | 返回数组中的当前元素. |
| extract()                   | 从数组中将变量导入到当前的符号表. |
| in_array()                  | 检查数组中是否存在指定的值. |
| key()                       | 从关联数组中取得键名. |
| krsort()                    | 对数组按照键名逆向排序. |
| ksort()                     | 对数组按照键名排序. |
| list()                      | 把数组中的值赋给一些变量. |
| range()                     | 创建包含指定范围单元的数组. |
| rsort()                     | 对数组逆向排序. |
| shuffle()                   | 将数组打乱. |
| sizeof()                    | count() 的别名. |
| sort()                      | 对数组排序. |
| uasort()                    | 使用用户自定义的比较函数对数组中的键值进行排序. |
| uksort()                    | 使用用户自定义的比较函数对数组中的键名进行排序. |
| usort()                     | 使用用户自定义的比较函数对数组进行排序. |




## 前端

* 可以设置类似`cookie` , 它可永久保存
    `localStorage.setItem(name, val)`
* 死链
    `<a href="javascript:void(0);">`
* CSS加载JS样式
    `behavior:url()`
* 选择div内第二个li,类似js
    `div li:first-child+li{}`
* 选择form下类型为text的input标签
    `form input[type="text"]{}`
* 居中CSS背景图片
    `background-position:center`
* 居中HTML图片
    `<div align="center"><img src="">`
* 居中div
    `{top:0;left:0;bottom:0;right:0;margin:auto}/margin:0 auto;`
* 单行文字水平/垂直居中
    `text-decoration:center;/line-height:10px;`
* 多行文字垂直居中:div1>div2>p
    `div1{display:table;/*转换成表格*/} div2{display:table-cell;vertical-align:middle;}`
* li中间向两边分布
    `ul{text-align:center}li{display: inline-block;}`
* li水平分布,ul宽度平均分割给li
    `ul{display:flex}li{flex:1;text-align:center}`
* li水平分布,类似最左和最右浮动两边的效果
    `div{overflow:hidden;width:11px} div ul{width:12px}   li{width:5px;margin_right:1px}`
* 轮播图不能有滚动条
    `在ul父级加overflow:hidden;width:100%;position: relative;//position用于ul居中`
* 清除浮动,class是my_name就可以清除浮动
    `.my_name:after{content:""; display:block; height:0; visibility:hidden; clear:both;}`
* 鼠标悬浮图片放大
    `.div:hover img{transform: scale(1.03);transition: all 1s ease 0s;-webkit-transform: scale(1.03);-webkit-transform: all 1s ease 0s;}`
* iframe高度
    ``` js
    var e = window,
        a = 'inner',
        $header = $(".header");//如果页面有头部,则要减去
    if (!('innerWidth' in window)) {
        a = 'client';
        e = document.documentElement || document.body;
    }
    var height = e[a + 'Height'] - $header.outerHeight();
    $(".iframe").css({
        height: height,
        width: "100%"
    });
    ````


**... 以上**
