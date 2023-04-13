# AWS OpenSearchServerless を AWS Lambda から使ってみる
この[記事](https://zenn.dev/robon/articles/57c464a5536c63)のサンプルリポジトリ

## デプロイ手順
1. OpenSearch Serverless デプロイ
```bash
make deploy-stack
```

2. Makefile の変数更新
- AossEndpoint: OpenSearch Serverless のエンドポイント
- LambdaRole: Lambda 用のロール ARN
```Makefile
AossEndpoint:=https://xxxxxxxxxxxxxxxxxxxx.ap-northeast-1.aoss.amazonaws.com
LambdaRole:=arn:aws:iam::123456789012:role/sample-aoss-role
```

3. Lambda デプロイ
```bash
make deploy-lambda
```

4. Lambda 実行
```bash
# 実行後、レスポンスの内容が response.json に出力される
make invoke
```

## リソース削除
作成したリソースを全て削除する
```bash
make delete
```
