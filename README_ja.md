[![go report status](https://goreportcard.com/badge/github.com/soracom/soracom-cli)](https://goreportcard.com/report/github.com/soracom/soracom-cli)
![build-artifacts](https://github.com/soracom/soracom-cli/actions/workflows/build-artifacts.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/soracom/soracom-cli.svg)](https://pkg.go.dev/github.com/soracom/soracom-cli)

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

  - 以下のようなエラーが起きたときは:

    ```
    -bash: __ltrim_colon_completions: command not found
    ```

    このエラーは macOS をお使いの場合に表示されることがあります。以下のいずれかの条件を満たす必要があるかもしれません：
    - `bash` のバージョン 4.0 以降を使用する
    - `brew install bash-completion` でインストールした bash-completion を使う（Xcode に付属の bash-completion では動作しない場合があります。）

      そしてこの場合、`.bash_profile` または `.profile` ファイルに以下を追加します:
      ```
      if [ -f "$(brew --prefix)/etc/bash_completion" ]; then
        . "$(brew --prefix)/etc/bash_completion"
      fi
      ```

- zsh completion（引数補完）に対応しています。以下のようなコマンドを実行して生成されるスクリプトを `_soracom` という名前で `$fpath` のどこかに配置してください。
  ```
  soracom completion zsh
  ```

# インストール方法

## macOS もしくは Linux をお使いで、homebrew によりインストールする場合

```shell
brew tap soracom/soracom-cli
brew install soracom-cli
brew install bash-completion
```

## それ以外の場合

以下に紹介するいずれかのコマンドを実行することで、最新版の `soracom` コマンドがダウンロードされてインストールされます。


もし `/usr/local/bin` にファイルを書き込む権限のあるユーザー（root など）でインストールする場合は以下のコマンドを実行してください。

```shell
curl -fsSL https://raw.githubusercontent.com/soracom/soracom-cli/master/install.sh | bash
```

`/usr/local/bin` に書き込む権限がない場合は以下のいずれかのコマンドを実行してください。

sudo コマンドを実行可能（ユーザーが sudoers に入っている）で、`soracom` コマンドを `/usr/local/bin` にインストールしたい場合：

```shell
curl -fsSL https://raw.githubusercontent.com/soracom/soracom-cli/master/install.sh | sudo bash
```

もしくは sudo コマンドを利用できないか、`soracom` コマンドを `$HOME/bin` など `/usr/local/bin` 以外の場所にインストールしたい場合：

```shell
mkdir -p "$HOME/bin"
curl -fsSL https://raw.githubusercontent.com/soracom/soracom-cli/master/install.sh | BINDIR="$HOME/bin" bash
```

`"$HOME/bin"` の部分はお好きなディレクトリに変更してください。

上記いずれかの方法でインストールした `soracom` コマンドをバージョンアップしたい場合は、同じコマンドを再度実行してください。

インストールした `soracom` コマンドをアンインストールしたい場合は、インストールした `soracom` コマンドの実行ファイルを手動で削除してください。（プロファイル情報も含めて完全に削除したい場合は `$HOME/.soracom/` ディレクトリも削除してください）

上記のコマンドがいずれもうまく行かない場合、もしくは古いバージョンの `soracom` コマンドをインストールしたい場合は、[Releases のページ](https://github.com/soracom/soracom-cli/releases) からターゲットの環境に合ったパッケージファイルをダウンロードして展開し、実行形式ファイルを PATH の通ったディレクトリに配置してください。


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
4. スイッチユーザー

選択してください (1-4) >
```

SAM ユーザーもしくはルートアカウントに対し、AuthKey（認証キー）を発行している場合は 1 を選択してください。
SAM ユーザーに対し認証キーを発行する方法についてはソラコムユーザーサイトの [認証キーを生成する](https://users.soracom.io/ja-jp/docs/sam/create-sam-user/#%e8%aa%8d%e8%a8%bc%e3%82%ad%e3%83%bc%e3%82%92%e7%94%9f%e6%88%90%e3%81%99%e3%82%8b) を参照してください

4 のスイッチユーザーを選択すると、スイッチ元のユーザーのプロファイルとスイッチ先のユーザーの Operator ID および SAM ユーザー名を指定することができます。
実行前にあらかじめスイッチ元のユーザーのプロファイルを作成しておいてください。
いったんスイッチユーザーのプロファイルを作成すると、各種サブコマンドを実行する際に soracom-cli が自動的にスイッチ元のプロファイルで認証し、スイッチ先の SAM ユーザーへスイッチしてから API の呼び出しを行うようになります。

スイッチユーザーの詳細については、[スイッチユーザー | ドキュメント | ソラコムユーザーサイト - SORACOM Users](https://users.soracom.io/ja-jp/docs/switch-user/) を参照してください。

以後、soracom コマンド実行時は、ここで入力した認証情報を使って API 呼び出しが行われます。



## 高度な使い方

### 複数のプロファイルを使い分ける

SORACOM アカウントを複数所有しているとか、複数の SAM ユーザーを使い分けたい場合は、configure に `--profile` オプションを指定し、プロファイル名を設定します。

```
soracom configure --profile user1
  :
  （user1 のための認証情報を入力）

soracom configure --profile user2
  :
  （user2 のための認証情報を入力）
```

このようにすると user1 および user2 という名前のプロファイルが作成されます。
プロファイルを利用する場合は通常のコマンドの後ろに `--profile` オプションを指定します。

```
soracom subscribers list --profile user1
  :
  （user1 に SIM の一覧を表示する権限があれば表示される）

soracom groups list --profile user2
  :
  （user2 にグループの一覧を表示する権限があれば表示される）
```

### API Sandbox 環境用のプロファイルを作成する

[SORACOM API Sandbox](https://dev.soracom.io/jp/docs/api_sandbox/) のセットアップなどにも soracom-cli を使うことができます。

`configure-sandbox` サブコマンドを用いてプロファイルを作成します。

```
soracom configure-sandbox
```
表示される質問に従って入力していくと、デフォルトでは `sandbox` という名前のプロファイルが作成されます。
その `sandbox` プロファイルを用いることで、以下のように API Sandbox に対してコマンドを発行することができます。

```
soracom subscribers list --profile sandbox
```

また、Sandbox 専用コマンドを利用することもできるようになります。

```
soracom sandbox subscribers create --profile sandbox
```

デフォルトとは異なるプロファイル名を使うこともできます。

```
soracom configure-sandbox --profile test
soracom sandbox subscribers create --profile test
```

シェルスクリプトなどから使いやすいように、プロファイル作成時に必要なパラメータをすべて引数で指定できるようになっています。

```
soracom configure-sandbox --coverage-type jp --auth-key-id="$AUTHKEY_ID" --auth-key="$AUTHKEY" --email="$EMAIL" --password="$PASSWORD"
```

### コマンドライン引数で指定する認証方法の優先順位

soracom-cli は SORACOM API を呼び出すために、通常は内部的に認証を行って API キーとトークンを取得し、API リクエストとともにそれらを送信しています。

認証を行ったり API キーとトークンを指定したりするためのオプションは複数あり、その使い方には以下のような方法があります。

1. 認証を事前に行って取得しておいた API キーとトークンを`--api-key` と `--api-token` オプションで直接指定し、それらを用いて API を呼び出す

2. 認証キーIDと認証キーを `--auth-key-id` と `--auth-key`オプションで指定して認証を行って API キーとトークンを取得し、それらを用いて API を呼び出す

3. `--profile-command` オプションで指定された外部コマンドを実行することでプロファイル（＝認証を行うための情報）を 生成し、そのプロファイルを用いて認証を行って API キーとトークンを取得し、それらを用いて API を呼び出す

4. 事前に構成されたプロファイルを`--profile` オプションで指定し、そのプロファイルを用いて認証を行って API キーとトークンを取得し、それらを用いて API を呼び出す

これらは 1 から 4 の順に優先されます。すなわち、1 が最も優先され、4 の優先度が最も低くなります。
たとえば、もし soracom-cli のユーザーが `--profile-command` オプションと `--profile` オプションを同時に指定してしまった場合、`--profile-command` オプションの内容が優先されます。

また、`--api-key` と `--api-token` や `--auth-key-id` と `--auth-key` のように、2 つ同時に指定する必要があるオプションを片方だけ指定した場合はエラーとなります。


### プロファイルで指定する認証方法の優先順位

プロファイルの中には以下のいずれかの認証方法を指定できます。

1. プロファイル情報を生成するための外部コマンドを指定する`profileCommand` フィールド

2. スイッチユーザー機能でスイッチ元のプロファイルを指定する `sourceProfile` フィールドおよびスイッチ先のオペレーター ID とユーザー名を指定する `operatorId` および `username` フィールド

3. 認証キーIDと認証キーを指定する `authKeyId` および `authKey` フィールド

4. ルートユーザーのメールアドレスとパスワードを指定する `email` および `password` フィールド

5. SAM ユーザーのオペレーター ID、ユーザー名およびパスワードを指定する `operatorId`、`username` および `password` フィールド

これらは 1 か 5 の順に優先されます。すなわち、1 が最も優先され、5 の優先度が最も低くなります。
たとえば、もし `profileCommand` フィールドと `authKeyId` および `authKey` フィールドを同時にプロファイル内に指定してしまった場合、`profileCommand` の内容が優先されます。

`sourceProfile` で参照されるプロファイルの中には `sourceProfile` を指定することはできません。


### Proxy 経由で API を呼び出したい場合

HTTP_PROXY 環境変数に `http://your-proxy-name:port` を設定した状態で soracom コマンドを実行してください。

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

### soracom-cli の AWS Lambda Layers を利用する

soracom-cli を AWS Lambda 上で利用しようと考えたことはありますか？
Zip パッケージやコンテナイメージに soracom-cli のバイナリを含めてデプロイすることで、あなたの Lambda 関数の中から soracom-cli を呼び出すことができます。

しかしながら、soracom-cli のバイナリは比較的大きいため、Zip パッケージやコンテナイメージの容量を圧迫してしまうかもしれません。

そこで、私達が提供するのが soracom-cli の Layer です。

以下のような ARN を指定することで、あなたの Lambda 関数の中から `soracom` コマンドを実行できるようになります。

- x86_64 アーキテクチャー：

  ```
  arn:aws:lambda:ap-northeast-1:717257875195:layer:soracom-cli-${ver}:1
  ```

- arm64 アーキテクチャー：

  ```
  arn:aws:lambda:ap-northeast-1:717257875195:layer:soracom-cli-${ver}-arm64:1
  ```

`${ver}` の部分には、対象となる soracom-cli のバージョン番号から `.` を取り除いたものが入ります。

たとえばバージョン `1.2.3` なら、`${ver}` は `123` となります。

バイナリは /bin/soracom にインストールされます。PATH が通っているので Lambda 関数の中では単に `soracom` コマンドとして実行できます。

Node.js 18.x ランタイムでは以下のようにして呼び出すことができます。（環境変数で AUTH_KEY_ID と AUTH_KEY を渡してください）

```
const execSync = require('child_process').execSync;
const jpBill = execSync(`soracom --auth-key-id ${process.env.AUTH_KEY_ID} --auth-key ${process.env.AUTH_KEY} bills get-latest --coverage-type jp`).toString();
```


### トラブルシューティング

もし、以下のようなエラーメッセージが表示されてしまったら、

```
Error: 認証情報ファイル 'path/to/default.json' へのアクセス権が十分に絞り込まれていません。
認証情報ファイルへは、soracom コマンドを実行しているユーザーのみがアクセス可能なように設定する必要があります。
```

以下のコマンドを実行して修復を試みてください。

```
soracom unconfigure
soracom configure
```

いったん `unconfigure` してから `configure` することにより、認証情報ファイルを適切なパーミッションで再作成します。



# ビルド/テスト方法

ソースからビルドしたい開発者の方や、バグ修正/機能追加等の Pull Request をしたい場合は以下のいずれかの方法でビルドおよびテストを行ってください。

## API 定義ファイル/ヘルプメッセージの更新

以下の API 定義ファイルを更新します。

- generators/assets/soracom-api.en.yaml
- generators/assets/soracom-api.ja.yaml
- generators/assets/sandbox/soracom-sandbox-api.en.yaml
- generators/assets/sandbox/soracom-sandbox-api.ja.yaml

`configure --help` で表示されるメッセージを更新するには以下のファイルを更新します。

- generators/assets/cli/en.yaml
- generators/assets/cli/ja.yaml

## ローカル環境でビルドする方法 (Linux / Mac OS X)

Go がインストールされている状態で、以下のようにビルドスクリプトを実行します。

```
./scripts/copy-apidef-files.sh # Before running this script, please copy API definitions before hand.
aws ecr-public get-login-password --profile {your AWS profile} --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws
VERSION=1.2.3
./scripts/build.sh $VERSION
```

ここで 1.2.3 はバージョン番号です。適当な番号を指定してください。

ビルドが成功したら、次にテストを実行します。

```
# API sandbox を利用するために、実在する SORACOM オペレーター（アカウント）の
# AuthKey ID と AuthKey を事前に環境変数に設定してください。
export SORACOM_AUTHKEY_ID_FOR_TEST=...
export SORACOM_AUTHKEY_FOR_TEST=...
./test/test.sh $VERSION
```

### ビルド時のトラブルシューティング

ビルド時に `go: could not create module cache: mkdir /go/pkg: permission denied` のようなエラーが表示されたときは、Docker container 内の /go/pkg の権限を確認してください。build.sh では、ホストの `${GOPATH:-$HOME/go}` を /go/pkg にマウントしているため、ホストの `$GOPATH` または `$HOME/go` の権限を確認します。ほとんどの場合は、以下のコマンドで解決できるはずです。

```
sudo chown -R $USER:$USER ${GOPATH:-$HOME/go}
```
