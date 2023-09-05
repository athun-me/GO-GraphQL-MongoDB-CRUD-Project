# GoLang GraphQL MongoDB CRUD Project

This project is a comprehensive example of building a GraphQL API in GoLang for performing CRUD (Create, Read, Update, Delete) operations with a MongoDB database. It serves as a practical reference for developers looking to create efficient and scalable GraphQL-based applications.


1. Create a new folder for the Project
`mkdir graphql`
2. Mod init your project, give it whatever name you like
`go mod init github.com/athunlal/graphql`
3. Get gql gen for your project
`go get github.com/99designs/gqlgen`
4. Add gqlgen to tools.go
`printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go`
5. Get all the dependencies
`go mod tidy`
6. Initialize your project
`go run github.com/99designs/gqlgen init`
7. After you've written the graphql schema, run this - `go run github.com/99designs/gqlgen generate`
8. After you've built the project, these are the queries to interact with the API - 


# Running the application

1. Install the dependencies

```bash
go mod tidy
```

2. Run the application

```bash
go run server.go  
```
3. Copy and paste into the browser.

```bash
    http://localhost:8080/
```

## Get All Jobs

```bash
query GetAllJobs{
     jobs{
         _id title description company url 
        }
    }
```

## Create Job

```bash
mutation CreateJobListing($input: CreateJobListingInput!){
    createJobListing(input:$input){
         _id title description company url 
        }
    }
```
Input
```bash
    {
        "input": {
            "title": "Software Development Engineer - I",
            "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt",
            "company": "Google", 
            "url": "www.google.com/" 
        } 
    }
```

## Get Job By Id

```bash
query GetJob($id: ID!){ 
    job(id:$id){ 
        _id title description url company 
    } 
}

```
Input
```bash
   { "id": "638051d7acc418c13197fdf7" }
```

## Update Job By Id

```bash
mutation UpdateJob($id: ID!,$input: UpdateJobListingInput!) { 
    updateJobListing(id:$id,input:$input){ 
        title description _id company url 
    } 
}

```
Input
```bash
    { 
        "id": "638051d3acc418c13197fdf6", 
        "input": { 
            "title": 
            "Software Development Engineer - III" 
        } 
    }
```


## Delete Job By Id

```bash
mutation DeleteQuery($id: ID!) { 
    deleteJobListing(id:$id){ 
        deletedJobId 
    } 
}

```
Input
```bash
   { "id": "638051d3acc418c13197fdf6" }
```