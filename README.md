# POC Serverless-Storage

Proyecto que tiene como objetivo demostrar cómo conseguir persistencia dentro de la tecnología serverless de una manera sencilla.
<br/>
La aplicación consiste en un monedero digital en el cual vamos guardando monedas. Podemos diferenciar 3 partes dentro de la aplicación:   
* Interfaz y frontend: desarrollado con el lenguaje **Go**.
* Backend y servicios: usando tecnología serverless implementado en **Amazon Web Services Lambda** en **Node.js**.
* Base de datos: almacenamiento simple en formato JSON en **Amazon Web Services S3**.

## Getting Started

Para comenzar a usar el proyecto y poder jugar con él se tiene que hacer es clonar el repositorio bajo la carpeta *src* del workspace de Go de nuestro ordenador. 
<br/>
Antes de ejecutarlo debemos sustituir los valores falsos de AWS en el archivo de [credenciales](credentials.go) por los nuestros propios.


Acabados los pasos previos ya se puede ejecutar directamente:
```
go run main.go aws.go constants.go fakecredentials.go functions.go 
```
O construir un ejecutable y ejecutarlo posteriormente desde donde queramos:
```
go install POCServerless
```
Típicamente este ejecutable se encuentra bajo la carpeta */bin* de nuestro workspace. Recordar que es muy útil incluir la carpeta */bin* en nuestro propio Path del ordenador y así tenemos acceso al ejecutable desde cualquier lugar:
```
POCServerless
```

## Prerequisites

### AWS prerequisites

* Tener una cuenta en AWS y obtener sus credenciales (*id* and *secret*).

* Crear un bucket de S3.

* Crear un rol de ejecución en el servicio IAM. Este rol es para la ejecución Lambda. Los servicios se deben permitir son S3 y CloudWatch Logs (solo para debug). 

### Code prerequisites

* Tener configurado todo el entorno de Go [Go Environment](https://golang.org/doc/code.html#Organization)

* Incluir la dependencia de [AWS](https://github.com/aws/aws-sdk-go) para Go:

```
go get github.com/aws/aws-sdk-go
```

## AWS Deployment

Add additional notes about how to deploy this on a live system

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

