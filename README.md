# Go-Web

A Simple framework for web development using go language
Database using MySQL

This web application that demonstrates how to use Golang with MySQL and Google login.

## Features

- Each webpage visits are logged in `logs/access.log`
- Fatal and panic errors are logged in `logs/error.log`
- User authentication
- Middleware
- Session
- Session flash message
- CSRF Token
- Users can sign up and sign in using their Google account
- Users can also sign up and sign in using his email account

## Installation

To run this web application, you need to have Golang, MySQL and Docker installed on your system. You also need to create a Google OAuth 2.0 client ID and secret for the Google login feature.

Follow these steps to install the web application:

1. Clone this repository to your local machine
2. Create a MySQL database named `sample_web_app` and a user named `sample_user` with password `sample_password`
3. Run the `schema.sql` file in the `db` folder to create the tables for the web application
4. Copy the `.env.example` file to `.env` and fill in the values for the environment variables
5. Build the Docker image for the web application using the command `docker build -t sample-web-app .`
6. Run the Docker container for the web application using the command `docker run -p 8080:8080 sample-web-app`

## Usage

To use the web application, open your browser and go to `http://localhost:8080`. You will see the home page of the web application.

To sign up or sign in, click on the `Sign in with Google` button and follow the instructions. You will be redirected to the dashboard page where you can see your tasks.

To create a new task, click on the `Create Task` button and fill in the details. To view, edit or delete a task, click on the corresponding buttons in the task list. To filter your tasks by status, use the dropdown menu on the top right corner. To mark a task as completed or uncompleted, click on the checkbox next to the task title.

To see your profile information or sign out, click on your name on the top right corner and select the appropriate option from the menu.

## Dependencies

This web application uses the following packages:

- [HTTP Router](https://github.com/julienschmidt/httprouter) - HttpRouter is a lightweight high performance HTTP request router
- [MySQL Driver](https://github.com/go-sql-driver/mysql) - A MySQL driver for Golang
- [Godotenv](https://github.com/joho/godotenv) - A package to load environment variables from `.env` files


- [Google API Go Client](https://github.com/googleapis/google-api-go-client) - A Go client library for Google APIs
- [OAuth2](https://github.com/golang/oauth2) - A package to provide OAuth2 support for Golang

## License

This web application is licensed under the MIT License. See the `LICENSE` file for more details.
