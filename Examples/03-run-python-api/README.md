# Ejemplo 03

En este tercer ejemplo vamos a utilizar `Dockefile` para crear una imane personalizada que se base en una imagen existente de python y le agregue nuestro archivo, instale librerias y corra nuestra api.

## Pasos

### Buildear Imagen
```sh
docker build . -t CurosDocker:curso:docker-api01
```
#### Resultado
```sh
Sending build context to Docker daemon  5.632kB
Step 1/8 : FROM python:3.8.4-alpine
 ---> fbfb63e3c6bb
Step 2/8 : WORKDIR /app
 ---> Using cache
 ---> bac9f905985d
Step 3/8 : ADD requirements.txt /app
 ---> Using cache
 ---> b6226a0c8ec6
Step 4/8 : RUN pip install -r requirements.txt
 ---> Using cache
 ---> d8bcb5626c44
Step 5/8 : EXPOSE 5000
 ---> Using cache
 ---> a01a3d19da82
Step 6/8 : COPY src/ /app/
 ---> Using cache
 ---> 27031b71dc1e
Step 7/8 : ENTRYPOINT ["/usr/local/bin/python"]
 ---> Using cache
 ---> ba960a0e2cb2
Step 8/8 : CMD ["app.py"]
 ---> Using cache
 ---> ac5da9780ea6
Successfully built ac5da9780ea6
Successfully tagged curso:docker-api01
```

### Correr el contendor de la imagen

En este caso van a ver que estamos agregandole el parametro `-p 5000:5000`, esto va a mapear el puerto interno del contedor con el puerto de nuestra maquina, para permitirno sacceder desde afuera

```sh
docker run -p 5000:5000 curso:docker-api01
```

#### Resultado

Exelente, ya tenemos nuestra api corriendo y escuchando peticiones.

```sh
 * Serving Flask app "app" (lazy loading)
 * Environment: production
   WARNING: This is a development server. Do not use it in a production deployment.
   Use a production WSGI server instead.
 * Debug mode: on
 * Running on http://0.0.0.0:5000/ (Press CTRL+C to quit)
 * Restarting with stat
 * Debugger is active!
 * Debugger PIN: 316-788-595
```
### Probar la api

para esto vamos a abrir un browser por ejemlo chrome, firefox e ingresamos a la siguiente URL `http://localhost:5000/say-hello/Pablo` y van a ver que les va a responder y en la terminal donde tenemos corriendo nuestra api se va a ver el log.


## Dockerfile explicado.

A continuacion les voy a comentar un poquito de que contiene el Dockerfile.

```Docker
FROM python:3.8.4-alpine

WORKDIR /app

ADD requirements.txt /app

RUN pip install -r requirements.txt

EXPOSE 5000

COPY src/ /app/

ENTRYPOINT ["/usr/local/bin/python", "[app py](app.py)"]
```

### [FROM] Especifica en que imagen se basa para crear la nuestra.

`FROM python:3.8.4-alpine`

### [WORKDIR] Nos define que directorio de trabajo vamos usar.

`WORKDIR /app`

### [ADD] Agrega un archivo requirements.txt del contexto a la imagen.

`ADD requirements.txt /app`

Add en realidad no es recomendable usar solamente para copiar archivos, ya que en realidad para es esta `COPY`. Add es mas completo ya que permite agregar contendi desde una URL entre otras cosas.

### [RUN] Ejecuta en la imagen la instruccion para instalar las librerias

`RUN pip install -r requirements.txt`

### [EXPOSE] Le informa a docker que el contenedor va a escuchar el puerto 5000

`EXPOSE 5000`

### [COPY] Copia el contenido local de la carpeta `src` a la carpeta de la imagen `/app`

`COPY src/ /app/`

### [ENTRYPOINT] Define el entrypoint que va a usar la imagen, permite que un contenedor se ejecute como un ejecutable que corra un programa especifico.

`ENTRYPOINT ["/usr/local/bin/python", "app.py"]`

## Limpiar la imágen

### 1 Limpiamos los contenedores que no estan estamos usando.

```sh
docker container prune
```

### 3 Borramos la imagen.

Una vez limpiados los contenedores borramos la imagen sin problema.

```sh
docker rmi curso:docker-api01
```
