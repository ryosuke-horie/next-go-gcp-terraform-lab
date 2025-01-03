# next-go-gcp-terraform-lab

このリポジトリは、Next.js / Go / GCP / Terraform / Kubernetes などを組み合わせた構成のハンズオン学習用プロジェクトです。
フロントエンド、バックエンド、インフラまで一貫して小規模アプリを構築することで、各技術の連携方法を学ぶことを目的としています。

## 概要

- **フロントエンド**: Next.js (pages router) + TypeScript + MUI + React Hook Form (zod連携)
- **バックエンド**: Go + chi + xo + PostgreSQL (CloudSQL)
- **インフラ:** Terraform で GCP (GKE Autopilot, CloudSQL, Artifact Registry など) を構築
- **CI/CD**: GitHub Actions を用いたビルド・テスト・デプロイ

簡単な TODOリストを題材に、ローカル環境からクラウド環境へのデプロイまでを通しで実践できるようにしています。

## ディレクトリ構成

```text
next-go-gcp-terraform-lab/
├─ frontend/ (Next.js + MUI + TypeScript + React Hook Form + zod)
├─ backend/ (Go + chi + xo によるAPI、PostgreSQL操作)
│ ├─ internal/models (xo生成コード)
│ └─ main.go (メイン処理やルーティングなど)
├─ infra/ (TerraformによるGCPリソース構築ファイル)
├─ .github/workflows (GitHub Actionsのワークフロー設定)
└─ README.md (このファイル)
```

## ローカル実行の流れ

1. リポジトリをクローンし、backendディレクトリでGoモジュールを初期化してからサーバーを起動し、ローカルのPostgreSQLと接続
2. frontendディレクトリでNode.js環境を整え、依存パッケージをインストール後に開発サーバーを起動
3. ブラウザから <http://localhost:3000> を開き、TODOリストの一覧・追加フォームが動作するか確認

## GCPデプロイの流れ

1. GCPプロジェクトとサービスアカウントを用意し、Terraform (infraディレクトリ) でGKE AutopilotやCloudSQL、Artifact Registryを作成
2. GitHub Actions (.github/workflows) を設定し、DockerイメージをビルドしてArtifact Registryにpush → GKEへデプロイ
3. Next.jsフロントエンドを同じGKE上で稼働させるか、Cloud RunやVercelなどへ分割してホスティングするなど構成を選択

## 使用技術・バージョン

| 領域 | 技術 / ツール |
|------------|-------------------------------------------------------------------------------|
| フロント   | Next.js, TypeScript, MUI, React Hook Form, zod  |
| バックエンド| Go , chi, xo, PostgreSQL, Docker |
| インフラ   | Terraform , Google cloud (GKE Autopilot, CloudSQL, Artifact Registry) |
| CI/CD  | GitHub Actions |

## フロントエンド

<https://opennext.js.org/cloudflare>
上記のテンプレートを利用し、Cloudflare Workersへデプロイ。

<https://next-go-gcp-terraform-lab.ryosuke-horie37.workers.dev>
