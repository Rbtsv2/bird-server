# Bird Language Server

Le serveur Bird est une plateforme légère pour interpréter et exécuter des scripts écrits dans le langage de programmation Bird, avec une intégration Web pour exécuter des fonctions personnalisées et servir des fichiers HTML.

## Fonctionnalités

- Interprétation du langage Bird, un langage simple pour des scripts rapides et des tâches web.
- Serveur HTTP intégré pour répondre aux requêtes Web avec le contenu dynamique généré par Bird.
- Support des fonctions personnalisées telles que `log()` pour le débogage et l'affichage de contenu.

## Installation

Suivez ces étapes pour configurer le serveur Bird sur votre système.

### Prérequis

- Go 1.15 ou plus récent.
- Accès à un terminal sur systèmes UNIX ou Windows.

### Instructions

1. **1er méthode : Cloner le dépôt Git Bird Server et obtenir un exécutable binaire nommé bird-server:**

```bash

mkdir -p /var/www/bird-server && cd /var/www/bird-server
git clone https://exemple.com/repo/bird-server.git .
cd cmd/bird-server/server
go build -o bird-server
chmod +x bird-server
./bird-server

```

2. **2eme méthode :Cloner le dépôt Git Bird Server et run bird-server:**

```bash

mkdir -p /var/www/bird-server && cd /var/www/bird-server
git clone https://exemple.com/repo/bird-server.git .
cd cmd/bird-server/server
go run server.go

```

Le serveur démarre sur le port 8080 par défaut. Visitez `http://localhost:8080` pour accéder au serveur.

3. **Cloner le dépôt Git Bird Template ( votre site web ):**

```bash

mkdir -p /var/www/bird && cd /var/www/bird
git clone https://exemple.com/repo/bird-template.git .

```

## Utilisation

Les scripts Bird sont placés dans le dossier `/var/www/server/birdfiles` avec l'extension `.bird`. Voici un exemple de script Bird :

 ```
 route("/")
    return "index.html", "PageTitle=Welcome to Bird Language"
 ```

Ce script associe la racine du serveur à `index.html`, en passant des paramètres à la template.

### Ajouter des Routes

Ajoutez de nouvelles routes en créant des fichiers `.bird` dans le dossier des scripts. Chaque fichier peut contenir plusieurs directives de route utilisant la syntaxe montrée ci-dessus.


## Contribution

Les contributions sont les bienvenues. Pour contribuer, veuillez forker le dépôt, créer une branche pour votre fonctionnalité, et soumettre une pull request.

## Licence

Ce projet est distribué sous la licence MIT. Voir le fichier `LICENSE` pour plus de détails.
