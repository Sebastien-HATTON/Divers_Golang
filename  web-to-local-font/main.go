// Copyright 2023 Sébastien-HATTON. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package main

import (
	"flag"          // Pour parser les arguments de ligne de commande.
	"fmt"           // Pour afficher des messages sur la console.
	"io"            // Pour copier les données de l'URL vers un fichier.
	"net/http"      // Pour envoyer des requêtes HTTP.
	"os"            // Pour créer des fichiers et des dossiers.
	"path/filepath" // Pour obtenir le nom de base d'un chemin.
	"regexp"        // Pour utiliser des expressions régulières.
	"strings"       // Pour remplacer les chaînes de caractères.
	"sync"          // Pour synchroniser les goroutines.
)

// downloadFile télécharge un fichier depuis une URL et le sauvegarde localement.
func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Erreur lors du téléchargement du fichier depuis l'URL %s: %s\n", url, err)
	}
	// Assurez-vous de fermer la réponse à la fin.
	defer resp.Body.Close()

	// Crée un fichier local pour enregistrer les données téléchargées.
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Erreur lors de la création du fichier %s: %s\n", filepath, err)
	}
	defer out.Close()

	// Copie les données de l'URL vers le fichier local.
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'écriture dans le fichier %s: %s\n", filepath, err)
	}
	return nil
}

// extractFontUrls extrait les URL des polices du CSS.
func extractFontUrls(css string) []string {
	// Utilisation d'une expression régulière pour trouver les URL.
	re := regexp.MustCompile(`url\(['"]?(.*?)['"]?\)`)
	matches := re.FindAllStringSubmatch(css, -1)
	var urls []string
	for _, match := range matches {
		urls = append(urls, match[1])
	}
	return urls
}

// extractFilename extrait le nom du fichier depuis l'URL.
func extractFilename(url string) string {
	return filepath.Base(url)
}

// La fonction principale du programme.
func main() {
	// Parsing des arguments de ligne de commande.
	urlPtr := flag.String("u", "", "URL pour télécharger le fichier Fonts")
	folderPtr := flag.String("f", "fonts", "Dossier pour enregistrer les fichiers téléchargés")
	flag.Parse()

	// Vérifie si l'URL est fournie.
	if *urlPtr == "" {
		fmt.Println("Veuillez fournir une URL en utilisant le flag -u")
		return
	}

	// Crée un dossier si il n'existe pas.
	dir := fmt.Sprintf("%s/files", *folderPtr)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	fmt.Println("Téléchargement du fichier Fonts...")
	// Télécharge le fichier CSS.
	err := downloadFile(fmt.Sprintf("%s/fonts.css", *folderPtr), *urlPtr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Lis le fichier CSS téléchargé.
	data, err := os.ReadFile(fmt.Sprintf("%s/fonts.css", *folderPtr))
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier Fonts: %s\n", err)
		return
	}

	content := string(data)
	// Extrayez les URL des polices du fichier CSS.
	fontUrls := extractFontUrls(content)

	ch := make(chan string, len(fontUrls))
	var wg sync.WaitGroup

	fmt.Printf("%d fichiers de polices chargés...\n", len(fontUrls))
	for _, url := range fontUrls {
		wg.Add(1)
		// Extrayez le nom du fichier depuis l'URL.
		filename := extractFilename(url)
		// Chemin où le fichier de police sera sauvegardé localement.
		filepath := fmt.Sprintf("%s/files/%s", *folderPtr, filename)
		go func(u, p string) {
			defer wg.Done()
			// Télécharge le fichier de police.
			err := downloadFile(p, u)
			if err != nil {
				ch <- err.Error()
			}
		}(url, filepath)
	}

	// Attend que toutes les goroutines soient terminées avant de fermer le canal.
	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		if msg != "" {
			fmt.Print(msg)
		}
	}

	for _, url := range fontUrls {
		filename := extractFilename(url)
		// Remplace les URL dans le fichier CSS par le chemin local du fichier de police.
		content = strings.ReplaceAll(content, url, fmt.Sprintf("'files/%s'", filename))
	}

	// Écrit le fichier CSS mis à jour.
	err = os.WriteFile(fmt.Sprintf("%s/fonts.css", *folderPtr), []byte(content), 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture du fichier CSS mis à jour: %s\n", err)
	}

	fmt.Println("Terminé!")
}
