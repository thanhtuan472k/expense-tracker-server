#!bin/bash

export DOMAIN_EXPENSE_TRACKER_ADMIN=localhost:8000
export DOMAIN_EXPENSE_TRACKER_APP=localhost:8001
export ENV=develop
export ZOOKEEPER_URI=127.0.0.1:2181
export ZOOKEEPER_PREFIX_EXTERNAL=/expense_tracker
export ZOOKEEPER_PREFIX_EXPENSE_TRACKER_COMMON=/expense_tracker/common
export ZOOKEEPER_PREFIX_EXPENSE_TRACKER_ADMIN=/expense_tracker/admin
export ZOOKEEPER_PREFIX_EXPENSE_TRACKER_APP=/expense_tracker/app


# make update-submodules branch=develop
update-submodules:
	git submodule update --init --recursive && \
	git submodule foreach git checkout $(branch) && \
	git submoduleforeach git pull origin $(branch)

run-admin:
	go run cmd/admin/main.go

run-app:
	go run cmd/app/main.go

swagger-admin:
	swag init -d ./ -g cmd/admin/main.go --pd --parseInternal \
    --exclude ./pkg/app \
    -o ./docs/admin

swagger-app:
	swag init -d ./ -g cmd/app/main.go --pd \
	--exclude ./pkg/admin \
	-o ./docs/app