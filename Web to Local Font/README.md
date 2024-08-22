
# Permet de télécharger des fichiers de polices en parallèle en utilisant des goroutines.

## Utilisation avec chemin de destination:

go run main.go -u "https://fonts.googleapis.com/css2?family=Roboto&display=swap" -f "./font"

## Utilisation sans chemin de destination:

go run main.go -u "https://fonts.googleapis.com/css2?family=Roboto&display=swap"

1. Après avoir extrait toutes les URLs des fichiers de polices du CSS, il crée une goroutine pour chaque URL pour télécharger le fichier correspondant.

2. Chaque goroutine appelle la fonction `downloadFile`, qui télécharge un fichier depuis une URL et le sauvegarde dans un dossier local.

3. Un `WaitGroup` est utilisé pour attendre que toutes les goroutines soient terminées avant de fermer le canal qui est utilisé pour envoyer des messages d'erreur depuis les goroutines.


## Voici un exemple de fonctionnement:

Supposons que votre fichier CSS contienne les URLs de trois fichiers de polices :

```css
@font-face {
  font-family: "FontA";
  src: url("http://example.com/fonts/fonta.ttf");
}
@font-face {
  font-family: "FontB";
  src: url("http://example.com/fonts/fontb.ttf");
}
@font-face {
  font-family: "FontC";
  src: url("http://example.com/fonts/fontc.ttf");
}
```

Lorsque vous exécutez le programme avec l'URL de ce fichier CSS, il démarrera trois goroutines, une pour chaque URL de fichier de police. Chaque goroutine téléchargera un fichier de police en parallèle des autres.

Pour exécuter le programme, vous pouvez utiliser la commande suivante (remplacez `http://example.com/path/to/your.css` par l'URL de votre fichier CSS) :

```bash
go run main.go -u http://example.com/path/to/your.css
```

Le programme téléchargera chaque fichier de police dans le dossier spécifié par le flag `-f` (par défaut, il s'agit du dossier `fonts`) et les messages d'erreur seront affichés sur la console si un téléchargement échoue.
