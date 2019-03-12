<!-- id -->

3

<!-- image -->

assets/cover3.jpg

<!-- title -->

REST API and MongoDB with Go

<!-- content -->

L'objectif de cette formation est de mettre en place une API REST simple en Go, qui effectue des opérations basiques (recherche, ajout, modification, suppression) sur une base de données MongoDB. Une approche en TDD (_Test Driven Development_) est adoptée pour la confection des routes de l'API.

## Préparation

### Installation de Go v11

Sous Ubuntu, pour installer la dernière version de Go, tu peux utiliser un _ppa_ :

```bash
sudo add-apt-repository ppa:gophers/archive
sudo apt update
sudo apt install golang-1.11
```

Ajoute ensuite la ligne suivante à ton .bashrc (ou ton .zshrc si tu utilises zsh) :

`export PATH=$PATH:/usr/lib/go-1.11/bin:$HOME/go/bin`

Si tu utilises un autre OS, tu peux toujours suivre les instructions d'installation sur le [site officiel](https://golang.org/doc/install).

<!-- guide -->

* Exécute la commande `go version` et vérifie le résultat (`go version go1.11.2 linux/amd64`)

<!-- content -->

### Environnement de développement

L'éditeur de code recommandé pour Go est Visual Studio Code. Télécharge-le depuis le [site officiel](https://code.visualstudio.com/).

Il te faudra aussi installer l'extension [**Go**](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go) pour VSCode.

*ProTip* : Pour pouvoir rapidement tester des programmes simples en Go, tu peux visiter [The Go Playground](https://play.golang.org/) qui te permet de compiler du code en ligne.

Pour le développement de serveurs, il est également recommandé d'utiliser un petit outil qui reconstruit ton projet dès que tu apportes des modification au code (un équivalent de _nodemon_ pour NodeJS). En Go, tu peux utiliser [gomon](https://github.com/c9s/gomon). Une fois installé, lance simplement la commande suivante :

```bash
gomon
```

Pour tester les routes du serveur, peu importe leur méthode, il sera nécessaire d'utiliser un outil tel que _Postman_.

<!-- content -->

### Les fondations

Afin de te familiariser avec Go ou d'en réviser les bases, regarde du coté de la première partie de la formation [Jeu de la Vie](/game-of-life)

<!-- guide -->

* Assure toi de connaître les bases du Go (variables, types, boucles)

<!-- content -->

### Les structures JSON

Go permet de définir des structures, assez semblables aux structures en C. Un avantage des structures en Go est la possibilité de décorer ses champs pour permettre la _sérialisation_ de la structure selon divers formats.

Soit la structure suivante :

```go
type Message struct {
	Content string `json:"content"`
}
```

L'instruction `json:"content"` permet à l'encodeur JSON de savoir que l'attribut `Content` de la structure doit être encodé dans le champ JSON `content`, et au décodeur JSON de savoir que le champ `content` du JSON doit être assigné à l'attribut `Content` de la structure.

## 1 - Le serveur

### Echo !

Go est fourni avec une librairie standard très complète, qui met à disposition un ensemble de fonctions et de types utiles à la programmation serveur. Cependant, des _frameworks_ tels qu'[Echo](https://echo.labstack.com/guide) ont vu le jour pour simplifier la programmation de serveurs HTTP, en proposant une interface proche d'_Express_ en NodeJS.

Le code d'exemple ci-dessous est très simple. Un serveur est instancié, une route GET "Hello, World!" basique lui est rattachée et le serveur est finalement lancé sur le port 1323.

```go
package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

<!-- guide -->

* Installe Echo `go get -u github.com/labstack/echo/...`
* Crée un fichier `server.go`
* Tape le code permettant d'afficher "Hello, World!"
* Vérifie que `gomon` a bien relancé le serveur
* Visite `localhost:1323` et vérifie que le message "Hello, World!" s'affiche

<!-- content -->

### Effectuer des tests

L'un des avantages du langage Go est qu'il n'est pas qu'un simple langage avec son compilateur. Go est fourni avec de nombreux outils (la commande `go` les liste). Un de ces outils est un pipeline de tests automatiques.

#### Ecrire un test

Les tests du fichier `foo.go` sont écrit dans un fichier nommé `foo_test.go`. Si la fonction à tester est `func Foo()`, sa fonction de test devra se être `func TestFoo(t *testing.T)`. Ces conventions permettent de s'y retrouver, tout simplement.

Imaginons que l'on souhaite tester la fonction `Square(x int) int` qui retourne, ou du moins est censé retourner, le carré de `x`. En utilisant un package tel que `github.com/stretchr/testify/assert`, on peut proposer le test suivant :

```go
func TestSquare(t *testing.T) {
	assert.Equal(t, 100, Square(10), "Must be equal")
	assert.Equal(t, 25, Square(-5), "Must be equal")
}
```

#### Lancer les test

Il suffit de lancer la commande suivante :

```bash
go test
```

L'outil de test détecte les fichiers de tests, et lance les tests qui s'y trouvent.

Depuis Visual Studio Code, l'intégration avancée du langage Go permet également de lancer un test en particulier en cliquant sur le lien `run test` qui apparaît en haut de chaque fonction de test.

<!-- guide -->

* Familiarise-toi avec les tests en Go
* Teste les fonctions de ton choix

<!-- content -->

### Application du TDD au développement serveur

Le TDD est un principe de développement qui consiste à tester une fonctionnalité avant de l'implémenter. L'implémentation est réussie quand tous les tests prévus sont passés.

Nous allons appliquer ce principe pour l'implémentation des routes de l'API. Il nous faut cependant préparer un peu le terrain pour rendre nos tests simples et efficaces.

#### Le routeur

Afin de pouvoir tester le routeur, ce dernier ne peut pas rester dans la fonction `main`.

Le routeur sera donc placé dans un fichier `routes.go`, dans une fonction de la sorte :

```go
func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TODO:
	e.GET("/", ...)
	e.POST("/", ...)
	e.PUT("/", ...)
	e.DELETE("/", ...)

	return e
}
```

Des middlewares ont été ajoutés pour faciliter le debug du serveur, et relancer ce denier en cas de problème.

#### Quelques fonctions utiles

Deux fonctions s'avèrent utiles pour faciliter l'implémentation des tests, en réduisant le nombre de lignes de code répétées :

```go

func request(method, url string, body interface{}) *httptest.ResponseRecorder {
	e := newRouter()
	j, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	return res
}

func response(res *httptest.ResponseRecorder, v interface{}) error {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
```

* `request` permet d'effectuer une requête de test à une certaine `url` avec une certaine `method` en spécifiant, dans le cas d'une méthode POST ou PUT un `body` qui sera envoyé sous forme de JSON.

* `response` permet de décoder la réponse obtenue du serveur sous format JSON dans une structure Go quelconque.

Ces deux fonctions doivent se trouver dans le fichier de test.

#### Un premier test

L'objectif ici est de tester la route `GET /`. Cette route prend en paramètre une _string_ `message` (représenté par l'URL `/?message=test_message`). Le serveur renvoie alors un JSON contenant le message adressé en paramètre :

```json
{
	"content": "test_message"
}
```

Donnons-nous la structure suivante, qui représente la réponse du serveur, ajoutée dans `routes.go` pour cette démonstration :

```go
type Message struct {
	Content string `json:"content"`
}
```

Voici alors un test très simple, implémenté dans `routes_test.go`. Une requête de test est effectuée. On vérifie d'abord qu'il n'y a pas eu d'erreur (code `200`). La réponse est alors décodée dans la structure `Message` prévue à cet effet, et son attribut `Content` est comparé avec le message envoyé au serveur.

```go
func TestGetMessage(t *testing.T) {
	// Make the request
	res := request("GET", "/?message=Salut", nil)
	assert.Equal(t, 200, res.Code)

	// Check content field
	var answer Message
	response(res, &answer)
	assert.Equal(t, "Salut", answer.Content)
}
```

La commande `go test` échoue alors, car le code d'erreur reçu de la part du serveur est `404` et non pas `200`. C'est assez normal, la route en question n'a tout simplement pas été implémentée.

Cette route peut être ajoutée simplement à `routes.go` :

```go
e.GET("/", func(c echo.Context) error {
	return c.JSON(200, Message{c.QueryParam("message")})
})
```

`go test` termine maintenant avec succès !

Nous étudierons les tests de routes plus en détails par la suite !

<!-- guide -->

* Déplace la création du routeur dans une fonction `newRouteur` dans un nouveau fichier `routes.go`
* Appelle la fonction `newRouter` dans `main` juste avant la ligne contenant `e.Start(":1323")`
* Crée un fichier `routes_test.go`
* Ajoute les deux fonctions `request` et `response` au fichier `routes_test.go`
* Code une fonction pour tester un route simple donnée de l'API
* Exécute `go test` et vérifie que le test échoue
* Implémente la route simple correspondant
* Exécute `go test` de nouveau et vérifie que le test passe (ou pas, si ton implémentation ne répond pas aux attentes de ton test)

<!-- content -->

## 2 - La base de données

Laissons le serveur de coté pour le moment, et concentrons nous sur la façon dont nous allons stocker les données que l'API REST va permettre de manipuler.

La technologie de base de données que nous allons explorer est MongoDB. Cette technologie a l'avantage d'être simple à utiliser et très bien intégrée au langage Go.

### Le modèle de données

Nous allons mettre en place le très classique modèle _User_. Ce _User_ possède les attributs suivants :

* Un nom `Name`
* Une adresse mail `Mail`
* La date à laquelle l'utilisateur est créé `CreatedAt`, de type `time.Time`
* N'oublions pas enfin l'ID du _User_ pour son stockage dans une base MongoDB, de type `bson.ObjectId`

L'attribut ID est défini de manière un peu particulière :

```go
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
```

Il est nécessaire de préciser comment chaque attribut est encodé en JSON (pour le serveur) et en BSON (pour la base de données). Par exemple

```go
	...
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
```

<!-- guide -->

* Crée un fichier `model.go`
* Ajoute une structure `UserModel` avec les quatre attributs explicités précédemment
* Spécifie le nom du champ JSON et BSON correspondant pour chaque attribut

<!-- content -->

### Connexion à la base de données

Le modèle défini ci-dessus représente en fait un _document_ en base de données MongoBD. Pour pouvoir ajouter, rechercher, modifier et supprimer des documents, il faut au préalable se connecter à un serveur MongoDB, et choisir une collection pour y stocker les documents.

Un "service" _User_, qui n'est autre qu'un pointeur vers la collection de _Users_, peut être défini comme suit :

```go
type UserService struct {
	col *mgo.Collection
}
```

Les différentes opérations effectuées sur le modèle _User_ en base de données seront simplement des méthodes de `UserService`, que l'on implémentera par la suite.

Pour créer une nouvelle session avec le serveur MongoDB, il suffit de ces quelques lignes :

```go
session, err := mgo.Dial("localhost:27017")
if err != nil {
	log.Panicln(err)
}
```

Ensuite, pour obtenir un pointeur vers la collection _user_ de la base de données _user-db_ afin d'y faire des opérations, il suffit de cette ligne :

```go
col := session.DB("user-db").C("user")
```

La logique liée à la gestion de la connexion à la base de données peut être implémentée dans une fonction :

```go
func init() {
	...
}
```

Cette fonction a comme propriété intéressante d'être exécutée au lancement du programme (pas besoin donc de l'appeler depuis le main ou depuis les tests).

<!-- guide -->

* Implémente la structure `UserService` à la suite de `model.go`
* Crée un fichier `db.go`
* Instancie `User` de type `*UserService` comme variable globale
* Crée une fonction `init()`
* Connecte-toi à la base de données et donne à l'attribut `col` de `UserService` la valeur du pointeur vers la collection `user`

<!-- content -->

### Définition des méthodes du modèle

Nous avons désormais un modèle de donnée (représentatif d'un _document_) ainsi qu'une connexion à une collection de la base de données MongoDB. L'étape suivante consiste à se donner des méthodes pour manipuler les documents en base de données.

Nous allons suivre le motif CRUD :

* _Create_ pour ajouter de nouveaux documents à la collection
* _Read_ pour trouver des documents existants dans la collection
* _Update_ pour modifier les documents de la collection
* _Delete_ pour supprimer des documents de la collection

Ces méthodes sont rattachées à la structure `UserService`, dont l'instance globale est `User`. Cela permet, depuis le reste du programme, et plus spécifiquement dans les _controllers_ où nous en auront besoin, d'appeler directement `User.MyMethod(...)` par exemple.

<!-- guide -->

* Positionne toi à la suite de `model.go`, en dessous de la déclaration de `UserService`
* _Create_ : Définie la fonction `func (s *UserService) Get(mail string) (UserModel, error)`
* _Read_ : Définie la fonction `func (s *UserService) Add(user UserModel) error`
* _Update_ : Définie la fonction `func (s *UserService) Update(mail string, user UserModel) error`
* _Delete_ : Définie la fonction `func (s *UserService) Delete(mail string) error`

<!-- content -->

### Implémentation des méthodes

Depuis l'intérieur de ces méthodes, puisqu'elles sont rattachées à `UserService`, tu as accès à `s.col` qui est un pointeur vers la collection qui nous intéresse. Pour connaître toutes les méthodes accessible via le _driver_ mgo, rendez-vous sur la [documentation](https://godoc.org/gopkg.in/mgo.v2) de celui-ci. Voici néanmoins les méthodes dont tu auras besoin :

```go
user := UserModel{}
err = s.col.Find(bson.M{"mail": mail}).One(&user)

err = s.col.Insert(user)

err = s.col.Update(bson.M{"mail": mail}, user)

err = s.col.Remove(bson.M{"mail": mail})
```

Par ailleurs, n'oublie pas, lors de la création d'un document, de lui attribuer un nouvel ID, et de marquer la date de sa création :

```go
user.ID = bson.NewObjectId()
user.CreatedAt = time.Now()
```

<!-- guide -->

* Implémente les quatre méthodes de `UserService`

<!-- content -->

## 3 - Les routes

Nous avons mis en place le serveur ainsi que le modèle de base de données _User_. Il ne manque plus qu'à implémenter le pont entre les deux, à savoir les routes et leurs _controllers_ qui, à partir des requêtes de l'extérieur, interagissent avec le modèle _User_ en conséquences.

### Définition

Les quatre routes (GET, POST, PUT, DELETE) sont rattachées au serveur Echo dans la fonction `newRouter()` :

```go
...
e.GET("/", UserGET)
e.POST("/", UserPOST)
e.PUT("/", UserPUT)
e.DELETE("/", UserDELETE)
```

Chaque définition de route prend en paramètre l'URL de la route (`/` ici) ainsi qu'une fonction de _callback_, qui traite la requête.

Une telle fonction de _callback_, aussi appelée _controller_ possède la signature suivante :

```go
func callback(c echo.Context) error {
	...
}
```

Nous implémenterons ces _controllers_ plus tard.

<!-- guide -->

* Définis les routes dans `routes.go`
* Déclare les différents _controllers_ à la suite, ou bien dans un nouveau fichier `controllers.go`

<!-- content -->

### Tests

Suivons les principes du TDD, et implémentons nos tests avant de compléter la logique de notre API.

Il te faut rédiger quatre tests, un par route. La structure du test est simple :

* Remise à zéro la base de données et ajoute les documents nécessaires au test courant
* Requête au serveur
* Test du retour serveur
* Test éventuel du contenu de la base de données

Tu peux définir des _Users_ globaux pour faciliter les tests :

```go
var johnDoe = UserModel{
	Name: "John Doe",
	Mail: "john@doe.fr",
}
```

_ProTip_ : Tester son application ne devrait pas affecter la base de données de production (ou du moins la base de données utilisée pour le développement). Avec Go, il est aisé de savoir si le programme est exécuté en conditions de test ou pas, en vérifiant la présence du flag `test.v`. Dès lors, il suffit de changer de nom de base de de données comme ceci :

```go
DBName := "user-db"
if flag.Lookup("test.v") != nil {
	DBName = "user-db_test"
}
```

<!-- guide -->

* Change de nom de base de données si le programme est lancé en conditions de tests
* Définis des _Users_ de test
* Implémente des fonctions de tests en s'inspirant du test simple effectué en première partie

<!-- content -->

### Implémentation

Chaque _controller_ prend en paramètre un contexte `c echo.Context`.
Ce contexte permet d'accéder au contenu de la requête :

```go
// Access query params
mail := c.QueryParam("mail")

// Put the query body into a struct
var s MyStruct
if err := c.Bind(&s); err != nil {
	return err
}
```

Mais également de définir la réponse du serveur :

```go
// Respond with string
return c.String(200, "Hello, World!")

// Respond with JSON
return c.JSON(200, user)

// Respond with status code only
return c.NoContent(401)
```

<!-- guide -->

* Code le contenu des _controllers_ en utilisant les fonctions mises à disposition par Echo et les méthodes de `User`
* Teste ton implémentation au fur et à mesure grâce aux tests rédigés précédemment

<!-- content -->

### Essai grandeur nature

Il t'est désormais possible d'interagir avec l'API depuis ton navigateur ou Postman.

Tu peux également déployer ce serveur très simplement sur une VM ou une machine physique en construisant l'exécutable et en le copiant sur la machine de destination.
