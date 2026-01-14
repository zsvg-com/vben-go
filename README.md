<div align="center"><h1 align="center">vben-python</h1></div>
<div align="center"><h3 align="center">一个开箱即用的通用开发平台GO语言版</h3></div>

# 🍿 相关地址

* PC端体验地址 ：[http://8.153.168.178/](http://8.153.168.178/)
* 移动端体验地址 ：[http://8.153.168.178/app/](http://8.153.168.178/unibest/)
* 联系方式（加微信后可找小狐狸拉进交流群）：

![输入图片说明](https://gitee.com/vben/vben-app/raw/master/docs/wx.jpg "微信联系方式")


# 🍟 概述
* 前端采用Vue3、Element Plus，基于vben5前端项目修改。
* 后端参考学习的`乐观开发者`的 [https://gitee.com/OptimisticDevelopers/Ruoyi-Go](https://gitee.com/OptimisticDevelopers/Ruoyi-Go)
* 后端采用go & gin & Jwt & gorm & mysql & copier & redis & gin-cache && xxl-job。
* 权限认证使用OAuth2 & Jwt，支持多终端认证系统。
* 支持加载动态权限菜单，多方式轻松权限控制。


# ⚡ 快速启动

```bash
# 拉取依赖
go mod tidy

# 配置环境
在config.yaml文件中配置开发环境的数据库和redis

# 运行sql文件
1.新建数据库vben-go(默认，可修改)
2.使用命令或数据库连接工具运行根目录下的vben-go.sql

# 运行后端
go run main.go
```


# 🍖 详细功能

1. 主控面板：控制台页面，可进行工作台，分析页，统计等功能的展示。
2. 部门管理：部门维护，支持多层级结构的树形结构。
3. 用户管理：用户维护，可设置用户部门，岗位，群组，职务，角色，数据权限等。
4. 岗位管理：岗位维护，岗位可作为用户的一个标签，岗位也可与权限等其他功能挂钩。
5. 群组管理：群组维护，群组可设置部门，用户，岗位，用于更广泛的权限设置。
6. 菜单管理：菜单目录，菜单，和按钮的维护是权限控制的基本单位。
7. 角色管理：角色绑定菜单后，可限制相关角色的人员登录系统的功能范围。
8. 字典管理：系统内各种枚举类型的维护。
9. 访问日志：用户的登录和退出日志的查看和管理。
10. 操作日志：用户的操作业务的日志的查看和管理。
11. 定时任务：定时任务的维护，通过cron表达式控制任务的执行频率。
12. 流程引擎：流程图展示，支持驳回、转办、废弃，跳转等功能
13. 在线表单：在线表单设计，配合流程可以实现表单数据流转
14. 消息机制：待办待阅功能，联通钉钉与企业微信接口
15. 代码生成：在线配置，一键生成前后端代码

# ⚡ 近期计划

- [ ] 完善文档
- [ ] 完善工作流

# 💐 特别鸣谢
- 👉 乐观开发者：[https://gitee.com/OptimisticDevelopers/Ruoyi-Go](https://gitee.com/OptimisticDevelopers/Ruoyi-Go)
- 👉 Vben-Admin：[https://vvbin.cn/doc-next](https://vvbin.cn/doc-next)
- 👉 vxe-table：[https://gitee.com/xuliangzhan_admin/vxe-table](https://gitee.com/xuliangzhan_admin/vxe-table)
- 👉 VForm：[https://www.vform666.com](https://www.vform666.com/)

```
如果对您有帮助，点击右上角⭐Star⭐关注 ，感谢支持开源！
```