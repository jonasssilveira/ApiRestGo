#O Que é Go?

- Go é uma linguagem de programação criada por Robert C. Pike (Unix, UTF8), Ken Thompson (Unix e UTF8)
- Criada pela e para a google, com a intenção de uma linguagem simples e robusta
- Standard library bem completa Multiplataforma. 
- Garbage collection (lightning fast!) Cross-compile.
- Compilada e de tipagem forte e estatica

##Quando Usar 
- Quando precisar de um serviço de escala
- Serviços da web, microservicos, crypto moedas
- Uma das mais rapidas

##Inicio 
- Muito parecido com o C
- Um hello world https://go.dev/play/
- Como que instanciamos uma variavel https://go.dev/play/p/9-SkNuEvaBZ
- Tambem podemos usar o operador curto de declaração, ou :=, ex: https://go.dev/play/p/SF-RmUxySG6
- Não é orientada a objetos, porém tem umas coisas que parecem como interfaces
- Não existem classes mas utilizamos structs como em C
```
type Pessoa struct{
    Nome string
    Idade int
    Altura float64
 }
```
- Só existe os modificadores public e private
- Acesso direto à memoria alocada
- - Não utiliza listas, mas arrays e uma implementação do array
```
sliceDeNumeros := []int{1, 2, 3, 4, 5}
arrayDeNumeros := [5]int{1, 2, 3, 4, 5}
```
- utilização de interfaces https://go.dev/play/p/yAgWWnN0nT6
- Sem while, o for tem diversos tipos de implementacao `for{}` loop infinito e
```
for int i = 0; i< 10; i++{}
```
- Loop com final
- Tambem existe a iteração por um array/slice com o range como demostrado em https://go.dev/play/p/yAgWWnN0nT6