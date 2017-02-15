[![wercker status](https://app.wercker.com/status/52243183472f21dc72756d12a649ee19/m/master "wercker status")](https://app.wercker.com/project/bykey/52243183472f21dc72756d12a649ee19)
[![go report status](https://goreportcard.com/badge/github.com/soracom/soracom-cli)](https://goreportcard.com/report/github.com/soracom/soracom-cli)

# soracom-cli

SORACOM API を呼び出すためのコマンドラインツール soracom を提供します。

# 特徴

soracom コマンドは以下のような特徴を備えています。

- soracom コマンドのバイナリファイルは API 定義ファイルから自動生成されますので、新しい API がリリースされた場合も迅速に対応が可能です。

- Go でクロスコンパイルされたバイナリファイルをターゲットの環境にコピーするだけで実行できます。環境を構築したり依存関係を解決したりする必要がありません。

- 指定された引数を元にリクエストを組み立て、SORACOM API を呼び出します。API からのレスポンス (JSON) をそのまま標準出力へ出力します。
  - これにより、soracom コマンドの出力を加工して他のコマンドへ渡したりすることが容易にできるようになります。

- bash completion（引数補完）に対応しています。以下のような行を .bashrc 等に記述してください。
  ```
  eval "$(soracom completion)"
  ```

  macOS をお使いの場合、以下のいずれかの条件を満たす必要があるかもしれません：
  1. `bash` のバージョン 4.0 以降を使用する
  2. `brew install bash-completion` でインストールした bash-completion を使う（Xcode に付属の bash-completion では動作しない場合があります。）

  以下のようなエラーが起きた場合は上記いずれかをお試し下さい。
  ```
  -bash: __ltrim_colon_completions: command not found
  ```

# インストール方法

## macOS をお使いで、homebrew によりインストールする場合

```
$ brew tap soracom/soracom-cli
$ brew install soracom-cli
```

## それ以外の場合
[Releases のページ](https://github.com/soracom/soracom-cli/releases) からターゲットの環境に合ったパッケージファイルをダウンロードして展開し、実行形式ファイルを PATH の通ったディレクトリに配置します。


# 使用方法

## 基本的な使い方

まずはじめに、プロファイルの作成をします。

```
soracom configure
```

このコマンドを実行すると、どのカバレッジタイプを使用するか質問されます。

```
カバレッジタイプを選択してください。

1. Global
2. Japan

選択してください (1-2) >
```

主に使用する方のカバレッジタイプを選択してください。（ここで選択したカバレッジタイプ以外も使用可能です）
よくわからない場合、日本在住の方が日本で SIM を使う場合には 2. Japan を選択し、それ以外の場合は 1. Global を選択してください。

次に認証方法について質問されます。

```
認証方法を選択してください。

1. AuthKeyId と AuthKey を入力する（推奨）
2. オペレーターのメールアドレスとパスワードを入力する
3. SAM ユーザーの認証情報を入力する（オペレーターID、ユーザー名、パスワード）

選択してください (1-3) >
```

SAM ユーザーもしくはルートアカウントに対し、AuthKey（認証キー）を発行している場合は 1 を選択してください。
（SAM ユーザーに対し認証キーを発行する方法については [SORACOM Access Managementを使用して操作権限を管理する](https://dev.soracom.io/jp/start/sam/) を参照してください）

以後、soracom コマンド実行時は、ここで入力した認証情報を使って API 呼び出しが行われます。



## 高度な使い方

### 複数のプロファイルを使い分ける

SORACOM アカウントを複数所有しているとか、複数の SAM ユーザーを使い分けたい場合は、configure に --profile オプションを指定し、プロファイル名を設定します。

```
soracom configure --profile user1
  :
  （user1 のための認証情報を入力）

soracom configure --profile user2
  :
  （user2 のための認証情報を入力）
```

このようにすると user1 および user2 という名前のプロファイルが作成されます。
プロファイルを利用する場合は通常のコマンドの後ろに --profile オプションを指定します。

```
soracom subscribers list --profile user1
  :
  （user1 に SIM の一覧を表示する権限があれば表示される）

soracom groups list --profile user2
  :
  （user2 にグループの一覧を表示する権限があれば表示される）
```


### Proxy 経由で API を呼び出したい場合

HTTP_PROXY 環境変数に http://your-proxy-nme:port を設定した状態で soracom コマンドを実行してください。

例）Linux/Mac の場合：
Proxy サーバーのアドレスを 10.0.1.2、ポート番号を 8080 だとすると
```
export HTTP_PROXY=http://10.0.1.2:8080
soracom subscribers list
```

もしくは

```
HTTP_PROXY=http://10.0.1.2:8080 soracom subscribers list
```


# ビルド/テスト方法

ソースからビルドしたい開発者の方や、バグ修正/機能追加等の Pull Request をしたい場合は以下のいずれかの方法でビルドおよびテストを行ってください。

## ローカル環境でビルドする方法 (Linux / Mac OS X)

Go がインストールされている状態で、以下のようにビルドスクリプトを実行します。

```
./scripts/build.sh 1.2.3
```

ここで 1.2.3 はバージョン番号です。適当な番号を指定してください。

ビルドが成功したら、次にテストを実行します。

```
./test/test.sh
```


## wercker を使ってビルドする方法

wercker の CLI をインストールし、以下のようにビルドを実行します。テストまで自動的に実行されます。

```
wercker build
```

TODO: 現状、ビルド結果はコンテナの中に出力されるので、マウントしたボリュームに出力できるように修正予定です
