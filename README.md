# Charge Notification

バッテリー残量が減るとSlackに教えてくれるGUI。

### 設定項目
- 通知する名前
- 通知するSlackの Webhook URL
- バッテリーを確認する間隔
- どのくらい以下になったら通知するか


## ダウンロードはこちらから（Mac用だけです）
https://github.com/mi11km/battery-notification/releases/tag/1.0.0


## git clone してからアプリ化する場合
クローン後に[こちら](https://wails.app/gettingstarted/installing/)を参考にwailsのcliをインストールし、プロジェクトルートで以下のコマンドを実行（前提としてGoがインストールされている必要があります。）
```
wails build -p
```
buildディレクトリにアプリのファイルが作成されます。
