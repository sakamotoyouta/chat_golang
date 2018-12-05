# このリポジトリについて
golangの勉強用。
dockerでの動作を想定。

## 動作方法
Dockerfileを用意したので、イメージを作成してコンテナ起動してあげれば動く。

docker image build -t ビルド後のイメージ名 Dockerfileのディレクトリ
docker container run --rm -d -v $PWD:/usr/src/myapp -p 3000:8080 イメージ名

3000でポートフォワードしているので、3000番でローカルホストを開けばいい。

※ 開発時は `go run` コマンドで動かした方が良いので、コンテナ起動時にコマンドを上書きすることを推奨

```
※ コンテナのポートをオプションにより切り替えることが可能。
docker image build -t ビルド後のイメージ名 Dockerfileのディレクトリ --build-arg port=<value>
```
