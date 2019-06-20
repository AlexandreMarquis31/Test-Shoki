# Test-Shoki

Pour installer et lancer le Front, il faut télécharger le dossier front, se rendre à la racine de ce dernier et taper les commandes suivantes :
- yarn install
- next build
- next start

Pour installer et lancer le Back, il faut avoir d'abord setup $GOPATH à la racine du dossier back, faire un go get pour installer les modules httpRouter et Colly (go get github.com/julienschmidt/httprouter et go get github.com/gocolly/colly), puis  run main.go (avec go run src/main.go).
