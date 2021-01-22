# clean-arch-with-caff

### Usage Example
```
mysql> create database hoge;
mysql> use hoge;
mysql> CREATE TABLE user ( id SERIAL PRIMARY KEY, name VARCHAR(64) NOT NULL, age INT NOT NULL DEFAULT 0 );
```

```
$ go get -u github.com/ngiyshhk/caff
$ caff --mysql_user root --mysql_dbname hoge --repo github.com/dev-sota/clean-arch-with-caff
```

### Acknowledgments
https://github.com/ngiyshhk/caff
