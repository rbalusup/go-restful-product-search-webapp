# Creating a Restful API with Go, AWS SDK for DynamoDB, Gin Framework for Backend service and ReactJS, NodeJS for the Front-End.

This is a sample web application to demonstrate following features:
1. Setup and interact with my own DynamoDB instance hosted in AWS using Restful API with routing.
2. Create and load sample data (eg: ProductId/TCIN, Title, Price, etc.,) 
3. The API should Query/Scan relevant information, sorting by price descending by default.
4. Populate the results in the web application based on the query string parameter.

# Considerations for this coding challenge exercise

1. This application assumes that the AWS CLI configuration is pre-configured `[i.e., credentials: (aws_access_key_id, aws_secret_access_key) and aws-region]`
2. Full text search on the `product title` is enabled for the users to enter any string and find matching products.
3. DynamoDB `Query` operations do not support text search well. So I have used `Scan` operations although it take longer time, expensive on larger data sets and do not make sorting easily for this exercise. 
4. Since `CloudSearch` is more optimal for the search operations, it was considered but is not implemented as part of this exercise.
5. API is read-only for this exercise.
6. Considered the search bar UI component for typing in the search string, but not implemented as part of this exercise.

# Key tasks accomplished

1. Implemented the Restful API using `Gin Framework` and `Golang` for routing and unit testing.
2. Configurable parameter `enableDataSetup` boolean `true` to setup data in AWS DynamoDB (creating table, load data from the csv file) during the start-up is implemented.
3. Configurable parameter `enableDataSetup` boolean `false` and provide respective table name `productsTableName`, If the table and data is already set-up in the AWS DynamoDB.
4. Implemented the `unit test cases` for the possible scenarios.
5. Implemented the user interface with minimal features using `React JS` to view the results based on the keyword search for `title` (eg: `http://localhost:3000/?query=Red`)
6. Implemented the search operation using `Query`(based on the partition key: `TCIN`) and `Scan` (based on the search keyword on product `Title`)
7. Sorting the product `Price` descending by default using the `Query` operation is done with parameter `ScanIndexForward` setting it to `false` along with the partition key `TCIN`.
8. Sorting the product `Price` descending by default using the `Scan` operation is done with `Golang` provided `sort.Slice` method on the product results.
9. Once the backend application is up and running, access the Restful API endpoint on the browser or `postman` (eg: `http://localhost:8080/api/v1/products/scan?q=Red`)
10. Once the frontend application is up and running, access the web application on the browser using `http://localhost:3000/?query=Red`, where "Red" is the `search string`. 

# Sample Response for Restful endpoint request: `http://localhost:8080/api/v1/products/scan?q=Red`
```
$ {
    data: [
      {
        tcin: "76695893",
        title: "iPhone Red 256GB",
        price: 976.23
      },
      {
        tcin: "76695895",
        title: "iPhone Red 256GB",
        price: 959.23
      },
      {
        tcin: "76695890",
        title: "iPhone Red 128GB",
        price: 767.19
      },
      {
        tcin: "76695886",
        title: "iPhone Red 64GB",
        price: 457.57
      }
    ]
  }
```

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

In order to run this project, ensure that the following software is all installed correctly:

* [Go](https://golang.org/)
* [Dep](https://golang.github.io/dep/)
* [Node.js](https://nodejs.org/en/)
* [Create React App](https://github.com/facebook/create-react-app)

Also ensure that this project is checked out in an appropriate place under the $GOPATH.

### Running the Backend Service

Ensure that Go and Dep are installed and set up on your machine. Download the necessary dependencies by executing `dep ensure`, and then run the backend by running `go run main.go`.

### Running the Web UI

Ensure that `Node.js` is installed on your machine. From the `webapp` directory execute `npm install` to download the dependencies and then `npm start` to run the application.

## Built With

* [Go](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
* [Create React App](https://github.com/facebook/create-react-app) - Create React apps with no build configuration.

