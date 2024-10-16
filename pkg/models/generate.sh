#!/usr/bin/env bash

if ! command -v goschema &> /dev/null; then
  gum style --foreground 196 "goschema is required to generate models. Please install it by running 'go get github.com/jbrewer/goschema'"
  exit 1
fi

if ! command -v gum &> /dev/null; then
  gum style --foreground 196 "gum is required to generate models. Please install it by running 'go get github.com/jbrewer/gum'"
  exit 1
fi

if ! command -v goimports &> /dev/null; then
  gum style --foreground 196 "goimports is required to generate models. Please install it by running 'go get golang.org/x/tools/cmd/goimports'"
  exit 1
fi

# Allow the user too select the model/s to generate
schemas=$(find ./schemas -type f -name "*.sql" | sed 's/\.\/schemas\///' | sed 's/\.sql//')

togen=$(echo "$schemas" | gum choose --no-limit --header "Select the model(s) to generate")

gum style "Are you sure you want to generate the following models?"

for model in $togen; do
  gum style --foreground 222 "  - $model"
done

gum confirm || exit 0

for model in $togen; do
  gum spin "$(goschema generate --templates=./templates/*tmpl --out=./ --sql=./schemas/"$model".sql)" --spinner dot --title "Generating model $model"
  go fmt ./"$model".go
  goimports -w ./"$model".go
done
