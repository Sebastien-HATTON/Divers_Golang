## Notes
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
