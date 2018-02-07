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
