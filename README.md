# yamato_delivery_status

ヤマト配達状況確認プログラム
ファイルに記載された伝票番号の配達状況を一括で出力する


使用言語，動作確認
Go 1.21.1
MacOS
改行コードを\nとしているためwindows未対応の可能性あり

1.ヤマトの伝票番号をinput.txtに入力して以下を実行
```
go mod run
```

2.output.txtが生成され各行の伝票番号に対応する状況が記入されて出力される

入力形式，出力形式はinput_sample.txtとoutput_sample.txtを参照