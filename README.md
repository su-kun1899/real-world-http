# Real World HTTP 読書ログ

## １章 HTTP/1.0のシンタックス：基本となる４つの要素

- HTTPの基本要素
  - メソッドとパス
  - ヘッダー
  - ボディ
- (1990) HTTP/0.9
  - シンプル
  - 現行との後方互換性はない
  - ボディの受信とパスだけ
- (1992) HTTP/1.0の前身
  - リクエスト
    - メソッドの追加
    - リクエスト時のHTTPバージョンの追加
    - ヘッダーの追加
    - リクエストにもコンテンツを含められる
      - メソッドによっては無視すべきものも
  - レスポンス
    - HTTPバージョンの追加
    - ステータスコードの追加
    - レスポンスヘッダの追加

### HTTPの仕様

- RFC: HTTPのルール文書
- IETF: RFCを作っている団体
- IANA: Webに関するデータベースを管理する団体
  - ポート番号やContent-type
- W3C: ウェブ周りの標準化団体
- WHATWG: ウェブ周りの企画を議論する団体
  - 🤔 W3Cと対立？
- 🤔 RFCとW3Cを参照するのが基本なのかな

### ヘッダ

- メールと同じ形式のヘッダ
  - フィールド名: 値
- リクエストヘッダ
  - User-Agent: クライアントアプリケーション名
  - Referer: リクエスト時に見ていたページのURL
  - Authorizaton: 認証情報
- レスポンス
  - Content-Type: ファイルの種類（MIMEタイプ）
  - Content-Length: ボディのサイズ
  - Content-Encoding: 圧縮されてる場合の圧縮形式
  - Date: ドキュメントの日時
- `X-` から始まるヘッダーは各アプリケーションが自由に使っていい
- IANAで登録済ヘッダーの一覧が見られる
- 同じヘッダは複数送れる
  - プログラミング言語やフレームワークに寄って正規化の方法等が異なるため、注意が必要

### メソッド

- GET: ヘッダーとコンテンツを要求
- HEAD: ヘッダーのみを要求
- POST: 新しいドキュメントを投稿
- PUT: 既存のドキュメントを更新
- DELETE: ドキュメントの削除
- 😇 いなくなったメソッドもいっぱい

### ステータスコード

- 100番台: 処理中
- 200番台: 成功
- 300番台: 正常処理としての命令。リダイレクトやキャッシュの利用など
  - 😳 ログイン後のページ遷移は303なのか
- 400番台: リクエスト不正
- 500番台: サーバー側のエラー

### URL

- 😵 URI/URNよく分からん
  - URLはドキュメントなどリソースの場所を特定する手段を提供する（住所
  - WebのシステムではURLとURIはほぼ同一
  - 😵 URLって言っておけばいいかな。。
- URLの構造: `スキーマ://ユーザ:パスワード@ホスト名:ポート/パス#フラグメント?クエリ`
- 仕様上長さ制限はないがIE/Edgeの制限あり
  - 2000文字が目安
  - `414 URI Too Long` がある
- 🐷 ﾌﾟﾆｺｰﾄﾞ

### その他

- ニュースグループという思い出（知らん
- `X-Content-Type-Options: nosniff` を使うとブラウザにContent-Typeの推測をさせなくできる
- curl、`-d @ファイル名` でファイルの中身をリクエストボディにできるのか

## ２章 HTTP/1.0のセマンティクス：ブラウザの基本機能の裏側

- URLエンコード(x-www-form-urlencode)
  - RFC3986とRFC1866で半角スペースのエンコードが異なる(%20と+)
  - ブラウザはRFC1866、curlの `--data-urlencode` はRFC3896
  - 🤔 昔ハマった記憶。。リライトルールとかだったかな。。
- ファイルの送信(multipart/form-data)
  - formのenctypeに指定する
- Formを使ったリダイレクト
  - パラメータをGETで引き回すのに難があるとき
  - formをonloadでsubmitする
  - 🤔 荒業感
- コンテントネゴシエーション
  - 言語やMIMEタイプをAccept系のリクエストヘッダに乗せる
  - charset
    - HTML5ならmetaタグでOK
- 圧縮
  - Brotli: Googleが公開した新しいアルゴリズム
  - 🧐 brとgzipとdeflateくらい抑えておけばOKだろうか
- Cookie
  - 😣 httponlyの使い所わかんね。。
- プロキシとゲートウェイ
  - 通信内容を改変する（できる）のがプロキシ
  - そのまま転送するのがゲートウェイ
- キャッシュ
  - GETとHEAD以外は基本的にキャッシュしない
  - 304 Not Modified
  - Last-Modifiedヘッダは通信が発生する、Expiresヘッダは通信しない
  - 🙄 ETagも結局更新日時（とファイルサイズ）が慣例になったのん？
- ⚠️ クライアントからのキャッシュコントロールは積極的に行うべきではない
  - `Pragma: no-cache` とか
  - セキュア通信だとプロキシは通信内容を監視できない
  - HTTPはステートレスをよしとしている
- リファラーポリシー
- robots.txt
  - RFCにはなっていない

## ３章 Go言語によるHTTP/1.0クライアントの実装

- 😵 `go build` 使わずに、いつも `go install` しちゃってるんだけど、だめかしら

## ４章 HTTP/1.1のシンタックス：高速化と安全性を高めた拡張

- SSLが進化してTLSになった
  - 🧐 SSLで通じるけど、TLSと呼ぶのが正しい？
- ファイルなどの破損をチェックするためのハッシュ値をチェックサムやフィンガープリントと呼ぶ
  - MD5やSHA-1はセキュリティ用途には非推奨だが、チェックサムとしてはまだ利用されている
- 共通鍵方式
  - 双方で鍵を共有する
- 公開鍵方式
  - 公開鍵は南京錠、秘密鍵がその鍵
  - 南京錠をかけて返送する
- デジタル署名
  - 南京錠の方が秘密鍵になる方式
  - 🙄 よく分からん
- CONNECTメソッド
  - 🤔 使い所はどんな時になるんだろう
- チャンク方式
  - 🤔 ブラウザから送れないのはどうしてかな

## ５章 HTTP/1.1のセマンティクス：広がるHTTPの用途

- HTTPはまさにインフラである

### ファイルのダウンロード

- `Content-Disposition` ヘッダを使うとファイル保存ダイアログを表示させられる
  - 日本語ファイル名も利用できる（RFCでエンコードルールが決まっている）
  - `inline` を指定するとブラウザのインライン表示を明示できる
  - metaタグのrefleshと組み合わせてダウンロード完了ページを表示できる
    - 「たいていこのページには広告がたくさん貼られています」
    - ﾜﾛﾀ
- 範囲指定をすることでダウンロードの中断・再開ができる
  - サーバーは範囲指定可能な場合 `Accept-Ranges` ヘッダをレスポンスに含める
  - クライアントは `Range` ヘッダで範囲指定する（ `Range: bytes=1000-1999` ）
  - ファイルが途中で変更される可能性があるので、 `Etag` ヘッダを利用する
    - `If-Range` ヘッダを使えば、ファイル片が無効になっている場合、通常のGETになる
  - 複数範囲指定もできる（ `Range: bytes=500-999,7000-7999` ）
    - 🤪 なんの意味があるのん？
  - Rangeヘッダを使えば並列ダウンロードが可能になるが、今はあんまりっぽい
    - サーバの負荷が高い
    - 途中の回線がボトルネックなら意味がない
    - CDNの普及
    - 動画は一部先読みでいい
    - 大きなファイルを高速にダウンロードする場合はBitTorrentなどを使うべき
      - 複数ノードに並列アクセスする方式

### XMLHttpRequest

- JavaScriptからリクエストする
- 2つバージョンがあるので注意(level2が新しい？)
- ブラウザ(HTML)との違い
  - 画面のリロードがない(Ajax)
  - GETとPOST以外も使える
  - JSONや画像などのバイナリ他、様々なフォーマットが送受信できる
  - セキュリティ上の制約がある
- セキュリティ
  - CookieはhttpOnlyがあるとアクセスできない
  - 別ドメインにはリクエストできない
  - CONNECT/TRACEメソッドは使えない
  - ヘッダにも制限がある
  - 🧐 悪意のある送受信を避けるためのキーワードといえるのかな

### COMET

- XMLHttpRequestを利用した双方向通信
- ポーリング
  - クライアントが頻繁に問い合わせる
- ロングポーリング
  - サーバー側でレスポンスを保留しておく
- レガシーな仕組み
  - 動作する環境は多いが、オーバヘッドが大きい

### Geo-Location

- 位置情報
- スマホのGPS
  - ユーザーの許可が必要
  - ユーザーが意識してくれるサービスでないと使いづらい
- Wifi
  - アクセスポイントのBSSIDのデータベースを利用する
  - 自動収集いかがなものか
- GeoIP
  - IPアドレスを地域と紐付けたデータベース
  - 正確性は劣るがユーザーの許可がいらない

### X-Powered-Byヘッダー

- サーバーがシステム名を返すヘッダー
- RFCにはないがデファクトスタンダードだった
- HTTP1.1ではRFCでServerヘッダーが定義されている
- 脆弱性の露呈に繋がるリスクがある
- 互換性の支援に利用できる
- HTTPの重要性は高まっているので今後はあまり必要ない？

### リモートプロシージャコール(RPC)

- 別のコンピュータにある機能を自分のコンピュータの機能のように呼び出す
- HTTPをベースにするRPCがいくつかある
- XML-RPC
  - メソッドはPOST
  - 引数も戻り値も text/xml
  - レスポンスは200で返しエラーはXMLで表現する流儀
- SOAP
  - XML-RPCの拡張
  - SOAPはデータ表現フォーマットでSOAP-RPCが定義されている
  - 可搬性を重視した結果、複雑化した
  - 😢 SOAPは人類の悲しい思い出と理解した
- JSON-RPC
  - シンプルに、という方針
  - 仕様は独自に定義している
  - TCP/IPソケットに使用することも想定している
  - 基本はPOST、冪等ならGETも可
  - ステータスコードは使い分ける
  - Batchモードで複数プロシージャ呼び出しが行なえる

### WebDAV

- HTTPを拡張したファイルシステム
- 同期型なので、ネットワークがないと使えない
  - Dropbox他はローカルにコピーがある
- Gitは転送用プロトコルにHTTPSとSSHをサポート
  - HTTPSの中ではWebDAVを使用
    - セットアップが簡単
  - SSHは独自のGitプロトコル
    - 通信速度で優る

### 認証と認可

- 認証
  - ユーザーが「何者か」を確認する
- 認可
  - 認証したユーザーにどこまで権限を与えるか決定する
- これからやるならOAuth2.0かOpenID Connect

#### シングルサインオン

- 一度サインインしたら各システムがすべて有効になるような仕組み
- プロトコルやルールがあるわけではなく、用途の分類名
- 各システムが認証サーバにアクセスしてログインを代行する方法
  - IDを一元管理する
  - 各システムでIDとパスワードは入力する
  - 😇 シングルサインオンじゃないやん
- プロキシが認証を代行する方法
- 各サービスに認証代行のエージェントを入れ、認証サーバに問い合わせる方法

#### Kerberos認証

- LDAP
  - RFCで定義されているユーザ管理の共通規格
  - 実装にはOpenLDAPやActive Directory
  - LDAPにはKerberos認証が広く使われている
  - 🤪 ここよく分かっていないので、話したい

#### SAML(Security Assertion Markup Language)

- HTTP/SOAPを前提としたシングルサインオンの仕組み
- octaやOneloginのようなSaSSがある
  - 🤔 導入が検討されていたような
- 規格はOASISで策定されている
- ドメインをまたいだサービス間でシングルサインオンできる
  - G Suite, Kintone, Office365, Dropbox と連携可能
- 実装には6種類の方法がある
  - SAML SOAP バインディング
  - リバース SOAP (PAOS) バインディング
  - HTTP リダイレクト バインディング
  - HTTP POST バインディング
  - HTTP アーティファクト バインディング
  - SAML URI バインディング
- HTTP POST バインディング
  - 認証プロバイダにサービスを登録
  - サービスにプロバイダを登録
  - 登録はメタデータXML
  - サービスは認証されていなかったら認証プロバイダにリダイレクト
  - 認証プロバイダの画面でログイン
  - 認証プロバイダはログイン情報をサービスにPOST

#### OpenID

- すでに登録されているWebサービスのユーザ情報を使い他のサービスにログインできる仕組み
- 多くは後続のOpenID Connectを利用している
- OpneIDプロバイダ(ユーザ情報を持っているサービス)はユーザの識別子をURL形式で提供する
- ユーザーは入力識別子でリライングパーティー（ログインしたいサービス）にログインする
- OpenIDプロバイダにリダイレクトし、許可を承認する
- 🤔 リライングパーティーが識別子で問い合わせるイメージなのよねきっと

#### OpenSocial

- Facebookに対抗するためにGoogleとMySpaceがソーシャルネットワークの共通APIとして開発
  - 😇 マイスペなつい
- 認証だけでなく、プラットフォームを指向
- mixi、モバゲー、GREEなど、日本で流行った
- UIはルールにそってXMLなど、提供側（SNS）と密結合
- 💀 今は虫の息？

#### OAuth

- 認証ではなく認可の仕組み
- 2012年に2.0がRFCに公開された
- 今も活発に使われている
- 登場人物
  - 認可サーバ
    - IDプロバイダー。アカウントがある
  - リソースサーバ
    - ユーザーが許可した範囲でアクセスできる対象。TwitterやFacebookだと認可サーバと一緒
  - クライアント
    - ユーザが使うサービスやアプリケーション
    - 認可サーバに登録が必要
      - クレデンシャル(client\_idとclient\_secret) を取得
- OpenIDよりも強力
  - クライアントは認可されたものは取得できるため
  - スコープは認可サーバ側で後から設定変更できる
- フロー
  - [ここ](https://qiita.com/TakahikoKawasaki/items/200951e5b5929f840a1f) が分かりやすかったかも
  - Authorization Code
    - 通常のフロー
    - クライアントがコンフィデンシャルで認可サーバに認可コードを発行してもらう
    - 認可コードとトークンを認可サーバに交換してもらう
    - トークンでリソースにアクセスする
  - Implicit Grant
    - アプリケーションがユーザのローカルなど外部にある場合に使うフロー
      - 🙄 コンフィデンシャルが晒されちゃうので？
      - 😱　クライアントが保証されないので危険？
    - スマホアプリとか、JavaScriptとか
  - Resource Owner Password Credentials Grant
    - クライアントがID/PASSWORDを入力させて認可サーバにリクエストする
      - クライアントがIDとパスワードに触れるということ
      - 認可サーバが信頼しているクライアントのみ
    iOS内臓のFacebookやTwitter連携など特殊なケース
  - Client Credentials Grant
    - クレデンシャルだけで認可サーバにトークンを発行してもらう
    - ユーザの認証はしないフロー
- 🤔 クレデンシャルで認可サーバがクライアントを認証するって感じかな

#### OpenID Connect

- OAuth2.0をベースに認可だけでなく認証もできるように拡張
- 今後のデファクトスタンダードになる
- ユーザ情報の取得が規格化された
  - アクセストークンとは別にIDトークンが発行される
- 認可エンドポイントが認可コードを返し、トークンエンドポイントがアクセストークンとIDトークンを返す
  - クライアント認証しない場合は、認可エンドポイントがトークンを返す
- フロー
  - Authorization Code Flow
    - OAuthのAuthorization Codeと一緒
  - Implicit Flow
    - 認可エンドポイントへのリクエストにnonceが必須になってOAuthより安全になった
    - 🙄 使い方が分からん。。アプリからのリクエストと、認可ページからのリクエストのnouncewを比較するとかかな。。？
  - Hybrid Flow
    - クライアントが複数あるような場合のフロー
      - スマホアプリのようなクライアント認証できないクライアント
      - バックエンドのウェブサービス
    - 😵 で、どんなフローになるのよ

## 7章 HTTP/2のシンタックス: プロトコルの再定義

- HTTP/2はALPNを使って、HTTP/1.1とはまったく別のプロトコルとして扱われる
  - 😳 マジか
- 高速化
- テキストベースからバイナリベースへ
  - 並列化に強くなった
- フレーム単位で送受信
  - 🤪 この辺難しい。。
- フローコントロール
  - 通信先のバッファサイズだけ送る
  - 通信速度差がありすぎる場合に、通信先を潰さないようにするため
- サーバープッシュ
  - 優先度の高いコンテンツを要求される前に送信
  - リクエストを送信するまでプッシュを検知できない
  - 🙄 結構無茶苦茶言ってるような気がするけど。。
- HPACK
  - ヘッダの圧縮方式
  - ヘッダは名前がよく被るので、辞書を事前に持っている
- SPDY
  - Googleが開発したHTTP/2の前身
- QUIC
  - UDPの上で動作するプロトコル
  - Google開発し、Chromeには搭載されている
  - 🤔 UDPの強いやつ考えればいいのかな。。
- Fetch API
  - XMLHttpRequestと同様のサーバーアクセスを行う関数
  - CORSが扱いやすくなった
    - Cross-Origin Resource Sharing
  - JSのモダン非同期処理であるPromiseに準拠
  - キャッシュが制御可能
    - default
      - 標準的なブラウザと同じ
    - no-store
      - キャッシュがないものとしてリクエスト、結果もキャッシュしない
    - reload
      - キャッシュがないものとしてリクエスト、可能なら結果もキャッシュする
    - no-chache
      - キャッシュがあってもリクエストを送る
      - ETagを送るので、304を受け取った場合キャッシュしたコンテンツを利用
    - force-cache
      - 期間外でもキャッシュがあれば利用する、キャッシュがない場合HTTPリクエストを送信
    - only-if-cached
      - 期間外でもキャッシュがあれば利用する、キャッシュがなければエラー
    - reload
  - リダイレクトが制御可能
    - follow
      - 最大20までリダイレクトをたどる
    - manual
      - リダイレクトを辿らず、リダイレクトがあった旨だけ伝える
    - error
      - リダイレクトがあった場合、ネットワークエラーにする
  - リファラーのポリシーを設定可能
  - Service Workerから利用可能
    - XMLHttpRequestとの一番大きな差
    - ブラウザが Web ページとは別にバックグラウンドで実行するスクリプト
    - 🙄 Service Workerなんだろ
  - 対応するデータ型
    - ArrayBuffer(固定長のバイナリ)
    - Blob（ファイル等のバイナリ）
    - FormData（HTMLフォーム）
    - Object（JSON）
    - string（Text文字列）
  - 利用可能なメソッド
    - CORS安全
      - GET,HEAD,POST
    - 禁止
      - CONNECT,TRACE,TRACK
  - CORSモードが指定可能
    - cors,same-origin,no-cors
  - 厳格な設定がデフォルトで、明示的に解除するという思想
    - ☺️ よい
  - 🤔 Fetch APIができたのに、nuxt.jsでaxios使うのが推奨されてるっぽいのはなんでだろー
- Server-Sent Events
  - HTML5の機能
  - サーバからイベントを通知する
    - Chunked形式の応用
  - MIMEタイプは `text/event-stream`
  - dataの内容はなんでもいいが、JSONがお多い
  - イベントストリームのタグ
    - id
    - event
    - data
    - retry
  - クライアントの再接続時はヘッダーにidを付けて送信
    - `Last-Event-ID`
    - 続きから受信できる
- WebSocket
  - サーバー/クライアント間の双方向通信
  - フレーム単位での送受信
    - 相手が決まっているのでボディだけ送っているようなもの
  - ステートフル
    - オンラインゲームのような遅延が許容されない場合に有効
  - 通信は必ずクライアントから始まる
    - Listen
      - サーバーが起動
    - Connect
      - クライアントが通信開始の宣言
    - Accept
      - クライアントからの接続依頼をサーバーが受け入れ
    - サーバーにソケットクラスのインスタンスが渡される
    - サーバーが受理、クライアントのソケットの送受信が有効になる
  - WebSocketはプロトコルのアップグレードを使用
    - HTTPでスタートした後WebSocketにアップグレード
  - Socket.IO
    - WebSocketを使うためのライブラリ
    - 後方互換性に強い
    - WebSocketのブラウザへの実装が進んだので、今後はあまり使われなさそう
- WebRTC(Web Real-Time Communication)
  - ブラウザ-サーバー間だけでなく、ブラウザ同士のP2P通信でも利用可
  - TCPではなくUDPがメイン
  - ユースケース
    - ビデオ通話
    - スクリーン共有
    - IP電話端末
    - etc
  - リアルタイム性が高いので1対多のリアルタイム映像配信にも利用可
    - HTTPのストリーミングは10~30秒程度遅延、WebRTCなら数秒で済む
  - P2PでのCDN
    - 同じ動画を見ているユーザとWebRTCのセッションを共有して配信する
  - RTCPeerConnection
    - WebRTCのベースはIP電話
    - IP電話で使われてる技術をまとめてJavaScriptのAPIを定めたもの
    - SDP(Session Description Protocol: セッション記述プロトコル)
      - P2Pのネゴシエーション時に使用
      - SDPのフォーマットで利用可能なコーデック情報やIP、ポートを共有
      - 渡す手段は定められていない
    - ICE (Interactive Connectivity Establishment)
      - NATを超えてP2Pの接続を確立する
        - NAT: プライベートIPをグローバルIPに変換する
        - 違うネットワークでP2Pする場合、NAT超えをする必要がある（ローカルIPは外部から分からないので）
        - STUNかTURNサーバどちらかが必要
          - STUNはアクセス元のグローバルIPを教えるサーバ
          - STUNにアクセスするとグローバルIPが分かるので、通信相手に伝えられる
          - TURNはSTUNが使えない時に、TCPで中継してくれるサーバ
          - 🤔 ローカルIPは分からなくていいのかな
  - メディアチャンネル
    - 実際の通信で音声やビデオを扱う
    - navigator.mediaDevices.getUserMedia()
      - ウェブカメラやオーディオデバイスの設定と取得を行うブラウザAPI
    - DTLS
      - WebRTCはDTLS上のデータ通信
      - D（データグラム）はUDPと同じ＝>暗号化したUDPがDTLS
      - UDPは再送も整列もしない分TLSよりも高速
        - WebRTCは遅くならない特性が必要
    - WebRTC MediaチャネルはDTLS上でSRTPというプロトコルを使用
      - セキュア(S)なリアルタイムプロトコル(RTP)
      - UDPに順序の整列を追加
  - データチャンネル
    - 音声以外をP2Pで送受信する
      - ファイルをチャット相手に送る、etc
    - SCTP(Stream Control Transmission Protocol)
      - DTLS上で使用
      - UDPよりかTCPよりか信頼性を設定できる
        - パフォーマンスはトレードオフ
        - ファイルなどは破損が問題になるため
      - 配送順序も設定で変更可能
- ORTC(Object Real-Time Communication)
  - WebRTCをより進化させようという活動
  - Edgeブラウザ版のSkypeで使われている
    - 他のブラウザは実装していない
  - WebRTC自体が活発に更新されている
  - ORTCは今のところ「気にする必要はない」
- HTTPウェブプッシュ
  - プッシュ通知を提供する仕組み
  - ブラウザが起動していなかったり、オフラインでも通知可能
  - ServiceWorkerがアプリケーションとの仲介になることで実現
  - サービスはオプトインで、ユーザの許可が必要
  - ChromeとFireFoxが実装済
    - RFCに必ずしも準拠していない
    - HTTPではなくWebSocketを使ったりしている
  - 流れ
    - ブラウザがプッシュサービスに購読を申し込む
      - レスポンスは201 Created
    - アプリケーションサーバーがプッシュサービスにメッセージを投稿
      - プッシュサービスが受け取ったら 201 Created(クライアントが受け取ったかは分からない)
      - `Prefer: respond-async` ヘッダをつけるとクライアントへの通知の成功まで待つ
        - レスポンスは 202 Acceptedになる
    - ブラウザがプッシュメッセージを受信
      - メッセージ受信のリクエスト
        - レスポンスにボディはない
        - HTTPサーバープッシュで、その後のレスポンスで受け取る
  - `Urgency` ヘッダでメッセージの緊急度を設定できる

## 8章 HTTP/2のセマンティクス: 新しいユースケース

- レスポンシブデザイン
  - かつての主流はユーザエージェントによるモバイル向けコンテンツ配信
  - レスポンシブデザインは様々なサイズのスクリーンやタブレット端末の縱橫表示に適切に表現できる
  - CSSピクセル
    - 実際の解像度とは異なるブラウザ上の論理解像度（スクリーンサイズ）
  - デバイスピクセルレシオ
    - デバイスピクセルとCSSピクセルの比率
  - JavaScriptやCSSで条件分けができる
  - モバイルブラウザはPC専用サイトも拡大縮小比率を変更しようとする
    - metaタグで無効にする
  - srcset属性
    - 画像をデバイスピクセルに合わせて選択できるimgタグの属性
    - CSSでも実現可能

### セマンティックウェブ

- ウェブを「テキスト」や「文書」ではなく「意味」を扱えるようにしようという運動
- 本来の目的はインターネット全体をウィキペディアのように繋いで、巨大なナレッジベースにすること
- 現在は検索エンジンのメタ情報提供のソリューションになっている
- RDF(Resource Description Framework)
  - セマンティックウェブのデータフォーマットのひとつ
    - 今はRDF/XMLと呼ばれる
    - 現在のインターネットではあまり見かけない
  - XMLに「主語」「術後」「目的語」で要素と属性と関係を表す
  - XMLなので、ページの外部に配置する必要があった
- ダブリンコア
  - 本・書籍・音楽などの著作物で汎用的に使えるメタデータ
  - 電子書籍のEPUBファイル等で使用されている
  - さまざまなセマンティックウェブ関連技術に利用されている
- RSS
  - RDFの応用例
  - ウェブサイトの更新履歴のサマリーに関するボキャブラリー
  - 一時期流行った
  - 🙋‍♀️ Googleリーダー使ってた
- マイクロフォーマット
  - HTMLのタグとクラスで表現する
    - RDFよりもポータビリティは高い
    - しかしCSSの名前衝突が起きやすい
  - schema.orgのフォーマットの方が推奨されている
- マイクロデータ
  - HTMLに埋め込み可能なセマンティックの表現形式
  - W3Cで定義
  - マイクロフォーマットと違い、HTMLの属性は独自
    - itemscope, itemtype, itemprop
  - 検索エンジンがサポートしてる
- RDFの逆襲
  - 多くの派生フォーマットが生まれた
  - マイクロフォーマットやマイクロデータより標準ボキャブラリが豊富
  - Google推奨のJSON-LDフォーマットが今後流行りそう
    - scriptタグでHTMLに記述可能

### オープングラフプロトコル

  - ソーシャルネットワークで使われるメタデータ
  - Facebookが開発した
  - リンクをシェアする時プレビューされるやつ
  - よく使う要素
    - タイトル
    - 種類
    - (canonicalな)URL
    - 画像
    - description
  - Twitterカード
    - オープングラフプロトコルをベースにTwitterの情報を追加したもの
    - オープングラフプロトコルと重複している要素もある

### AMP(Accelerated Mobile Pages)

- モバイル高速化の仕組み
- カルーセル形式でダイジェスト表示されたりもする
- コンテンツ（ニュース、レシピ、製品ページ、ブログなど）で大きな力を発揮
  - 動的単一ページ（地図の経路案内、ソーシャルネットワークなど）ではあまり効果的ではない
- ページの構成を固定化する
- コンテンツをCDNに乗せる
- Googleはmetaタグを書き換えてさらに高速化（AMPの仕様外）

### HTTPライブストリーミング（HLS）

- 動画のライブストリーミング再生の仕組み
- 2009年にAppleが提唱
- 標準化作業は停滞している
- モバイル向けではデファクトスタンダード
- videoタグにマスターのマニフェストファイルを定義する（.m3u8ファイル）
  - 先頭で推奨の帯域幅と対応するファイルのリストを列挙
    - 音声だけのストリーミングに切り替えさせることも可能
  - クライアントは最初に書かれたインデックスを使用
    - 回線の余裕を見て切り替える
    - 最初以外は順序で動作に影響はない
- 字幕用の.m3u8ファイル
  - ルートの.m3u8ファイルから字幕ファイルを束ねる.m3u8ファイルがリンクされる
    - WEBVIT
      - W3Cで定義されているウェブ標準の字幕用スクリプト
- 動画ファイル
  - マスターの.m3u8ファイルから参照される
  - ライブ配信できる
  - ビデオ・オン・デマンド（VOD:最初から再生）できる
  - 中継のみの場合、再生が終わった動画を削除する
- HLSのメリット
  - HTTPで動作するので専用サーバがいらない
  - 設定はMIMEタイプとキャッシュの生存期間のみ
- HLSのデメリット
  - 実際にはストリーミングではなくプログレッシブダウンロード
    - 遅延が30秒ほど発生する
  - サポートされる環境が少ない
    - 特にデスクトップ
- HLS前後の歴史
  - 専用アプリケーションによるストリーミング
    - Windows Media Player
    - QuickTime
  - FlashプレイヤーがサポートするRTMP（リアルタイム・メッセージ・プロトコル）
    - HTTPとは異なる
    - Youtube
    - ニコニコ動画
  - MicrosoftのSilverlight
  - AppleによるHLSリリースとHTTPプログレッシブダウンロード
    - MicrosoftのSmooth Streamingプロトコル（Silverlightに追加）
    - AdobeのHDS（HTTPダイナミックストリーミング）
    - FlashプレイヤーにHLSサポート追加
    - モバイルAndroidにHLSサポート追加

### MPEG-DASHによる動画ストリーミング再生

- MPEG-DASH
  - Dynamic Adaptive Streaming over HTTP
  - Apple以外のブラウザベンダーによる標準化
  - 目指す方向性はHLSとほぼ同じ（プログレッシブダウンロードを核としたストリーミング）
- video.js
  - リファレンス実装
- MPEG-DASHとHLSの再生方法の違い
  - HLSはブラウザが.m3u8ファイルを解釈して再生するシステム
  - MPEG-DASHはデータの解析をJavaScriptで行い、再生はHTML5 Media Source Extensionsを利用
  - HLSとの差はなくなってきている
  - 2016年、AppleがMPEG-DASHに譲歩する発表を行った
- Media Presentation Description(MPD)
  - MPEG-DASHのマニフェストファイル
    - 拡張子は.mpd
    - 実態はXML
  - 複数ファイルに分割可能
  - HLSはシンプルだが、MPEG-DASHはかなり複雑
  - MP4BoxというツールでMPDファイルを作れる
  - AWSのElastic TranscoderでMPEG-DASH、HLSの両方に変換できる

## 10章 セキュリティ: ブラウザを守るHTTPの機能

### 従来型の攻撃

- ブラウザを狙った攻撃以外の攻撃、の意
- マルウェア
  - コンピューターウイルス
  - ワーム
- ユーザーが意図せず起動することを目論む
- キーロガー
  - キー操作を記録してパスワード等を盗む
- トロイの木馬
  - バックドアを設けて外部から操作する
- 今でも重大な脅威の1つである

### ブラウザを狙う攻撃の特徴

- ブラウザ自体はOSの領域で何かをするわけではない
- 他のサービスとの窓になっている
  - ログイン情報
  - 他人との会話やプライバシー
  - インターネットバンク
- ブラウザの脆弱性だけでなく、サービス側のHTMLやJavaScriptも標的になる

### クロスサイトスクリプティング(XSS)

- 入力内容をそのまま出力する場合、スクリプトを埋め込んで実行される
- 他の攻撃の起点になりうるため、もっとも危険
  - クッキーの漏洩
  - セッションハイジャック
  - 入力内容の転送
- サニタイズで対策する
  - 出力長駆前にエスケープする
    - HTMLでの表示
    - SQLや外部コマンドの引数
  - 適切なライブラリを使えばミスを減らせる

#### 漏洩を防ぐためのクッキーの設定

- httpOnly属性を付与すればJavaScriptからはアクセスできなくなる

#### X-XSS-Protectionヘッダー

- 明らかに怪しいパターンを検出するヘッダー
- 非公式（X-）のヘッダー
  - IE, Chrome, Safariが対応
  - FireFoxはコンテントセキュリティポリシーがあるから対応しない方針
- パターンマッチによる判定のため偽陽性となる可能性がある

#### Content-Security-Policyヘッダー

- ウェブサイトで使える機能を細かくON/OFFできる
  - JSからの通信先やiframeで利用できるURLなど
  - インラインスクリプトを制限すればXSSを抑制できる
- W3Cが定義
- 検査はブラウザが行い、結果はサーバー側で分からない
  - report-uriを指定すると結果が送信される
- 一括設定も可能
- 利便性を落とすリスク
  - Content-Security-Policy-Report-Onlyヘッダー
  - チェックだけするモード

#### Content-Security-PolicyとJavaScript製テンプレートエンジン

- JavaScriptが発展して、ブラウザ環境でやることが増えた
- いろいろなテンプレートライブラリ
  - Hogan.js
  - Vue.js
  - Riot.js
- 動的な関数生成はCSPのunsafe-evalに引っかかる
  - プリコンパイルが必要

#### Mixed Content

- HTTP素材が混ざっている Mixed Content
  - 広告や外部サービスのコンテンツ
- HTTPSに修正しないで、Content-Security-Policyヘッダで対処できる
  - upgrade-insecure-requestsディレクティブ
    - HTTPと書かれていてもHTTPSで取得する
  - block-all-mixed-content
    - Mixed Contentをエラーにする
- Content-Security-PolicyはHTMLのメタタグにも書ける

#### クロスオリジンリソースシェアリング（CORS）

- ドメインをまたいでリソースを共有する作法
  - W3Cで定義
  - XMLHttpRequestやFetch API
- クライアントからサーバにアクセする直前までの権限確認
- simple cross-origin request
  - HTTPメソッドがシンプル
    - GET, POST, HEAD
  - ヘッダがシンプルヘッダのみで構成される
    - Accept, Accept-Language, Content-Language, Content-Type
  - Content-Typeの中身がいずれか
    - application/x-www-form-urlencoded, multipart/form-data, text-plain
- 複雑な通信ではプリフライトリクエストを実行
  - OPTIONSメソッド
- クロスオリジンの通信はデフォルトではCookieを送受信しない
  - 明示的に設定する
- プリフライトリクエストはキャッシュ機構がある

### 中間者攻撃（MITM攻撃）

- プロキシが通信を中継する際に通信内容を抜き取られて情報漏えいする
- HTTPS（TLS）を利用する
  - 通信経路が信頼できない状態でもセキュリティが維持される

#### HSTS（HTTP Strict Transport Security）

- サーバ側からHTTPSでの接続を要請するレスポンスヘッダー
  - RFCで定義されている
  - ブラウザ内部のDBで保持する
  - 有効期限やサブドメインを指定できる
- 初回がHTTPになってしまうリスク
  - 事前に申請すると最新のブラウザでされるようになる

### セッションハイジャッキング

- セッショントークンを盗み出して攻撃者がログインする
- セッショントークンはクッキーで持つ
  - XSSや中間者攻撃を踏み台にして盗まれる
- HTTPS化やCookieのhttpOnly,secure属性で守れる

#### クッキーインジェクション

- HTTPのサブドメインなどを利用して、HTTPSを迂回する
- ChromeとFirefoxは対策されている
- RFCのドラフトは作成されている

### クロスサイトリクエストフォージェリ（CSRF）

- 意図しない操作を強制する攻撃
- CSRF対策トークン
  - 隠しフィールドにランダムなトークンを埋め込み、サーバ側でチェックする
  - セッショントークンとは別物
  - Webアプリケーションフレームワークの提供している機能を使って対策する
- SameSite属性
  - ChromeのCookieの属性
  - 同一サイトのみでCookieを送信する
  - RFCのドラフトが作成されており、Firefoxもサポートに前向き

### クリックジャギング

- IFRAMEを使用して、意図しない操作をさせる
  - 有名サイトの上に透明な悪意のあるページを重ねる
  - 悪意のあるページの上に透明な有名サイトを重ねる
- X-Frame-Optionsヘッダで対策
  - DENY: フレーム内の使用を拒否
  - SEMEORIGIN: 同一URL以外の使用を拒否
  - ALLOW-FROM: 指定されたURLからの呼び出し以外を拒否

### リスト型アカウントハッキング

- 脆弱なWebサーバから流出したIDとパスワードで、流用しているサービスのセキュリティが無効になる
- 二段階認証
  - RFCで規格化されている
- Geo-Location
  - 生活拠点から離れた場所のアクセスを再認証させる
- 時間あたりのアクセス制限
  - IPあたりのログインアクセス数制限
  - ログイン失敗毎に待ち時間を延ばす
  - reCAPTHAなどでボットによる大量のログイン試行を防ぐ
- ユーザがサービスごとに別のパスワードを設定するのが一番強固
