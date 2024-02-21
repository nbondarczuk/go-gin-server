NETWORK=test_network

all: create start logs stop rm

create:
	docker network create $(NETWORK)

start:
	docker-compose up -d

stop:
	docker-compose down

logs:
	docker-compose logs -f

rm:
	docker network rm $(NETWORK)
