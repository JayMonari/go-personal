#!/bin/sh

docker run -d --rm \
  -e POSTGRES_HOST_AUTH_METHOD=trust -e POSTGRES_DB=complex_pipeline \
  -p 5432:5432 --name complex_pipeline_2 postgres:12.3-alpine

sleep 5

docker exec -it complex_pipeline_2 psql -U postgres -d complex_pipeline \
  -c 'CREATE TABLE "public"."names" (
	"nconst"              varchar(255),
	"primary_name"        varchar(255),
	"birth_year"          varchar(4),
	"death_year"          varchar(4) DEFAULT '\'''\'',
	"primary_professions" varchar[],
	"known_for_titles"    varchar[]
);'

DATABASE_URL="postgres://postgres:@127.0.0.1:5432/complex_pipeline" go run main.go
