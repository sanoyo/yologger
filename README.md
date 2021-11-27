# yologger
ログの実装

## how to use

```go
logger := yologger.New(os.Stdout)
logger.Info("failed to fetch URL").Out()

$ [info] 2021-11-27 17:00:56.250237 +0900 JST m=+0.000191919  failed to fetch URL
```

## 参考リポジトリ
https://github.com/uber-go/zap
