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
4. Copy the `application.yaml.example` file to `application.yaml` and fill in the values for the environment variables
5. Set the Google OAuth redirect URL to `http://localhost:8080/google-user/login` in your Google Cloud Console
6. Build the Docker image for the web application using the command `docker build -t sample-web-app .`
7. Run the Docker container for the web application using the command `docker run -p 8080:8080 sample-web-app`

## Usage

To use the web application, open your browser and go to `http://localhost:8080`. You will see the home page of the web application.

To sign up or sign in, click on the `Sign in with Google` button and follow the instructions. You will be redirected to the dashboard page where you can see your tasks.

To see your profile information or sign out.

## Dependencies

This web application uses the following packages:

- [github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt) - A package for creating and verifying JSON Web Tokens (JWT)
- [github.com/gorilla/csrf](https://github.com/gorilla/csrf) - A middleware for CSRF protection
- [github.com/gorilla/sessions](https://github.com/gorilla/sessions) - A package for managing user sessions
- [github.com/xhit/go-simple-mail/v2](https://github.com/xhit/go-simple-mail) - A package for sending emails
- [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2) - A package for parsing and generating YAML

## License

This web application is licensed under the MIT License. See the `LICENSE` file for more details.
