#!/bin/bash

buildApp=""
toPush=false
forced=false

function fail() {
  gum style --foreground 196 "Error: $1"
  exit 1
}

function print_usage() {
  printf "Usage: docker_build.sh [OPTIONS]\n"
}

while getopts 'abfp' flag; do
  case "${flag}" in
  a)
    buildApp="build-all"
    ;;
  b)
    if [ "$buildApp" != "" ]; then
      print_usage
      exit 1
    fi
    shift
    buildApp="$1"
    ;;
  f)
    forced=true
    ;;
  p)
    toPush=true
    ;;
  *)
    print_usage
    exit 1
    ;;
  esac
done

function handle_subdirectory() {
  hash=$(git rev-parse --short HEAD)
  DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ')

  echo "$hash"
  echo "$DATE"

  # Build binary for the subdirectory
  gum style --foreground 10 "Building binary for $1"
  cd ./"$1" || fail "Unable to cd to cmd/$1"

  GOOS=linux GOARCH=amd64 go build -o bin/app -ldflags "-X main.Commit=$hash -X main.Date=$DATE" . || fail "Unable to build binary for $1"

  # Move app to bin directory from the root of the repo. Create the bin directory if it doesn't exist
  mkdir -p ../../bin || fail "Unable to create bin directory"
  mv bin/app ../../bin/app || fail "Unable to move binary to bin/app"

  rm -rf bin

  cd ../../ || fail "Unable to cd to root of repo"

  # Remove the "./cmd/" prefix from the directory name
  appName=${1#"./cmd/"}
  appName=${appName#"cmd/"}

  # Build the docker image
  gum style --foreground 10 "Building docker image for $appName"
  registry="ghcr.io/jacobbrewer1/league-manager-$appName"

  if [ "$forced" == true ]; then
    gum style --foreground 10 "Forcing build"
    docker build -t "$registry:$hash" -t "$registry:latest" . || fail "Unable to build docker image for $appName"
  else
    gum spin --spinner dot --title "Building image" -- docker build -t "$registry:$hash" -t "$registry:latest" . || fail "Unable to build docker image for $appName"
  fi

  if [ "$forced" == false ]; then
    gum style --foreground 10 "Would you like to push the docker image to the GitHub Container Registry?"
    gum confirm && toPush=true
  fi

  # Check if the toPush variable is set to true. If yes, push the docker image to the GitHub Container Registry
  if [ "$toPush" != true ]; then
    gum style --foreground 10 "Docker image built for $appName"

    # Cleanup the binary
    rm -rf bin
    rm -rf ./cmd/"$appName"/bin

    return 0
  fi

  # Push the docker image to the GitHub Container Registry
  gum style --foreground 10 "Pushing docker image to GitHub Container Registry"

  docker push "$registry:$hash" || fail "Unable to push docker image to GitHub Container Registry"
  gum style --foreground 10 "Docker image pushed to $registry:$hash"

  docker push "$registry:latest" || fail "Unable to push docker image to GitHub Container Registry"
  gum style --foreground 10 "Docker image pushed to $registry:latest"

  gum style --foreground 10 "Docker image pushed to $registry:$hash"

  # Cleanup the binary
  rm -rf bin
  rm -rf ./cmd/"$appName"/bin
}

# If the toPush variable is not set, set it to false
if [ -z "$toPush" ]; then
  toPush="false"
fi

subdirectories=$(ls -d ./cmd/*)

if [ "$buildApp" == "build-all" ]; then
  gum style --foreground 196 "Building all apps"
elif [ "$buildApp" == "" ]; then
  apps=$(find cmd -mindepth 1 -type d -exec basename {} \; | sed 's/^cmd\///g')
  buildApp=$(echo "$apps" | gum choose --no-limit --header "Select the app(s) to generate")
  if [ -z "$buildApp" ]; then
    gum style --foreground 196 "No app selected"
    exit 0
  fi
  subdirectories="cmd/$buildApp"
else
  gum style --foreground 196 "Building $buildApp"
  handle_subdirectory "cmd/$buildApp"
  exit 0
fi

# For each subdirectory of the cmd directory and run the subdirectory function
for dir in $subdirectories; do
  if [ -d "$dir" ]; then
    gum style --foreground 10 "Building $dir"
    handle_subdirectory "$dir"
  fi
done
