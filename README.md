    /my_project/
    |-- /cmd/
    |   |-- /server/
    |   |   |-- main.go      # Point d'entrée pour démarrer le serveur
    |
    |-- /pkg/
    |   |-- /bird/
    |   |   |-- lexer.go       # Analyse lexicale pour Bird
    |   |   |-- parser.go      # Analyse syntaxique pour Bird
    |   |   |-- interpreter.go # Logique d'exécution pour Bird
    |   |
    |   |-- /server/
    |       |-- handler.go   # Gestionnaire des requêtes HTTP
    |
    |-- /internal/
    |   |-- /config/
    |       |-- config.go    # Gestion de la configuration du serveur
    |
    |-- /scripts/
    |   |-- setup.sh         # Script pour la mise en place ou le déploiement
    |
    |-- /static/
    |   |-- /css/
    |   |-- /js/
    |   |-- /images/
    |
    |-- /templates/
    |   |-- index.html       # Template HTML pour le rendu côté serveur
    |
    |-- /logs/
    |   |-- server.log       # Journaux pour le serveur
    |
    |-- go.mod               # Gestion des dépendances de Go
    |-- go.sum               # Somme de contrôle des dépendances
    |-- README.md            # Documentation du projet
