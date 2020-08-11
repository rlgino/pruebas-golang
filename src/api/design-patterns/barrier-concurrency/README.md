### Barrier Concurrency Pattern

La idea del patrón es poder hacer N cantidad de actividades y devolver un resulado cuando todas estas actividades hayan concluido.
En nuestro ejemplo de HTTP Agregattor vamos a hacer 10 peticiones a jsonplaceholder y se va a devolver el resultado cuando esto finalice.
El patrón no sólo es útil para realizar solicitudes de red; también podríamos usarlo para dividir algunos problemas complejos en simples tareas que ejecutará cada goroutine.

## Fuente:
* https://adrianalonso.es/arquitectura-del-software/patrones-de-diseno/patrones-de-concurrencia-en-go-barrier-concurrency-pattern/