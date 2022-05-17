# **Documentación de los Requerimientos, Criterios de Aceptación y Casos de Pruebas**

- Este programa será capaz de convertir números decimales a números romanos. Dependiendo de las entradas que le de el usuario obtendrá buenas salidas o otros tipos de salidas que indique el error realizado por el usuario

## **Requerimientos**

- Req-1: El sistema debe convertir de números decimales a romanos.

## **Criterios de Aceptación**

- Cri-1-1: Solo se acepta números enteros

## **Casos de Pruebas**

| Escenarios | Resumen                       | Prioridad | Precondiciones | Datos de entrada | Pasos                                                  | Resultado Esperado | Resultado Obtenido |
| ---------- | ----------------------------- | --------- | -------------- | ---------------- | ------------------------------------------------------ | ------------------ | ------------------ |
| Esc-1-1-1  | Convertir de decimal a romano | Alta      | N/A            | Ingresar 1       | Abrir el programa y luego digitar los datos de entrada | I                  | I                  |
| Esc-1-1-2  | Convertir de decimal a romano | Alta      | N/A            | Ingresar 2.3     | Abrir el programa y luego digitar los datos de entrada | Error              | Error              |
| Esc-1-1-3  | Convertir de decimal a romano | Alta      | N/A            | Ingresar ha      | Abrir el programa y luego digitar los datos de entrada | Error              | Error              |
| Esc-1-1-4  | Convertir de decimal a romano | Alta      | N/A            | Ingresar 3z      | Abrir el programa y luego digitar los datos de entrada | Error              | Error              |
| Esc-1-1-5  | Convertir de decimal a romano | Alta      | N/A            | Ingresar 49      | Abrir el programa y luego digitar los datos de entrada | XLIX               | XLIX               |
