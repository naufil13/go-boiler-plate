# Go Boilerplate

This repository provides a boilerplate for Go projects, including a variety of useful libraries and frameworks to help you get started quickly.

## Pre-requisites

Before you begin, ensure you have the following installed on your system:

1. **Go Programming Language**
   - Install Go from the official [Go website](https://golang.org/dl/).

2. **Verify Go Installation**
   - Verify the installed version and path by using the following command:
     ```sh
     go env
     ```

3. **Initialize Go Module**
   - Create a new Go module by running:
     ```sh
     go mod init go-boilerplate
     ```

4. **Install CompileDaemon**
   - Install the CompileDaemon package to automatically recompile your Go application when source files are changed:
     ```sh
     go get github.com/githubnemo/CompileDaemon
     go install github.com/githubnemo/CompileDaemon
     ```

5. **Install godotenv**
   - Install the godotenv package to load environment variables from a `.env` file:
     ```sh
     go get github.com/joho/godotenv
     ```

6. **Install Gin Framework**
   - Install the Gin web framework for building web applications and microservices:
     ```sh
     go get -u github.com/gin-gonic/gin
     ```

7. **Install GORM**
   - Install GORM, the ORM library for Golang:
     ```sh
     go get -u gorm.io/gorm
     go get -u gorm.io/driver/mysql
     ```

8. **JWT Token**
   - Use the golang-jwt package for handling JWT tokens:
     [golang-jwt/jwt](https://github.com/golang-jwt/jwt)

9. **Email Libraries**
   - For sending emails, install Gomail and SendGrid packages:
     - [go-gomail/gomail](https://github.com/go-gomail/gomail)
     - [sendgrid/sendgrid-go](https://github.com/sendgrid/sendgrid-go)

10. **AWS SDK**
    - Install the AWS SDK for Go for interacting with AWS services:
      ```sh
      go get github.com/aws/aws-sdk-go-v2
      go get github.com/aws/aws-sdk-go-v2/config
      go get github.com/aws/aws-sdk-go-v2/service/s3
      ```

11. **Redis**
    - Install the Redis client for Go:
      ```sh
      go get github.com/redis/go-redis/v9
      ```

By following these steps, you will have a solid foundation for building robust and scalable Go applications. Happy coding!
