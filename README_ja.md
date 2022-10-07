[![go report status](https://goreportcard.com/badge/github.com/soracom/soracom-cli)](https://goreportcard.com/report/github.com/soracom/soracom-cli)
![build-artifacts](https://github.com/soracom/soracom-cli/actions/workflows/build-artifacts.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/soracom/soracom-cli.svg)](https://pkg.go.dev/github.com/soracom/soracom-cli)

# soracom-cli

SORACOM API を呼び出すためのコマンドラインツール soracom を提供します。

# 特徴

soracom コマンドは以下のような特徴を備えています。

- soracom コマンドのバイナリファイルは、./generators/assets/ 以下に保存された API 定義ファイルから自動生成されます。新しい API がリリースされた場合も、API 定義ファイルを更新すれば迅速に対応できます。

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
    そしてこの場合、`.bash_profile` または `.profile` ファイルに以下を追加します:
    ```
    if [ -f $(brew --prefix)/etc/bash_completion ]; then
      . $(brew --prefix)/etc/bash_completion
    fi
    ```

  以下のようなエラーが起きた場合は上記いずれかをお試し下さい。
  ```
  -bash: __ltrim_colon_completions: command not found
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

## Linux をお使いで、snap によりインストールする場合

```shell
sudo snap install soracom
```

snap を使って `soracom` コマンドをインストールし、`$HOME/.soracom` ディレクトリに保存されたプロファイル情報を利用したい場合は `dot-soracom` インターフェースを `soracom` の snap パッケージに connect してください。

```shell
snap connect soracom:dot-soracom
```

さらに、snap でインストールした場合のデフォルトのプロファイルディレクトリ `$SNAP_USER_DATA/.soracom`（すなわち `$HOME/snap/soracom/<revision>/.soracom`）の代わりに `$HOME/.soracom` を利用するために以下のような行を `.bashrc` などに追加してください。

```bash
export SORACOM_PROFILE=$HOME/.soracom
```

## それ以外の場合

以下に紹介するいずれかのコマンドを実行することで、最新版の `soracom` コマンドがダウンロードされてインストールされます。


もし `/usr/local/bin` にファイルを書き込む権限のあるユーザー（root など）でインストールする場合はいかのコマンドを実行してください。

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

## 認証情報を準備する

あらかじめ、[ステップ 2: 認証情報を準備する](https://users.soracom.io/ja-jp/tools/cli/getting-started/#ステップ-2-認証情報を準備する) の手順に従って、SORACOM CLI を実行するための認証情報を準備します。

## 認証情報をプロファイルに保存する

soracom コマンドを実行すると、ここでプロファイルに保存した認証情報を使って API が呼び出されます。

なお、環境変数 LANG に `ja` を設定すると、soracom コマンドのメッセージが日本語に切り替わります。

Linux の場合:

```
export LANG=ja
```

プロファイルを作成するために、以下のコマンドを実行します。

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

今後、SORACOM CLI を使用するときに、主に使用するカバレッジタイプを選択してください。

たとえば、「1. Global」を選択した場合は、`soracom sims list` と入力すると、グローバルカバレッジの IoT SIM (plan01s など) の一覧を取得できます。

「2. Japan」を選択した場合は、`soracom sims list` と入力すると、日本カバレッジの IoT SIM (plan-D など) の一覧を取得できます。

ここで選択しなかったカバレッジタイプを使用する場合は、`soracom sims list --coverage-type g` や `soracom sims list --coverage-type jp` のように、毎回カバレッジタイプを指定する必要があります。

カバレッジタイプについて詳しくは、[カバレッジタイプ (Coverage Type) の概要](https://users.soracom.io/ja-jp/guides/basic-knowledge/coverage-type/) を参照してください。

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

最後に、soracom コマンドが実行できることを確認します。

```
soracom sims list
```

IoT SIM の一覧が出力されたら成功です。


なお、SAM ユーザーの認証情報を使用している場合は、以下のエラーメッセージが表示されることがあります。この場合は、[SAM ユーザーの権限を設定する](https://users.soracom.io/ja-jp/docs/sam/set-permissions/) の手順に従って SAM ユーザーの権限を設定してください。詳しくは、エラーメッセージ全文を添えて、管理者にお問い合わせください。

```
Error: {"code":"SEM0049","message":"SAMユーザー 'samuser' にこの操作を行うための権限 'Sim:listSims' が許可されていません。SAMユーザーの権限設定については管理者にお問い合わせください。なお、SAMユーザーは管理者が適切な権限を設定したあとで再度ログインしなおす必要があります。"}
```

管理者が権限を設定したら、もう一度同じコマンドを実行してください。IoT SIM の一覧が出力されたら成功です。

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
soracom sims list --profile user1
  :
  （user1 に SIM の一覧を表示する権限があれば表示される）

soracom groups list --profile user2
  :
  （user2 にグループの一覧を表示する権限があれば表示される）
```

### API Sandbox 環境用のプロファイルを作成する

[SORACOM API Sandbox](https://users.soracom.io/ja-jp/tools/api-sandbox/) のセットアップなどにも soracom コマンドを使用できます。

`configure-sandbox` サブコマンドを用いてプロファイルを作成します。

```
soracom configure-sandbox
```

Sandbox 用のプロファイルを作成するために必要な情報は、以下のとおりです。

- 本番環境の authKeyId
- 本番環境の authKey
- Sandbox 環境のルートユーザーのメールアドレス
- Sandbox 環境のルートユーザーのパスワード

表示される質問に従って入力していくと、デフォルトでは `sandbox` という名前のプロファイルが作成されます。
その `sandbox` プロファイルを用いることで、以下のように `--profile sandbox` を追加して、API Sandbox に対してコマンドを発行することができます。

```
soracom sims list --profile sandbox
```

また、Sandbox 専用コマンドを利用することもできるようになります。

```
soracom sandbox subscribers create --profile sandbox
```

デフォルトの `sandbox` とは別のプロファイル名も使用できます。

```
soracom configure-sandbox --profile test
soracom sandbox subscribers create --profile test
```

シェルスクリプトなどから使いやすいように、プロファイル作成時に必要なパラメータをすべて引数で指定できるようになっています。

```
soracom configure-sandbox --coverage-type jp --auth-key-id="$AUTHKEY_ID" --auth-key="$AUTHKEY" --email="$EMAIL" --password="$PASSWORD"
```

### Proxy 経由で API を呼び出したい場合

HTTP_PROXY 環境変数に `http://your-proxy-name:port` を設定した状態で soracom コマンドを実行してください。

例）Linux/Mac の場合：
Proxy サーバーのアドレスを 10.0.1.2、ポート番号を 8080 だとすると
```
export HTTP_PROXY=http://10.0.1.2:8080
soracom sims list
```

もしくは

```
HTTP_PROXY=http://10.0.1.2:8080 soracom sims list
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

## ローカル環境でビルドする方法 (Linux / Mac OS X)

### 前提条件

- Go 言語

    詳しくは、[Download and install](https://go.dev/doc/install) を参照してください。

- Docker

    Docker をインストールし、非 root ユーザーで Docker を実行できるように設定してください。詳しくは、[Install Docker Engine](https://docs.docker.com/engine/install/) を参照してください。

- jq

    テスト用に jq コマンドをインストールします。詳しくは、[Download jq](https://stedolan.github.io/jq/download/) を参照してください。

### ローカル環境でビルドする

ビルドスクリプトを実行します。

```
VERSION=1.2.3
./scripts/build.sh $VERSION
```

ここで指定した 1.2.3 はバージョン番号です。任意のバージョン番号を指定してください。

ビルドが成功すると、`./soracom/dist/$VERSION` に各 OS 用のファイルが作成されます。作成されたファイルは、以下のコマンドで確認できます。

```
ls ./soracom/dist/$VERSION
```

次に、テストスクリプトを実行します。

テストは API Sandbox を利用します。本番環境の authKeyId と authKey を事前に環境変数に設定してください。

```
export SORACOM_AUTHKEY_ID_FOR_TEST=...(本番環境の authKeyId を指定する)
export SORACOM_AUTHKEY_FOR_TEST=...(本番環境の authKey を指定する)
```

テストを実行します。

```
./test/test.sh $VERSION
```

最後に `TEST OK` と表示されたら成功です。
