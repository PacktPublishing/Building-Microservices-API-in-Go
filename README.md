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

## Swagger documentation
You need to install go-swagger in your path. I have used [v0.28.0](https://github.com/go-swagger/go-swagger/releases/tag/v0.28.0)
Download the version compatible with your OS and rename the file to `swagger`. This file should be available in your path. If you are on `linux` or `osx` you might also need to add `execute permission` to the downloaded file using `chmod +x swagger`

Generate swagger json file along with models spes

```shell
$ swagger generate spec -o ./swagger.json --scan-models
```

Run the swagger documentation, you can also specify port
```shell
$ swagger serve swagger.json --port 5000
```
