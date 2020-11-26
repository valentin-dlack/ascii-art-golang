# ACSII ART --- GOLANG

## Sommaire :

- [Groupe de développement](#groupe-de-développement)
- [Mettre en place le programme](#mettre-en-place-le-programme)
- [Cas d'utilisation](#utilisation-du-programme)
- [Licence](#license)

## Groupe de développement

Ce projet a été réalisé par plusieurs personnes :

- Costa Reype
- Arthur Abadie
- Valentin Dautrement

Nous nous sommes donc réparti les tâches selon nos domaines de compétence dans les fonctionnalités à développer.

## Mettre en place le programme

Premierement il faut s'assurer que la dernière version de Golang est installée sur votre ordinateur. Pour ça faites la commande suivante :

![Command : go version](https://i.imgur.com/6efPnSg.png)

Si une erreur apparait vous devez installer correctement golang dans sa version la plus récente sur votre machine.  
Vous pouvez le télécharger ici [-- Cliquez Ici --](https://golang.org/dl/)

Ensuite, il faut cloner le repository sur votre machine en local, pour faites la commande suivante : 

![Command : git clone (lien du repo)](https://i.imgur.com/hA549iR.png)

*NB : Le lien du repo est à récuperer au niveau du bouton "code" en haut de la page*

Vérifiez que tout les fichiers on bien été transférés sur votre machine. Il doit y avoir les fichiers suivants *(Sauf le .git, il est surement caché)* :

![Fichiers présents : ascii-art.go; ascii-table-template.txt ; License ; shadow.txt; standart.txt; thinkertoy.txt](https://i.imgur.com/QGNqsdG.png)

Quand tout les fichiers sont présent, vous pouvez build le programme avec la commande suivante, puis vérifer que vous avez bien le fichier `ascii-art` :

![Command : go build ascii-art.go](https://i.imgur.com/NVrnef2.png)

Tout est pret ! Vous pouvez enfin utiliser le programme !

## Utilisation du programme  

*Le `|` est une sépration entre plusieurs choix. Attention à bien respecter la casse des arguments. Dans les exemples si dessous, nous prenons en compte le fait qu'il y'a `./ascii-art` avant les argument.*

Le programme possède plusieurs arguments :

| Arguments (NB : il faut mettre ./ascii-art avant les arguments.) | Utilisation/Description      |
| ---------| -------------------------- |
| `--help` | Affiche l'aide de la commande|
| `"text"--color=couleur`  | Permet de coloriser le texte. N'est disponible qu'avec ces couleurs : `Red, Green, Yellow, Blue, Purple, Cyan, White`. Attention a respecter la casse ! |
| `"text" template`   | Permet de changer le résultat de l'ascii art selon les templates disponibles. Templates disponibles : `standart|shadow|thinkertoy`|
| `"text" --align=alignement`   | Permet de changer l'alignement du texte dans le terminal. Alignements disponibles : `right|left|center`|
| `"text" --output=filename.txt` | Permet d'envoyer le résultat du programme dans un fichier texte. `filename` est le nom du fichier ou sera stocké le résultat de la commande |

---
Les arguments sont stackables, vous pouvez en utiliser plusieurs dans une commande. Exemples si dessous :

Commande : `./ascii-art "Hello" --color=Red "World" --color=Blue --align=center`

- Résultat :

![](https://i.imgur.com/QlYX38X.png)

Commande : `./ascii-art "Hello there !" shadow --color=Cyan --align=right --output=banner.txt`

- Résultat dans le terminal :

![](https://i.imgur.com/LGqzT5f.png)

- Résultat dans le fichier `banner.txt` :

![](https://i.imgur.com/6a4bqQ7.png)

Dernier exemple :

Commande : `./ascii-art "Thanks" standart --color=Purple --align=right "For" thinkertoy --color=Green --align=center "Watching !" shadow --color=Yellow --align=left`

- Résultat : 

![](https://i.imgur.com/ArvJwki.png)

## License

This program is under [MIT License](https://opensource.org/licenses/MIT)  
You are free to use and reuse everything of this code.
