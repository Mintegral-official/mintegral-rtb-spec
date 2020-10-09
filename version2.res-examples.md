# Mintegral RTB Specification ver 2.5: Response Examples

* [Banner](#banner)
* [Native](#native)
* [Rewarded-Video](#rewarded-video)
* [Interstitial-Video](#interstitial-video)
* [Splash](#splash)


### Banner

```json
{
  "cur": "USD",
  "id": "0000fbb5-0000-4218-8551-4ea553000000",
  "seatbid": [
    {
      "bid": [
        {
          "adid": "40734",
          "adm": "<script>(function(){var e=this||self;})()</script><div><img src=\"https://adsystem/?q=eD3TPt6ibXC7FHb5U3gXH6ylGqZ5kc355we1OGn87GXzNth1O7M1hYDg&pr=${AUCTION_PRICE}\" border=0 width=1 height=1 alt=\"\" style=\"display:none\"></div>",
          "adomain": [
            "ecnavi.jp"
          ],
          "cid": "75580",
          "crid": "40734",
          "h": 250,
          "id": "0",
          "impid": "00007d85-89e3-41e9-9e09-96e3a5000000",
          "language": "ja",
          "price": 0.014828,
          "w": 300
        }
      ]
    }
  ]
}
```

### Native

```json
{
  "cur": "JPY",
  "id": "78a69e4a-0000-4e94-8e14-a6e815000000",
  "seatbid": [
    {
      "bid": [
        {
          "adid": "9Kfj2l",
          "adm": "{\"native\":{\"ver\":\"1.1\",\"assets\":[{\"id\":1,\"title\":{\"text\":\"今すぐお届けできます\"}},{\"id\":4,\"data\":{\"value\":\"東陽2丁目店\"}},{\"id\":3,\"data\":{\"value\":\"出前\"}},{\"id\":2,\"img\":{\"url\":\"https://adsystem/img.png\",\"w\":160,\"h\":160}}],\"link\":{\"url\":\"https://adsystem/landing-url\"},\"imptrackers\":[\"https:///wp=${AUCTION_PRICE}\"]}}",
          "adomain": [
            "ecnavi.jp"
          ],
          "attr": [],
          "cid": "JUgtXt",
          "crid": "9Kfj2l",
          "id": "201912",
          "impid": "0000b361-02ce-4a93-873e-02d943000000",
          "price": 16.281675070724084
        }
      ]
    }
  ]
}
```

### Rewarded-Video
``` json
```

### Interstitial-Video
``` json
```

### Splash
``` json
```