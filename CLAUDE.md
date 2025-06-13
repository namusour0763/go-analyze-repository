# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 開発のルール

- 日本語で受け答えすること
- 編集後はテストをパスするか確認すること

## プロジェクト概要

このプロジェクトはGoで開発されたCLIツール「anarepo」です。リポジトリ（フォルダ）内のファイルを解析し、拡張子ごとの行数と割合を表示します。

## 開発コマンド

```bash
# ビルド
go build

# 実行
go run main.go [folder]

# テスト実行
go test ./...

# コードフォーマット
go fmt ./...

# 静的解析
go vet ./...

# モジュール管理
go mod tidy
```

## 作成するツール

### 目的

リポジトリ（フォルダ）内のファイルを読み込み、拡張子ごとの行数と割合を表示するCLIツール

### 入出力

以下の形式を想定

```text
$ anarepo folder

| FileType   | Lines | Percent |
| ---------- | ----: | ------: |
| TypeScript |   720 |  72.0 % |
| JavaScript |   182 |  18.2 % |
| Markdown   |    80 |   8.0 % |
| .settings  |    18 |    1.8% |
| ========== | ===== |  ====== |
| Total      |  1000 |   100 % |
```

### 技術的仕様

- コマンドの引数にリポジトリ（フォルダ）を指定する
- 表形式（整形する）でアウトプットを出力する
- FileTypeは .ts や .tsx であればTypeScriptと表示するような変換をする
- FileTypeの変換はメジャーな拡張子について変換テーブルを持っておく
- 上記変換テーブルに該当しない拡張子は .settings のようにそのまま表示してよい
- デフォルトでは上位5種類のファイルについて表示し、オプション --all をつけた場合のみ全ファイルタイプについて表示する
- node_modules等の一般的には解析が不要なフォルダはスキップする

## テスト

- 単体テストおよび結合テストを実施する
- 結合テストについては、testフォルダに実際にファイルを配置し、想定した出力になるか確認する
