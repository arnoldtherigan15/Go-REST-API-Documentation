# Go-REST-API-Documentation

## Step 1: Basic Setup

We will be using third party packages in this application. If you have never installed them before, you can run the following commands:

```
go get github.com/badoux/checkmail
go get github.com/jinzhu/gorm
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/jinzhu/gorm/dialects/mysql" //If using mysql 
go get github.com/jinzhu/gorm/dialects/postgres //If using postgres
go get github.com/joho/godotenv
go get gopkg.in/go-playground/assert.v1
```

Next, create an api and tests directory

```
mkdir api && mkdir tests
```

Create the .env file to house our environmental variables

```
touch .env
```

At this point this is the structure we have:

```
your_repository
├── api
├── tests
├── .env
└── go.mod
```

## Step 2: Env Details

```
# Postgres Live
DB_HOST=127.0.0.1
DB_DRIVER=postgres
API_SECRET=98hbun98h #Used for creating a JWT. Can be anything 
DB_USER=steven
DB_PASSWORD=
DB_NAME=fullstack_api
DB_PORT=5432 

# Postgres Test
TestDbHost=127.0.0.1
TestDbDriver=postgres
TestApiSecret=98hbun98h
TestDbUser=steven
TestDbPassword=
TestDbName=fullstack_api_test
TestDbPort=5432
```

## Step 3: Create Models

Inside the api create the models directory

```
mkdir models
```

Since we are building a blog, we will have two model files for now: **User** and **Post**

The User model:
In the path ```api/models``` create the user model:

```
touch User.go
```
