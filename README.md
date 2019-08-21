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

### Code prerequisites

* Tener configurado todo el entorno de Go [Go Environment](https://golang.org/doc/code.html#Organization)

* Incluir la dependencia de [AWS](https://github.com/aws/aws-sdk-go) para Go:

```
go get github.com/aws/aws-sdk-go
```

## AWS Deployment

1. **Crear un bucket privado de S3.** Importante quedarse con su nombre y su zona, necesistaremos configurarlo en las credenciales. 

2. **Incluir archivo de almacenamiento en S3.** Cargar el archivo [data.json](aws/data.json) dentro del bucket.

3. **Crear un rol de ejecución en el servicio IAM.**  Este rol es para la ejecución Lambda. Los servicios que se deben permitir son S3 y CloudWatch Logs (solo para debug). 
 
4. **Crear las funciones Lambda.** Crear las funciones [saveWallet.js](aws/saveWallet.js) y [getWallet.js](aws/getWallet.js) dentro de Lambda usando el rol creado anteriormente. Importante sustituir el nombre de bucket existente en las funciones por el creado anteriormente.

5. **Sustituir creedenciales.** A parte de nuestro *idkey* y *secretkey* en [credentials.go](credentials.go) también debemos reemplazar los valores de arn de nuestras funciones Lambda y la región donde esten alojadas.

## Running
```
MBP-Dani:~ daniortiz$ POCServerless 

*****************************
**Welcome to POC Serverless**
*****************************

Choose one of the following options:
1. View the current status of the wallet
2. Save coins
3. Exit
> 
```
```
> 1
The wallet has 0.00 coins

Choose one of the following options:
1. View the current status of the wallet
2. Save coins
3. Exit
> 
```
```
> 2
Coins to save:
5
Saving 5 coins...
5 coins saved in your wallet!

Choose one of the following options:
1. View the current status of the wallet
2. Save coins
3. Exit
> 
```
```
> 2
Coins to save:
8
Saving 8 coins...
8 coins saved in your wallet!

Choose one of the following options:
1. View the current status of the wallet
2. Save coins
3. Exit
> 
```

```
> 1
The wallet has 13.00 coins

Choose one of the following options:
1. View the current status of the wallet
2. Save coins
3. Exit
> 3
See you soon!
```


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

