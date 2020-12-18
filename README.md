# リポジトリ情報

| Back end | Front end | Auth |
----|----|----
| Gin | Nuxt.js | Firebase Authentication SDK |

## Nuxt.js + Firebaseの参考

[Base Repository](https://github.com/davidroyer/nuxt-ssr-firebase-auth.v2)

[Baseの解説日本語](https://qiita.com/basho/items/acd6a17bb6e2a2f7a932#%E5%89%8D%E6%8F%90%E3%81%AE%E7%A2%BA%E8%AA%8D)

[一部参考記事](https://qiita.com/ririli/items/d0d3a6ae78c1b6e827fc#%E3%83%AD%E3%82%B0%E3%82%A4%E3%83%B3%E6%83%85%E5%A0%B1%E3%82%92%E8%A1%A8%E7%A4%BA%E3%81%97%E3%81%A6%E3%81%BF%E3%82%88%E3%81%86)

## Gin + Firebaseの参考

[全体像の把握で参考にした](https://qiita.com/po3rin/items/d3e016d01162e9d9de80)

[middlewareの設定参考](https://qiita.com/kannan_xiao4/items/31a42313e98a3e919772#firebase%E3%81%AE%E5%B0%8E%E5%85%A5)

[JWTリクエストのLocalStorage / Cookieの書き方参照](https://qiita.com/yoshinori_hisakawa/items/a2254edc9f9784f47453)

## ログインに関して
- email + passwordアカウントは、usernameとavatar情報を取得することができません。
- 上記をすべて満たせるのは、googleログインだけである。

