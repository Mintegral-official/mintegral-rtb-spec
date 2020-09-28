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
    * [source Object](#source-object)
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
    * [pmp Object](#pmp-object)
    * [deal Object](#deal-object)
    * [skadn Request Object](#skadn-request-object)
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
 | ~~site~~  | site object  |     定义site信息|
 | app     | app object   |   定义app信息；具体见[Object App参数说明](#object-app)|
 | device  | device object; required   |   定义设备信息；具体见[Object Device参数说明](#object-device)|
 | ~~user~~| object  |    定义用户信息|
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