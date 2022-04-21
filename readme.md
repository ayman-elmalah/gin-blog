## Installation guide

- Clone repo
`` git clone https://github.com/ayman-elmalah/gin-blog && cd gin-blog``
- ``go mod tidy``
- ``cp .env.example .env``
- Change values in .env
- Create database and put it's credentials to .env

## Commands

#### To run migrate
``go run main.go migrate``

#### To run seeder
``go run main.go seed``

#### To start app
``go run main.go serve``