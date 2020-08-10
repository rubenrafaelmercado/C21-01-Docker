# Exercice 02

En este caso vamos a crear un sitio para dar de alta alumnos y listarlos, compuesto por 3 services:

1. `proxy`
2. `web`
3. `api`

Este sitio se va a manejar con docker-compose e imagenes de docker. En las carpetas estan los archivos necesarios para poder levantar los docker y de los servicios `proxy` y `web`

Para el servicio de la api, tenes que crear un archivo Dockerfile con la imagen similar al del ejercicio `01-make-dockerfile-api`.

## Proxy

El servicio debe contemplar las siguientes condiciones:

1. Basarse en la imagen `nginx:lates`
2. Hacer una redireccion de puerto del `8080` de la maquina local al `80` del contendor.
3. Mapear el archivo `./proxy/nginx/site.conf` del proyecto al `/etc/nginx/conf.d/default.conf` del contendor.

## Web

El servicio debe contemplar las siguientes condiciones:

1. Basarse en la imagen del dockerfile de la carpeta `./web`

## Api

El servicio debe contemplar las siguientes condiciones:

1. Basarse en la imagen del dockerfile de la carpeta `./api`


# Ejercicio extra

En el caso de que les quede tiempo y quieran experimentar. unifiquen el serivio `proxy` y `web`, en un solo contendor.
