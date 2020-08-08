# Ejemplo 04

En este cuarto ejemplo vamos a utilizar docker-compose definiendo 3 sevicios los cuales estaran comprendidos en distintos contenedores:

Servicios:

1. **proxy:** Utilizaremos un servidor NGINX para redirigir el trafico a los otros 2 servicios
2. **web**: Utilizando un servidor NGINX serviremos de forma statica un front creado en angular que se comunicara con una api interna
3. **api**: Este servicio contendra una api en flask que permite guardar alumnoes y listarlos.

## Settings docker-compose

En este caso para configurar todo utilizaremos un archivo docker-compose.yml que lo comentaremos a continuacion


```yaml
version: '3'

services:
    proxy:
        image: nginx:latest
        ports:
            - "8080:80"
        volumes:
            - ./nginx/site.conf:/etc/nginx/conf.d/default.conf
        networks:
            - code-network
    web:
        build: ./publisher
        networks:
            - code-network
    api:
        build: ./api
        networks:
            - code-network

networks:
    code-network:
        driver: bridge
```

### Version

Los archivos docker-file son versionados por lo que primero debemos espeficiar que version vamos autilizar.


```yaml
version: '3'
```

### Servicios

Luego vamos a declarar los servicios que comprenderan nuestro docker-compose
```yaml
services:
``` 
#### Servicio: Proxi

En el caso del proxy como utilizamos una imagen base estandar sin modificar definimos directamente la imagen `nginx` esto levantara la imagen, para luego aplicarle 
```yaml
services:
    proxy:
    image: nginx:latest
        ports:
            - "8080:80"
        volumes:
            - ./nginx/site.conf:/etc/nginx/conf.d/default.conf
        networks:
            - code-network
```

#### Service: web

En el caso del service web vemos que no tenemos un tag image, sino que tenemos en su lugar un build en el que pasamo el contexto de de un Dockerfile para que buildee una imagen y la incopore en nuestro proyecto.

```yaml
services:
    web:
        build: ./publisher
        networks:
            - code-network
``` 

#### Service: api

En el caso del service api vamos a usar un Dockerfile para crear una imagen que luego se levantara con nuestra api y montara volumenes, mapea el puerto.

```yaml
services:
    api:
        build: ./api
        networks:
            - code-network
```
