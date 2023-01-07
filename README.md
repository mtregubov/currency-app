# About 
This is a test task for a full stack developer.

The task is to create a cli-tool and a  web application that will display the exchange rate of the Czech crown.

Cli-tool should be able to download the exchange rate from the Czech National Bank website and save it to the database.
URL — https://www.cnb.cz/en/financial_markets/foreign_exchange_market/exchange_rate_fixing/year.txt?year=2018

Web application should be able to display the exchange rate from the database created by the cli-tool.
Web consists of SPA (reactjs) and go backend.

# Usage

## Build docker image
    $ make build-docker

## CLI 
Run the following command:

    $ make run-docker-cli 

After running the sqlite db will be created in the current directory — ./data/data.db
Data will be saved in the table `currencies` with the following fields:
1. date
1. curr
1. quantity
1. rate

## Web 
Run the following command:

    $ make run-docker-srv

Open the following url in the browser:

    http://localhost:8080

### Stop the web server

    $ make stop-docker-srv

##  TODO
1. Add tests
1. Add validation for the input data
