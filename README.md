## Pruebas en Golang

### Nativos de Go

### Algoritmos de ordenamiento

#### Buble Sort
Es el método de ordenamiento más común y más fácil de hacer. Consta de ir recorriendo posición por posición 
e ir viendo si uno es mayor que otro. La complejidad de este algorítmo termina siendo O(n^2) debido a que 
recorrería el arreglo 2 veces.

#### Merge Sort
Esté es un método divide and conquer, esto significa que separa al array de tal forma de ir ordenandolo 
por "mitades". La idea de este método, es ir ordenando y mergeando mitades. La complejidad de esté 
método de ordenamiento es de T(n) = 2T(n/2) + θ(n) que al ser recursivo puede ser expresado así.

#### Quick Sort
Al igual que el Merge Sort, tambien es un divide and conquer algorithm. Esté algoritmo consta de tener 
un pivot e ir acomodando los valores a la izquierda y derecha. La complejidad de esté método es de 
O(n) debido a que es un algoritmo recursivo. La única desventaja es la complejidad a la hora de armar 
dicho algoritmo.

### Patrones de diseño

### Arquitectura en Go
* RPC (Remote Procedure Call): Llamada a procedimiento remoto o RPC por sus siglas en inglés (Remote Procedure Call) es una técnica que utiliza el modelo cliente-servidor para ejecutar tareas en un proceso diferente como podría ser en una computadora remota. A veces solamente se le llama como llamada a una función o subrutina remota.

### Fuentes

* [Concurrencia](https://adrianalonso.es/arquitectura-del-software/patrones-de-diseno/patrones-de-concurrencia-en-go-barrier-concurrency-pattern/)
* [Tour de Go](https://tour.golang.org/concurrency/1)
* [Design patterns in go](http://www.designpatternsingo.com/)
* [Que es RPC](https://codingornot.com/que-es-rpc-llamada-a-procedimiento-remoto)
* [Merge Sort](https://www.geeksforgeeks.org/merge-sort/)
* [Quick Sort](https://www.geeksforgeeks.org/quick-sort/)