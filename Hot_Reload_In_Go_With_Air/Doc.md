## Fichier de configuration pour [Air](https://github.com/cosmtrek/air) au format TOML

### Répertoire de travail

. ou chemin absolu, veuillez noter que les répertoires suivants doivent être sous la racine.

root = "."
tmp_dir = "tmp"

[build]

### Tableau de commandes à exécuter avant chaque construction

pre_cmd = ["echo 'hello air' > pre_cmd.txt"]

### Juste une vieille commande shell classique. Vous pourriez aussi utiliser `make`.

cmd = "go build -o ./tmp/main ."

### Tableau de commandes à exécuter après ^C

post_cmd = ["echo 'hello air' > post_cmd.txt"]

### Fichier binaire résultant de `cmd`.

bin = "tmp/main"

### Personnaliser le binaire, peut configurer les variables d'environnement lors de l'exécution de votre application.

full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"

### Surveillez ces extensions de nom de fichier.

include_ext = ["go", "tpl", "tmpl", "html"]

### Ignorez ces extensions de nom de fichier ou répertoires.

exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]

### Surveillez ces répertoires si vous les avez spécifiés.

include_dir = []

### Surveillez ces fichiers.

include_file = []

### Exclure les fichiers.

exclude_file = []

### Exclure des expressions régulières spécifiques.

exclude_regex = ["_test\\.go"]

### Exclure les fichiers inchangés.

exclude_unchanged = true

### Suivre le lien symbolique pour les répertoires

follow_symlink = true

### Ce fichier journal se trouve dans votre tmp_dir.

log = "air.log"

### Sondez les fichiers pour les changements au lieu d'utiliser fsnotify.

poll = false

### Intervalle de sondage (par défaut à l'intervalle minimum de 500 ms).

poll_interval = 500 # ms

### Il n'est pas nécessaire de déclencher la construction chaque fois que le fichier change si c'est trop fréquent.

delay = 0 # ms

### Arrêtez d'exécuter l'ancien binaire lorsque des erreurs de construction se produisent.

stop_on_error = true

### Envoyer le signal d'Interruption avant de tuer le processus (windows ne prend pas en charge cette fonctionnalité)

send_interrupt = false

### Délai après l'envoi du signal d'Interruption

kill_delay = 500 # nanoseconde

### Relancer ou non le binaire

rerun = false

### Délai après chaque exécution

rerun_delay = 500

### Ajouter des arguments supplémentaires lors de l'exécution du binaire (bin/full_bin). Exécutera './tmp/main hello world'.

args_bin = ["hello", "world"]

[log]

### Afficher l'heure du journal

time = false

### Afficher uniquement le journal principal (silencieux pour le surveillant, la construction, le coureur)

main_only = false

[color]

### Personnaliser la couleur de chaque partie. Si aucune couleur trouvée, utilisez le journal d'application brut.

main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]

### Supprimer le répertoire tmp à la sortie

clean_on_exit = true

[screen]
clear_on_rebuild = true
keep_scroll = true
