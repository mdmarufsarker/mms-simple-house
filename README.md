### Repo Structure
---
- Create Some Directory's

        mkdir templates public public/gallary public/icons scss scss/abstracts scss/base scss/components scss/layout scss/pages

- Create Some File's

        touch .gitignore LICENSE main.go README.md scss/style.scss scss/abstracts/_functions.scss scss/abstracts/_mixins.scss scss/abstracts/_variables.scss scss/base/_base.scss scss/base/_utilities.scss scss/layout/_header.scss scss/layout/_footer.scss scss/pages/_index.scss

- Add those lines in the style.scss file

        @import "abstracts/functions";
        @import "abstracts/mixins";
        @import "abstracts/variables";
        @import "base/base";
        @import "base/utilities";
        @import "layout/header";
        @import "layout/footer";
        @import "pages/home";

- Your Style color variable should like this in the scss/_variables.scss file

        $color-white: #fff;

- Add those lines in the scss/_base.scss file

        *,
        *::before,
        *::after{
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        html{
            font-size: 62.5%;
        }
        
- Add those lines in the scss/_utilities.scss file

        a{
            text-decoration: none;
        }

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