# Go-reloaded

Le but du projet est de crée un programme qui modifie un text rentré dans le fichier `input.txt` pour le ressortir dans `output.txt`, pour cela...

## chaque modifications à son propre fichier, respectivement : 

- **`bin.go`** : modifie le mot precedant "(bin)" et le convertit en sa valeur binaire exemple : `"100" --> "4"`

- **`cap.go`** : modifie le mot precedant "(cap)" et met en majuscule sa premiere lettre exemple : `"bonjour" --> "Bonjour"`

- **`hex.go`** : modifie le mot precedant "(hex)" et le convertit en sa valeur hexadecimal exemple : `"100" --> "256"`

- **`low.go`** : modifie le mot precedant "(low)" et le rends entièrement en minuscule exemple : `"AU-REVOIR" --> "au-revoir"`

- **`occurenceA.go`** : sert pour l'anglais et rajoute une occurence de "A" quand le prochain mot commence par une voyelle ou un "H" exemple : `"A exemple" --> "An exemple"`

- **`ponctuation.go`** : sert a coller les differentes ponctuation au mot precedant exemple : `"Salut ." --> "Salut."`

- **`up.go`** : modifie le mot precedant "(up)" et le rends entièrement en majuscule exemple : `"au-revoir" --> "AU-REVOIR"`

---

et pour finir le **`main.go`** sert de fichier executable et gère les erreurs
