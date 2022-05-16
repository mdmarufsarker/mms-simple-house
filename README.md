### Repo Structure
---
- Create Some Directory's

        mkdir templates public public/gallary public/icons scss scss/abstracts scss/base scss/components scss/layout scss/pages

- Create Some File's

        touch .gitignore LICENSE main.go README.md scss/style.scss

- Now Initialize go

        go mod init github.com/mdmarufsarker/project-name

- Now install mux for handling your route

        go get github.com/gorilla/mux

- Now Initialize NPM

        npm init

- Now install node-sass package

        npm i node-sass

- Now add those line in the scripts

        "compile:scss": "node-sass ./scss/style.scss -o ./public/ -w"

---
### Now you are ready to go
---

- Want to run go server?

        go run main.go

- Want to compile your scss code to vanilla css code

        npm run compile:scss