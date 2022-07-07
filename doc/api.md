# ncloud_pcdn api
## 1 历史记录
|  版本   | 修改  | 作者  |     时间     |
|:-----:|:---:|:---:|:----------:|
| 1.0.0 | 创建  | 包月兴 | 2021/12/31 |

## 2 参数说明
### 2.1 系统(header)参数
| 序号  |      参数       | 选项  |  数据类型  |        说明        |
|:---:|:-------------:|:---:|:------:|:----------------:|
|  1  | Content-Type  |  Y  | string | application/json |
|  2  | Authorization |  Y  | string |      鉴权字符串       |

### 2.2 sign 鉴权说明
#### 2.2.1 变量
```
ak string类型, 授权ak[########]
sk string类型 授权sk[########]
path string类型 接口path
sign string类型 签名
t string类型 时间戳字符串
authorization string类型 鉴权字符串
```
#### 2.2.2 签名加密方法
```
加密算法: HmacSHA1
加密字符串: {ak}{path}{t}
eg:
ak = "abc123"
path = "/refresh"
t = 1568953410

加密前字符串: "abc123/refresh1234561568953410"
加密秘钥: sk
加密后字符串: SIGNED = b41ec844042b68cd0093d9b54018350231fa3087

```


#### 2.2.3 鉴权计算方法
```
加密算法: base64
加密字符串: ak={ak}&&sign={sign}&t={t}
eg:
ak = "abc123"
sign= "b41ec844042b68cd0093d9b54018350231fa3087"
t = 1568953410

加密前字符串: "ak=abc123&sign=b41ec844042b68cd0093d9b54018350231fa3087&t=1568953410"
authorization = base64("ak=abc123&sign=b41ec844042b68cd0093d9b54018350231fa3087&t=1568953410")

```
#### 2.2.4 请求头
```
Content-Type = "application/json"
Authorization = {authorization}
```
## 3 接口列表
### 3.1 创建刷新预热任务
```
1.path:/v1/cdn/refresh
2 协议:POST 
3 系统参数:Content-Type Authorization
```
##### 业务参数
|    参数    |   数据类型   | 选项  |   说明   |
|:--------:|:--------:|:---:|:------:|
| taskType |  string  |  Y  | 刷新预热类型 |
|  domain  |  string  |  Y  |   域名   |
|  files   | []string |  Y  |  文件列表  |
```
备注：
刷新预热类型,(url:url刷新,dir:目录刷新,prefetch:预热)
目前支持单域名
```
##### 响应数据
|    参数    |  数据类型  | 选项  |       说明        |
|:--------:|:------:|:---:|:---------------:|
|    ok    |  bool  |  Y  | true 成功 false失败 |
|   msg    | string |  Y  |      信息输出       |
|   data   | objetc |  Y  |       任务        |
| --taskId | string |  Y  |      任务id       |
|  --date  | string |  Y  |       日期        |

##### 响应数据返回实例
```
{
    "ok": true,
    "msg": "请求成功",
    "data": {
            "taskId": "9328ea7c-66a8-41ff-b636-8398f798557e",
            "date": "20220330",
    }
}
```
### 3.2 根据taskId查询刷新预热任务
```
1.path:/v1/cdn/refreshTask
2 协议:POST 
3 系统参数:Content-Type Authorization
```
##### 业务参数
|   参数   |  数据类型  | 选项  |  说明  |
|:------:|:------:|:---:|:----:|
| taskId | string |  Y  | 任务id |
|  date  | string |  Y  |  日期  |
##### 响应数据
|      参数       |   数据类型   | 选项  |       说明        |
|:-------------:|:--------:|:---:|:---------------:|
|      ok       |   bool   |  Y  | true 成功 false失败 |
|      msg      |  string  |  Y  |      信息输出       |
|     data      | []objetc |  Y  |       任务        |
|     --url     |  string  |  Y  |       url       |
|   --status    |   int    |  Y  |       状态        |
| --create_time |   int    |  Y  |      创建时间       |

##### 响应数据返回实例
```
{
    "ok": true,
    "msg": "请求成功",
    "data": [
            {
            "url": "http://download-v1.xyuncloud.com/site/15847292.mp4.f30.mp4",
            "status": 1,
            "create_time":1648609894
            },
    }
}
```

### 3.3 查询cdn日志
```
1.path:/v1/cdn/downLog
2 协议:POST 
3 系统参数:Content-Type Authorization
```
##### 业务参数
|     参数     | 选项  |  数据类型  |     说明     |             备注             |
|:----------:|:---:|:------:|:----------:|:--------------------------:|
|   domain   |  Y  | string | 用户ping权限域名 ||
| startTime  |  Y  | string |    起始时间    | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|  endTime   |  Y  | string |    结束时间    | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|  Interval  |  Y  | string |   日志时间间隔   | hour，day，fiveMin 目前只支持hour |
|  pageSize  |  Y  |  int   |  分页 每页大小   ||
| pageNumber |  Y  |  int   |   分页 页数    |            大于1             |
##### 响应数据
|      参数      |   数据类型   | 选项  |       说明        |  备注  |
|:------------:|:--------:|:---:|:---------------:|:----:|
|      ok      |   bool   |  Y  | true 成功 false失败 ||
|     msg      |  string  |  Y  |      信息输出       ||
|     data     |  objetc  |  Y  |       --        ||
|   --total    |   int    |  Y  |       总数        ||
|  --pageSize  |   int    |  Y  |     分页 每页大小     ||
| --pageNumber |   int    |  Y  |      分页 页数      ||
|    --urls    | []objetc |  Y  |        -        ||
|   ----url    |  string  |  Y  |     日志下载地址      ||
| ----fileName |  string  |  Y  |     日志文件名称      ||
|   ----size   |   int    |  Y  |     日志文件大小      | 单位 B |

##### 响应数据返回正确实例
```
{
    "data": {
        "total": 4225,
        "pageSize": 2,
        "pageNumber": 1,
        "urls": [
            {
                "url": "https://oss1.y6.xyuncloud.com/pcdn/download-v1.xyuncloud.com,20220401103001.zip",
                "fileName": "pcdn/download-v1.xyuncloud.com,20220401103001.zip",
                "size": 254,
                "domain": "download-v1.xyuncloud.com"
            },
            {
                "url": "https://oss1.y6.xyuncloud.com/pcdn/download-v1.xyuncloud.com,20220401093001.zip",
                "fileName": "pcdn/download-v1.xyuncloud.com,20220401093001.zip",
                "size": 254,
                "domain": "download-v1.xyuncloud.com"
            }
        ]
    },
    "msg": "下载成功",
    "ok": true
}
```
##### 响应数据返回错误实例
```
{
    "data": {
        "total": 0,
        "pageSize": 0,
        "pageNumber": 0,
        "urls": []
    },
    "msg": "域名未配置源,请联系工作人员",
    "ok": false
}
```
### 3.4 查询用户ping权限下所有域名
```
1.path:/v1/cdn/domain
2 协议:POST 
3 系统参数:Content-Type Authorization
4 无业务参数
```
##### 响应数据
|  参数  |   数据类型   | 选项  |       说明        | 备注  |
|:----:|:--------:|:---:|:---------------:|:---:|
|  ok  |   bool   |  Y  | true 成功 false失败 ||
| msg  |  string  |  Y  |      信息输出       ||
| data | []string |  Y  |  用户ping权限下所有域名  ||
##### 响应数据返回正确实例
```
{
    "data": [
        "download-v1.xyuncloud.com"
    ],
    "msg": "请求成功",
    "ok": true
}
```
##### 响应数据返回错误实例
```
{
    "data": [],
    "msg": "域名未配置,请联系工作人员",
    "ok": false
}
```
### 3.5 查询所有运营商
```
1.path:/v1/cdn/isp
2 协议:POST 
3 系统参数:Content-Type Authorization
4 无业务参数
```
##### 响应数据
|  参数  |   数据类型   | 选项  |       说明        | 备注  |
|:----:|:--------:|:---:|:---------------:|:---:|
|  ok  |   bool   |  Y  | true 成功 false失败 ||
| msg  |  string  |  Y  |      信息输出       ||
| data | []string |  Y  |      运营商列表      ||
##### 响应数据返回正确实例
```
{
    "data": [
        "移动",
        "联通",
        "电信"
    ],
    "msg": "请求成功",
    "ok": true
}
```
### 3.6 查询所有地区
```
1.path:/v1/cdn/region
2 协议:POST 
3 系统参数:Content-Type Authorization
4 无业务参数
```
##### 响应数据
|  参数  |   数据类型   | 选项  |       说明        | 备注  |
|:----:|:--------:|:---:|:---------------:|:---:|
|  ok  |   bool   |  Y  | true 成功 false失败 ||
| msg  |  string  |  Y  |      信息输出       ||
| data | []string |  Y  |      地区列表       ||
##### 响应数据返回正确实例
```
{
    "data": [
        "安徽",
        "北京",
        "重庆"
    ],
    "msg": "请求成功",
    "ok": true
}
```

### 3.7 查询域名带宽统计
```
1.path:/v1/cdn/statisticsData/domain
2 协议:POST 
3 系统参数:Content-Type Authorization
```
#### 业务参数
|    参数     | 选项  |   数据类型   |      说明      |             备注             |
|:---------:|:---:|:--------:|:------------:|:--------------------------:|
|  domains  |  Y  | []string | 用户ping权限域名集合 |       空数组 ping权限所有域名       |
| startTime |  Y  |  string  |     起始时间     | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|  endTime  |  Y  |  string  |     结束时间     | 北京时间 [yyyy-mm-dd HH:mm:ss] |
#### 发送业务参数范例
```
{
    "domains":[],
    "startTime":"2022-04-01 10:20:00",
    "endTime":"2022-04-01 10:35:00"
}
```

##### 响应数据
|      参数       |   数据类型   |   选项   |       说明        |             备注             |
|:-------------:|:--------:|:------:|:---------------:|:--------------------------:|
|      ok       |   bool   |   Y    | true 成功 false失败 ||
|      msg      |  string  |   Y    |      信息输出       ||
|     data      | []object |   Y    |                 ||
|  --startTime  |    Y     | string |      起始时间       | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|   --endTime   |    Y     | string |      结束时间       | 北京时间 [yyyy-mm-dd HH:mm:ss] |
| --statistics  |    Y     | object |                 |                            |
| ----startTime |    Y     | string |      起始时间       | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|  ----endTime  |    Y     | string |      结束时间       | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|  ----domain   |    Y     | string |       域名        ||
|    ----isp    |    Y     | string |       运营商       ||
|  ----region   |    Y     | string |       地区        ||
| ----bandwidth |    Y     |  int   |       带宽        |            单位 B            |

##### 响应数据返回正确实例
```
{
    "data": [
        {
            "startTime": "2022-04-01 10:25:000",
            "endTime": "2022-04-01 10:30:00",
            "statistics": [
                {
                    "startTime": "2022-04-01 10:25:000",
                    "endTime": "2022-04-01 10:30:00",
                    "domain": "download-v1.xyuncloud.com",
                    "isp": "",
                    "region": "",
                    "bandwidth": 213331147
                }
            ]
        }
    ],
    "msg": "请求成功",
    "ok": true
}
```
##### 响应数据返回错误实例
```
{
    "data": [
        {
            "startTime": "",
            "endTime": "",
            "statistics": []
        }
    ],
    "msg": "非法域名请求",
    "ok": false
}
```
### 3.8 查询运营商带宽统计
```
1.path:/v1/cdn/statisticsData/isp
2 协议:POST 
3 系统参数:Content-Type Authorization
```
#### 业务参数
|    参数     |  选项  |   数据类型   |      说明      | 备注 |
|:---------:| :----: |:--------:|:------------:|:--:|
|  domains  |Y| []string | 用户ping权限域名集合 |  空数组 ping权限所有域名  |
|  isp  |Y| []string | 用户ping权限运营商集合 |  空数组 ping权限所有运营商  |
| startTime |Y|  string  |     起始时间     | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|  endTime  |Y|  string  |     结束时间     | 北京时间 [yyyy-mm-dd HH:mm:ss] |
#### 发送业务参数范例
```
{
    "domains":[],
    "isp":[],
    "startTime":"2022-04-01 10:20:00",
    "endTime":"2022-04-01 10:35:00"
}
```

##### 响应数据
|      参数       |   数据类型   |   选项   |       说明        |             备注             |
|:-------------:|:--------:|:------:|:---------------:|:--------------------------:|
|      ok       |   bool   |   Y    | true 成功 false失败 ||
|      msg      |  string  |   Y    |      信息输出       ||
|     data      | []object |   Y    |                 ||
|  --startTime  |    Y     | string |      起始时间       | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|   --endTime   |    Y     | string |      结束时间       | 北京时间 [yyyy-mm-dd HH:mm:ss] |
| --statistics  |    Y     | object |                 |                            |
| ----startTime |    Y     | string |      起始时间       | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|  ----endTime  |    Y     | string |      结束时间       | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|  ----domain   |    Y     | string |       域名        ||
|    ----isp    |    Y     | string |       运营商       ||
|  ----region   |    Y     | string |       地区        ||
| ----bandwidth |    Y     |  int   |       带宽        |            单位 B            |

##### 响应数据返回正确实例
```
{
    "data": [
        {
            "startTime": "2022-04-01 10:20:000",
            "endTime": "2022-04-01 10:25:00",
            "statistics": [
                {
                    "startTime": "2022-04-01 10:20:000",
                    "endTime": "2022-04-01 10:25:00",
                    "domain": "download-v1.xyuncloud.com",
                    "isp": "联通",
                    "region": "",
                    "bandwidth": 31376999
                },
                {
                    "startTime": "2022-04-01 10:20:000",
                    "endTime": "2022-04-01 10:25:00",
                    "domain": "download-v1.xyuncloud.com",
                    "isp": "电信",
                    "region": "",
                    "bandwidth": 122029671
                }
            ]
        },
        {
            "startTime": "2022-04-01 10:25:000",
            "endTime": "2022-04-01 10:30:00",
            "statistics": [
                {
                    "startTime": "2022-04-01 10:25:000",
                    "endTime": "2022-04-01 10:30:00",
                    "domain": "download-v1.xyuncloud.com",
                    "isp": "联通",
                    "region": "",
                    "bandwidth": 57617613
                },
                {
                    "startTime": "2022-04-01 10:25:000",
                    "endTime": "2022-04-01 10:30:00",
                    "domain": "download-v1.xyuncloud.com",
                    "isp": "电信",
                    "region": "",
                    "bandwidth": 155713534
                }
            ]
        }
    ],
    "msg": "请求成功",
    "ok": true
}
```
##### 响应数据返回错误实例
```
{
    "data": [
        {
            "startTime": "",
            "endTime": "",
            "statistics": []
        }
    ],
    "msg": "非法域名请求",
    "ok": false
}
```
### 3.9 查询运营商带宽统计
```
1.path:/v1/cdn/statisticsData/region
2 协议:POST 
3 系统参数:Content-Type Authorization
```
#### 业务参数
|    参数     |  选项  |   数据类型   |      说明      | 备注 |
|:---------:| :----: |:--------:|:------------:|:--:|
|  domains  |Y| []string | 用户ping权限域名集合 |  空数组 ping权限所有域名  |
|  isp  |Y| []string | 用户ping权限运营商集合 |  空数组 ping权限所有运营商  |
|  regions  |Y| []string | 用户ping权限地区集合 |  空数组 ping权限所有地区  |
| startTime |Y|  string  |     起始时间     | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|  endTime  |Y|  string  |     结束时间     | 北京时间 [yyyy-mm-dd HH:mm:ss] |

#### 发送业务参数范例
```
{
    "domains":[],
    "isp":[],
    "regions":[],
    "startTime":"2022-04-01 10:20:00",
    "endTime":"2022-04-01 10:35:00"
}
```


##### 响应数据
|      参数       |   数据类型   |   选项   |       说明        |             备注             |
|:-------------:|:--------:|:------:|:---------------:|:--------------------------:|
|      ok       |   bool   |   Y    | true 成功 false失败 ||
|      msg      |  string  |   Y    |      信息输出       ||
|     data      | []object |   Y    |                 ||
|  --startTime  |    Y     | string |      起始时间       | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|   --endTime   |    Y     | string |      结束时间       | 北京时间 [yyyy-mm-dd HH:mm:ss] |
| --statistics  |    Y     | object |                 |                            |
| ----startTime |    Y     | string |      起始时间       | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|  ----endTime  |    Y     | string |      结束时间       | 北京时间 [yyyy-mm-dd HH:mm:ss] |
|  ----domain   |    Y     | string |       域名        ||
|    ----isp    |    Y     | string |       运营商       ||
|  ----region   |    Y     | string |       地区        ||
| ----bandwidth |    Y     |  int   |       带宽        |            单位 B            |

##### 响应数据返回正确实例
```
{
    "data": [
        {
            "startTime": "2022-04-01 10:20:000",
            "endTime": "2022-04-01 10:25:00",
            "statistics": [
                {
                    "startTime": "2022-04-01 10:20:000",
                    "endTime": "2022-04-01 10:25:00",
                    "domain": "download-v1.xyuncloud.com",
                    "isp": "联通",
                    "region": "广东",
                    "bandwidth": 31376999
                },
                {
                    "startTime": "2022-04-01 10:20:000",
                    "endTime": "2022-04-01 10:25:00",
                    "domain": "download-v1.xyuncloud.com",
                    "isp": "电信",
                    "region": "广东",
                    "bandwidth": 122029671
                }
            ]
        },
        {
            "startTime": "2022-04-01 10:25:000",
            "endTime": "2022-04-01 10:30:00",
            "statistics": [
                {
                    "startTime": "2022-04-01 10:25:000",
                    "endTime": "2022-04-01 10:30:00",
                    "domain": "download-v1.xyuncloud.com",
                    "isp": "联通",
                    "region": "广东",
                    "bandwidth": 57617613
                },
                {
                    "startTime": "2022-04-01 10:25:000",
                    "endTime": "2022-04-01 10:30:00",
                    "domain": "download-v1.xyuncloud.com",
                    "isp": "电信",
                    "region": "广东",
                    "bandwidth": 155713534
                }
            ]
        }
    ],
    "msg": "请求成功",
    "ok": true
}
```
##### 响应数据返回错误实例
```
{
    "data": [
        {
            "startTime": "",
            "endTime": "",
            "statistics": []
        }
    ],
    "msg": "非法域名请求",
    "ok": false
}
```
