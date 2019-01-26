
# Whiteboard easy stuff

Estos problemas son muy simples, en 30 minutos puedes meter 3 de ellos fácil, y tienen buena argumentación de _y esta solución es más eficiente que esta otra?_

## Find missing number of a set of 99 numbers ranging 1..100.
El famoso. Este lo hemos comentado varias veces, es bueno, podría ser sustituido por el de **"Sum all numbers from 1..10.000 without using a loop or a lib function"**


## Iterative and recursive Fibonacci
Hay una puta implementación de Fibonacci usando generadores que me encanta, es imbatible.

```
# iterative
def fib_iter():
    a, b = 0, 1
    while True:
        yield a
        a, b = b, a + b

# recursive, param `n` is number in fib sequence to show
def fib_rec(n):
    if n == 0:
        return 0
    elif n == 1
        return 1
    return fib_rec(n-1) + fib_rec(n-2)

```
Y este permite variantes tipo:

_Si cualquera de estas dos funciones fueran una función de la libreria, que no puedes modificarla, como se implementaria una función `fsum(x, y)` que devolviera la suma de los numeros de fib `x` e `y`?_

Claro, si el tio responde que seria `fib(x) + fib(y)` le das con el laptop en la cabeza. Pero bien, eh?


## Print a x-mas tree of N heigth.

Éste, que parece una broma, me gustó bastante por la forma de abarcarlo cuando lo hace en la pizarra delante tuyo, o sea argumentando el algoritmo.

El problema en si es una chorrada hay que dibujar un arbolico, por ejemplo si n = 4:

```
   *
  ***
 *****
*******
```
Fijate que lo más facil es _pensarlo_ en 2 pasos:

```
# [1, 3, 5, 7...]       # left pad = n - i
1    *                  3        *
3    ***                2       ***
5    *****              1      *****
7    *******            0     *******
```
Me refiero que si lo hace en la pizarra y lo encara así es bien... en realidad es muy fácil, pero si lo encaras calculando ya desde el principio los espacios que tienes que dejar a la izquierda y el ancho del layer te puedes liar facilmente.

## Find closing parens
Se trata de buscar si una cadena de parentesis (o brackets en general) está bien cerrada
```
(()())  # good
())(()  # bad
```
El problema se suele solucionar haciendo una lista (un stack, de hecho) de los parentesis que vayan saliendo, cuando encuentras uno de apertura `(` pa'dentro, y cuando encuentras el `)` lo sacas de la lista. Si vas a sacarlo y la lista está vacia, muerte, y si al acabar, la lista está llena, muerte tambien

Hay una implementación muy chula que es almacenar en la lista el inverso al que te encuentras, con lo que la resolución es más directa, [parens.py](parens.py)

Si lo hace bien se puede cambiar el `()` por brackets en general `()[]{}` y el mismo algoritmo vale, basta con meter esto en un `dict`.

## Find sum of two
El famoso que da nombre a nuestro grupo de Telegram y que me destrozasteis, hijos de mala madre.
Es ese en que dada una lista `[5, 2, 89, 11, 1, -54]` y una N tienes que encontrar 2 valores de la lista que sumados sean igual a N.
Si juegas con las variantes de ordenada, no ordenada etc le vas metiendo miedo.



## Bonus track
## Sort a list of `a`, `b`, `c`.
Este tiene mucha gracia es especial para gente como tu que domináis los arrays con mucha soltura y veís a la primera la solución en tiempo y espacio más eficiente. 

Yo te paso `['c', 'b', 'a', 'b', 'c']`, una lista de 3 elementos distintos repetidos al tun-tun

y tu me devuelves `['a', 'b', 'b', 'c', 'c']` es decir la lista ordenada

Imagina la cantidad de respuestas posibles si preguntas a 5 developers
- Reducir en 3 listas y concatenarlas
- Iterar mientras construyes 3 listas y concatenarlas
- Quicksort
- bla bla bla
- Y la solución O(n) que te adjunto. El truco es que al ser sólo 3 valores los puedes agrupar con punteros `left`, `middle` y `right`, recorriendo la lista si encuenas una `a` la metes al principio y pones el `left` a `+1`.

La solución para 2 también tiene gracia. Se puede cambiar `a`, `b` por `1`, `2` segun lo veas.

O quizá la gracia es hacer el de 2 primero y el de 3 después 

Te lo he dejado [sort_list_one_loop.go](sort_list_one_loop.go)
