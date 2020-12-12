```
$ docker-compose up -d --build
```

```
$ docker exec -it original-code-db bash
root@ec19d85976f4:/# mysql -u root -h db -p
Enter password:
```

```
mysql> CREATE DATABASE originalcode;
```

```
$ docker exec -it original-code bash
root@fe385569a625:/go/src/app# go run main.go
```
