# 概要

ソフトウェア設計のための様々な原則やアーキテクチャ勉強用のベースとなるコード  
（あえてすべてのコードを main.go で実装）

# 主な使用技術

echo  
gorm  
goose

# 使用方法

## 1．Docker イメージのビルド&コンテナの起動

```
$ docker-compose up -d --build
```

## 2．データベースの作成

① DB コンテナ内へ移動

```
$ docker exec -it original-code-db bash
```

② DB 接続

```
root@ec19d85976f4:/# mysql -u root -h db -p
Enter password:
```

③ DB 作成

```
mysql> CREATE DATABASE originalcode;
```

## 3．マイグレーションファイルの実行

① アプリケーションコンテナ内へ移動

```
$ docker exec -it original-code bash
```

② マイグレーションファイルの実行

```
root@fe385569a625:/go/src/app# goose up
```

## 4．アプリケーションの起動

```
root@fe385569a625:/go/src/app# go run main.go
```

# 注意点

・gorm のバージョンが古いと使用できない可能性があるため注意
（ gorm.io/gorm v1.20.4 では使用できず、go.mod の内容を gorm.io/gorm v1.20.7 に手動で書き換えたらできた。）
