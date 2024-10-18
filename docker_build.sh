#!/bin/bash

buildApp=$1
toPush=$2

function fail() {
  gum style --foreground 196 "Error: $1"
  exit 1
}

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

  gum spin --spinner dot --title "Building image" --show-output -- docker build -t "$registry:$hash" -t "$registry:latest" . || fail "Unable to build docker image for $appName"

  # Check if the toPush variable is set to true. If yes, push the docker image to the GitHub Container Registry
  if [ "$toPush" != "true" ]; then
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

if [ "$buildApp" == "build-all" ]; then
  gum style --foreground 196 "Building all apps"
elif [ -z "$buildApp" ]; then
    apps=$(find cmd -mindepth 1 -type d -exec basename {} \; | sed 's/^cmd\///g')
    buildApp=$(echo "$apps" | gum choose --no-limit --header "Select the app(s) to generate")
    if [ -z "$buildApp" ]; then
      gum style --foreground 196 "No app selected"
      exit 0
    fi
else
  gum style --foreground 196 "Building $buildApp"
  handle_subdirectory "cmd/$buildApp"
  exit 0
fi

# Get all the subdirectories in the cmd directory
subdirectories=$(ls -d ./cmd/*)

# For each subdirectory of the cmd directory and run the subdirectory function
for dir in $subdirectories; do
  if [ -d "$dir" ]; then
    gum style --foreground 10 "Building $dir"
    handle_subdirectory "$dir"
  fi
done