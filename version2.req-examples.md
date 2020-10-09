# Mintegral RTB Specification ver 2.5: Request Examples

* [Web](#web)
* [App](#app)
* [Native (image)](#native-image)
* [Native (video)](#native-video)
* [Rewarded-Video](#rewarded-video)
* [Interstitial-Video](#Interstitial-video)
* [Banner](#banner)
* [Splash](#splash)


### Web

```json
{
  "at": 1,
  "badv": [
    "blocked-advertiser1.com",
    "blocked-advertiser2.com",
    "blocked-advertiser3.com"
  ],
  "bapp": [],
  "bcat": [
    "IAB23",
    "IAB23-1",
    "IAB23-10",
    "IAB23-2",
    "IAB23-3",
    "IAB23-4",
    "IAB23-5",
    "IAB23-6",
    "IAB23-7",
    "IAB23-8",
    "IAB23-9",
    "IAB24",
    "IAB25-2",
    "IAB25-3",
    "IAB25-4",
    "IAB25-5",
    "IAB26",
    "IAB26-1",
    "IAB26-2",
    "IAB26-3",
    "IAB26-4"
  ],
  "cur": [
    "JPY",
    "USD"
  ],
  "device": {
    "connectiontype": 0,
    "devicetype": 2,
    "dnt": 0,
    "geo": {
      "city": "Shibuya",
      "country": "JPN",
      "ipservice": 3,
      "region": "JP-13",
      "type": 2,
      "utcoffset": 540,
      "zip": "150-0043"
    },
    "hwv": "",
    "ifa": "",
    "ip": "210.168.46.254",
    "ipv6": "",
    "js": 1,
    "language": "ja",
    "lmt": 0,
    "make": "",
    "model": "",
    "os": "Windows NT",
    "osv": "10.0",
    "pxratio": 0.0,
    "ua": "Mozilla\/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit\/537.36 (KHTML, like Gecko) Chrome\/51.0.2704.79 Safari\/537.36 Edge\/14.14393"
  },
  "id": "86e40f49-f710-4293-a229-35e0fbf82e7a",
  "imp": [
    {
      "banner": {
        "api": [],
        "battr": [],
        "btype": [],
        "ext": {},
        "format": [
          {
            "h": 250,
            "w": 300
          }
        ],
        "h": 250,
        "pos": 1,
        "w": 300
      },
      "bidfloor": 2.430134,
      "bidfloorcur": "USD",
      "ext": {},
      "id": "",
      "instl": 0,
      "metric": [],
      "pmp": {
        "deals": [],
        "ext": {},
        "private_auction": 0
      },
      "secure": 1,
      "tagid": "1111:22222222"
    }
  ],
  "site": {
    "cat": [
      "IAB1"
    ],
    "domain": "magazine.fluct.jp",
    "ext": {},
    "id": "123",
    "mobile": 1,
    "name": "Example Site",
    "page": "https:\/\/magazine.fluct.jp\/category\/news",
    "pagecat": [],
    "publisher": {
      "cat": [],
      "domain": "",
      "ext": {},
      "id": "234",
      "name": ""
    },
    "ref": "https:\/\/google.com\/search?",
    "sectioncat": []
  },
  "source": {
    "ext": {
      "schain": {
        "complete": 1,
        "nodes": [
          {
            "asi": "adingo.jp",
            "ext": {},
            "hp": 1,
            "rid": "86e40f49-f710-4293-a229-35e0fbf82e7a",
            "sid": "234"
          }
        ],
        "ver": "1.0"
      },
      "stype": ""
    },
    "fd": 0,
    "tid": "86e40f49-f710-4293-a229-35e0fbf82e7a"
  },
  "tmax": 120,
  "user": {
    "buyeruid": "synced-buyeruid",
    "ext": {},
    "id": "9c73dc20-712e-4d94-8255-6de93accf040"
  },
  "wseat": []
}
```

### App

```json
{
  "app": {
    "bundle": "1234567890",
    "cat": [
      "IAB1"
    ],
    "ext": {},
    "id": "123",
    "name": "Example App",
    "pagecat": [],
    "publisher": {
      "cat": [],
      "domain": "",
      "ext": {},
      "id": "234",
      "name": ""
    },
    "sectioncat": [],
    "storeurl": "https:\/\/apps.apple.com\/jp\/app\/hogehoge-app\/id1234567890"
  },
  "at": 1,
  "badv": [],
  "bapp": [
    "12345",
    "23456",
    "34567"
  ],
  "bcat": [
    "IAB23",
    "IAB23-1",
    "IAB23-10",
    "IAB23-2",
    "IAB23-3",
    "IAB23-4",
    "IAB23-5",
    "IAB23-6",
    "IAB23-7",
    "IAB23-8",
    "IAB23-9",
    "IAB24",
    "IAB25-2",
    "IAB25-3",
    "IAB25-4",
    "IAB25-5",
    "IAB26",
    "IAB26-1",
    "IAB26-2",
    "IAB26-3",
    "IAB26-4"
  ],
  "cur": [
    "JPY",
    "USD"
  ],
  "device": {
    "connectiontype": 2,
    "devicetype": 2,
    "dnt": 0,
    "geo": {
      "city": "Shibuya",
      "country": "JPN",
      "ipservice": 3,
      "region": "JP-13",
      "type": 2,
      "utcoffset": 540,
      "zip": "150-0043"
    },
    "hwv": "",
    "ifa": "e2249a97-8609-4072-9aaf-b4a15129873f",
    "ip": "210.168.46.254",
    "ipv6": "",
    "js": 1,
    "language": "ja",
    "lmt": 0,
    "make": "Apple",
    "model": "iPhone",
    "os": "iOS",
    "osv": "9.1",
    "pxratio": 0.0,
    "ua": "Mozilla \/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit\/601.1.46 (KHTML, like Gecko) Version\/9.0 Mobile\/13B5110e Safari\/601.1"
  },
  "id": "66a31e28-d140-4483-b61d-b435159725ca",
  "imp": [
    {
      "banner": {
        "api": [],
        "battr": [],
        "btype": [],
        "ext": {},
        "format": [
          {
            "h": 250,
            "w": 300
          }
        ],
        "h": 250,
        "pos": 1,
        "w": 300
      },
      "bidfloor": 2.430134,
      "bidfloorcur": "USD",
      "ext": {},
      "id": "",
      "instl": 0,
      "metric": [],
      "pmp": {
        "deals": [],
        "ext": {},
        "private_auction": 0
      },
      "secure": 1,
      "tagid": "1111:22222222"
    }
  ],
  "source": {
    "ext": {
      "schain": {
        "complete": 1,
        "nodes": [
          {
            "asi": "adingo.jp",
            "ext": {},
            "hp": 1,
            "rid": "66a31e28-d140-4483-b61d-b435159725ca",
            "sid": "234"
          }
        ],
        "ver": "1.0"
      },
      "stype": ""
    },
    "fd": 0,
    "tid": "66a31e28-d140-4483-b61d-b435159725ca"
  },
  "tmax": 120,
  "user": {
    "buyeruid": "",
    "ext": {},
    "id": "e2249a97-8609-4072-9aaf-b4a15129873f"
  },
  "wseat": []
}
```

### Native (image)

```json
{
  "at": 1,
  "badv": [
    "blocked-advertiser1.com",
    "blocked-advertiser2.com",
    "blocked-advertiser3.com"
  ],
  "bapp": [],
  "bcat": [
    "IAB23",
    "IAB23-1",
    "IAB23-10",
    "IAB23-2",
    "IAB23-3",
    "IAB23-4",
    "IAB23-5",
    "IAB23-6",
    "IAB23-7",
    "IAB23-8",
    "IAB23-9",
    "IAB24",
    "IAB25-2",
    "IAB25-3",
    "IAB25-4",
    "IAB25-5",
    "IAB26",
    "IAB26-1",
    "IAB26-2",
    "IAB26-3",
    "IAB26-4"
  ],
  "cur": [
    "JPY",
    "USD"
  ],
  "device": {
    "connectiontype": 2,
    "devicetype": 2,
    "dnt": 0,
    "geo": {
      "city": "Shibuya",
      "country": "JPN",
      "ipservice": 3,
      "region": "JP-13",
      "type": 2,
      "utcoffset": 540,
      "zip": "150-0043"
    },
    "hwv": "",
    "ifa": "",
    "ip": "210.168.46.254",
    "ipv6": "",
    "js": 1,
    "language": "ja",
    "lmt": 0,
    "make": "",
    "model": "SOV39",
    "os": "Android",
    "osv": "9",
    "pxratio": 0.0,
    "ua": "Mozilla\/5.0 (Linux; Android 9; SOV39 Build\/52.0.C.1.119) AppleWebKit\/537.36 (KHTML, like Gecko) Chrome\/68.0.3440.91 Mobile Safari\/537.36"
  },
  "id": "75a40036-f9d6-4f8c-8240-d37b28c2dbc7",
  "imp": [
    {
      "bidfloor": 300.0,
      "bidfloorcur": "JPY",
      "ext": {},
      "id": "",
      "instl": 0,
      "metric": [],
      "native": {
        "api": [],
        "battr": [],
        "ext": {},
        "request": "{\"native\":{\"assets\":[{\"id\":1,\"required\":1,\"title\":{\"len\":10}},{\"id\":2,\"img\":{\"hmin\":10,\"type\":1,\"wmin\":10},\"required\":1},{\"data\":{\"len\":30,\"type\":1},\"id\":3},{\"data\":{\"len\":50,\"type\":2},\"id\":4}],\"layout\":3,\"plcmtcnt\":2,\"ver\":\"1.0\"}}",
        "ver": "1.0"
      },
      "pmp": {
        "deals": [],
        "ext": {},
        "private_auction": 0
      },
      "secure": 1,
      "tagid": "1111:22222222"
    }
  ],
  "site": {
    "cat": [
      "IAB1"
    ],
    "domain": "magazine.fluct.jp",
    "ext": {},
    "id": "123",
    "mobile": 1,
    "name": "Example Site",
    "page": "https:\/\/magazine.fluct.jp\/category\/news",
    "pagecat": [],
    "publisher": {
      "cat": [],
      "domain": "",
      "ext": {},
      "id": "234",
      "name": ""
    },
    "ref": "https:\/\/google.com\/search?",
    "sectioncat": []
  },
  "source": {
    "ext": {
      "schain": {
        "complete": 1,
        "nodes": [
          {
            "asi": "adingo.jp",
            "ext": {},
            "hp": 1,
            "rid": "75a40036-f9d6-4f8c-8240-d37b28c2dbc7",
            "sid": "234"
          }
        ],
        "ver": "1.0"
      },
      "stype": ""
    },
    "fd": 0,
    "tid": "75a40036-f9d6-4f8c-8240-d37b28c2dbc7"
  },
  "tmax": 120,
  "user": {
    "buyeruid": "synced-buyeruid",
    "ext": {},
    "id": "2c7b94bfca27de03ce862bb6c9cb7a1e7093197d"
  },
  "wseat": []
}
```

### Native (video)
``` json
```

### Rewarded-Video
``` json
```

### Interstitial-Video
``` json
```

### Banner
``` json
```

### Splash
``` json
```