## PROJECT OVERVIEW

The project is a web application for managing a farm. It allows users to create, list, and delete farms, as well as view and creating crops related to the farm.

## HOW TO SET UP AND RUN THE APPLICATION LOCALLY

To set up and run the application locally, follow these steps:
1. Clone the repository to your local machine.
2. Clone the .env.model file to a file called .env.
3. Install the required dependencies by running `go mod tidy`.
4. Build the application by running `go build -o main cmd/server/main.go`.
5. Run the application by executing the `./main` command.
6. You should be able to access the application at http://localhost:4000.

## DEPLOYMENT INSTRUCTIONS

To deploy the application, you can use the containerization solution that uses Docker.
1. Run the command `docker-compose up -d` to start the containers.
2. Access the application at http://localhost:4000.

## HOW TO RUN TESTS

To run tests, you can use the `go test ./...` command.

## TO HAVE ACCESS TO THE ROUTES

If you want to test it running the routes and testing inputs you can copy the file called `docs.json` that is located in the root of the project and import it into your postman/insomnia application.