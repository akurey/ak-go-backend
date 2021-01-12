# README #

### What is this repository for? ###

Use this repository whenever you need to start a Golang backend proyect.

### How do I get set up? ###

* Change .env config to a valid configuration
* Run go ```mod download```
* Run ```go run main.go```

### Things to improve ###

* How the responses are built, especially when the response is a collection of objects.
* Decorators: not sure if this can be done in go but we can think of a way to refactor 
  some overhead done inside the controllers.
* Scaffolding: I added a use case directory to remove some responsibilities from the 
  controllers. Also, added a repository directory to store all files that communicate with external
  services.
* Middlewares: improve how we load middlewares into the routes.

### Who do I talk to? ###

* Heiner León