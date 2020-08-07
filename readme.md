简单的说就是一个前后端分离的web版的资源管理器。

目标支持平台：win,linux,mac.  [mac未经过测试，理论上类unix系统都能正常兼容]

In short, it is a resource manager based on web with the front and back ends separated.

language: vue,go

Unfinished...

已知BUG：打开txt可能乱码

使用：
  从源码：前后端分离项目
      前端基于vue，需要node环境，一次执行npm install，npm run serve
      后端基于golang，使用gin框架，请直接 go build 然后运行
  直接使用：
        下载release即可

