=== IMAGE ===

https://pilsniak.com/wp-content/uploads/2017/04/golang.jpg

=== CONTENT ===

# Game Of Life with Go

## Préparation

### Installation de Go v11

Sous Ubuntu, pour installer la dernière version de Go, tu peux utiliser un _ppa_ :

    sudo add-apt-repository ppa:gophers/archive
    sudo apt update
    sudo apt install golang-1.11

Ajoute ensuite la ligne suivante à ton .bashrc (ou ton .zshrc si tu utilises zsh) :

`export PATH=$PATH:/usr/lib/go-1.11/bin:$HOME/go/bin`

Si tu utilises un autre OS, tu peux toujours suivre les instructions d'installation sur le [site officiel](https://golang.org/doc/install).

=== GUIDE ===

* Exécute la commande `go version` et vérifie le résultat (`go version go1.11.2 linux/amd64`)

=== CONTENT ===

### Environnement de développement

L'éditeur de code recommandé pour Go est Visual Studio Code. Télécharge-le depuis le [site officiel](https://code.visualstudio.com/).

Il te faudra aussi installer l'extension [**Go**](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go) pour VSCode.

*ProTip* : Pour pouvoir rapidement tester des programmes simples en Go, tu peux visiter [The Go Playground](https://play.golang.org/) qui te permet de compiler du code en ligne.

=== CONTENT ===

## 1 - Les Fondations

### Hello, World!

Voici les lignes de codes du programme "Hello World" en Go :

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

Tout comme en C, un programme Go doit contenir une fonction `main` qui joue le rôle de point d'entrée.

Le package `fmt`, importé en troisième ligne, met à disposition plusieurs fonctions. La fonctions `Println` permet ici d'afficher le texte qu'elle prend en paramètre dans le terminal.

Une fois le code prêt, deux choix sont possibles pour lancer le programme. Depuis un terminal, lancer l'une des deux commandes suivantes :

* `go run main.go`: compile le programme à la volée et l'exécute directement, parfait pour une exécution rapide
* `go build main.go` puis `./main`: compile le programme dans le repertoire local, puis l'exécute

=== GUIDE ===

* Crée un répertoire de travail
* Ajoute un fichier `main.go`
* Tape le code permettant d'afficher "Hello, World!"
* Lance la commande `go run main.go`

=== CONTENT ===

### Variables & Types

#### Types basiques

Go est un langage typé, dont les types se rapprochent du C++ (`string`, `int`, `int64`, `float32`, `float64`, `bool`, ...). Contrairement à tous les autres langages, en Go, le type des variables et paramètres est donné *à droite* de ces derniers (ex: `func f(x int)`).
Supposons que tu souhaites crée une variable `x` de type `float32`. Tu peux le faire de différentes manières :

* `var x float32`: déclare seulement la variable (la met à 0)
* `var x float32 = 10.0`: donne la valeur 10.0 à la nouvelle variable `x`
* `x := 10.0`: infère le type de `x` à partir de la valeur qui lui est assignée.

#### Pointeurs

Toutes les variables et objets sont passés en copie entre les fonctions. Cependant, tout comme le C, le Go possède un système de pointeurs, ce qui rend possible la modification de la valeur à laquelle le pointeur réfère.

```go
x := 4
y := x
fmt.Println(x, y) // 4 4

ptr := &x
*ptr = 10
fmt.Println(ptr, x, y) // 0x416020 10 4
```

Si `ptr` est un objet possédant un attribut `attr`, alors, pour accéder à la valeur de cet attribut, pas besoin de la notation `->` du C++, un point suffit !

```go
x := ptr.attr
```

#### Types avancés

Il est possible de définir des _arrays_ (appelées _slices_) et des _maps_ en Go. Ces types utilisent une quantité de mémoire variable, et doivent ainsi être construits explicitement, contrairement aux types de base.

##### Arrays

Un _array_ se note `[]`, ainsi un _array_ d'entier est de type `[]int`.
Pour construire un _array_ d'entier `arr` de 40 éléments, il te suffit d'écrire :

```go
arr := make([]int, 40)
```

Si la taille est fixe et connue à la compilation, tu peux simplement utiliser cette syntaxe :

```go
var arr [40]int
```

Tu peux accéder au 10ème élément de cet _array_ comme ceci:

```go
x := arr[10]
```

##### Maps

Une _map_ permet d'organiser des données sous la forme de couples clé-valeur.
Pour construire une _map_ dont les clés sont des `string` et les valeurs sont des booléens, il suffit d'écrire :

```go
m := make(map[string]bool)
```

Tu peux insérer une valeur avec sa clé de la manière suivante :

```go
m["key"] = true
```

Tu peux accéder à la valeur correspondant à la clé `"key"` de cette _map_ comme ceci:

```go
b := m["key"]
```

=== GUIDE ===

* Crée plusieurs variables de types différents
* Affiche ces variables grâce au package fmt
* Crée un tableau de nombres flottants sur 64 bits de taille 40
* Crée une map de clés entières à valeurs booléennes
* Crée un tableau d'entier de taille N, avec N un entier défini auparavant

=== CONTENT ===

### Conditions, Boucles et Fonctions

#### Conditions

La syntaxe de la condition est très classique. Elle est tout à fait analogue à celle du C ou du Javascript, mais sans les parenthèses :

```go
if x == 10 {
	...
} else {
	...
}
```

#### Boucles

Go supporte plusieurs méthodes d'itérations dans les boucles `for` :

##### Classique

```go
for i := 0; i < 10; i++ {
	...
}
```

##### Range

Le mot-clé `range` permet d'itérer sur un tableau ou sur une _map_.

```go
arr := make([]int, 10)
for key, value := range arr {
	...
}
```

Les boucles `while` n'existent pas en Go. Il suffit seulement d'utiliser un `for` :

```go
i := 0
for i < 10 {
	i++
}
```

#### Fonctions

Une fonction se définit grâce au mot-clé `func` en précisant son nom, suivi de ses paramètres entre parenthèses, puis de son ou ses types de retours :

```go
func f(x int) int {
	...
	return x
}

func g(x, y int) (int, float32) {
	...
	return x+y, float32(x)+10
}

func h(x int) {
	...
	return 1####0
}
```

=== GUIDE ===

* Utilise les différents types de boucles sur des exemples basiques
Définis des fonctions simples à plusieurs paramètres et valeurs de retour


=== CONTENT ===

### Structures

Go n'est pas, à strictement parler, un langage orienté objet (pas de classes ou d'héritage par exemple), mais propose un ensemble de concepts qui donnent au langage une saine et élégante apparence orienté objet.

Sont donnés ci-dessous quelques éléments pour manipuler des "objets" en Go, à savoir des `structs` et leurs méthodes associées.

#### Les structures

Les structures en Go sont tout à fait semblables à leur cousine en C :

```go
type Person struct {
	Name string
	Age int
}
```

Les attributs de `structs` commençant par une majuscule sont _exportés_, c'est à dire qu'ils sont publiquement accessible depuis une base de code différente qui importerai le package dans lequel la `struct` est défini. Les attributs commençant par une minuscule sont eux invisibles.
Tu peux tout de même accéder à tous les attributs, exportés ou pas, depuis le code situé dans le même répertoire.

#### Les méthodes

Contrairement aux classes en C++, la structure ne contient que l'état de l'objet, mais pas de méthodes s'y appliquant.
Ces méthodes sont définies en dehors du corps de la `struct`, à l'aide d'une parenthèse précisant le type auquel la méthode s'applique et son alias dans le corps de la fonction (l'équivalent du `self` en Python) :

```go
func (p *Person) getName() string {
	return p.Name
}

func (p *Person) happyBirthday() {
	p.Age++
}
```

Depuis le reste du programme, il est alors possible d'utiliser ces méthodes :

```go
p := Person{Name: "John", Age: 10}
fmt.Println(p.getName())
p.happyBirthday()
```

=== GUIDE ===

* Crée une structure simple, possédant des attributs de plusieurs types
* Construis quelques méthodes (getters, setters, autres) qui complètent la structure
* Instancier la structure et faire usage des différentes méthodes codées précédement

=== CONTENT ===

## 2 - La Grille

### Mise en place

La grille du jeu de la vie est une matrice de cellules de taille _NxM_.
Tu peux stocker cette grille dans une structure `board` contenant la matrice et ses dimensions. Un tel object te permettra d'accéder facilement à la matrice par la suite, et d'en connaitre ses dimensions.

En Go, tu peux découper ton code en plusieurs fichiers. Dès lors que ces fichiers se situent *dans le même répertoire* et qu'ils arborent la meme ligne `package <mon-package>` en tête du fichier, ces fichiers partagent les mêmes définitions (comme si tu avais tout codé dans le même fichier). Pas besoin de `.h` !

=== GUIDE ===

* Crée un fichier `board.go` dans le même répertoire
* Crée une structure `type board struct {}`
* Ajoute les attributs de largeur `width` et de hauteur `height` de la matrice, de type `int`
* Ajoute la matrice de booléen `grid` (`[][]bool`) en attribut

=== CONTENT ===

### Création

#### Fonction de création de la grille de jeu

Pour faciliter l'instanciation de la `board`, tu peux encapsuler ce processus dans une fonction, que l'on appellera par exemple `newBoard`. Cette fonction se contentera des dimensions de la matrice en paramètre, et retournera un pointeur vers la `board` nouvellement créée.

```go
func newBoard(N, M int) *board {
	...
}
```

#### Instanciation de la grille

Un pointeur vers une `board` peut être créé comme suit :

```go
b := &board{}
```

#### Construction de la matrice de booléens

Pour construire une matrice de taille _NxM_, il te faut construire le tableau de tableau d'abord de taille _N_, puis construire chacun des _N_ tableaux de taille _M_.

=== GUIDE ===

* Crée la fonction `newBoard`
* Instancie la variable contenant la grille
* Construis la matrice de cellules
* Retourne le pointeur vers la grille

=== CONTENT ===

### Méthodes

Pour interagir avec la `board`, il faudra mettre en place plusieurs méthodes rattachées à la `board` :

* Une méthode de mise à jour de la grille de jeu
* Une méthode d'accès à une cellule de la grille, selon ses coordonnées _i, j_
* Une méthode qui permette de modifier une cellule de la grille, selon ses coordonnées _i, j_, en lui assignant un certain booléen

```go
func (b *board) update()
func (b *board) at(i, j int) bool
func (b *board) set(i, j int, v bool)
```

=== GUIDE ===

* Définis les méthodes (en laissant leur contenu vide)
* Complète le contenu de la méthode `at`
* Complète le contenu de la méthode `set`

=== CONTENT ===

### Instanciation

Il est temps d'utiliser la Grille !

L'objet `board` peut être instancié dans la fonction `main` de manière simple grâce à la fonction `newBoard` codée précédemment.

=== GUIDE ===

* Définis deux variables entières `width` et `height` au début de la fonction main en leur donnant une valeur
* Crée un objet board à partir des dimensions de la grille de jeu
* Utilise une combinaison de `board.at(i, j)` et `board.set(i, j, v)` pour tester ton implémentation
* Utilise la commande `go run *.go` pour lancer tous les fichiers sources du répertoire

=== CONTENT ===

### Effets de bord

Le jeu de la vie se déroule étape par étape, en simulant à chaque itération le nouvel état de la grille à partir du précédent. Une première observations peut être faite :

* Le calcul du nouvel état d'une cellule se base uniquement sur les voisins directs de cette dernière.

Afin de simplifier le parcours de la grille et de la recherche des voisins, sans avoir à tester pour les conditions au bord, il est malin de construire une grille en laissant une bordure vide d'une cellule de large que l'on ne touche pas.

Si la grille visible est de taille _NxM_, la grille en mémoire fait _(N+2)x(M+2)_.

=== GUIDE ===

* Changer les dimensions de la grille créée pour prendre en compte le bord dans `newBoard`
* Changer les fonctions `board.set` et `board.at` en conséquences

=== CONTENT ===

### Swap de la Grille

Une seconde observation est la suivante :

* L'état des cellules à une itération donnée ne dépend que de l'état de l'itération immédiatement précédente.

Afin d'optimiser l'espace mémoire du programme et réduire le nombre de créations et destructions de grilles inutile, on peut initier deux grilles dès le départ. On affiche l'état du jeu en se basant sur la première grille, puis on met à jour la deuxième grille en se basant sur la première. A l'itération suivante, on affiche la seconde grille, et on met à jour la première grille à partir de la deuxième.
Ce système de va-et-vient est plus optimal.

=== GUIDE ===

* Ajoute un attribut `swap` à la structure `board` qui désigne l'indice (0 ou 1) de la grille couramment dessinée
* Change le type de la grille pour `[2][][]bool` (`grid` est un tableau de deux matrices)
* Corrige la fonction `newBoard` pour prendre en compte l'ajout d'une dimension à la grille
* Change les méthodes de la `board` en conséquence

=== CONTENT ===

## 3 - La Logique

### Comptage des voisins

Le jeu de la vie est un automate cellulaire dont la fonction de changement d'état dépend de l'entourage immédiat des cellules de la Grille.

Afin de connaitre le nouvel état de chaque cellule, il est nécessaire d'implémenter une fonction permettant de compter le nombre de voisins d'une cellule _(i, j)_ donnée.

Cette fonction doit accéder à l'état de la grille, et pourra donc être ajoutée comme méthode de l'objet `board`.

De part la prise en compte des effets de bords auparavant, il n'est pas utile de tester ces conditions dans la fonction de comptage, tu peux supposer que les cases mémoire auxquelles tu accèdes sont valides !

=== GUIDE ===

* Crée une fonction `(b *board) countNeighbors(i, j int) int` qui compte les voisins autour d'une cellule _(i, j)_
* Implémente la logique de cette fonction

=== CONTENT ===

### État des cellules

Consulte les [règles du jeu](https://fr.wikipedia.org/wiki/Jeu_de_la_vie) qui s'appliquent à chacune des cellules de la Grille.

Il est désormais temps d'implémenter ces règles, sachant la valeur d'une case et le nombre de ces voisins.

=== GUIDE ===

* Crée une fonction `rule(n int, v bool) bool` qui, sachant le nombre de voisins _n_ et l'état d'une cellule _v_, retourne le nouvel état de cette cellule
* Implémente les règles du jeu de manière succinte

=== CONTENT ===

### Mise à jour de la Grille

Cette étape consiste à compléter la méthode `update` de la `board`.
Cette méthode est appelée à chaque itération du jeu de la vie.

Tu feras appel aux fonctions de décompte des voisins et des règles du jeu.

Par la suite, on supposera que la grille d'indice `b.swap`est la grille actuellement affichée, et que la grille d'indice `1-b.swap` est la grille qui doit être calculée pour la prochaine itération.

=== GUIDE ===

* Réinitialise à `false` l'état de la nouvelle grille
* Parcours les cellules pertinentes de la grille actuelle, compte le nombre de voisin de chacune d'entre elle et mets le résultat de la fonction `rule` dans la nouvelle grille
* Echange les grilles (attribut `b.swap`)

=== CONTENT ===

### Configuration de la grille initiale

La grille de jeu est pour le moment vide. Le Jeu de la Vie prend tout son intérêt lorsqu'on initialise l'état du jeu avec certaines structures particulières (un [planeur](https://fr.wikipedia.org/wiki/Planeur_(jeu_de_la_vie)) par exemple).

Décris une configuration initiale de la grille avec les motifs de ton choix.

=== GUIDE ===

* Renseigne toi sur les motifs intéressants du jeu de la vie
* Utilise la méthode `board.set` pour les implémenter depuis la fonction `main`

=== CONTENT ===

## 4 - L'affichage graphique

### Installation des dépendances graphiques

Le Jeu de la Vie sera affiché dans une fenêtre graphique avec un contexte GLFW

L'affichage graphique sera réalisé grâce au package Go `github.com/faiface/pixel`.

Pour installer cette librairie et la rendre importable dans le programme, tout en n'étant pas forcé de placer le répertoire de travail dans le GOPATH, on utilisera une fonctionnalité récente de Go v1.11 : les modules.

Lance la commande suivante à la racine de ton répertoire de travail :


    go mod init go-game-of-life

Installe ensuite le package graphique :

    go get github.com/faiface/pixel/...

Tu pourras désormais importer le package en toute sérénité (les imports sont automatiques si tu as bien mis en place ton environnement de développement Go) :

```go
import "github.com/faiface/pixel"
...
```

=== CONTENT ===

### Ouvrir une fenêtre

L'ouverture d'une fenêtre graphique demande toujours un certain nombre de lignes de code.

Le code ci-dessous crée une fenêtre de taille _WxH_, vérifie qu'il n'y ait pas d'erreurs, puis boucle à l'infini tant que la fenêtre n'est pas fermée par l'utilisateur. Dans cette boucle infinie, la fenêtre est remplie de blanc puis le programme attend 50 ms avant de reboucler.

```go
var win *pixelgl.Window

pixelgl.Run(func() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Go - Game Of Life",
		Bounds: pixel.R(0, 0, W, H),
		VSync:  true,
	})
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Clear(colornames.White)

		// TODO: Update grid state
		// TODO: Display grid

		win.Update()
		time.Sleep(50 * time.Millisecond)
	}
})
```

=== GUIDE ===

* Crée un fichier `window.go` dans lequel les fonctions graphiques seront codées
* Entoure la fonction `pixelgl.Run` de ta propre fonction `run` qui prenne en paramètre `b` un pointeur vers la `board`, et `blockSize` la taille en pixel d'une cellule affichée à l'écran
* Appelle ta fonction `run` depuis `main.go`
* Teste le programme (`go run *.go`) et vérifie qu'une fenêtre blanche apparaît à l'écran

=== CONTENT ===

### Afficher la Grille

La grille est accessible à l'intérieur de la fonction `run` car elle y a été passée en paramètre.
Il est donc possible depuis l'intérieur de la boucle infinie de mettre à jour la `board`, puis d'afficher la grille actuelle cellule par cellule.

Pour afficher les cellules, il te faudra le package _imdraw_, qui s'utilise comme suit :

```go
\\ Initialization
imd := imdraw.New(nil)
imd.Clear()

\\ Set drawing color
imd.Color = colornames.Black

\\ Add a rectangle between the points (10, 10) and (20, 20)
imd.Push(pixel.V(10.0, 10.0))
imd.Push(pixel.V(20.0, 20.0))
imd.Rectangle(0)

// Draw all geometry
imd.Draw(win)
```

=== GUIDE ===

* Crée l'objet `imd` avant l'entrée dans la boucle infinie
* Appelle la méthode `update` de la `board` avant d'afficher les cellules
* `Clear` l'afficheur avant de dessiner les cellules
* Itère sur les cellules de la grille et affiche un rectangle bien placé de la couleur qu'il faut (blanc = cellule morte, noir = cellule vivante)
* Appelle la méthode `imd.Draw` une fois que toutes les cellules ont été ajoutée à l'afficheur

=== CONTENT ===

## 5 - Aller plus loin

### Remplissage aléatoire

De manière facultative, tu peux remplir la grille de manière aléatoire.

Le package `math/rand` met à disposition plusieurs fonctions de génération de nombre aléatoire.

Il est nécessaire de commencer par initialiser le générateur :

```go
rand.Seed(time.Now().UnixNano())
```

On pourra ensuite utiliser la fonction `rand.Float32()` dans l'optique de générer des booléens aléatoires.

=== GUIDE ===

* Crée une méthode `func (b *board) randomInit()` dans `board.go`
* Initialise le générateur aléatoire
* Parcours les _NxM_ cellules de la grille et attribue leur un booléen aléatoire
* Appelle la fonction `randomInit()` depuis `main`

=== CONTENT ===

### Multithreading

Go est aussi connu pour sa gestion simple et puissante de la concurrence. Il est en effet très simple d'exécuter plusieurs parties de son code en parallèle, à l'aide des [goroutines](https://gobyexample.com/goroutines), qui sont une alternative très légères aux _threads_.

#### Les goroutines

Si `f` est une fonction définie dans le programme que l'on souhaite lancer en parallèle de l'exécution du programme, il suffit d'utiliser la commande `go` :

```go
go f()
```

C'est tout.

Le programme continuera son exécution, alors que le contenu de `f` sera exécuté en parallèle. Par défaut, en cas de besoin, Go peut étaler ses calculs sur tous les cœurs du PC.

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

#### Application au Jeu de la Vie

Afin d'accélérer le calcul du nouvel état de la grille, il est intéressant de mettre à jour les lignes de la grille en parallèle en lançant _N_ goroutines, puis d'attendre la fin de tous les calculs avant de passer à la suite.

=== GUIDE ===

* Crée une méthode `func (b *board) updateRow(ch chan int, i int)` qui prend en paramètre un channel entier, et le numéro de la ligne à mettre à jour
* Déplace le code de mise à jour de la ligne de la méthode `update` dans cette nouvelle méthode
* Signalise la fin du calcul de la ligne en envoyant un entier quelconque dans le channel (`ch <- 0`)
* Crée un channel entier dans la méthode `update`
* Lance une goroutine pour `updateRow` par ligne de la grille
* Attends ensuite _N_ signaux de fin de calcul ( `<-ch`) avant de continuer l'exécution normale du programme

=== CONTENT ===

### Tests

L'un des avantages du langage Go est qu'il n'est pas qu'un simple langage avec son compilateur. Go est fourni avec de nombreux outils (la commande `go` les liste). Un de ces outils est un pipeline de tests automatiques.

#### Ecrire un test

Les tests du fichier `foo.go` sont écrit dans un fichier nommé `foo_test.go`. Si la fonction à tester est `func Foo()`, sa fonction de test devra se être `func TestFoo(t *testing.T)`. Ces conventions permettent de s'y retrouver, tout simplement.

Imaginons que l'on souhaite tester la fonction `Square(x int) int` qui retourne, ou du moins est censé retourner, le carré de _x_.
Une fonction de test possible est la suivante :

```go
func TestSquare(t *testing.T) {
	var res int
	res = Square(10)
	if res != 100 {
		t.Error("Expected 100, got ", res)
	}
	res = Square(-5)
	if res != 25 {
		t.Error("Expected 25, got ", res)
	}
}
```

Certains packages tels que `github.com/stretchr/testify/assert` permettent de quelque peu condenser les tests :

```go
func TestSquare(t *testing.T) {
	assert.Equal(t, 100, Square(10), "Must be equal")
	assert.Equal(t, 25, Square(-5), "Must be equal")
}
```

#### Lancer les test

Il suffit de lancer la commande suivante :

    go test

L'outil de test détecte les fichiers de tests, et lance les tests qui s'y trouvent.

Depuis Visual Studio Code, l'intégration avancée du langage Go permet également de lancer un test en particulier en cliquant sur le lien `run test` qui apparaît en haut de chaque fonction de test.

=== GUIDE ===

* Teste les fonctions de ton choix
