# E-commerce backend setup

## Installation

Clone the repository to your laptop

```bash
git clone https://github.com/Kreedciik/e-commerce-backend.git
```

Then through terminal move to this directory and create a branch name.
For example:

```bash
git switch -c sardor
```

## Run application

Create database in postgresql and create the tables which are written in migration/query.sql.
Open the terminal in project folder and type this command

```bash
go mod tidy
```

Configure your postgres database credentials through constants. It's written in cmd/main.go. If your credentials same with these then just copy and paste this:

```go
const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "root"
	dbName   = "ecommerce"
	sslMode  = "disable"
)
```

Create a new database ecommerce if it does not exist.
To run application

```bash
go run cmd/main.go
```

## Pushing to github branch

According to your personal github branch, please commit meaningful message. Push only to your branch and notify the responsible person to make merge request from your branch to main branch. ilfat and sardor branches only for development purpose. main branch is production branch. If you have any question feel free to contact!
