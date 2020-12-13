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

## 3．アプリケーションの起動

① アプリケーションコンテナ内へ移動   

```
$ docker exec -it original-code bash
```

② アプリケーションの起動   

```
root@fe385569a625:/go/src/app# go run main.go
```