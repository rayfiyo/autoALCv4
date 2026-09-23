# 契約による設計 を意識したドキュメント

- 対応する処理名: 関数・メソッド・ハンドラーなどの名前
- 事前条件: その処理の開始時に保証されるべき引数や状態の条件
- 事後条件: その処理の正常終了時に保証されるべき返り値や状態の条件
- 不変条件: その処理の開始時と正常終了時に共通して保証されるべき値や状態の条件
- 操作 (usecase)

- 対応する処理名: ``
- 事前条件
  - URL: ``
- 事後条件
- 不変条件
- 操作 (usecase)

# 1. フラグの処理

- 対応する処理名: ``
- 事前条件
  - 各フラグに対応する変数: 型に合った初期値が代入
- 事後条件
  - 各フラグに対応する変数: フラグで指定された値 XOR そのフラグのデフォルト値
- 不変条件
- 操作 (usecase)
  1. 実行時にフラグ（実行時オプション）が指定されていれば、それを設定する
     - 指定されていなければ、そのフラグのデフォルト値を設定する
     - 各フラグについては `overall_design.md` に記載している

# 2. ログイン

- 対応する処理名: ``
- 事前条件
  - URL: `https://nanext.alcnanext.jp/anetn/student/stlogin/index/nit-ariake/`
    - アクセス先の URL
- 事後条件
  - URL: `https://nanext.alcnanext.jp/anetn/Student/StTop`
    - 操作によって遷移し、URL が変化する
- 不変条件
  - クレデンシャル: 操作時に参照する
    - ID: `overall_design.md` に記載
    - Password: `overall_design.md` に記載
- 操作 (usecase)
  1. CSS セレクター `#AccountId` に、ID を入力
  2. CSS セレクター `#Password` に、Password を入力
  3. CSS セレクター `#BtnLogin` を押す

# 3. セッション ID の取得

- 対応する処理名: ``
- 事前条件
  - 取得するセッション ID を格納する変数: 型に合った初期値が代入
- 事後条件
  - 取得するセッション ID を格納する変数: 24 文字の英数小文字
    - 例: `ajzzl37oxovqicmnhiv1rkrr`
- 不変条件
  - URL: `https://nanext.alcnanext.jp/anetn/Student/StTop`
    - `1. ログイン` の事後条件を引き継ぐため、新規セッションでアクセスする必要はない
    - 操作は HTML 要素の属性を取得するだけなので、URL 遷移は発生しない（不変）
- 操作 (usecase)
  1. CSS セレクター `#HidSessionId` が割り当てられている要素の `value` 属性値を取得
     - JS の DOM 操作なら `document.querySelector("#HidSessionId").value;` に該当する

# 4. 処理するコース

- 対応する処理名: ``
- 事前条件
  - URL: ``
- 事後条件
- 不変条件
- 操作 (usecase)

# 5.

- 対応する処理名: ``
- 事前条件
  - URL: ``
- 事後条件
- 不変条件
- 操作 (usecase)

# 6.

- 対応する処理名: ``
- 事前条件
  - URL: ``
- 事後条件
- 不変条件
- 操作 (usecase)
