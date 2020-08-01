# Ejemplo 02

En este segundo ejemplo vamos a correr un contenedor con la imagen ubuntu en la cual vamos meternos de forma interactiva en el contendor y ejecutar comandos.

## Pasos
```sh
docker run -ti ubuntu bash
```

## Resultado

```sh
Unable to find image 'ubuntu:latest' locally
latest: Pulling from library/ubuntu
3ff22d22a855: Pull complete 
e7cb79d19722: Pull complete 
323d0d660b6a: Pull complete 
b7f616834fd0: Pull complete 
Digest: sha256:5d1d5407f353843ecf8b16524bc5565aa332e9e6a1297c73a92d3e754b8a636d
Status: Downloaded newer image for ubuntu:latest
root@86469f917fea:/# 
```

## Pruebas

### 01) Prueba dentro del contenedor de ubuntu ver la version de linux.

```sh
cat /proc/version
```

### 02) Prueba dentro del contenedor de ubuntu crear un archivo.

```sh
touch nuevo_archivo_prueba
```

### 03) Prueba dentro del contenedor de ubuntu instalar un paquete y ejecutarlo.
```sh
apt update && apt install screenfetch
screenfetch
```

Limpiar la imágen

En el caso que quieras borrar la imagen ubuntu de tu maquina para que esta no ocupe espacio podes hacerlo con el siguiente pasos.

### 1 Verificamos si tenemos contenedores creados con el siguiente comando:

Esto nos va alistar los contedores que tenemos, para ellos vamos a buscar los que pertenescan a la imagen `ubuntu`

```sh
docker container ls --all
```

#### Salida
```sh
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                      PORTS               NAMES
4c31ea22f04d        ubuntu              "bash"              25 minutes ago      Exited (1) 12 minutes ago                       upbeat_heyrovsky
```

### 2 Borramos el contenedor que no usamos mas.

Vamos a tomar el nombre del contenedor que lo vemos en la salida anterior y vamos a ejecutar el siguiente comando que borra el contenedor.

```sh
docker rm upbeat_heyrovsky
```

### 3 Borramos la imagen.

Una vez que no tengamos mas contenedores que pertenescan a nuestra imagen borramos la imagen sin problema.

```sh
docker rmi ubuntu
```
