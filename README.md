# autoALCv4

- ALC を自動でやる 4 版
- This is the 4th edition of ALC with automation

# Documentation

```
docs/
├── design_by_contract.md
│   - 契約による設計 を意識したドキュメント
│   - 具体的には次が記載されている
│     - 対応する処理名: 関数・メソッド・ハンドラーなどの名前
│     - 事前条件: その処理の開始時に保証されるべき引数や状態の条件
│     - 事後条件: その処理の正常終了時に保証されるべき返り値や状態の条件
│     - 不変条件: その処理の開始時と正常終了時に共通して保証されるべき値や状態の条件
│     - 操作 (usecase)
└── overall_design.md
    - アプリ全体の設計書
```
