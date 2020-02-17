# Angular + Go Auth App for OneCause
This is an application that has log in form in Angular that hits a Go web API.

# Setup
First, the following Go packages will need to be installed (assuming you have a working Go envirnoment and `GOPATH/bin` is in your `PATH`).

[gin](https://github.com/codegangsta/gin) which is used for a live reloading utility for go servers.
```
go get github.com/codegangsta/gin
```

[gorilla/mux](https://github.com/gorilla/mux) to handle routing
```
go get -u github.com/gorilla/mux
```

[cors](https://github.com/rs/cors) for handling Cross Origin Requests
```
go get github.com/rs/cors
```

[jwt-go](https://github.com/dgrijalva/jwt-go) which is a Golang implementation of JSON Web Tokens (JWT)
```
go get github.com/dgrijalva/jwt-go
```

Using a package manager, install all the packages needed for the Angular app.

# Run The Project
After installing, using `npm start` will run both the Go server (by default on port 4201) and Webpack Dev Server (on port 4200). `npm start` runs the `serve.sh` bash script which in turn starts `ng serve` and `gin --port 4201 --path . --build ./src/web-api/ --i --all` which runs the Go server parallely. You can then access the app by visiting `http://localhost:4200`.

# Usage
The web app will automatically load the sign in page were it will ask for an email, password, and token. Upon a successful sign in, a JSON Web Token (JWT) is returned to keep track of an authenticated user and the user will then be redirected to the [OneCause](https://www.onecause.com) website. Once returning to the website after signing in, the email will be displayed in the upper right corner of the screen and the log in page will redirect back to the base URL since the user has already signed in. Once the local storage is cleared to remove the JWT Token, the sign in screen will once again appear.

#### For a successful sign in
* Email: c137@onecause.com
* Password: #th@nH@rm#y#r!$100%D0p#
* Token: The 2 digit hour (using the 24 hour clock) and the 2 digit minute at the time of submission to create a 4 digit token

### CHEERS!
