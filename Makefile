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
	swag init -d ./ -g cmd/admin/main.go \
    --exclude ./pkg/app \
    -o ./docs/admin --pd

swagger-app:
	swag init -d ./ -g cmd/app/main.go \
	--exclude ./pkg/admin \
	-o ./docs/app