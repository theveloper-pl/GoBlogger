# Welcome to the Blogger
This project is a simple blog website written in Go, using the Gin web framework and GORM for data modeling and database interactions. It is designed to be with the PostgresSQL database running in Docker container.

## Prerequisites
To run this project, you will need to have the following installed on your local machine:

Go
Docker

## Installation
Clone the repository to your local machine:
```
git clone https://github.com/theveloper-pl/GoBlogger.git
```

Navigate to the project directory:
```
cd GoBlogger
```

Build the Docker containers:
```
docker-compose build
```

Run the Docker container:
```
docker-compose up
```

Run the Website:
```
go run main.go
```

The website will be running on localhost:8080.

## Usage
To view a specific post, click on the title of the post on the homepage. This will take you to the post's detail page.

## Contributing
We welcome contributions to this project! If you have any ideas for improvements or have found a bug, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License.


## TODO
1. Form for adding new posts - compulsory
2. Registration + loginning for users + commenting section for them  - additional

### This README.md file was generated by chatgpt
