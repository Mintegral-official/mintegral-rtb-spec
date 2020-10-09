# Mintegral RTB Specification ver 2.5

This specification complies with OpenRTB Version 2.5.

See [IAB OpenRTB API Specification Version 2.5](https://www.iab.com/wp-content/uploads/2016/03/OpenRTB-API-Specification-Version-2-5-FINAL.pdf) for details.

### Attention

This specification may be incomplete. If you find any inadequacy or encounter confusion, please contact the person in charge: [developer@mintegral.com](mailto:developer@mintegral.com).

Moreover, this specification does not contain description of general RTB protocol (e.g., OpenRTB).

## Table of Contents

* [1. cookie sync](#1-cookie-sync)
* [2. Request/Response encoding](#2-requestresponse-encoding)
  * [a. Request encoding](#a-request-encoding)
  * [b. Response encoding](#b-response-encoding)
* [3. Request specification](#3-request-specification)
  * [a. Endpoint URL](#a-endpoint-url)
  * [b. OpenRTB Version HTTP Header](#b-openrtb-version-http-header)
  * [c. Bid request parameters](#c-bid-request-parameters)
    * [Bid Request Object (Top Level)](#bid-request-object-top-level)
    * [imp Object](#imp-object)
    * [site Object](#site-object)
    * [app Object](#app-object)
    * [publisher Object](#publisher-object)
    * [user Object](#user-object)
    * [device Object](#device-object)
    * [geo Object](#geo-object)
    * [banner Object](#banner-object)
    * [format Object](#format-object)
    * [video Object](#video-object)
    * [audio Object](#audio-object)
    * [native Object](#native-object)
* [4. Response specification](#4-response-specification)
  * [a. Bid response parameters](#a-bid-response-parameters)
    * [Bid Response Object (Top Level)](#bid-response-object-top-level)
    * [seatbid Object](#seatbid-object)
    * [bid Object](#bid-object)
    * [skadn Response Object](#skadn-response-object)
  * [b. impression/click beacon](#b-impressionclick-beacon)
    * [endpoint for impression beacon](#endpoint-for-impression-beacon)
    * [fluct transmits according to the following conditions (imp beacon):](#fluct-transmits-according-to-the-following-conditions-imp-beacon)
    * [endpoint for click beacon](#endpoint-for-click-beacon)
    * [fluct transmits according to the following conditions (click beacon):](#fluct-transmits-according-to-the-following-conditions-click-beacon)
  * [c. Macro substitution](#c-macro-substitution)
  * [d. Click Measuring](#d-click-measuring)
* [5. Code table](#5-code-table)
* [6. Bid Request/Response Examples](#6-bid-requestresponse-examples)



## 1. cookie sync

Usually the following sync will be performed, however we are also able to provide specified flows separately.

1. SSP delivers sync image tags of DSP to ad inventories.
2. DSP redirects to SSP’s sync URL when received access to sync tag.
3. DSP adds an ID parameter to sync URL as a notification to SSP.

DSP should prepare sync tags for SSP in advance. (DSP reponsibility)

Currently we only support image sync tags and https requests.

Below is an example of sync URLs of fluct. HTTPs sync is also supported.

    https://cs.adingo.jp/sync/?from=[DSP_NAME]&id=[DSP_USER_ID]

The DSP_NAME parameter will be provided by SSP beforehand (SSP responsibility), please contact SSP for that.

The expiry date of sync is default to 30 days, however it can also be customized to any length. Below is an example of a sync with 90 days of expiry date.

    https://cs.adingo.jp/sync/?from=your_dsp&id=XXXXXX&expire=90

## 2. Request/Response encoding

To minimize HTTP traffic exchanged between DSP and Fluct, it is **recommended** to have **both** bid request and response bodies **compressed**.

### a. Request encoding

Bid requests can be gzip-compressed.
If DSPs wish to receive gzip-compressed bid requests, contact an alliance representative.

Following header is added to a request when its body is gzip-compressed:

    Content-Encoding: gzip

### b. Response encoding

A response body can be compressed when its request has a header similar to:

    Accept-Encoding: gzip

When returning a compressed response body, such response should have a header similar to:

    Content-Encoding: gzip

## 3. Request specification

### a. Endpoint URL

The endpoint URL which will be used by bid request should be passed to SSP. (DSP reponsibility)

SSP conducts bid request with POST method to the specified endpoint URL.

For the optimal performance, DSP should enable HTTP keep-alive whenever possible.

https request is also supported.

Please let us know if we need to set the specific Port.

### b. OpenRTB Version HTTP Header

The OpenRTB Version of the request will be passed in bid request HTTP header as:

    X-OpenRTB-Version: 2.5

The version is in format `<major>.<minor>`, and patch version is not included.

Fluct will follow backward-compatible OpenRTB 2 minor version updates, and the version in HTTP header will be updated, accordingly.

### c. Bid request parameters

Serialize format: JSON only.


#### Bid Request Object (Top Level)

|  字段    | 类型|   描述|
|---|---|---|
 | id |      string; required |         MintegralADX提供的bidrequest唯一标识|
 | imp| array of imp objects; required  |一次请求包含一个或以上的impressionobject；每个impressionobject代表一个广告位的请求；具体见imp参数说明 |
 | site | site object  |     定义site信息|
 | app     | app object   |   定义app信息；具体见[Object App参数说明](#object-app)|
 | device  | device object; required   |   定义设备信息；具体见[Object Device参数说明](#object-device)|
 | user |user object  |    定义用户信息|
 | ~~test~~| integer |    是否为测试模式|
 | at      |   integer; required|   拍卖结算类型，值为1表示一价结算，值为2表示二价结算；|
 | tmax    |   integer |    DSP广告返回超时时限以ms为单位 |
 | ~~wseat~~|  array of strings|    买方席位白名单|
 | ~~bseat~~|  array of strings|    买方席位黑名单|
 | ~~allimps~~|   integer  |     本次展示机会是否覆盖上下文所有广告展示机会|
 | ~~cur~~    | array of strings|   返回允许的货币类型|
 | ~~wlang~~  | array of strings|   素材语言白名单|
 | bcat    | array of strings   | 广告主IAB category黑名单|
 | badv    | array of strings   |    广告主域名黑名单|
 | bapp    | array of strings   |     App类广告包名黑名单;  安卓示例 com.amazon.mShop; ios示例：*907394059*
 | ~~source~~| source object |   ~~流量来源信息~~ |
 | regs    | regs object  |     政策法规要求；具体见[Object Regs](#object-regs) |

 * Only one of the "site" and “app” sections will be necessary.


 #### imp Object

| 字段	| 类型	| 描述 |
| ---|---|---|
|id	|string; required |请求中唯一标识本次出售展示的标识；Mintegral ADX 一次请求只出售一个展示|
|~~metric~~	|array of metric object |	返回允许的货币类型|
|banner	|banner object;banner流量必传|	见 Object Banner|
|video	|video object;video 流量必传|	见 Object Video|
|~~audio~~	|audio object;audio流量必传|	见 Object Audio|
|native	|native object;native流量必传|	见 Object Native|
|~~pmp~~ |pmp object |	私有市场交易|
|~~displaymanager~~	|string |	媒体合作伙伴的名字，|
|~~displaymanagerver~~	|string |	媒体合作伙伴的版本|
|instl	|integer |	0-非插屏；1-广告为全屏或者插屏|
|tagid	|string; required|广告位的唯一标识|
|bidfloor|	float; required	|本次展示的cpm底价|
|bidfloorcur| string; required |DSP 竞价货币单位；默认为美元 USD|
|secure | integer; required|	标识展示请求是否需要使用HTTPS，默认0, 0-标识不需要使用安全链路； 1-标识需要使用安全链路|
|~~iframebuster~~|array of strings |特定交易支持的iframe buster的名字数组|
|~~exp~~|	integer|	建议广告展示有效时间窗口|


#### banner Object

| 字段	| 类型 	| 描述 |
| ---|---|---|
|format	|array of format object|	若广告位可兼容多组宽高，则通过该字段列举可兼容的多组宽高|
|w	|integer |	广告位宽度，单位为设备独立像素|
|h	|integer |	广告位高度，单位为设备独立像素|
|id	|string  |banner 对象的唯一标识；在一个 Ad 中包含 Banner 与 Video 的时候使用|
|btype|	array of integer |	限制不可投放的banner类型，枚举值参考附录Banner Ad Types|
|battr|	array of integer |	限制不可投放的物料属性，枚举值参考附录Creative Attributes|
|pos|	integer |	广告位置，枚举值参考附录Ad Position|
|mimes|	array of strings;required|	支持的内容 mime-type；包括但不仅限于“application/javascript”, “image/jpg”, and “image/gif”|
|~~topfram~~|	integer |	banner是在顶层frame中而不是iframe中，0-不是; 1-是|
|~~expdir~~|	array of integer|	banner可以扩展的方向，参考表附录Expandable Direction|
|api| array of integer |	该次展示可支持的 API 框架；枚举值释义参考附录API Frameworks；该字段缺省表示所有枚举值均不支持|
|~~vcm~~|	integer	|	当banner object作为video objecthttp://dev.mintegral.com/doc/adx/cn/#frc12_1的伴随广告时，标识伴随广告的渲染模式；0-concurrent, 1-endcard|

#### format Object

| 字段	| 类型 	| 描述 |
| ---|---|---|
|w	|integer;required |	广告位宽度，单位为设备独立像素|
|h	|integer;required |	广告位高度，单位为设备独立像素|


#### video Object

| 字段	| 类型 	| 描述 |
|---|---|---|
|mimes|array of strings; required |支持的内容 mime-type；包括但不仅限于“video/mp4”；|
|minduration	|integer |	最小的视频长度， 以秒为单位；|
|maxduration	|integer |	最大的视频长度， 以秒为单位；|
|protocols	|array of integer |	支持的视频竞价响应协议；枚举值参考附录Protocols|
|w|	integer|视频播放器的宽度，单位为像素|
|h|	integer	|视频播放器的高度，单位为像素|
|~~startdelay~~	|integer |视频前，中及之后的广告位中视频广告的启动延时，以秒为单位,枚举值参考附录Start Delay|
|~~placement~~	|integer |广告位类型；枚举值参考Video Placement Types；|
|linearity	|integer |展示是否是线性的，1-线性；2-非线性；|
|skip	|integer |传 0 表示不支持用户 skip，传 1 表示支持用户 skip；|
|~~skipmin~~|	integer	|最小可跳过的广告总时长；超过该时长的广告才支持跳过；|
|~~skipafter~~	|integer |广告播放多少秒后可跳过；|
|sequence	|integer |如果在同一个竞价请求中提供了多个展示， 则需要考虑多个物料传输的顺序|
|battr	|array of integer |限制不可投放的物料属性，枚举值参考附录Creative Attributes|
|~~maxextended~~	|integer |	最大的视频广告延长时间长度（如果支持延长）-1：表示允许延时，且没有时间限制空或者0：表示不允许延长大于0：表示可以延长的时间长度比maxduration大的值|
|minbitrate	|integer |最小的比特率，以 Kbps 为单位；缺省表示不限制|
|maxbitrate	|integer |最大的比特率，以 Kbps 为单位；缺省表示不限制|
|~~boxingallowed~~	|integer |是否允许将4：3的内容展示在16：9的窗口， 0表示不允许，1表示允许|
|playbackmethod	|array of integer|	允许的播放方式，播放方式和对应的枚举值如下：1: 自动播放（有声）2: 自动播放（静音）3: 点击播放 如果该字段缺省，表示支持全部|
|~~playbackend~~	|integer |	导致视频播放中断的原因；|
|delivery	|array of integer|	支持的传输方式，传输方式和对应的枚举值如下：1-Streaming；2-Progressive如果没有指定，表示全部支持|
|~~pos~~	|integer |	广告位置，枚举值参考附录Ad Position|
|companionad	| array of banner object |	如果视频支持伴随广告，该字段表示一组 Banner 对象；|
|api	|array of integer |	该次展示可支持的 API 框架；枚举值参考附录API Frameworks该字段缺省表示所有枚举值均不支持；|
|companiontype	|array of integer | 否 |	支持的 VAST companion广告类型；如果在companionad 中填充了 Banner 对象则推荐使用；枚举值参考附录Companion Types|
|ext	|video-ext Object|	否|	见video-ext object|

#### video-ext Object
| 字段	| 类型 	| 描述 |
|---|---|---|
|is_rewarded | bool | 是否是激励视频 |


#### native Object

| 字段	| 类型	| 描述 |
|---|---|---|
| request | string; required |  表示 Native Ad Specification 定义的 native 标签协议； 具体见下文的 [NativeRequest 参数说明](#object-nativerequest)
|ver |  string | 采用的 Dynamic Native Ads API 的版本|
|api |  array of integers |        该次展示可支持的 API 框架；枚举值参考[附录API Frameworks](appendix.md#api-frameworks) |
|battr    |array of integers|        限制的物料属性，枚举值参考[附录Creative Attributes](appendix.md#creative-attributes)|


#### app Object

| 字段      | 类型        | 描述                                                            |
|---------------|----------|------------|
| id            | string | 表示 adx 内部的 APP id；                                        |
| name          | string | 表示 adx 内部的 APP 名称；                                      |
| bundle        | string | 商店应用的包名；安卓示例 com.foo.mygame；ios示例907394059;      |
| domain        | string | App 的域名，例如 mygame.foo.com                                 |
| storeurl      | string | APP 的商店链接地址；                                            |
| cat           | array of strings | APP 的 IAB category；见附录Content Categories                   |
| ~~sectioncat~~    | array of strings | 应用当前部分的IAB内容类型数组，枚举值参考附录Content Categories |
| ~~pagecat~~       | array of strings | 应用当前视图的IAB内容类型数组，参考附录Content Categories       |
| ver           | string | APP 的版本号；                                                  |
| ~~privacypolicy~~ | integer  | 表示该应用是否有隐私策略， 0-没有；1-有；                       |
| paid          | integer | 应用是否需要付费， 0-免费；1-付费；                             |
| publisher     | publisher object | 表示开发者的具体信息，具体 Object Publisher                     |
| content       | content object  | Content对象， 该应用内容的详细信息                              |
| keywords      | string  | 逗号分隔的应用的关键字信息                                      |

#### site Object

| 字段      | 类型        | 描述                                                            |
|---------------|----------|------------|
| id            | string | 表示 adx 内部的 Site ID；                                        |
| name          | string | Site 名称，也可能是开发者的别名；                                      |
| domain        | string | Site 的域名                                 |
| storeurl      | string | APP 的商店链接地址；                                            |
| cat           | array of strings | Site 的 IAB category；见附录Content Categories                   |
| ~~sectioncat~~    | array of strings | 应用当前部分的IAB内容类型数组，枚举值参考附录Content Categories |
| ~~pagecat~~       | array of strings | 应用当前视图的IAB内容类型数组，参考附录Content Categories       |
| page           | string | 发生展示的页面的URL；                                                  |
| ref           | string | 浏览到当前页面的Referrer URL；                                                  |
| search           | string | 浏览到当前页面的搜索字符串；                                                  |
| mobile           | integer | Site页面的布局是否针对于移动设备做了优化 0=no，1=yes                                                  |
| ~~privacypolicy~~ | integer  | 表示该应用是否有隐私策略， 0-没有；1-有；                       |
| publisher     | publisher object | 表示开发者的具体信息，具体 Object Publisher                     |
| content       | content object  | Content对象， 该应用内容的详细信息                              |
| keywords      | string  | 逗号分隔的应用的关键字信息                                      |


#### user Object

| 字段 | 类型  | Description                                          |
|-----------|------|------------------------------------------------------|
| id        | string | 交易用户的ID, 推荐id和buyeruid 至少传一个              |
| buyeruid     | string  | buyer的ID                                         |
| yob       | integer | 用户出生年份的四位数字 |
| gender    | string  | 用户的性别，M=male, F=female, O=known to be other        |
| keywords  | string  | 逗号分割的关键字，兴趣或者意向列表 ｜
| customdata | string | 可选特性，用于传递给竞拍者的信息，在交易平台的cookie中设置，字符串必须base85编码的cookie|
| geo       | geo object | 用户Geo的信息 |
| data   | array of data object | 附加的用户信息，每个Data对象表示一个不同的数据源 |


#### data Object

| 字段 | 类型  | Description                                          |
|-----------|------|------------------------------------------------------|
| id        | string | 交易特定的数据提供者ID              |
| name     | string  | 交易特定的数据提供者名称                                         |
| segment      | array of segment object | 包含数据内容的一组Segment对象 |

#### segment Object

| 字段 | 类型  | Description                                          |
|-----------|------|------------------------------------------------------|
| id        | string | 数据提供者的特定数据段的ID              |
| name     | string  | 数据提供者的特定数据段的名称                                         |
| value      | string | 表示数据字段的值的字符串 |


#### publisher Object
| 字段 | 类型  | Description                                          |
|-----------|------|------------------------------------------------------|
| id        | string | 表示 ADX 内部的开发者 id；                           |
| ~~name~~      | string  | 开发者名称；                                         |
| cat       | array of strings | 发布者的IAB内容类型数组， 参考附录Content Categories |
| ~~domain~~    | string  | 发布者的顶级域名（例如， "publisher.com" )；         |

#### device Object

| 字段      | 类型  | 描述                                                                                                                                                              |
|----------------|-----|-------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| ua             | string | 用户设备 HTTP 请求头中的 User-Agent 字段；                                                                                                                        |
| geo        | geo object | 表示用户当前位置；                                                                                                                                                |
| dnt            | integer  | 浏览器在 HTTP 头中设置的标准的“Do Not Track"标识， 0-不限制追踪；1-限制（不允许）追踪 注意，该字段值类型定义与open RTB 2.4版本协议有所不同                        |
| lmt            | integer  | “限制广告追踪”表示用户对商业追踪行为的授权， 值为 0-不限制追踪；1-限制追踪 注意，该字段值类型定义与open RTB 2.4版本协议有所不同                                   |
| ip             | string; required | 最接近设备的 IPv4 地址；                                                                                                                                          |
| ~~ipv6~~           | string   | 最接近设备的 Ipv6 地址；                                                                                                                                          |
| devicetype     | integer; required | 设备类型；枚举值参考附录Device Type                                                                                                                               |
| make           | string; required   | 设备制造商，比如“Apple”，未知传 unknown                                                                                                                           |
| model          | string; required  | 1） android 设备:调用系统接口android.os.Build.MODEL 直接获得； 2） ios 设备：对系统接口返回原始值做转换后得到，取值例如 iPhone5 、 iPhone6s 、 iPhone 6sPlus 等； |
| os             | string; required  | 操作系统；未知传 unknown                                                                                                                                          |
| osv            | string; required  | Os 版本；三段式或两段式版本号；                                                                                                                                   |
| hwv            | string   | 设备硬件版本， 例如 “5S”；                                                                                                                                        |
| h              | integer  | 屏幕的物理高度， 以像素为单位；                                                                                                                                   |
| w              | integer  | 屏幕的物理宽度，以像素为单位；                                                                                                                                    |
| ~~ppi~~            | integer | 以像素每英寸表示的屏幕尺寸；                                                                                                                                      |
|~~pxratio~~        | float | 设备物理像素与设备无关像素的比率；                                                                                                                                |
| js             | integer  | 支持javascript, 0-不支持；1-支持；                                                                                                                                |
| ~~geofetch~~       | integer | 表示该广告位是否为JavaScript代码提供geolocaion API， 0-不提供；1-提供                                                                                             |
| ~~flashver~~       | string  | 浏览器支持的Flash版本；                                                                                                                                           |
| language       | string; required  | 设备语言；使用 ISO-639-1-alpha-2；未知传unknown                                                                                                                   |
| carrier        | string; required | 运营商；字段值采用 MCC 和 MNC 结合的代码， 如46001 表示中国联通；未知传 unknown；                                                                                 |
| ~~mccmnc~~         | string   | 运营商mcc-mnc代码；                                                                                                                                               |
| connectiontype | integer;required | 网络连接类型；枚举值参考附录Connection Type                                                                                                                       |
| ifa            | string  | 广告主标识， 明文表示； Ios 传 idfa，必传； Android 国外传 gaid，国内不传；                                                                                       |
| imei           | string | 硬件设备 ID，安卓传 IMEI                                                                                                                                          |
| android_id     | string  | 设备平台 ID，安卓传 Android ID                                                                                                                                    |
| oaid     | string  | OAID                                                                                                                                   |
| didsha1        | string  | 硬件设备 ID，安卓传 IMEI，使用 SHA1 哈希算法；                                                                                                                    |
| didmd5         | string  | 硬件设备 ID，安卓传 IMEI，使用 md5 哈希算法；                                                                                                                     |
| dpidsha1       | string | 设备平台 ID，安卓传 Android ID，使用 SHA1 哈希算法；                                                                                                              |
| dpidmd5        | string | 设备平台 ID，安卓传 Android ID，使用 md5 哈希算法；                                                                                                               |

#### geo Object

| 字段    | 类型  | 描述                                                                                                                                                                                                               |
|---------------|----|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| lat           | float | 纬度信息，取值范围-90.0到+90.0， 负值表示南方；                                                                                                                                                                   |
| lon           | float | 经度信息，取值返回-180.0到+180.0， 负值表示西方；                                                                                                                                                                |
| type          | integer | 位置信息的源；值为 1 表示 GPS/定位服务；值为 2 表示 IP 地址； 值为 3 表示用户提供；                                                                                                                                |
| ~~accuracy~~      | integer | 精度，详细到米；当经纬度是通过定位服务获取时，上报该字段；                                                                                                                                                         |
| ~~lastfix~~       | integer | 生成竞价请求的时间和设备最后一次获取地理位置的时间之间的差别，这个时间差的单位：秒 请注意：设备可能缓存多个不同时间获取的地理位置数据。 在理想状况下，这个数值应该是从此竞价请求里包含的地理位置获取时间开始计算的 |
| ~~ipservice~~    | integer | 从IP地址分析地理位置的服务方或者供应商（当type = 2的时候）;                                                                                                                                                        |
| country       | string | ISO-3166-1 Alpha-3 国家码；                                                                                                                                                                                  |
| region        | string | ISO-3166-2 区域码；                                                                                                                                                                                              |
|~~regionfips10~~ | string | 国家的区域，使用FIPS 10-4 notation 编码表示，也可用 ISO 3166-2编码表示；                                                                                                                                       |
| metro         | string  | 谷歌metro code Metro codes 与 Nielsen DMAs 相似，但有一些差异                                                                                                                                                      |
| city          | string  | 城市名称使用联合国贸易与运输位置码                                                                                                                                                                                 |
| zip           | string  | 邮政编码或者邮递区号；                                                                                                                                                                                             |
|~~utcoffse~~     | integer | 本地时间，用比标准UTC时间多加或减少的分钟数来表示；                                                                                                                                                                |

#### regs Object

| 字段 | 类型    | 描述                                                                                                                                  |
|----------|-------- |---------------------------------|
| coppa    | integer | 表示该次展示是否遵从 COPPA 法案， 0-不遵从；1-遵从； 对于遵从 COPPA 法案的展示，DSP 必须保证返回的广告的内容和素材符合 COPPA 广告规定 |

## 4. Response specification
### a. Bid response parameters
DSP 应该使用JSON格式序列化Bid的信息

HTTP 204 No Content 表示不bid

#### Bid Response Object (Top Level)
| 参数名称 | 类型 | 描述                                                      |
|----------|-------------|-----------------------------------------------------------|
| id       | string;required    | 竞价请求的标识，即请求带的 request id；                   |
| seatbid  | array of seatbid object;required    | 一组 SeatBid 对象， 如果出价，则至少应该填充一个seatbid； |
| bidid    | string    | 竞拍者生成的响应 ID, 辅助日志或者交易追踪；               |
| cur      | string   | 出价货币单位，使用 ISO-4217 码；不传默认 USD              |
| nbr      | integer   | 不出价原因；                                              |


#### seatbid Object
| 参数名称 | 类型 | 描述               |
|----------|-------|--------------------------------------------|
| bid | array of bid object; required| 至少一个 Bid 对象的数组，每个对象关联一个展示。   |
| seat     | string     | 出价者席位标识， 代表本次出价的出价人；               |
| group | integer| 1 = 出价方要求对所有展示的出价必须整组胜出，或失败; 0 = 对某一展示的一次出价可以独立胜出，默认值 = 0 ; |


#### bid Object
| 参数名称       | 类型   | 描述                         |
|----------------|--------------|-------------------------------------------------------|
| id  | string; required  | 竞拍者生成的竞价 ID，用于记录日志或行为追踪；                     |
| impid | string; required    | 关联的竞价请求中的 Imp 对象的 ID；          |
| price          | float; required  | 对该次展示的出价，以 CPM 表示, 默认美元；      |
| nurl    | string;required   | 胜出通知链接；MTG adx将在广告成功展示时调用该链接；    |
| burl    | string     | 可计费展示回调；      |
| lurl    | string     | 竞价失败回调；  |
| adm     | string;required | 广告素材标记；Native广告形式返回native response； 视频广告形式返回VAST XML;Banner广告形式返回xhtml；|
| adid    | string     | 竞价的广告的ID， 如果交易胜出，该广告会被发送给媒体；    |
| adomain | array of strings | 广告主域名， 用于过滤检测；   |
| bundle  | string; required（下载类广告必传） | 应用的包名信息； 安卓包名示例 com.foo.mygame；ios 包名示例907394059  |
| iurl    | string  | 用于质量或者安全监测的表示广告活动内容的图像地址； |
| cid     | string | 广告 id，辅助广告审核；iurl 代表的一组素材   |
| crid    | string | 一组素材的 id；辅助广告审核  |
| tactic  | string | 广告投放策略id； |
| cat     | array of strings  | creative 的 IAB 内容类型；枚举值参考附录Content Categories  |
| attr    | array of integers  | 描述 creative 的属性集合；枚举值参考附录Creative Attributes   |
| api     | integer | 该次展示可支持的 API 框架；枚举值参考附录API Frameworks   |
| protocols | integer  | 支持的视频竞价响应协议；枚举值参考附录Protocols  |
| qagmediarating | integer  | 表示根据 IAB IGQ 标准的素材内容等级； 枚举值参考附录IQG Media Ratings   |
| language  | string | 素材语言；设备语言；使用 ISO-639-1-alpha-2；|
| dealid    | string | 如果该竞价从属于某个私有市场交易， 这个参数包含这个私有市场交易的交易ID； 如果竞拍的展示从属于某个私有市场交易， 那么该竞价必须包含相同的私有市场交易的ID |
| w    | integer   | 广告的宽度，单位：像素。 |
| h    | integer   | 广告的高度，单位：像素。 |
| wratio  | integer    | 广告的相对宽度，单位：像素。  |
| hratio  | integer   | 广告的相对高度，单位：像素。  |
| exp     | integer   | 广告从返回到实际展示的有效延迟时间，单位为秒；默认值为 3600；     |
| ext     | bid-ext object          | 具体见 ext object  |

### b. impression/click beacon
DSP在bidresponse中向Mintegral ADX返回nurl和展示监测链接（其中，展示监测链接可选，DSP可不返回）。
当广告成功在客户端展示时，MTG ADX向DSP同时调nurl和展示监测链接。

为了缩小展示数据差异，建议对于所有广告形式，DSP都返回独立的展示监测链接回收展示数据。

Native image和Native video，DSP通过Native Ads协议的imptrackers字段返回展示监测链接；

Rewarded video 和 Interstitial video，DSP通过VAST的\<impression>标签返回展示监测链接；

Interative ads和Banner，DSP通过 Bidresponse.Seatbid.Bid.Ext.Imptrakers 返回展示监测链接。

### c. Macro substitution
Mintegral 支持以下宏替换

| 宏名 | 含义 | 备注 |
|---|---|---|
| ${AUCTION_ID} |	竞价请求 ID |	mtg-adx 请求dsp 的request.id |
| ${AUCTION_BID_ID}	| DSP 竞价响应的 ID	| DSP返回的mtg-adx的BidResponse.BidID |
|${AUCTION_IMP_ID} |展示 ID |	DSP 返回的 Bid.ImpId|
|${AUCTION_SEAT_ID} |席位 ID |	DSP返回的BidResponse.SeatBid.Seat|
|${AUCTION_AD_ID}	 |广告 ID |	广告 ID，Bid.Adid|
|${AUCTION_PRICE}	 |结算价格 |	加密后的结算价格，具体请参考结算价格解密|
### d. Click Measuring

## 5. Code table

## 6. Bid Request/Response Examples

### Request Examples
See [version2.req-examples.md](./version2.req-examples.md) for details.

### Response Examples
See [version2.res-examples.md](./version2.res-examples.md) for details.