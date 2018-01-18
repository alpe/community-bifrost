# Bifrost
This is a hard fork from the stellar.org Go [monorepo](https://github.com/stellar/go) for the Bifrost service.
More about Bifrost can be found in the [package docs](https://github.com/stellar/go/tree/master/services/bifrost) or the [design document](https://docs.google.com/document/d/1lxn5YXuDWMpX2m9DvKNPHGgZHSG1mviPET_Dm55xRKI/edit#heading=h.ipw1sk4jppwc)

## Integration tests

* Start [Postgresql](https://www.postgresql.org/) Database server.

For example via docker
```bash
docker run -it --rm  --name stellar-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -p5432:5432 circleci/postgres:9-alpine
```

* Run integration tests

Pass the DB url via parameter

```bash
export bifrostTestDBAddress="postgres://root:mysecretpassword@127.0.0.1:5432/circle_test?sslmode=disable"
 go test -v ./database -tags=integration
```

* Manually access Docker DB

```bash
docker exec -it stellar-postgres  /usr/local/bin/psql -d circle_test
```
