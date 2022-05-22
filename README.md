# openservice

#### 介绍
一个openstore的管理程序，主要用于openstore、应用服务的管理。包含open-center open-node open-service三大组件。

    1. open-center: 部署在一个中心节点上，主要包含mysql数据库服务、restful接口服务、open-node节点管理。
    1. open-node: 部署在每个分布式节点上，主要用于响应open-center的管理指令，包含cdm-instance、pool、set、block的管理功能。
    1. open-service: 部署在每个分布式节点上，主要用于响应open-center的管理指令，包含VMware、阿里云等虚拟化平台，oracle，mysql等数据库，ftp/http等应用服务相关数据源收集插件和恢复插件的管理功能。

#### 软件架构
软件架构说明

接口说明：

1. 查询存储列表

    请求参数：
    GET /openstore_v1/pools
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|

    返回数据：
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---| 
    | pools |  array of pool | ****  |  创建存储池类型 |  
      
2. 查询存储详细信息
    
    GET /openstore_v1/pools/{pool_id}
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   
 
    返回数据：
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---| 
    | pool |  pool | ****  |  创建存储池类型 |  

3. 创建存储
    
    PUT /openstore_v1/pools
    请求参数：
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|
    |  type | string  | 是  | local  | 存储池类型,取值：local,remote,s3,dedup |
    |  host | string  | 否  | 192.168.2.1  | 存储池地址  |
    |  port | int | 否  | 9527  | 存储池端口 |
    | s3_scheme| string | 否 | oss |对象存储类型，取值：aws,oss,cos,other |
    | bucket| string| 否 | s3_storage |对象存储桶 |
    | region| string| 否 | cn-north-1 |对象存储region |
    | access_key| string| 否 | PF47C3UA95UUGTTA88FM |对象存储访问key |
    | secret_key| string| 否 | fPzXzJwGc2y1VzWZ025 |对象存储访问密钥 |
    | secret_key| string| 否 | fPzXzJwGc2y1VzWZ025 |对象存储访问密钥 |
    | https| bool| 否 | true|是否启用HTTPS连接 |

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---| 
    | pool-id |  string | pool-123456789  |  创建存储池类型 |

4. 删除存储
    
    DELETE /openstore_v1/pools
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---| 

5. 查询磁盘集列表
    
    GET /openstore_v1/sets
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   
    |filter|string|否|oracle|null|

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---| 

6. 查询磁盘集详细信息
    
    GET /openstore_v1/sets/{set_id}
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

7. 创建磁盘集
    
    PUT /openstore_v1/sets
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

8. 删除磁盘集
    
    DELETE /openstore_v1/sets/{set_id}
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

9. 查询磁盘
    
    GET /openstore_v1/blocks
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

10. 查询磁盘详细信息
    
    GET /openstore_v1/blocks/{block_id}
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

11. 创建磁盘
    
    PUT /openstore_v1/blocks
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

12. 删除磁盘
    
    DELETE /openstore_v1/blocks/{block_id}
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

13. 上传磁盘
    
    POST /openstore_v1/blocks/{block_id}?action=upload
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

14. 复制磁盘
    
    POST /openstore_v1/blocks/{block_id}?action=copy
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

15. 移动磁盘
    
    POST /openstore_v1/blocks/{block_id}?action=move
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

16. 恢复磁盘
    
    POST /openstore_v1/blocks/{block_id}?action=restore
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

17. 查询cdm实例列表
    
    GET /openstore_v1/cdm_instances
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

18. 查询cdm实例详细信息
    
    GET /openstore_v1/cdm_instances/{ins-id}
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

19. 加载cdm实例
    
    PUT /openstore_v1/cdm_instances
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

20. 移除cdm实例
    
    DELETE /openstore_v1/cdm_instances/{ins-id}
    | 名称   |  类型 |  必选 | 示例  |  描述值 |
    |---|---|---|---|---|   

    返回数据
    | 名称   |  类型 | 示例  |  描述值 |   
    |---|---|---|---|

#### 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
