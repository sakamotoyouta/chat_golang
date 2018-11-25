# 概要

『Go言語によるWebアプリケーション開発』のログ情報出力用パッケージ
こちらもDockerでの実行を想定


## テストするとき
1. cd tracer_test.goがあるディレクトリへ
2. docker container run --rm -it -v $PWD:/usr/src/myapp -w /usr/src/myapp golang:1.8 go test -cover
