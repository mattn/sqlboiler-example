# sqlboiler-example

Example application using sqlboiler

## Initialize

Install sqlboiler and sqlboiler-sqlite3

```
$ go get github.com/volatiletech/sqlboiler
$ go get github.com/volatiletech/sqlboiler-sqlite3
```

## Create table

```
$ sqlite3 models.sqlite < ddl.sql
```

## Generate models

```
$ go generate
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
