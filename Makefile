include ${PWD}/.env

USER:=$(shell id -u)
GROUP:=$(shell id -g)

init:
	ansible-playbook -i deploy/hosts.yml deploy/local.yml -t configuration -e @deploy/vars/local.yml -e "USER=$(USER)" -e "GROUP=$(GROUP)"
build:
	docker-compose run --rm app sh -c "CGO_ENABLED=0 go build -o tmp/app cmd/main.go"
up:
	docker-compose up -d && make log
down:
	docker-compose stop
exec:
	docker-compose exec app bash
exec.root:
	docker-compose exec -u root app bash
log:
	docker-compose logs -f app
