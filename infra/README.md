# Google Cloud側（バックエンド）

## ローカルで利用するコマンド

```bash
terraform init
terraform apply 

# applyの前後でADCがないなら
gcloud auth application-default login
```

## Terraformの管理外のリソース

- Remote Backend用のCloudStorage
  - Terraformの状態管理ファイル格納用
- Dockerイメージ格納用のArtifact Registry
  - CI/CDで管理化にあるとフローが複雑化するため

## OIDC構築

利用したコマンドは以下参照。
GitHub ActionsでTerraformの生成/削除を行うためTerraformでは実装しない。

```bash
gcloud iam workload-identity-pools create "github-pool" \
  --project="plasma-renderer-446307-u5" \
  --location="global" \
  --display-name="GitHub Actions Pool"

gcloud iam workload-identity-pools providers create-oidc "github-provider" \
  --project="plasma-renderer-446307-u5" \
  --location="global" \
  --workload-identity-pool="github-pool" \
  --display-name="GitHub Provider" \
  --attribute-mapping="google.subject=assertion.sub,attribute.actor=assertion.actor,attribute.repository=assertion.repository,attribute.repository=assertion.repository,attribute.repository_owner=assertion.repository_owner" \
  --attribute-condition="attribute.repository == assertion.repository && attribute.repository_owner == assertion.repository_owner" \
  --issuer-uri="https://token.actions.githubusercontent.com"

gcloud iam service-accounts create "terraform" \
  --project="plasma-renderer-446307-u5" \
  --display-name="Terraform Service Account"

gcloud projects add-iam-policy-binding "plasma-renderer-446307-u5" \
  --member="serviceAccount:terraform@plasma-renderer-446307-u5.iam.gserviceaccount.com" \
  --role="roles/editor"

gcloud iam workload-identity-pools describe "github-pool" --project="plasma-renderer-446307-u5" --location="global" --format="value(name)"
; projects/191083186598/locations/global/workloadIdentityPools/github-pool

gcloud iam service-accounts add-iam-policy-binding "terraform@plasma-renderer-446307-u5.iam.gserviceaccount.com" \
  --project="plasma-renderer-446307-u5" \
  --role="roles/iam.workloadIdentityUser" \
  --member="principalSet://iam.googleapis.com/projects/191083186598/locations/global/workloadIdentityPools/github-pool/attribute.repository/ryosuke-horie/next-go-gcp-terraform-k8s-lab"
```
