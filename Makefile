include .envrc

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api -db-dsn=${GREENLIGHT_DB_DSN} -jwt-secret=${JWT_SECRET}

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	psql ${GREENLIGHT_DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${GREENLIGHT_DB_DSN} up

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy and vendor dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ==================================================================================== #
# PRODUCTION
# ==================================================================================== #

production_host_ip = 54.167.116.69

## production/connect: connect to the production server
.PHONY: production/connect
production/connect:
	ssh -i learnam.pem ubuntu@${production_host_ip}

## production/deploy/migrations: copy and run migrations on production
.PHONY: production/deploy/migrations
production/deploy/migrations:
	@echo 'Copying migration files to production...'
	rsync -rP --delete -e "ssh -i learnam.pem" ./migrations ubuntu@${production_host_ip}:~/greenlight/
	@echo 'Running migrations on production...'
	ssh -i learnam.pem ubuntu@${production_host_ip} 'cd ~/greenlight && source /etc/environment && migrate -path=./migrations -database=$$GREENLIGHT_DB_DSN up'

## production/deploy/api: deploy the api to production
.PHONY: production/deploy/api
production/deploy/api:
	@echo 'Syncing files to production...'
	rsync -rP --delete -e "ssh -i learnam.pem" ./cmd ubuntu@${production_host_ip}:~/greenlight/
	rsync -rP --delete -e "ssh -i learnam.pem" ./internal ubuntu@${production_host_ip}:~/greenlight/
	rsync -rP --delete -e "ssh -i learnam.pem" ./migrations ubuntu@${production_host_ip}:~/greenlight/
	rsync -P -e "ssh -i learnam.pem" ./go.mod ./go.sum ubuntu@${production_host_ip}:~/greenlight/

## production/status: check production service status
.PHONY: production/status
production/status:
	ssh -i learnam.pem ubuntu@${production_host_ip} 'sudo systemctl status api'

## production/logs: view production logs
.PHONY: production/logs
production/logs:
	ssh -i learnam.pem ubuntu@${production_host_ip} 'sudo journalctl -u api -f'