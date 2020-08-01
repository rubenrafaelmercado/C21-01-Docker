# Ejemplo 02

En este segundo ejemplo vamos a correr un contenedor con la imagen archlinux en la cual vamos meternos de forma interactiva en el contendor y ejecutar comandos.

## Pasos
```sh
docker run -ti archlinux bash
```

## Resultado

```sh
Unable to find image 'archlinux:latest' locally
latest: Pulling from library/archlinux
52d0bcceb754: Pull complete 
6b70b2065185: Pull complete 
d3e55ff511d0: Pull complete 
25f632385a0c: Pull complete 
e4b6fe943aed: Pull complete 
482970d479e4: Pull complete 
Digest: sha256:e543fcbafadece75d0129ac04484b1cb2c36c18847c8609ae7634fe11c688651
Status: Downloaded newer image for archlinux:latest
[root@4c31ea22f04d /]# 
```

## Pruebas

### 01) Prueba dentro del contenedor de Arch ver la version de linux.

```sh
cat /proc/version
```

### 02) Prueba dentro del contenedor de Arch crear un archivo.

```sh
touch nuevo_archivo_prueba
```

### 03) Prueba dentro del contenedor de Arch instalar un paquete y ejecutarlo.
```sh
pacman -Syu
pacman -S archey3
archey3
```

## Limpiar la imágen

En el caso que quieras borrar la imagen archlinux de tu maquina para que esta no ocupe espacio podes hacerlo con el siguiente pasos.

### 1 Verificamos si tenemos contenedores creados con el siguiente comando:

Esto nos va alistar los contedores que tenemos, para ellos vamos a buscar los que pertenescan a la imagen `archlinux`

```sh
docker container ls --all
```

#### Salida
```sh
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                      PORTS               NAMES
4c31ea22f04d        archlinux           "bash"              25 minutes ago      Exited (1) 12 minutes ago                       upbeat_heyrovsky
```

### 2 Borramos el contenedor que no usamos mas.

Vamos a tomar el nombre del contenedor que lo vemos en la salida anterior y vamos a ejecutar el siguiente comando que borra el contenedor.

```sh
docker rm upbeat_heyrovsky
```

### 3 Borramos la imagen.

Una vez que no tengamos mas contenedores que pertenescan a nuestra imagen borramos la imagen sin problema.

```sh
docker rmi arhclinux
```
