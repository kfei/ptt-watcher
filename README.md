# ptt-watcher

訂閱 PTT 文章, 發送至 Slack.

[Demo](#screenshot)

## Highlights

 * 可自訂多組關鍵字查詢條件過濾感興趣的文章
 * 針對不同板面分別設定訂閱條件
 * 透過 RSS feed 更新資料無需帳號密碼

## Quickstart

直接透過 `go get` 安裝:

```bash
go get github.com/kfei/ptt-watcher
```

編輯設定檔於運行目錄下, 參考[範例](#sample-config):

```bash
vim config.json
```

跑起來:

```bash
ptt-watcher
```


## Sample config

```json
{
  "slackToken": "YOUR_SLACK_TOKEN",
  "slackChannel": "#ptt",
  "slackUserName": "ptt-watcher",
  "subscriptions": [
    {
      "name": "HardwareSale",
      "feedUrl": "http://rss.ptt.cc/HardwareSale.xml",
      "refreshTime": 60,
      "filters": [
        "賣 DDR3L",
        "賣 i7"
      ]
    },
    {
      "name": "DC_SALE",
      "feedUrl": "http://rss.ptt.cc/DC_SALE.xml",
      "refreshTime": 60,
      "filters": [
        "賣 5D2"
      ]
    },
    {
      "name": "Beauty",
      "feedUrl": "http://rss.ptt.cc/Beauty.xml",
      "refreshTime": 600,
      "filters": [
        "正妹"
      ]
    }
  ]
}
```

  - **slackToken**: 可在 Slack [網頁介面](https://api.slack.com/web)自行生成
  - **slackChannel**: 訊息要發送到的 channel 名稱 (注意 `#` 符號)
  - **slackUserName**: 要顯示的發送者名稱
  - **feedUrl**: PTT 各板面自己的 RSS feed 網址 (注意大小寫)
  - **refreshTime**: 更新周期 (單位秒)
  - **filters**: 查詢方式與 PTT `/` 語法相同

## Screenshot

![demo](http://i.imgur.com/d2hZdOF.jpg?1)
