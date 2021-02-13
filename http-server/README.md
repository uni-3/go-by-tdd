
### api definition

```
GET /players/{name} should return a number indicating the total number of wins
POST /players/{name} should record a win for that name, incrementing for every subsequent POST
```

### テスト作成の方向性

GETをテストするためにはplayerとそのスコア(PlayerStore)が必要になる
intergaceとして、作ると、stubでテストできる（実際のデータやstorage機構を気にしなくてもよくなる

POSTにはPlayerStoreがあることで、spyとして、仮の値を格納することができる
この、保存に関しては、GETとは関係なく実装できる

->storeの対象をmemoryにする

### integraton test

システム的に動作するかどうか
test pyramid的な意味でunit testの上位にある

