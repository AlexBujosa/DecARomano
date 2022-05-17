# Requerimientos, Criterios, Casos - Cuestionario

### 1. **¿Qué es un Coding Dojo?**

- Es una reunión en donde un grupo de programadores, trabajan en la solución de uno o varios desafíos de programación, con la finalidad de divertirse y sobre todo mejorar sus habilidades de programación. No es un ambiente en donde humillas al otro programador porque lo haces más rápido que él como si fuera una competencia, sino que colaboras para el desarrollo de todos. Lo que se necesita es una PC y un proyecto digital, en donde se proyecta el desafío, además de que se puede discutir posibles soluciones para resolver el problema. Su duración está entre las 2hrs 30 min a 3hrs.

### 2. **Diferencia entre Requerimientos, Criterios de Aceptación y Escenarios de Prueba. Dar ejemplos a partir de un problema distinto a la referencia.**

- Según lo leído, se puede decir que estos tres conceptos son tres niveles de detalles de lo mismo y todo con el objetivo de describir lo que esperamos del sistema, pero sus diferencias radican en el nivel de detalle de cada uno, ya que un requerimiento define las restricciones y necesidades de un producto de software, pero a gran escala.

Por ejemplo: Un sistema de cálculo de índice académico, vamos a definir algunos requisitos que componen el Iniciar Sesión, básicamente

- Req-1: El sistema debe permitir el acceso al sistema al usuario que esté registrado
- - Ya luego tenemos los criterios de aceptación son un conjunto de reglas que se definen para satisfacer ciertos requisitos
- - Cri-1-1: Se necesita la combinación de letras mayúscula, minúscula y números en la contraseña para poder accesar al sistema.
- - Cri-1-2: Se necesita mínimo de 8 dígitos para aceptar la contraseña.

- - Ahora definiremos los escenarios que son ejemplos concretos que se definen para la validación de un único criterio de aceptación:
    Escena-1-1-1: C12fc28SJ39 - Válido
    Escena-1-1-2: A27Bj12cc - Válido
    Escena-1-1-3: B21JC2817 - No válido

### 3. **De dos ejemplos de requerimientos no-funcionales, y especifique cuál es su categoría (seguridad, performance, portabilidad, etc.)**

1. 1. La aplicación de Facebook genera alrededor de 500 TB de información cada día, y responde de manera rápida a las búsquedas de los usuarios. (categoría de requerimiento no funcional: Rendimiento)
2. 2. La aplicación de Facebook se accede desde cualquier plataforma y dispositivo dígase Celulares, Tablets, Computadoras y sistemas operativos tales como Linux, MacOs, Windows, etc. (Categoría de requerimiento no funcional: Portabilidad)

### 4. **¿Qué es TDD?**

- Es una metodología de diseño que se basa en primero realizar las pruebas antes del proceso de desarrollo de la funcionalidad, que es lo contrario a lo que muchas veces hacemos que posponemos las pruebas para el final. En su mayoría estás están compuestas por pruebas unitarias que son la base para probar funcionalidad por funcionalidad.

### 5. **¿Qué son las pruebas unitarias automatizadas?**

- Son un tipo de pruebas unitarias, que se basan en revisar unidades diminutas del programa ara así encontrar errores de este, lo único que, con ayuda de software, estás controlan su tiempo de ejecución y comparar como cualquier otra prueba los resultados esperados con los obtenidos.

### 6. **¿Cuál fue el primer framework de pruebas unitarias y para cual lenguaje fue creado?**

- El primer framework de pruebas unitarias es SUnit, desarrollando casos de prueba en Smalltalk, originalmente escrito por el padre Extreme Programming (Kent Beck), es un lenguaje reflexivo, es decir, un lenguaje que se inspecciona así mismo mientras se ejecuta. El padre del XP, Kent Beck, describió sobre este framework en su libro “Simple Smalltalk Testing: with Patterns”, en el año 1989.

### 7. **_ Describa los componentes de la arquitectura xUnit _**

- **Corredor de pruebas:** Es un programa binario que ejecuta pruebas implementadas usando una xUnit, informando los resultados de esta.
- **Caso de prueba** Son los escenarios que nos definen si el comportamiento del sistema es el adecuado.
- **Accesorios de prueba:** Son las condiciones previas o requisitos para poder ejecutar una prueba. De forma tal, que después de ejecutarlas vayan a su estado original.
- **Suites de prueba:** Son un conjunto de pruebas que comparten los mismos fines.
- **Ejecucion de pruebas:** Las pruebas unitarias se ejecutan como: Configuración (); _ Cuerpo de la prueba _ y teardown (); Setup() y Teardown() se encargan de iniciar y limpiar el cuerpo de la prueba. De esta forma, no se perturban los resultados de las pruebas futuras.
- **Formateador de resultados:** El formateado de resultados, básicamente produce los resultados en uno o más formato de salida, ósea que no solamente puedes hacer que sea legible para los humanos, sino que también puedes formatear esa información y guardarla un archivo XML o Json, que puede usarse para presentar esa información de otras formas o por ejemplo puede ser usado por aplicaciones que den a mostrar en formas de gráficos esa información.
- **Afirmaciones:** Es una función que verifica el comportamiento de la unidad bajo la prueba. Por lo general, una afirmación expresa una condición lógica que es verdadera que se cumple cuando es algo que esperamos, mientras que sus fallas son llamados excepciones.

### 8. **Indique al menos tres desventajas de las pruebas automatizadas**

- **Mayor cantidad de código a escribir:** Esto es lo más lógico que puede pasar porque si es independiente y se auto ejecuta, significa que hay que dar pasos adicionales para configurarla programarla y que nos lancen las respuestas de manera estructurada como nosotros queremos.
- **Aumento de los esfuerzos de mantenimiento:** Por más automatizado que uno se encuentra, está super claro que el esfuerzo de desarrollo inicial hay que sumarle alrededor del 45% de tiempo adicional y más el tiempo que conlleva de hacer ajustes.
- **Mas costosas:** Son mucho más costosas que las manuales, ya que se deben invertir en herramientas y/o softwares especializados.

### 9. **Indique al menos tres ventajas de las pruebas automatizadas**

- **Eliminan el error humano:** Como se guardan los resultados para futuras comparaciones, las pruebas automatizadas registran los resultados automáticamente, por lo que el error de escribir una conclusión desaparecería.
- **Estabilidad del código:** Mientras más pruebas se realicen ya que son automatizadas, mayor confiabilidad habrá en el código.
- **Agregar otras funciones con facilidad:** Al realizarse todas las pruebas posibles a profundidad, el equipo podrá añadir mas funciones nuevas, sin que estropee todo aquello codificado anteriormente.
