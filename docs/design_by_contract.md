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

# 1. ログイン

- 対応する処理名: ``
- 事前条件
  - URL: `https://nanext.alcnanext.jp/anetn/student/stlogin/index/nit-ariake/`
    - アクセス先の URL
- 事後条件
  - URL: `https://nanext.alcnanext.jp/anetn/Student/StTop`
    - 操作によって、遷移先の URL になる
- 不変条件
  - クレデンシャル: 操作時に参照する
    - ID: `overall_design.md` に記載
    - Password: `overall_design.md` に記載
- 操作 (usecase)
  1. CSS セレクター `#AccountId` に、ID を入力
  2. CSS セレクター `#Password` に、Password を入力
  3. CSS セレクター `#BtnLogin` を押す
