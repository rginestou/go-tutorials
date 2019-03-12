<!-- id -->

2

<!-- image -->

assets/cover2.jpg

<!-- title -->

Distributed Password Brute Force with Go

<!-- content -->

Cette formation à pour objectif de réaliser un programme qui puisse trouver le plus rapidement possible, par une attaque _brute force_, le mot de passe dont on connaît le hash. Il s'agira ici du hash MD5 du mot de passe.

Le programme, écrit en Go, parcourt tous les mot de passes possibles, en calcule leurs hash MD5, puis compare chacun de ces hash avec le hash à attaquer, jusqu'à trouver le hash identique, et donc le mot de passe.

Dans une optique de performance, l'espace de recherche du mot de passe sera réparti sur plusieurs coeurs de la machine de calcul, voire à terme sur plusieurs machines, divisant d'autant le temps nécessaire à trouver le mot de passe.

## Préparation

### Installation de Go v11

Sous Ubuntu, pour installer la dernière version de Go, tu peux utiliser un _ppa_ :

```bash
sudo add-apt-repository ppa:gophers/archive
sudo apt update
sudo apt install golang-1.11
```

Ajoute ensuite la ligne suivante à ton .bashrc (ou ton .zshrc si tu utilises zsh) :

```bash
export PATH=$PATH:/usr/lib/go-1.11/bin:$HOME/go/bin
```

Si tu utilises un autre OS, tu peux toujours suivre les instructions d'installation sur le [site officiel](https://golang.org/doc/install).

### Environnement de développement

L'éditeur de code recommandé pour Go est Visual Studio Code. Télécharge-le depuis le [site officiel](https://code.visualstudio.com/).

Il te faudra aussi installer l'extension [**Go**](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go) pour VSCode.

*ProTip* : Pour pouvoir rapidement tester des programmes simples en Go, tu peux visiter [The Go Playground](https://play.golang.org/) qui te permet de compiler du code en ligne.

<!-- guide -->

* Exécute la commande `go version` et vérifie le résultat (`go version go1.11.2 linux/amd64`)
* Mets en place ton environnement de travail

<!-- content -->

### Les fondations

Afin de te familiariser avec Go ou d'en réviser les bases, regarde du coté de la première partie de la formation [Jeu de la Vie](/game-of-life)

La présente formation est centrée sur l'usage de la concurrence au sein d'un programme Go. Une explication détaillée du fonctionnement des éléments de langage s'y rapportant sera apportée tout au long de la formation.

<!-- guide -->

* Assure toi de connaître les bases du Go (variables, types, boucles)

<!-- content -->

## Fonctions de base

### Manipulation de hashs

Plusieurs fonctions basiques pour manipuler les hash de mots de passe doivent être implémentées pour commencer. Ces fonctions permettent de générer un hash à partir d'un mot de passe, d'obtenir un hash sous forme de string (représentation hexadécimale) et d'obtenir un hash sous forme de tableau de _bytes_ à partir de sa représentation hexadécimale.

Par la suite, le **hash** d'un mot de passe sera implicitement supposé être un **tableau de _byte_** (`[]byte`), et non sa représentation hexadécimale sous forme de string.

Pour ce faire, les packages `crypto/md5` et `encoding/hex` sont de bons alliés. Ils proposent, entre autres, les fonctions suivantes.

Obtenir hash à partir d'un `pass`, donné comme un tableau de _byte_ (`[]byte`) :

```go
var hasher = md5.New()
hasher.Write(pass)
hash := hasher.Sum(nil)
```

Obtenir la représentation hexadécimale d'un hash :

```go
hex.EncodeToString(hash []byte) string
```

Obtenir le hash à partir de sa représentation hexadécimale (`string`) :

```go
hex.DecodeString(str string) ([]byte, error)
```

<!-- guide -->

* Crée un fichier `hash.go`
* Implémente une fonction simple `generateHash(pass []byte) []byte`
* Implémente une fonction simple `hashToString(hash []byte) string`
* Implémente une fonction simple `stringToHash(str string) []byte`

<!-- content -->

### Incrémentation du mot de passe

La méthode du _brute force_ repose sur le principe d'essais successifs de tous les mots de passe possibles de l'espace de recherche. Afin d'explorer cet espace des possibles sans oublier de candidats, une méthode simple est d'incrémenter le mot de passe, tout simplement.

On représente le mot de passe comme un tableau de _bytes_, où chaque _byte_ est le code ASCII du caractère du mot de passe. Ainsi, `[48 73 97]` correspond au mot de passe _0Ia_.

L'entièreté de la table ASCII n'est pas intéressante, cependant. Aussi, tu peux définir une borne inférieure `lb` et supérieure `up` pour réduire le nombre de caractères disponibles pendant la recherche du mot de passe. Par exemple :

```go
const (
	lb byte = 40
	ub byte = 126
)
```

Le mot de passe est incrémenté en partant de la droite (`[48 73 97]` devient ainsi `[48 73 98]`). Lorsque le _byte_ en position $i$ atteint la borne supérieure, ce _byte_ prend la valeur `lb` et le _byte_ en position $i-1$ est incrémenté à son tour. Et ainsi du suite jusqu'à atteindre le mot de passe `[ub ub ... ub ub]`.

L'appel à la fonction d'incrémentation devra renvoyer `true` si l'incrémentation s'est bien passée, `false` si le mot de passe a atteint la borne supérieure.

<!-- guide -->

* Propose une implémentation de `incrementPass(pass []byte) bool` qui respecte les directives listées ci-dessus
* Affine l'implémentation pour la rendre courte et efficace (8 lignes, qui dit mieux ?)
* Crée un fichier `main.go` et la fonction `main` associée.
* Teste tes fonctions sur des exemples de ton choix

<!-- content -->

## Interface utilisateur

### _Flags_ de la ligne de commande

Afin de permettre à l'utilisateur d'interagir avec le programme, une bonne solution est d'avoir recours à des _flags_. Ce sont les fameux arguments passés à la ligne de commande : `my_program -flag1 -flag2`.

Bonne nouvelle, la gestion de ces _flags_ en Go est simple. On donne ci-dessous un usage basique pour des _flags_ booléens, mais [d'autres types sont possibles](https://gobyexample.com/command-line-flags) :

```go
flag1Ptr := flag.Bool("flag1", false, "My first flag")
flag2Ptr := flag.Bool("flag2", false, "My second flag")

flag.Parse()

if *flag1Ptr {
    fmt.Println("flag1!")
}

if *flag2Ptr {
    fmt.Println("flag2!")
}

// Display flag help
flag.Usage()
```

Le programme aura, pour le moment, deux usages :

* **Générateur de hash** : renvoie le hash (sous forme hexadécimal) d'un mot de passe entré par l'utilisateur
* **Worker** : lance une attaque sur le hash rentré par l'utilisateur

<!-- guide -->

* Crée un premier _flag_ booléen `gen-hash` dans le fonction `main`
* Crée un second _flag_ booléen `worker` dans le fonction `main`

<!-- content -->

### Générateur de hashs

Il s'agit de la partie simple du programme. Lorsque le _flag_ `gen-hash` est spécifié, l'utilisateur est invité à entrer un mot de passe. Pour ce faire, le code suivant est intéressant, puisqu'il cache ce que l'utilisateur tape sur son clavier.

```go
pass, _ := terminal.ReadPassword(0)
```

<!-- guide -->

* Récupère le mot de passe entré par l'utilisateur
* Utilise la fonction `generateHash` pour générer le hash du mot de passe
* Utilise la fonction `hashToString` pour obtenir la `string` correspondante au hash
* Affiche le hash du mot de passe grâce à `fmt.Println`

<!-- content -->

### Worker

L'objectif est désormais de construire, en quelques lignes de code, une première fonction simple de _brute-force_ de mot de passe, qui s'exécute dès lors que le _flag_ `worker` est passé en argument du programme.

L'attaque à implémenter est relativement simple : on se donne une longueur $\ell$ qui caractérise les mots de passe à explorer. Partant du mot de passe `[lb lb ... lb lb]`, on compare le hash du mot de passe courant au hash cible entré par l'utilisateur. Tant que les hash ne sont pas égaux et que l'espace n'a pas été totalement exploré, le mot de passe est incrémenté à l'aide de la fonction `incrementPass` codée précédemment.

La fonction `Scanf` du package `fmt` permet de récupérer l'entrée utilisateur (ici le hash à attaquer). Attention, l'entrée est la représentation hexadécimale du hash, et non directement le tableau de _bytes_ exploitable.

```go
var str string
fmt.Scanf("%s", &str)
```

Pour **comparer deux tableaux de _bytes_**, il est recommandé d'utiliser la fonction suivante :

```go
bytes.Equal(hash1, hash2)
```

<!-- guide -->

* Récupère le hash entré par l'utilisateur
* Crée un fichier `worker.go` pour y implémenter
* Crée une nouvelle fonction `forcePass(targetHash []byte, l int) []byte` qui prend en paramètres le hash entré par l'utilisateur et la taille des mots de passe à tester
* Génère le mot de passe initial (composé de _bytes_ égaux à `lb`)
* Implémente l'attaque _brute-force_ à l'aide de `generateHash` et `incrementPass`.
* Retourne le mot de passe trouvé (ou `nil` si aucun candidat ne convient)

<!-- content -->

## Répartition du calcul sur plusieurs _threads_

### Division de l'espace de recherche

Afin d'accélérer la recherche du mot de passe à attaquer, on souhaite désormais répartir les calculs sur $n$ processus qui s'exécuteraient en parallèle, et diviseraient d'autant le temps de l'attaque.

Il te faut donc répartir l'espace de recherche en une partition de $n$ sous espaces. Une manière simple d'y arriver, conceptuellement, est de commencer par calculer le nombre total de mots de passe candidats (notons ce nombre $N$), de diviser ce nombre par $n$ pour obtenir le nombre $k$ de mots de passe testés par processus ($k = N/n$). On donne alors au processus 0 les $k$ premiers mots de passes de l'espace de recherche, au processus 1 les mots de passes $k+1$ jusqu'à $2k$, etc.

Le nombre total de mots de passe à tester est potentiellement très grand. On le stockera donc sur un `uint64`.

<!-- guide -->

* Crée une fonction `func dividePass(l int, n int) ([][]byte, uint64, uint64)` dans le fichier `hash.go`. Cette fonction prend en paramètre la longueur $\ell$ des mots de passe à tester et le nombre $n$ de processus sur lesquels répartir l'attaque. La fonction renvoie un tableau de $n$ mots de passes initiaux donc chaque entrée sera fournie à chaque processus. La fonction retourne également le nombre de mots de passes à tester par processus, ainsi que le nombre total de mots de passe possibles.
* Calcule `p` le nombre de mots de passe possibles, en stockant les puissances successives nécessaires au calcul dans un tableau `powers`
* Calcule `r` le nombre de mots de passes à tester par processus
* Construis `passes` un tableau de mots de passe
* Remplis les mots de passe initiaux à partir du compteur de mots de passes courant

<!-- content -->

#### Les _channels_

Comment discuter entre les différents codes exécutés en parallèles ? La solution proposée par le langage Go, très élégante, est l'usage de [_channels_](https://gobyexample.com/channels). Ces channels sont des tuyaux de communication inter goroutine.

Supposons que tu souhaites lancer un calcul dans une fonction `f`, et rapatrier le résultat du calcul dans le fil d'exécution du programme principal.

Tu peux envoyer le résultat du calcul dans un channel (bloquant), et récupérer ce résultat à l'autre extrémité du _channel_ (bloquant) à l'aide de la notation `<-` :

```go
func f(ch chan int, x int) {
	ch <- x * x // Envoi du résultat dans le channel
}

func main() {
	ch := make(chan int) // Création du channel
	go f(ch, 10) // Lancement de f en parallèle

	result := <- ch // Récupération du résultat

	fmt.Println("Result :", result)
}
```

<!-- content -->

### Modifications de la fonction de _brute-force_

La concurrence implique de modifier quelque peu la définition de la fonction `forcePass`, pour qu'elle puisse accéder aux informations supplémentaires suivantes :

* Le mot de passe initial à partir duquel explorer
* Un _channel_ vers lequel envoyer le mot de passe trouver (ou `nil` le cas échéant)
* Une variable booléenne signifiant si un autre processus a trouvé le mot de passe, auquel cas le calcul du processus courant est interrompu

La nouvelle signature de `forcePass` ressemble donc à cela :

```go
func forcePass(ch chan<- []byte, found *bool, initPass, targetHash []byte, n uint64) {
    ...
}
```

Lorsque le mot de passe est trouvé, il suffit de l'envoyer à travers le _channel_ et de stopper la fonction :

```go
ch <- pass
return
```

<!-- guide -->

* Mets à jour le signature de la fonction `forcePass` existante
* Renvoie le mot de passe trouvé (ou `nil`) à travers le channel `ch`

<!-- content -->

### Lancements des goroutines

Une fois l'espace de recherche divisé en $n$ parties, les $n$ processus correspondants doivent être lancés en concurrence depuis la fonction `main`.

Leurs résultats sont ensuite attendu, et le mot de passe trouvé est affiché.

<!-- guide -->

* Lance, pour chacun des $n$ mot de passes initiaux issus de la division de l'espace de recherche, la fonction `forcePass` dans une goroutine
* Attends ensuite les $n$ résultats
* Affiche le mot de passe trouvé
