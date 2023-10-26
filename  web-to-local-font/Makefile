.DEFAULT_GOAL := help

.PHONY: help
help:
	clear
	@echo " Usage:"
	@echo "		make <commande>"
	@echo "       "
	@echo " Avec commandes:"
	@echo "       "
	@echo "		init 	Initialisation du projet Go"
	@echo "		get 	Récupération des dépendances"
	@echo "		dev 	lance l'application sans compilation"
	@echo "       "
	@echo "		build 	Création du binaire"
	@echo "		slim 	Création d'un binaire plus léger"
	@echo "		run 	Lance l'application via le binaire"
	@echo "       "

.PHONY: init
init:
	@echo "Initialisation du projet Go"
	@echo "-----------------------------------------------------"	
	go mod init main

.PHONY: get
get:
	@echo "Récupération des dépendances"
	@echo "-----------------------------------------------------"
	go get

.PHONY: dev
dev:
	@echo "Lancement en mode dev sans compilation"
	@echo "-----------------------------------------------------"
	go run .

.PHONY: build
build: 
	@echo "Création du binaire"
	@echo "-----------------------------------------------------"
	CGO_ENABLED=0 go build -o main

.PHONY: slim
slim:
	@echo "Création du binaire taille reduite"
	@echo "-----------------------------------------------------"
	go build -ldflags="-s -w" -o main2

.PHONY: run
run:
	./main
