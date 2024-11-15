# Golang / Gin JWT Auth Tutorial

Install
```
$ go get -u github.com/gin-gonic/gin

$ go get github.com/githubnemo/CompileDaemon

$ go install github.com/githubnemo/CompileDaemon

$ go get -u gorm.io/gorm

$ go get -u gorm.io/driver/postgres

$ go get github.com/joho/godotenv

$ go get github.com/gin-contrib/cors

$ go get -u github.com/golang-jwt/jwt/v5

$ go get -u golang.org/x/crypto/bcrypt
```

# DB Setup:

```
$ psql -U postgres
```

```
$ CREATE USER myuser WITH PASSWORD 'mypassword';
```

```
$ CREATE DATABASE mydb;
```

```
$ GRANT ALL PRIVILEGES ON DATABASE mydb TO myuser;
```

Navigate into specific DB
```
$ \c vintagejazzrecord
```

Check Table Existence
```
$ \d+ albums
```

Drop Table
```
$ DROP TABLE IF EXISTS albums;
```

Tutorial Video:
- [CRUD App to build JSON Web API using GOLANG , Postgres, React Redux, GIN web framework, GORM - 2023](https://www.youtube.com/watch?v=pIZSR7u_NlI)
- [JWT Authentication, Cookie using GOLANG web API , Postgres, GIN web framework, GORM - 2023 Pt- 4](https://www.youtube.com/watch?v=HAZCtEyhPWA)

43:37