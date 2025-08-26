### Testing CDC (change data capture) in golang

We will store games and let updated moves generate monotonic version bumps.

### Good to know:
* Postgres 16 Docker image comes with migration on bootup in Docker container.

### Useful commands
```
docker compose down -v
```