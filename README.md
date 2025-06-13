# anarepo

リポジトリ（ディレクトリ）内のファイルを解析し、拡張子ごとの行数と割合を表示するGoで開発されたCLIツールです。

## 機能

- ディレクトリ内のファイルを再帰的に解析
- 拡張子ごとの行数統計を表示
- メジャーな拡張子の表示名変換（.ts → TypeScript等）
- 上位5種類の表示（デフォルト）または全ファイルタイプ表示
- node_modulesなどの不要なディレクトリを自動除外

## インストール・ビルド

```bash
# ビルド
go build

# 実行可能ファイルの作成
go build -o anarepo
```

## 使用方法

```bash
# 基本的な使用方法
anarepo <ディレクトリパス>

# 全てのファイルタイプを表示
anarepo --all <ディレクトリパス>
```

## 実行例

### 基本実行（上位5種類のみ表示）

```bash
$ anarepo test
| FileType     | Lines | Percent |
| ------------ | ----- | ------- |
| CSS          |    29 |   25.0% |
| JavaScript   |    26 |   22.4% |
| TypeScript   |    26 |   22.4% |
| Markdown     |    20 |   17.2% |
| JSON         |    15 |   12.9% |
| ============ | ===== | ======= |
| Total        |   116 |  100.0% |
```

### 全ファイルタイプ表示

```bash
$ anarepo --all test
| FileType     | Lines | Percent |
| ------------ | ----- | ------- |
| CSS          |    29 |   25.0% |
| TypeScript   |    26 |   22.4% |
| JavaScript   |    26 |   22.4% |
| Markdown     |    20 |   17.2% |
| JSON         |    15 |   12.9% |
| ============ | ===== | ======= |
| Total        |   116 |  100.0% |
```

### 現在のディレクトリ解析

```bash
$ anarepo .
| FileType     | Lines | Percent |
| ------------ | ----- | ------- |
| Go           |   662 |   76.1% |
| Markdown     |    92 |   10.6% |
| JSON         |    30 |    3.4% |
| CSS          |    29 |    3.3% |
| JavaScript   |    28 |    3.2% |
| ============ | ===== | ======= |
| Total        |   870 |  100.0% |
```

## サポートしている拡張子変換

主要なプログラミング言語の拡張子を認識し、わかりやすい表示名に変換します：

- `.ts`, `.tsx` → TypeScript
- `.js`, `.jsx` → JavaScript  
- `.py` → Python
- `.go` → Go
- `.java` → Java
- `.c` → C
- `.cpp`, `.cc`, `.cxx` → C++
- `.cs` → C#
- `.php` → PHP
- `.rb` → Ruby
- `.rs` → Rust
- `.html`, `.htm` → HTML
- `.css` → CSS
- `.md` → Markdown
- その他多数

## 除外されるディレクトリ

以下のディレクトリは自動的に解析対象から除外されます：

- `node_modules`
- `.git`, `.svn`, `.hg`, `.bzr`
- `vendor`, `build`, `dist`, `target`
- `bin`, `obj`, `out`
- `.next`, `.nuxt`
- `.vscode`, `.idea`
- `__pycache__`, `.pytest_cache`
- `tmp`, `temp`, `logs`
- その他

## 開発

### テスト実行

```bash
go test ./...
```

### コードフォーマット

```bash
go fmt ./...
```

### 静的解析

```bash
go vet ./...
```

### モジュール管理

```bash
go mod tidy
```
