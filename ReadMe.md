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

To run application type

```bash
go run cmd/main.go
```

## Pushing to github branch

According to your personal github branch, please commit meaningful message. Push only to your branch and notify the responsible person to make merge request from your branch to main branch. ilfat and sardor branches only for development purpose. main branch is production branch. If you have any question feel free to contact!
