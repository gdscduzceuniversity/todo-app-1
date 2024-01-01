# todo-app-1

Welcome to the repository of the most amazing Todo application in history! This application is designed with the aim of
providing the best user experience and helping you manage your tasks efficiently.

Just kidding, this is a simple todo application that we developed to learn the basics of team working. I hope you
will find it useful and instructive.

## Features

- User authentication
- Add a task
- Mark a task as to-do, in-progress or done
- set specific deadlines
- View all tasks
- edit a task and show the change history
- delete a task

## Tech Stack

- React
- Golang
- MongoDB
- Docker
- Swagger

## For Stand Up On Your Local

### Prerequisites

- MongoDB account
- Docker

### Steps

#### Backend

* Clone the repository

```bash
$ git clone https://github.com/gdscduzceuniversity/todo-app-1.git
```

* Go to your MongoDB account and create a new `todo-app-1` cluster.
* Create collections as follows:
    * `users` collection
    * `tasks` collection (specify as "clustered index collection" while creating)
    * `activites` collection
* Go to `Database Access` and create a new user with `readWrite` permissions.
* Go to `Network Access` and add your IP address to the whitelist.
* Go to `Clusters` and click on `Connect` button.
* Click on `Connect your application` and copy the connection string.
* Go to the backend folder and create a `connection.env` file.
* Paste the connection string to the `connection.env` file as follows:

```bash
MONGO_URI=<your-connection-string>
```

* To build the docker image, run the following command:

```bash
$ docker build -t todo-app-1 .
```

* To run the backend docker image, run the following command:

```bash
$ docker run -p 3000:3000 todo-app-1
```

* Congratulations! You have stood up the backend.

#### Frontend

* coming soon...

### Swagger

* For connect to swagger ui go to http://127.0.0.1:3000/swagger/index.html

### Contributing

* Fork and create a pull request. We will review it as soon as possible.
* If you have any questions, feel free to open an issue.

### License

* This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.