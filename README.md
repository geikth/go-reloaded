# Go-reloaded

## Présentation

Go-reloaded est un outil en ligne de commande qui complète, édite et corrige automatiquement un texte



## Fonctionnalités

- Convertir un nombre hexadécimal en décimal avec `(hex)`
- Convertir un nombre binaire en décimal avec `(bin)`
- Passer le mot précédent en majuscules, en minuscules ou avec une majuscule initiale avec `(up)`, `(low)` et `(cap)`
- Appliquer ces trois transformations à plusieurs mots avec `(up, n)`, `(low, n)` et `(cap, n)`
- Corriger la ponctuation (`. , ! ? : ;`), y compris les groupes comme `...` et `!?`
- Placer correctement les apostrophes autour d'un ou plusieurs mots
- Transformer `a` en `an` devant une voyelle ou un `h`

## Technologies

- Go

## Prérequis

- Go 1.20 ou plus récent (vérifier avec `go version`)

## Installation

```bash
git clone <url-du-depot>
cd go-reloaded
```

## Lancement

```bash
go run . <fichier_entrée> <fichier_sortie>
```

## Architecture

```
go-reloaded/
├── main.go
├── transform.go
├── casse.go
├── nombres.go
├── articles.go
├── ponctuation.go
├── apostrophes.go
├── transform_test.go
├── go.mod
└── README.md
```


Projet individuel 
