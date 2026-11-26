# リリース手順

soracom-cli の新しいバージョンをリリースするメンテナー向けの手順です。

リリースには、soracom/soracom-cli と soracom/homebrew-soracom-cli への書き込み権限、および Lambda Layer を公開している AWS アカウントでの Layer 公開権限が必要です。

## 配布チャネル

| チャネル | リリース時の作業 |
|---|---|
| GitHub Releases | タグを push し、作成された pre-release を昇格させる |
| `install.sh` | 不要（最新の Release を自動で参照する） |
| Homebrew（soracom/homebrew-soracom-cli） | formula を更新する PR を作成する |
| AWS Lambda Layers | 各リージョンに Layer を公開する |
| Snap | 不要（非推奨。[Snap について](#snap-について) を参照） |

## 必要なツール

- Docker
- [GitHub CLI](https://cli.github.com/)（`gh`）
- AWS CLI、`jq`、`curl`、`zip`

## 1. バージョンを決める

前回のリリースからの変更を確認し、[セマンティックバージョニング](https://semver.org/lang/ja/) に従ってバージョンを決めます。

```bash
git log --oneline --no-merges v<前回のバージョン>..origin/master
```

## 2. API 定義を最新にする（必要な場合）

最新の API 定義を取り込む場合は、`scripts/copy-apidef-files.sh` で API 定義ファイルを `generators/assets/` にコピーし、ビルド（手順 3）で生成されたコードとあわせて PR を作成します。
## 3. ビルドしてテストする

master の最新状態で、README の [ローカル環境でビルドする方法](../README_ja.md#ローカル環境でビルドする方法-linux--mac-os-x) に従ってビルドとテストを実行します。ここでビルドしたファイルはテスト用です。リリースするファイルは GitHub Actions がビルドします（手順 4）。

ビルドの途中でソースコードが再生成されます。ビルド後に `git status` で差分が出ていないことを確認してください。差分が残っていると、リリース時のビルドが失敗します。差分が出た場合は、PR にしてマージしてから、ビルドをやり直してください。

## 4. タグを push する

テストしたコミットに `v<バージョン>` 形式のタグを付けて push します。

```bash
VERSION=1.2.3
git tag v$VERSION
git push origin v$VERSION
```

タグが push されると、GitHub Actions の Release ワークフロー（`.github/workflows/release.yml`）が次の処理を行います。

1. リリース用のファイルをビルドする
2. ビルドしたファイルに [build provenance attestation](https://docs.github.com/ja/actions/security-for-github-actions/using-artifact-attestations) を付与する
3. リリースノートを自動生成し、pre-release として GitHub Release を作成する

ワークフローの run を確認し、完了を待ちます。

```bash
gh run list --workflow release.yml --branch v$VERSION
gh run watch <run ID>
```

> [!WARNING]
> GitHub Release を手動で作成したり、ローカルでビルドしたファイルをアップロードしたりしないでください。同じタグの Release が既にあるとワークフローは失敗します。また、ファイルを差し替えると attestation と一致しなくなります。

## 5. pre-release を昇格させる

ワークフローが作成した Release は pre-release です。昇格させるまで、`install.sh` は前のバージョンをインストールします。

自動生成されたリリースノートには PR のタイトルしか載らないため、主な変更点や挙動の変更を補足してください。master にマージ済みの変更は PR でレビュー済みのため、他のメンテナーによるリリースのレビューは不要です。

リリースノートを整えたら、[`gh release edit`](https://cli.github.com/manual/gh_release_edit) で最新のリリースに昇格させます。

```bash
gh release edit v$VERSION --prerelease=false --latest
```

## 6. Homebrew の formula を更新する

昇格後に実行します。

```bash
./scripts/update-homebrew-formula.sh $VERSION "<git user.name>" "<git user.email>"
```

soracom/homebrew-soracom-cli にバージョンと SHA256 を更新する PR が作成されます。

formula の更新に使うスクリプトは、アセットのダウンロードが 404 などで失敗してもエラーにならず、誤った SHA256 のまま PR を作成します。マージする前に、PR の SHA256 が公開済みアセットの値と一致することを確認してください。

```bash
gh api repos/soracom/soracom-cli/releases/tags/v$VERSION \
  -q '.assets[] | select(.name | test("_(darwin_(amd64|arm64)|linux_(amd64|arm64|386|arm))$")) | "\(.name) \(.digest)"'
```

マージ後、`brew update && brew upgrade soracom-cli && soracom version` で新しいバージョンがインストールされることを確認します。

## 7. AWS Lambda Layers を公開する

昇格後に実行します。

```bash
./scripts/build-lambda-layer.sh $VERSION
./scripts/release-lambda-layer.sh $VERSION <AWS プロファイル名>
```

`build-lambda-layer.sh` は、GitHub Release からバイナリをダウンロードして Layer の zip を作成します。`release-lambda-layer.sh` は、AWS アカウントで有効なリージョン（ap-east-1 を除く）に、`soracom-cli-<バージョン>`（x86_64）と `soracom-cli-<バージョン>-arm64` の 2 つの Layer を公開し、誰でも参照できる権限を付与します。AWS プロファイル名を省略するとテスト用のプロファイルが使われるため、本番用のプロファイルを必ず指定してください。

> [!WARNING]
> README に記載している Layer の ARN は、バージョン番号 `:1` を前提にしています。スクリプトが途中で失敗しても再実行しないでください。再実行すると、公開済みのリージョンに `:2` が追加で作成されます。スクリプトにはリージョンを指定する機能がないため、残りのリージョンには、スクリプト内の `aws lambda publish-layer-version` と `aws lambda add-layer-version-permission` を手動で実行してください。

完了後、README に記載している ARN で Layer を取得できることを確認します。Layer を公開した AWS アカウント以外の認証情報で実行すると、誰でも参照できることもあわせて確認できます。

```bash
aws lambda get-layer-version-by-arn --region ap-northeast-1 \
  --arn "arn:aws:lambda:ap-northeast-1:717257875195:layer:soracom-cli-${VERSION//./}:1"
aws lambda get-layer-version-by-arn --region ap-northeast-1 \
  --arn "arn:aws:lambda:ap-northeast-1:717257875195:layer:soracom-cli-${VERSION//./}-arm64:1"
```

## Snap について

Snap パッケージ（Snap Store の `soracom`）は非推奨で、リリース時に更新しません。
