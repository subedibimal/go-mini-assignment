# go-mini-assignment
This is [Golang](https://go.dev/) project which registers users, authenticates them and fetches their profile using [`GraphQL`](https://graphql.org/) implemented with Gqlgen and Gorm. 

## Getting Started

Follow the following steps:

```bash
1. git clone the repo
2. cd to the main project
3. copy .env.example and make a .env file
4. add database credentials
5. docker-compose up -d
6. go run server.go
```

## Usage

To make use of the created GraphQL API:
You can either test from web by the help of GraphiQL or through Postman. 


Endpoint = http://localhost:8080/query (for local)

# Registration

```bash
mutation{
  auth{
    register(input: {
      name:"Test User"
      email:"test@mail.com"
      password:"12345678"
      location:"Test place"
    })
  }
}
```

# Authentication/Login

```bash
mutation{
  auth{
    login(email:"test@mail.com", password:"12345678")
  }
}
```

# Fetch user profile

```bash
query{
    getuser{
        id
        name
        email
        location
    }
}
```