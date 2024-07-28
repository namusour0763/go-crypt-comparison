# Go Crypt Comparison

Go を使って各種暗号化・ハッシュ関数の処理時間を計測してみるリポジトリ。

## 実行結果

### 実行環境

- CPU: Core-i7 6700K @ 4GHz
- Memory: 16 GB
- OS: Ubuntu 20.04 LTS (on WSL2)
- Go: version 1.21.3

### 実行時間

| ハッシュ関数 | 実行時間 (100回平均) |
| ------------ | -------------------- |
| MD5          | 0.139 ms             |
| SHA-256      | 0.221 ms             |
| Bcrypt       | 55.2 ms              |
| Scrypt       | 76.7 ms              |

## Go プロジェクトの始め方

Go 初心者のメモ。

```bash
mkdir my-go-project
cd my-go-project
go mod init github.com/yourusername/my-go-project
touch main.go
# コードを書く
go mod tidy # モジュールの追加・削除をする場合
go run main.go
go build # ビルドする場合
```
