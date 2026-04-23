# ax
ディレクトリに名前を付けておいて、その名前でカレントディレクトリの変更を可能にするツールです。

# 使い方
```bash
ax [name or path]

ax @name # 名前衝突防止用に@を先頭に付けることが出来ます。
```

# インストール
1.以下のコマンドを入力。

```bash
go install github.com/midry3/ax/cmd/ax-lookup@latest
```

2.`~/.bashrc`に以下を追記。
```bash
eval "$(ax-lookup)"
```

3.`~/.axconf`を編集、json形式です。
```json
{
    "name": "path"
}
```
の書き方で追記していってください。
