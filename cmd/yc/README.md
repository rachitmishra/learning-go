### Podman commands

#### List all containers

```
podman ps -a
```

#### Compose up

```
podman compose up
```

#### Compose down

```
podman compose down -v
```

#### Delete all containers

```
podman container rm -f $(docker container ls -aq)
```

#### Delete all volumes

```
podman volume rm -f $(docker volume ls -q)
```

#### Delete all volumes

```
podman volume rm -f $(docker volume ls -q)
```

#### Prune system

```
podman system prune
```

#### Prune all containers

```
podman container prune
```

#### Prune all volumes

```
podman volume prune
```
