# Building Microservices API in Go

###  Master Go programming with an in-depth course covering advanced topics such as authentication, authorization, JWT tokens, and refresh tokens. Learn how to write reliable code with effective unit testing techniques, while exploring concepts such as logging, error handling, and modularization. Build secure and scalable web applications with Go to take your coding expertise to the next level.

This is the code repository for [Building Microservices API in Go](https://www.packtpub.com/product/building-microservices-api-in-go/9781805124429), published by [Packt](https://www.packtpub.com/?utm_source=github). It contains all the supporting project files necessary to work through the book from start to finish.

## About Course
If you are looking to take your programming skills to the next level, Go programming is a must-know language in today’s tech landscape. This powerful language has gained widespread popularity due to its simplicity, performance, and scalability. In this comprehensive course, you will dive deep into Go programming and learn how to build efficient and scalable microservices using REST-based APIs.  
  
In this course, you will start with an in-depth exploration of security, covering best practices for securing your applications. You will learn about authentication and authorization strategies, how to implement role-based access control, and generate JWT tokens for secure user management.  
  
Next, you will delve into unit testing techniques for writing reliable and robust code. You will understand state-based testing, testing routes, and services, and learn the difference between mocks and stubs. You will also explore how to generate mocks and rewire applications for effective unit testing.  
  
In addition, this course covers logging and error handling, critical aspects of building robust applications. You will learn how to extract logger and AppError packages from your project to create a Go module for better code organization and reusability.  
  
You will also discover how to modularize your code with the banking-lib module, integrating it into your banking API and banking-auth API. You will refactor your code to use Claims domain objects for parsing JWT tokens and understand the concept of refresh tokens for maintaining secure access to your applications.  
  
By the end, you will gain the skills needed to build secure and reliable Go applications, optimize code organization, and elevate your coding expertise.

## Target Audience

This course is for developers looking to level up their skills in Go programming and learn about REST-based microservices API development. It’s also suitable for college students learning Golang and experienced developers new to Go and microservices. Whether you are a software developer, web developer, or IT professional, this course provides the knowledge, examples, and best practices to excel in your career.

## How To Use This Repo:

<li>For **banking** package:  Refer to the [main branch](https://github.com/PacktPublishing/Building-Microservices-API-in-Go/tree/main)
<li>For **banking-auth** package:  Refer to the [banking-auth branch](https://github.com/PacktPublishing/Building-Microservices-API-in-Go/tree/banking-auth)
<li>For **banking-lib** package:  Refer to the [banking-lib branch](https://github.com/PacktPublishing/Building-Microservices-API-in-Go/tree/banking-lib)
<li>For **banking-swagger** package:  Refer to the [banking-swagger branch](https://github.com/PacktPublishing/Building-Microservices-API-in-Go/tree/banking-swagger)


# banking
##### Run `./start.sh` to download the dependencies and run the the application

To run the application, you have to define the environment variables, default values of the variables are defined inside `start.sh`

- SERVER_ADDRESS    `[IP Address of the machine]`
- SERVER_PORT       `[Port of the machine]`
- DB_USER           `[Database username]`
- DB_PASSWD         `[Database password]`
- DB_ADDR           `[IP address of the database]`
- DB_PORT           `[Port of the database]`
- DB_NAME           `[Name of the database]`

# mysql database
You can use any one of the following procedure to make a database instance, and make the changes to your `start.sh` file accordingly 
1. `docker-compose.yml` file. This contains the init script to generate and tables and insert the default data. You just need to bring the container up

    To start the docker container, run the `docker-compose up` inside the `resources/docker` folder
 
2. `resources/database.sql` this contains the SQL for generating the tables. In case you dont want to use the docker-compose file you can use this file to generate tables and insert the default data

# mocks generator
`./generate-mocks.sh`

# run unit tests
  `./run-tests.sh`
