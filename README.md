## Getting Started
To get started with the mini lakehouse, simply start the Docker containers:
```
docker compose up -d
```

## Nessie CLI
To start a Nessie CLI container, run the following command
```
docker run -it --network=mini-lakehouse_minilake --name=nessie-cli ghcr.io/projectnessie/nessie-cli
```
In the CLI, run the following command to connect to the Nessie server in the Docker network.:
```
CONNECT TO http://nessie:19120/api/v2
```

After connecting, you can then run commands to interact with the Nessie server.  For example, to list all branches/references, run:
```
LIST REFERENCES
```
Use `help` to see all available commands.

## Configurations
### Trino
You may want to run multiple Trino containers to test out different nessie references. To do that, you can follow this steps:
- Duplicate the `trino` configuration directory and name it accordingly (e.g. `trino-dev`).
- Copy the `trino` service in the [docker-compose.yml](docker-compose.yml) file, change the container name, and update the volume to point to the new configuration directory.
- In the new configuration directory, update the `iceberg.nessie-catalog.ref` property in the `iceberg.properties` file to the desired Nessie reference.
- Start the new Trino container using `docker compose up -d <new-container-name>`.
