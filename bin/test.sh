#!/bin/bash
docker-compose exec postgres dropdb -U docker testdb \
    && echo "deleted test db" \
    && docker-compose exec postgres createdb -U docker testdb \
    && echo "created test db, running tests..." \
    && docker-compose exec web go test ./... 