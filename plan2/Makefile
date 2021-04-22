LOCAL_DB_USER := user
LOCAL_DB_PASS := password
LOCAL_INSTANCE_CONNECTION_NAME := tcp\(127.0.0.1:3306\)/sample

.PHONY: generate
generate:
	go run github.com/99designs/gqlgen generate

.PHONY: boil
boil:
	sqlboiler mysql

.PHONY: run
run:
	DB_USER=$(LOCAL_DB_USER) \
	DB_PASS=$(LOCAL_DB_PASS) \
	INSTANCE_CONNECTION_NAME=$(LOCAL_INSTANCE_CONNECTION_NAME) \
	realize start
