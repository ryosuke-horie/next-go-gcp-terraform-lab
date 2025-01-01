# Google Cloud側（バックエンド）

## ローカルで利用するコマンド

```bash
terraform init \
  -backend-config="credentials=key.json"

terraform apply 

# applyの前後でADCがないなら
gcloud auth application-default login
```

## Terraformの管理外のリソース

- Remote Backend用のCloudStorage
  - Terraformの状態管理ファイル格納用
- Dockerイメージ格納用のArtifact Registry
  - CI/CDで管理化にあるとフローが複雑化するため
