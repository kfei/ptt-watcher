# ptt-watcher

> 訂閱 PTT 文章, 發送至 Slack 或 Line.

## Demo

![demo-slack](http://i.imgur.com/d2hZdOF.jpg)
![demo-line](http://i.imgur.com/YnFL7rK.jpg)

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
  "notifications": {
    "slack": {
      "token": "YOUR_SLACK_TOKEN",
      "channel": "#ptt",
      "userName": "ptt-watcher"
    },
    "line": {
      "channelSecret": "YOUR_LINE_CHANNEL_SECRET",
      "channelAccessToken": "YOUR_LINE_CHANNEL_ACCESS_TOKEN",
      "toUserId": "USER_ID (you should get this from the webhook service)"
    }
  },
  "subscriptions": [
    {
      "name": "HardwareSale",
      "feedUrl": "http://rss.ptt.cc/HardwareSale.xml",
      "refreshTime": 60,
      "notifyMethods": ["slack"],
      "filters": [
        "賣 DDR3L",
        "賣 i7"
      ]
    },
    {
      "name": "DC_SALE",
      "feedUrl": "http://rss.ptt.cc/DC_SALE.xml",
      "refreshTime": 60,
      "notifyMethods": ["slack", "line"],
      "filters": [
        "賣 5D2"
      ]
    },
    {
      "name": "Beauty",
      "feedUrl": "http://rss.ptt.cc/Beauty.xml",
      "refreshTime": 600,
      "notifyMethods": ["line"],
      "filters": [
        "正妹"
      ]
    }
  ]
}
```

### General

  - **name**: 訂閱的板面名稱
  - **feedUrl**: PTT 各板面自己的 RSS feed 網址 (注意大小寫)
  - **refreshTime**: 更新周期 (單位秒)
  - **notifyMethods**: 推送通知的管道
  - **filters**: 查詢方式與 PTT `/` 語法相同

### Slack

  - **token**: 可在 Slack [網頁介面](https://api.slack.com/web)自行生成
  - **channel**: 訊息要發送到的 channel 名稱 (注意 `#` 符號)
  - **userName**: 要顯示的發送者名稱

### Line

  - **channelSecret**
  - **channelAccessToken**
  - **toUserId**: 接收通知的帳號 User ID（非帳號名）

> Line 的開發者網頁介面比較鳥，若有設定問題請開 issue。
>
> 取得 User ID 最快的方法是跑一個 server 接 webhook，一但有使用者跟 bot
> 說話就可以從 webhook request log 裡找到 User ID。連 web server
> 都懶得開的話也可以用 [RequestBin](https://requestb.in)。
>
> 若有更快的方法還請分享給我。:smile:
