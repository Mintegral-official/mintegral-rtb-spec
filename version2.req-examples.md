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
    "id": "5d5e0a705779c1436f58db0x",
    "imp": [
        {
            "id": "1",
            "native": {
                "request": "{\"ver\":\"1.1\", \"assets\":[{\"id\":7, \"required\":1, \"img\":{\"type\":3, \"w\":1200, \"h\":627}}, {\"id\":1, \"required\":1, \"img\":{\"type\":1, \"w\":300, \"h\":300}}, {\"id\":3, \"required\":1, \"title\":{\"len\":20}}, {\"id\":4, \"required\":1, \"data\":{\"type\":2}}, {\"id\":5, \"required\":1, \"data\":{\"type\":3}}, {\"id\":6, \"required\":1, \"data\":{\"type\":12}}}]}",
                "ver": "1.1"
            },
            "tagid": "1510001443",
            "bidfloor": 1,
            "bidfloorcur": "USD",
            "secure": true,
            "support_download": false,
            "request_type": 7,
            "displaymanagerver": "mi_3.3.2",
            "ext": {
                "native_type": 3
            },
            "ext_channel": ""
        }
    ],
    "app": {
        "id": "1510000000",
        "cat": [
            "IAB10",
            "IAB22",
            "IAB14",
            "IAB1"
        ],
        "ver": "8.3.4",
        "bundle": "ni.criteo.1",
        "publisher": {
            "id": "12441"
        },
        "storeurl": "https://itunes.apple.com/app/id925015459"
    },
    "device": {
        "ua": "Mozilla/5.0%20(iPhone%3B%20CPU%20iPhone%20OS%2010_2_1%20like%20Mac%20OS%20X)%20AppleWebKit/602.4.6%20(KHTML%2C%20like%20Gecko)%20Mobile/14D27%20Camera360/8.3.4",
        "geo": {
            "country": "SG",
            "city": "390"
        },
        "dnt": false,
        "ip": "39.109.124.93",
        "devicetype": 4,
        "make": "0",
        "model": "iphone 6s",
        "os": "ios",
        "osv": "11.1.2",
        "language": "",
        "carrier": "",
        "connectiontype": 0,
        "ifa": "9601B92C-0A21-4FD6-B176-6AAAF241B579",
        "didsha1": "",
        "didmd5": "",
        "imei": ""
    },
    "at": 2,
    "tmax": 2000,
    "bcat": [
        "IAB26",
        "IAB24",
        "IAB25"
    ]
}
```

### Native (video)
``` json
{
    "id": "5d5e0a705779c1436f58db0x",
    "imp": [
        {
            "id": "1",
            "native": {
                "request": "{\"ver\":\"1.1\", \"assets\":[{\"id\":7, \"required\":1, \"img\":{\"type\":3, \"w\":1200, \"h\":627}}, {\"id\":1, \"required\":1, \"img\":{\"type\":1, \"w\":300, \"h\":300}}, {\"id\":3, \"required\":1, \"title\":{\"len\":20}}, {\"id\":4, \"required\":1, \"data\":{\"type\":2}}, {\"id\":5, \"required\":1, \"data\":{\"type\":3}}, {\"id\":6, \"required\":1, \"data\":{\"type\":12}}, {\"id\":2, \"required\":1, \"video\":{\"mimes\":[\"video/mp4\"], \"w\":720, \"h\":1280, \"ext\":{\"orientation\":1}}}]}",
                "ver": "1.1"
            },
            "tagid": "1510001443",
            "bidfloor": 1,
            "bidfloorcur": "USD",
            "secure": true,
            "support_download": false,
            "request_type": 7,
            "displaymanagerver": "mi_3.3.2",
            "ext": {
                "native_type": 3
            },
            "ext_channel": ""
        }
    ],
    "app": {
        "id": "1510000000",
        "cat": [
            "IAB10",
            "IAB22",
            "IAB14",
            "IAB1"
        ],
        "ver": "8.3.4",
        "bundle": "ni.criteo.1",
        "publisher": {
            "id": "12441"
        },
        "storeurl": "https://itunes.apple.com/app/id925015459"
    },
    "device": {
        "ua": "Mozilla/5.0%20(iPhone%3B%20CPU%20iPhone%20OS%2010_2_1%20like%20Mac%20OS%20X)%20AppleWebKit/602.4.6%20(KHTML%2C%20like%20Gecko)%20Mobile/14D27%20Camera360/8.3.4",
        "geo": {
            "country": "SG",
            "city": "390"
        },
        "dnt": false,
        "ip": "39.109.124.93",
        "devicetype": 4,
        "make": "0",
        "model": "iphone 6s",
        "os": "ios",
        "osv": "11.1.2",
        "language": "",
        "carrier": "",
        "connectiontype": 0,
        "ifa": "9601B92C-0A21-4FD6-B176-6AAAF241B579",
        "didsha1": "",
        "didmd5": "",
        "imei": ""
    },
    "at": 2,
    "tmax": 2000,
    "bcat": [
        "IAB26",
        "IAB24",
        "IAB25"
    ]
}
```

### Rewarded-Video
``` json
{
	"id": "916ac896-4028-4564-a73b-e2c6ac86187f",
	"app": {
		"id": "115263",
		"name": "采矿大亨：掘金之旅 (Idle Miner Tycoon)",
		"cat": [
			"IAB1"
		],
		"ver": "3.16.1",
		"bundle": "1116645064",
		"publisher": {
			"id": "17702"
		},
		"storeurl": "https://apps.apple.com/us/app/idle-miner-tycoon/id1116645064?mt=8",
		"storeid": "com.fluffyfairygames.idleminertycoon"
	},
	"at": 2,
	"tmax": 600,
	"bcat": [
		"IAB26",
		"IAB24",
		"IAB25",
		"IAB11-4"
	],
	"source": {
		"ext": {
			"schain": {
				"ver": "1.0",
				"complete": 1,
				"nodes": [{
					"asi": "mintegral.com",
					"sid": "17702",
					"rid": "916ac896-4028-4564-a73b-e2c6ac86187f",
					"hp": 1
				}]
			}
		}
	},
	"imp": [{
		"id": "1",
		"video": {
			"mimes": [
				"video/mp4"
			],
			"maxduration": 30,
			"protocols": [
				2,
				3,
				5,
				6
			],
			"w": 720,
			"h": 1280,
			"linearity": 1,
			"skip": 0,
			"battr": [
				16
			],
			"maxbitrate": 2000,
			"delivery": [
				2
			],
			"pos": 7,
			"companionad": [{
				"w": 720,
				"h": 1280,
				"id": "1",
				"battr": [
					16
				],
				"pos": 7,
				"mimes": [
					"application/javascript",
					"image/jpeg",
					"image/jpg",
					"text/html",
					"image/png",
					"text/css",
					"image/gif"
				],
				"api": [
					3,
					5,
					6
				],
				"ext": {
					"orientation": 1,
					"is_rewarded": true
				}
			}],
			"api": [
				3,
				5,
				6
			],
			"companiontype": [
				1,
				2,
				3
			],
			"ext": {
				"orientation": 1,
				"videotype": 1,
				"videoendtype": [
					1,
					2
				],
				"endcardonly": true,
				"is_rewarded": true
			}
		},
		"tagid": "211967",
		"bidfloorcur": "USD",
		"displaymanagerver": "mi_6.3.3",
		"ext": {},
		"instl": 1,
		"secure": 1,
		"bidfloor": 3
	}],
	"device": {
		"ua": "Mozilla/5.0 (iPhone; CPU iPhone OS 14_0_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
		"ip": "112.146.94.2",
		"devicetype": 4,
		"make": "apple",
		"model": "iphone 7 plus",
		"os": "ios",
		"osv": "14.0.1",
		"carrier": "45008",
		"connectiontype": 2,
		"ifa": "00000000-0000-0000-0000-000000000000",
		"didmd5": "",
		"dpidmd5": "",
		"geo": {
			"city": "namchang-dong",
			"country": "KOR"
		},
		"dnt": 0,
		"lmt": 0,
		"language": "ko"
	},
	"regs": {
		"coppa": 0,
		"ext": {
			"gdpr": 0
		}
	}
}
```

### Interstitial-Video
``` json
{
	"id": "9148c67b-ff15-4be3-aaf1-6417b206be39",
	"app": {
		"id": "130199",
		"name": "AMAZE!!!",
		"cat": [
			"IAB1"
		],
		"ver": "2.9.5",
		"bundle": "1452526406",
		"publisher": {
			"id": "18837"
		},
		"storeurl": "https://apps.apple.com/us/app/amaze/id1452526406",
		"storeid": "com.amaze.game"
	},
	"at": 2,
	"tmax": 600,
	"bcat": [
		"IAB26",
		"IAB14-1",
		"IAB1",
		"IAB24",
		"IAB25",
		"IAB11-4"
	],
	"bapp": [
		"id1411754473",
		"id1448852425",
		"id1077137248",
		"id1374403536",
		"id1041545937",
		"id1139379843",
		"id1370346970",
		"id1089836344",
		"id1424349196",
		"id970417047",
		"id1328359160",
		"id1403848965",
		"id969316884",
		"id524731580",
		"id741183306",
		"id1182341536",
		"id1447322621",
		"id529652920",
		"id1438385422",
		"id558207346",
		"id1277029359",
		"com.candywriter.bitlife",
		"com.sail.advanced.booster",
		"com.randomsaladgames.ginrummydeluxe",
		"com.aim.ginrummy",
		"ca.lbcstudios.hempire",
		"com.tutapp.idlezombies",
		"com.cmcm.live",
		"com.admarket.passionpuzzle",
		"com.neonplay.casualrollersplat2",
		"com.joinroot.root",
		"com.nextgames.android.twd",
		"com.nextgames.android.ourworld",
		"com.scopely.headshot",
		"com.telltalegames.walkingdead100",
		"com.telltalegames.walkingdead200",
		"veeu.watch.funny.video.vlog.moment",
		"com.fiogonia.yatzy",
		"com.noodlecake.zombieroadtrip",
		"net.mobigame.zombietsunami",
		"jp.co.cybird.appli.android.sgk",
		"com.vng.android.mps.dead.zombie2",
		"com.mars.avgchapters",
		"id1411754473",
		"id1448852425",
		"id1077137248",
		"id1374403536",
		"id1041545937",
		"id1139379843",
		"id1370346970",
		"id1089836344",
		"id1424349196",
		"id970417047",
		"id1328359160",
		"id1403848965",
		"id969316884",
		"id524731580",
		"id741183306",
		"id1182341536",
		"id1447322621",
		"id529652920",
		"id1438385422",
		"id558207346",
		"id1277029359",
		"com.candywriter.bitlife",
		"com.sail.advanced.booster",
		"com.randomsaladgames.ginrummydeluxe",
		"com.aim.ginrummy",
		"ca.lbcstudios.hempire",
		"com.tutapp.idlezombies",
		"com.cmcm.live",
		"com.admarket.passionpuzzle",
		"com.neonplay.casualrollersplat2",
		"com.joinroot.root",
		"com.nextgames.android.twd",
		"com.nextgames.android.ourworld",
		"com.scopely.headshot",
		"com.telltalegames.walkingdead100",
		"com.telltalegames.walkingdead200",
		"veeu.watch.funny.video.vlog.moment",
		"com.fiogonia.yatzy",
		"com.noodlecake.zombieroadtrip",
		"net.mobigame.zombietsunami",
		"jp.co.cybird.appli.android.sgk",
		"com.vng.android.mps.dead.zombie2",
		"com.mars.avgchapters"
	],
	"source": {
		"ext": {
			"schain": {
				"ver": "1.0",
				"complete": 1,
				"nodes": [{
					"asi": "mintegral.com",
					"sid": "18837",
					"rid": "9148c67b-ff15-4be3-aaf1-6417b206be39",
					"hp": 1
				}]
			}
		}
	},
	"imp": [{
		"id": "1",
		"video": {
			"mimes": [
				"video/mp4"
			],
			"maxduration": 30,
			"protocols": [
				2,
				3,
				5,
				6
			],
			"w": 720,
			"h": 1280,
			"linearity": 1,
			"skip": 1,
			"maxbitrate": 2000,
			"delivery": [
				2
			],
			"pos": 7,
			"companionad": [{
				"w": 720,
				"h": 1280,
				"id": "1",
				"battr": [
					16
				],
				"pos": 7,
				"mimes": [
					"application/javascript",
					"image/jpeg",
					"image/jpg",
					"text/html",
					"image/png",
					"text/css",
					"image/gif"
				],
				"api": [
					3,
					5,
					6
				],
				"ext": {
					"orientation": 2
				}
			}],
			"api": [
				3,
				5,
				6
			],
			"companiontype": [
				1,
				2,
				3
			],
			"ext": {
				"orientation": 1,
				"videotype": 2,
				"videoendtype": [
					1,
					2
				],
				"endcardonly": true
			}
		},
		"tagid": "270606",
		"bidfloorcur": "USD",
		"displaymanagerver": "mi_6.2.0",
		"ext": {},
		"instl": 1,
		"secure": 1,
		"bidfloor": 3
	}],
	"device": {
		"ua": "Mozilla/5.0 (iPhone; CPU iPhone OS 14_0_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
		"ip": "180.211.19.226",
		"devicetype": 4,
		"make": "apple",
		"model": "iphone xs",
		"os": "ios",
		"osv": "14.0.1",
		"carrier": "45005",
		"connectiontype": 2,
		"ifa": "00000000-0000-0000-0000-000000000000",
		"didmd5": "",
		"dpidmd5": "",
		"geo": {
			"city": "dal-dong",
			"country": "KOR"
		},
		"dnt": 0,
		"lmt": 0,
		"language": "ko"
	},
	"regs": {
		"coppa": 0,
		"ext": {
			"gdpr": 0
		}
	}
}
```

### Banner
``` json
{
	"id": "2853cfd1-b043-4476-bfd3-f725ab2309df",
	"app": {
		"id": "120098",
		"name": "斗地主经典版-单机游戏",
		"cat": [
			"IAB9-7"
		],
		"ver": "1.9",
		"bundle": "1102002812",
		"publisher": {
			"id": "19316"
		},
		"storeurl": "https://apps.apple.com/cn/app/id1102002812?mt=8",
		"storeid": "com.zongyi.ndoudizhu"
	},
	"at": 2,
	"tmax": 600,
	"bcat": [
		"IAB26",
		"IAB24",
		"IAB25",
		"IAB11-4"
	],
	"source": {
		"ext": {
			"schain": {
				"ver": "1.0",
				"complete": 1,
				"nodes": [{
					"asi": "mintegral.com",
					"sid": "19316",
					"rid": "2853cfd1-b043-4476-bfd3-f725ab2309df",
					"hp": 1
				}]
			}
		}
	},
	"imp": [{
		"id": "1",
		"banner": {
			"w": 320,
			"h": 50,
			"battr": [
				17
			],
			"mimes": [
				"application/javascript",
				"image/jpeg",
				"image/jpg",
				"text/html",
				"image/png",
				"text/css"
			],
			"api": [
				3,
				5,
				6
			],
			"ext": {
				"orientation": 0,
				"adtype": 1
			}
		},
		"tagid": "159637",
		"bidfloorcur": "USD",
		"displaymanagerver": "mi_5.8.3",
		"ext": {},
		"secure": 1,
		"bidfloor": 0.15
	}],
	"device": {
		"ua": "Mozilla/5.0 (iPhone; CPU iPhone OS 12_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
		"ip": "117.136.79.94",
		"devicetype": 4,
		"make": "apple",
		"model": "iphone 6",
		"os": "ios",
		"osv": "12.4.1",
		"carrier": "46000",
		"connectiontype": 6,
		"ifa": "BC5EA04D-F3F9-470F-A1FA-885314C946E5",
		"didmd5": "",
		"dpidmd5": "",
		"geo": {
			"city": "shenzhen",
			"country": "CHN"
		},
		"dnt": 0,
		"lmt": 0,
		"language": "zh"
	},
	"regs": {
		"coppa": 0,
		"ext": {
			"gdpr": 0
		}
	}
}
```

### Splash
``` json
{
    "id": "5d9bf80a35304353b9ddb5fx",
    "imp": [
        {
            "id": "1",
            "banner": {
                "w": 1280,
                "h": 720,
                "battr": [
                    17
                ],
                "mimes": [
                    "application/javascript",
                    "image/jpeg",
                    "image/gif",
                    "text/css",
                    "text/html",
                    "image/png"
                ],
                "api": [
                    3,
                    5,
                    6
                ],
                "ext": {
                    "orientation": 2,
                    "adtype": 3
                }
            },
            "tagid": "1510000297",
            "bidfloor": 1,
            "bidfloorcur": "USD",
            "secure": true,
            "support_download": false,
            "request_type": 7,
            "displaymanagerver": "mal_99.0.0"
        }
    ],
    "app": {
        "id": "1510000000",
        "ver": "5.4.0",
        "bundle": "test.aladdin.1510000000",
        "publisher": {
            "id": "1500000000"
        },
        "storeurl": ""
    },
    "device": {
        "ua": "Mozilla/5.0 (Linux; Android 7.0; SM-N9208 Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.108 Mobile Safari/537.36",
        "geo": {
            "country": "CN",
            "city": "7391"
        },
        "dnt": false,
        "ip": "112.65.61.82",
        "devicetype": 4,
        "make": "",
        "model": "honor",
        "os": "android",
        "osv": "11.1.2",
        "language": "zh-Hans-US",
        "carrier": "46002",
        "connectiontype": 2,
        "ifa": "A0635584-FCB1-4106-B924-A80C29150E4D"
    },
    "at": 1,
    "tmax": 600
}
```