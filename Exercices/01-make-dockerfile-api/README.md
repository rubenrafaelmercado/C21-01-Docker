Ejercicio 01

En este ejercicio solamente vamos a crear un contenedor docker con el tag `exericie01:go-web` basado en la imagen `golang:1.14` para correr nuestro servicio web creado en go en el archivo `main.go`.

A diferencia de python para go vamos a necesitar los siguientes comando como pasos previos antes de ejecutar el comando que deja corriendo nustros servidor

1. Situado en la carpeta de nuestro codigo `main.go` vamos a instalar las dependencias necerias con el siguiente comando.

```sh
go get -d -v ./...
```
2. Vamos a instalar nuestro proyecto como app ejecutable.

```sh
go install -v ./...
```

3. Ejecutar nuestro servidor.
```sh
app
```

## Correr el contendor

```sh
docker run -p 5000:5000 exericie01:go-web
```

Luego ingresar en su browser a `http://localhost:5000/`
