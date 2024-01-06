# Description
![Static Badge](https://img.shields.io/badge/BUPT-OAuth2.0-orange)
![Static Badge](https://img.shields.io/badge/Go-1.16.5-blue)
![Static Badge](https://img.shields.io/badge/ProVerif-2.02-red)



## 协议流程
![Design-OAuth_2_0_Sequence_Diagram.png](doc%2FDesign-OAuth_2_0_Sequence_Diagram.png)


BUPT 安全协议设计与分析课程附录，包括如下几个g部分：
- 协议实现
  - 使用**API工具**来实现大致认证、授权、保护资源访问流程，（上图 1-6、9-10流程）
  - 使用Go语言来实现刷新令牌步骤，（上图7-8流程）
- 协议安全性分析
  - 使用**ProVerif**来分析协议的安全性 [协议分析](./doc/oauth2simple.pv)


## 项目结构
- [doc](./doc): 项目文档
  - [oauth2simple.pv](./doc/oauth2simple.pv): 协议分析代码
  - [Design-OAuth_2_0_Sequence_Diagram.png](./doc/Design-OAuth_2_0_Sequence_Diagram.png): 协议流程图
  - [Design.puml](./doc/Design.puml): 协议流程图源文件
- [http](./http): http协议相关代码
- [middleware](./middleware): 中间件
- [test](./test): 测试代码
- [main.go](./main.go): 项目入口
## 协议实现
参照官方文档使用授权码模式实现了`Google OAuth2.0` 和 `Github OAuth2.0` 协议，具体测试接口可以查看：[接口实现](https://apifox.com/apidoc/shared-532c54cd-e26d-4e22-a34d-cc47ac473efa/api-139404017)