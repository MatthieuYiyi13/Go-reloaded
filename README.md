# Go-reloaded

Le but du projet est de crée un programme qui modifie un text rentré dans le fichier `input.txt` pour le ressortir dans `output.txt`, pour cela...

## ce qu'il faut installer avant de commencer

Il faut avoir **Visual Studio Code** et **Go** installé sur son ordinateur. Le projet utilise Go **1.25.0 minimum**, comme indiqué dans le fichier `go.mod`.

### 1. installer et lancer Visual Studio Code

- aller sur le site de [Visual Studio Code]
- une fois installé, lancer Visual Studio Code

### 2. installer Go sur son ordinateur

- aller sur le site de [Go](https://go.dev/dl/) et télécharger une version **1.25.0 ou plus récente** pour son ordinateur
- sur Windows, ouvrir le fichier `.msi` téléchargé et suivre les étapes de l'installation
- pour les autres systèmes, suivre les étapes de la [documentation de Go](https://go.dev/doc/install)
- fermer puis relancer VS Code après l'installation pour qu'il puisse retrouver Go

Go est installé sur l'ordinateur, l'extension dans VS Code sert ensuite à travailler avec les fichiers `.go`.

### 3. installer l'extension Go dans VS Code

- dans VS Code, ouvrir l'onglet **Extensions** sur la gauche, ou faire `Ctrl + Shift + X` sur Windows
- rechercher **Go** et installer l'extension publiée par **Go Team at Google**,
- quand le projet sera ouvert, ouvrir un fichier `.go` pour activer l'extension
- si VS Code propose d'installer les outils Go manquants, accepter pour avoir l'aide à l'écriture du code

## ouvrir le projet dans VS Code

- récupérer le dossier du projet sur son ordinateur, si c'est un fichier `.zip` il faut d'abord l'extraire
- Le terminal doit aussi être placé dans le dossier du projet, celui où se trouve `go.mod`. Si ce n'est pas le cas, utiliser `cd` suivi du chemin du dossier entre guillemets :
```
cd "C:\chemin\vers\Go-reloaded"
```

## préparer le texte dans input.txt

Ouvrir `input.txt`, rentrer le texte qu'on souhaite modifier 

Pour `up`, `low` et `cap`, on précise après la virgule le nombre de mots précédents à modifier. Même pour un seul mot il faut mettre `1`, exemple : `(up,1)`.

On peut écrire `(up,3)` ou `(up, 3)`, les deux fonctionnent. C'est pareil pour `low` et `cap`.

Il faut mettre un nombre entier positif et avoir au moins autant de mots avant l'instruction, sinon le programme peut s'arrêter avec une erreur. Les formes sans nombre comme `(up)` ou `(up,)` ne sont pas traitées.

Exemple de texte à mettre dans `input.txt` :

```text
bonjour tout le monde (up,3) . BONNE JOURNEE (low,2) ! salut les amis (cap,3) . 100 (bin) et 100 (hex) . A exemple .
```

## lancer le projet

Dans le terminal de VS Code, rentrer cette commande puis appuyer sur Entrée :

```sh
go run . input.txt output.txt
```

- **`go run .`** : compile et lance le programme du dossier actuel
- **`input.txt`** : le premier fichier est celui qui contient le texte à modifier
- **`output.txt`** : le deuxième fichier est celui où le programme écrit le résultat

Il faut bien mettre les deux noms de fichiers et dans cet ordre.

Une fois la commande terminée, ouvrir `output.txt` dans VS Code pour voir le texte modifié. Le programme n'affiche pas le résultat dans le terminal.

Avec l'exemple au dessus, on obtient :

```text
bonjour TOUT LE MONDE. bonne journee! Salut Les Amis. 4 et 256. An exemple.
```

Le fichier `output.txt` est créé s'il n'existe pas, sinon son contenu est remplacé à chaque lancement.

Pour essayer un autre texte, modifier `input.txt`, enregistrer puis relancer la même commande. Les retours à la ligne et les espaces en trop sont regroupés par le programme.

## chaque modifications à son propre fichier, respectivement :

- **`bin.go`** : modifie le mot precedant "(bin)" et convertit sa valeur binaire en nombre décimal exemple : `"100 (bin)" --> "4"`

- **`cap.go`** : modifie les mots precedant "(cap,nombre)" et met en majuscule leur premiere lettre et le reste en minuscule exemple : `"bonjour les amis (cap,3)" --> "Bonjour Les Amis"`

- **`hex.go`** : modifie le mot precedant "(hex)" et convertit sa valeur hexadecimal en nombre décimal exemple : `"100 (hex)" --> "256"`

- **`low.go`** : modifie les mots precedant "(low,nombre)" et les rends entièrement en minuscule exemple : `"AU-REVOIR LES AMIS (low,3)" --> "au-revoir les amis"`

- **`occurenceA.go`** : sert pour l'anglais et transforme "A" en "An", ou "a" en "an", quand le prochain mot commence par une voyelle ou un "H" exemple : `"A exemple" --> "An exemple"`

- **`ponctuation.go`** : sert a coller les differentes ponctuation au mot precedant exemple : `"Salut ." --> "Salut."`

- **`up.go`** : modifie les mots precedant "(up,nombre)" et les rends entièrement en majuscule exemple : `"au-revoir les amis (up,3)" --> "AU-REVOIR LES AMIS"`

et pour finir le **`main.go`** sert de fichier executable et gère les erreurs de lecture et d'écriture des fichiers.

Il applique les modifications dans cet ordre : `bin`, `hex`, `up`, `low`, `cap`, les changements de `a` en `an`, puis la ponctuation.

## si le projet ne se lance pas

- **`go` n'est pas reconnu** : vérifier que Go est installé sur l'ordinateur, puis fermer et relancer VS Code et son terminal. Si le problème continue, vérifier que le dossier d'installation de Go contenant `go` est bien dans le `PATH`, la liste des dossiers où le terminal cherche les commandes
- **le terminal ne trouve pas `go.mod`** : vérifier qu'on est dans le dossier qui contient `main.go` et `go.mod`
- **le terminal affiche une erreur avec `input.txt`** : vérifier que le fichier existe, que son nom est correct et qu'il se trouve dans le dossier du projet
- **le terminal affiche `index out of range`** : vérifier qu'on a mis les deux noms de fichiers dans la commande et qu'on ne demande pas plus de mots qu'il y en a avant une instruction
- **le texte ne change pas** : enregistrer `input.txt` avant de relancer, vérifier les instructions et regarder le résultat dans `output.txt`
