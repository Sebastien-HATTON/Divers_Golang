## Hashe
Pour hasher les mots de passe en utilisant bcrypt, on peut utiliser la fonction `bcrypt.GenerateFromPassword`.

Ci-dessous, un exemple de la façon dont on peut le faire pour obtenir des mots de passe hashés à partir de mots de passe en clair, puis les afficher.

```go
package main

import (
	"fmt"
	"log"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashed)
}

func main() {
	password1 := "password1"
	password2 := "password2"

	hashedPassword1 := hashPassword(password1)
	hashedPassword2 := hashPassword(password2)

	fmt.Println("Hash pour password1:", hashedPassword1)
	fmt.Println("Hash pour password2:", hashedPassword2)
}
```

Exécutez ce programme, il générera et affichera les versions hashées de `password1` et `password2`.

bcrypt produit un hash différent à chaque fois, même pour le même mot de passe en entrée, en raison du sel intégré. Cependant, la fonction `bcrypt.CompareHashAndPassword` sait comment gérer cela, donc tant qu'on utilise la même bibliothèque pour hasher et vérifier, tout fonctionnera comme prévu.

## Stocker les hashes
La décision de stocker les hashes de mots de passe sous forme de tableau de bytes (`[]byte`) ou de chaîne de caractères (`string`) dépend de plusieurs facteurs, notamment la manière dont le système gère et stocke les données. Voici quelques considérations pour vous aider à choisir entre les deux :

### Stockage en tant que `[]byte` :

**Avantages :**

- **Sécurité :** En mémoire, les slices de bytes peuvent être effacées après utilisation, ce qui peut potentiellement réduire la fenêtre pendant laquelle un mot de passe hashé pourrait être extrait par un processus malveillant.
- **Efficacité :** Les opérations cryptographiques travaillent naturellement avec des bytes, donc il n'y a pas besoin de convertir entre les types avant de faire des opérations de hashage ou de vérification.

**Inconvénients :**

- **Non human-readable :** Les données ne sont pas lisibles par un humain, ce qui peut rendre le débogage plus difficile.
- **Stockage :** Certains systèmes de stockage ou bases de données pourraient ne pas gérer aussi bien les bytes bruts que les chaînes de caractères, ce qui peut nécessiter une conversion ou un encodage (comme Base64) avant le stockage.

### Stockage en tant que `string` :

**Avantages :**

- **Compatibilité :** Les chaînes de caractères sont un type de données universellement géré par presque toutes les bases de données et systèmes de stockage.
- **Facilité de manipulation :** Les chaînes de caractères sont plus faciles à manipuler et à transmettre entre différentes parties d'une application ou entre services.
- **Human-readable :** Un hash sous forme de chaîne est plus facile à lire et à copier manuellement si nécessaire.

**Inconvénients :**

- **Sécurité :** Les chaînes de caractères sont immuables en Go, ce qui signifie qu'une fois qu'un hash est créé sous forme de chaîne, il ne peut pas être effacé de la mémoire jusqu'à ce que le ramasse-miettes (garbage collector) l'exécute, potentiellement laissant le hash sensible en mémoire plus longtemps.
- **Efficacité :** Les chaînes en Go sont des slices d'octets immuables avec un surcoût en mémoire pour le descripteur de chaîne. Pour le hashage et la vérification, ils doivent être convertis en slices d'octets, ce qui ajoute un petit surcoût de traitement.

### Conclusion :

En général, si vous travaillez directement avec des bibliothèques cryptographiques ou si vous avez besoin d'effacer le hash de la mémoire dès que possible pour des raisons de sécurité, le stockage en tant que `[]byte` pourrait être préférable. En revanche, si votre système gère mieux les chaînes de caractères et que vous avez besoin de les stocker ou de les transmettre facilement, il peut être plus pratique de les stocker sous forme de `string`.

La plupart des systèmes modernes de gestion de bases de données gèrent bien les chaînes de caractères et leur manipulation est plus simple dans de nombreux contextes. Assurez-vous simplement de ne pas compromettre la sécurité de vos utilisateurs quelle que soit l'approche choisie.

### Exemple :

Hash pour password1: $2a$10$7gYpjUd9PS6W5Gkv85LICuDDvw2jsVGh1oac6/mnZkPKM/xtjFM/y

Hash pour password2: $2a$10$pTtv2AQKgOp6KgeYASJEje1wFTyYfhog2Q0CU9bPhWvgzRjW/VmOy

Hash []byte pour password1: [36 50 97 36 49 48 36 102 46 112 99 111 102 107 101 47 76 56 74 56 112 66 115 70 115 46 115 69 101 51 80 85 78 47 67 55 50 101 85 106 80 108 114 68 80 49 49 112 109 70 74 84 117 70 115 98 116 77 74 87]

Hash []byte pour password2: [36 50 97 36 49 48 36 74 108 100 120 56 84 117 103 88 83 98 110 101 73 117 85 108 66 76 83 51 101 53 115 116 72 89 81 66 87 119 86 119 79 89 85 101 49 122 80 101 112 89 107 116 84 75 55 109 104 116 113 101]

## Base de donnée qui gére très bien les bytes bruts

Les bases de données modernes sont généralement assez flexibles pour gérer différents types de données, y compris les bytes bruts. Voici une liste de systèmes de gestion de base de données (SGBD) qui peuvent stocker efficacement des données binaires, comme les bytes bruts :

### Bases de données relationnelles:

- **PostgreSQL**: Utilise le type de donnée `bytea` pour stocker des données binaires.
- **MySQL / MariaDB**: Propose le type `BLOB` (Binary Large Object) pour stocker des données binaires de grande taille.
- **Oracle Database**: Offre également un type `BLOB` pour les données binaires.
- **Microsoft SQL Server**: Utilise les types `binary` et `varbinary` pour stocker des données binaires fixes ou variables en longueur.
- **SQLite**: Supporte le type `BLOB` pour stocker n'importe quelle donnée directement dans son format binaire.

### Bases de données NoSQL:

- **MongoDB**: Permet de stocker des données binaires avec le type `BinData`.
- **Cassandra**: Offre le type `blob` pour stocker des données en format binaire.
- **Riak**: Permet le stockage de données binaires à travers ses buckets en tant que valeurs sans traiter celles-ci comme des chaînes de caractères.
- **Couchbase**: Possède des types de documents binaires qui permettent de stocker et de gérer des données binaires.

### Bases de données de séries temporelles:

- **InfluxDB**: Peut stocker des données binaires dans des champs spécifiques, bien qu'elle soit optimisée pour les métriques sous forme numérique.

### Bases de données orientées objet:

- **ObjectDB**: Une base de données orientée objet qui peut stocker des objets Java ou .NET, y compris des tableaux de bytes.

### Bases de données en mémoire:

- **Redis**: Bien qu'étant une base de données clé-valeur en mémoire, elle peut stocker des chaînes de caractères et des valeurs binaires de manière efficace.
- **Memcached**: Principalement utilisé pour la mise en cache d'objets en mémoire, il peut stocker des données binaires.

### Bases de données de graphes:

- **Neo4j**: Peut stocker des données binaires en tant que propriétés des nœuds ou des relations.

Il est important de noter que bien que ces bases de données puissent gérer les données binaires, la manière dont vous interagissez avec ces données peut varier. Lorsque vous travaillez avec des données binaires, il est essentiel de prendre en compte les questions de performance, d'encodage (par exemple, Base64 si nécessaire), de manipulation et de récupération des données. Assurez-vous également que le pilote ou l'ORM que vous utilisez pour interagir avec la base de données prend correctement en charge les données binaires.
