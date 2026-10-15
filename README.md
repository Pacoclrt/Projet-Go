<div align="center">

# ✍️ go-reloaded

**Un outil de correction et de mise en forme de texte, écrit en Go.**

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![Tests](https://img.shields.io/badge/tests-passing-brightgreen)
![Dépendances](https://img.shields.io/badge/d%C3%A9pendances-aucune-blue)

*Projet réalisé dans le cadre de la formation Ynov.*

</div>

---

## 📖 Présentation

`go-reloaded` lit un fichier texte, applique une série de transformations (conversions numériques, changements de casse, ponctuation, guillemets, articles) puis écrit le résultat corrigé dans un nouveau fichier.

```text
Entrée  : it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom
Sortie  : It was the best of times, it was the worst of TIMES, it was the age of wisdom
```

Le projet n'utilise **que la bibliothèque standard** de Go.

## 🚀 Démarrage rapide

### Prérequis

- [Go](https://go.dev/dl/) 1.22 ou plus récent

### Utilisation

```bash
go run . sample.txt result.txt
```

| Argument     | Description                              |
| ------------ | ---------------------------------------- |
| `sample.txt` | Fichier texte à corriger (lecture)       |
| `result.txt` | Fichier dans lequel écrire le résultat   |

Sans les deux arguments, le programme affiche simplement :

```text
Utilisation : go run . sample.txt result.txt
```

### Compiler un exécutable

```bash
go build -o go-reloaded .
```

```bash
./go-reloaded sample.txt result.txt
```

## 🛠️ Fonctionnalités

### 🔢 Conversions numériques

Le marqueur convertit **le mot précédent** en nombre décimal.

| Marqueur | Effet                         | Exemple                  | Résultat     |
| -------- | ----------------------------- | ------------------------ | ------------ |
| `(hex)`  | Hexadécimal → décimal         | `1E (hex) files`         | `30 files`   |
| `(bin)`  | Binaire → décimal             | `10 (bin) years`         | `2 years`    |

> Si le mot n'est pas un nombre valide dans la base donnée (ex. `zz (hex)`), il est laissé tel quel et le marqueur est supprimé.

### 🔠 Changements de casse

| Marqueur     | Effet                                   | Exemple                          | Résultat                     |
| ------------ | --------------------------------------- | -------------------------------- | ---------------------------- |
| `(up)`       | Met le mot précédent en MAJUSCULES      | `go (up)`                        | `GO`                         |
| `(low)`      | Met le mot précédent en minuscules      | `SHOUTING (low)`                 | `shouting`                   |
| `(cap)`      | Met une majuscule au mot précédent      | `bridge (cap)`                   | `Bridge`                     |
| `(up, n)`    | Applique à **n** mots précédents        | `so exciting (up, 2)`            | `SO EXCITING`                |
| `(low, n)`   | Applique à **n** mots précédents        | `I AM HERE (low, 2)`             | `I am here`                  |
| `(cap, n)`   | Applique à **n** mots précédents        | `is so exciting (cap, 3)`        | `Is So Exciting`             |

> Si `n` dépasse le nombre de mots disponibles, la transformation s'applique simplement à tous les mots précédents.

### ❗ Ponctuation

Les signes `.` `,` `!` `?` `:` `;` sont collés au mot précédent et suivis d'un espace. Les groupes comme `...` ou `!?` restent ensemble.

| Entrée                                | Sortie                              |
| ------------------------------------- | ----------------------------------- |
| `over there ,and then BAMM !!`        | `over there, and then BAMM!!`       |
| `I was thinking ... You were right`   | `I was thinking... You were right`  |

### 💬 Apostrophes

Les paires de `'` se collent aux mots qu'elles entourent.

| Entrée                    | Sortie                |
| ------------------------- | --------------------- |
| `describe me: ' awesome '` | `describe me: 'awesome'` |
| `' I am the best '`       | `'I am the best'`     |

### 🅰️ Articles

`a` devient `an` (et `A` devient `An`) lorsque le mot suivant commence par une voyelle (`a`, `e`, `i`, `o`, `u`) ou un `h`.

| Entrée                    | Sortie                    |
| ------------------------- | ------------------------- |
| `A amazing rock`          | `An amazing rock`         |
| `a untold story`          | `an untold story`         |

## 🔄 Pipeline de traitement

Chaque ligne du fichier est traitée indépendamment, en passant par les étapes suivantes dans cet ordre :

```text
  ligne brute
       │
       ▼
  découpage en mots   (strings.Fields)
       │
       ▼
  applyTags          → (hex) (bin) (up) (low) (cap) (x, n)
       │
       ▼
  fixPunctuation     → . , ! ? : ;
       │
       ▼
  fixQuotes          → ' ... '
       │
       ▼
  fixArticles        → a → an
       │
       ▼
  ligne corrigée
```

## 📁 Structure du projet

```text
.
├── main.go           # Point d'entrée : lecture/écriture des fichiers, pipeline
├── tags.go           # Marqueurs (hex), (bin), (up), (low), (cap) et (x, n)
├── punctuation.go    # Rattachement de la ponctuation
├── quotes.go         # Gestion des apostrophes
├── articles.go       # Correction a → an
├── main_test.go      # Tests unitaires
└── go.mod            # Module go-reloaded
```

## 🧪 Tests

Lancer la suite de tests :

```bash
go test -v ./...
```

Les tests couvrent :

- ✅ les changements de casse (`TestCase`)
- ✅ les conversions numériques (`TestNumbers`)
- ✅ la ponctuation (`TestPunctuation`)
- ✅ les apostrophes (`TestQuotes`)
- ✅ les articles (`TestArticles`)
- ✅ les exemples complets du sujet (`TestExamples`)
- ✅ les cas limites : texte vide, marqueur en début de ligne, `n` trop grand, nombre invalide (`TestStrangeCases`)

## 💡 Exemple complet

**`sample.txt`**

```text
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
There it was. A amazing rock !
As he said: ' this is so exciting (up, 3) '
```

```bash
go run . sample.txt result.txt
```

**`result.txt`**

```text
It was the best of times, it was the worst of TIMES, it was the age of wisdom
Simply add 66 and 2 and you will see the result is 68.
There it was. An amazing rock!
As he said: 'this IS SO EXCITING'
```

## ⚠️ Limites connues

- Les marqueurs à nombre doivent contenir un espace après la virgule : `(up, 2)` fonctionne, `(up,2)` n'est pas reconnu.
- Les espaces multiples à l'intérieur d'une ligne sont réduits à un seul espace.
- Un marqueur placé en début de ligne (sans mot avant lui) est simplement supprimé.

---
