#!/bin/bash
set -eo pipefail

SCRIPT_NAME=$(basename "$0")
SCRIPT_PATH=$(dirname "$(realpath "$0")")

IMAGE_BASE_TAG="hobbyfarm"
PLATFORM="linux/amd64"
SERVICE_BASE_PATH="$SCRIPT_PATH/v3/services"
SERVICE_DOCKERFILE="$SCRIPT_PATH/v3/Dockerfile"
GARGANTUA_DOCKERFILE="$SCRIPT_PATH/Dockerfile"

function print_help() {
  echo "Usage:"
  echo "  $SCRIPT_NAME [options] [service_name...]"
  echo ""
  echo "Examples:"
  echo "$0                    - builds all images"
  echo "$0 costsvc gargantua  - builds costsvc and gargantua image (names of services match services in $SERVICE_BASE_PATH)"
  echo "$0 -t my-hobbyfarm    - builds all images with image base tag my-hobbyfarm"
  echo "$0 -p linux/arm64     - builds all images with for linux/arm64"
  echo ""
  echo "Options:"
  echo "  -t, --tag IMAGE_BASE_TAG          define the image base tag (default is $IMAGE_BASE_TAG)"
  echo "  -p, --platform OS/ARCH[/VARIANT]  define the OS/ARCH[/VARIANT] of the image (default is $PLATFORM)"
  echo "  -h, --help                        print help"
}

SERVICES=()

while [[ $# -gt 0 ]]; do
  case $1 in
    -t|--tag)
      IMAGE_BASE_TAG="$2"
      shift # past argument
      shift # past value
      ;;
    -p|--platform)
      PLATFORM="$2"
      shift # past argument
      shift # past value
      ;;
    -h|--help)
      print_help
      exit 0
      ;;
    -*)
      echo "Unknown option $1"
      exit 1
      ;;
    *)
      SERVICES+=("$1") # save positional arg
      shift # past argument
      ;;
  esac
done

set -- "${SERVICES[@]}" # restore positional parameters

DOCKER_COMMAND="docker"

# check if podman should be used
if ! command -v "$DOCKER_COMMAND" > /dev/null 2>&1
then
    DOCKER_COMMAND="podman"
fi
# check if podman exists
if ! command -v "$DOCKER_COMMAND" > /dev/null 2>&1
then
    >&2 echo "neither \"docker\" nor \"podman\" exists on the system"
    exit 1
fi

cd "$SCRIPT_PATH"

# no services specified, populate SERVICES array
if (( ${#SERVICES[@]} == 0 )); then
  SERVICES=("gargantua")
  while IFS='' read -r service_name; do SERVICES+=("$service_name"); done < <(find "$SERVICE_BASE_PATH"/* -maxdepth 0 -type d -exec basename {} \;)
fi

# check if all services exist
for service_name in "${SERVICES[@]}"; do
  if [ "$service_name" != "gargantua" ]; then
    service_path="$SERVICE_BASE_PATH/$service_name"
    if [ ! -d "$service_path" ]; then
      echo "service \"$service_name\" does not exist"
      exit 1
    fi
  fi
done

# build specified services
for service_name in "${SERVICES[@]}"; do
  echo "building \"$service_name\""
  echo ""

  if [ "$service_name" = "gargantua" ]; then
    $DOCKER_COMMAND build --platform "$PLATFORM" --file "$GARGANTUA_DOCKERFILE" --tag "${IMAGE_BASE_TAG}/${service_name}:latest" .
    echo ""
  else
    $DOCKER_COMMAND build --platform "$PLATFORM" --build-arg SERVICE_NAME="$service_name" --file "$SERVICE_DOCKERFILE" --tag "${IMAGE_BASE_TAG}/${service_name}:latest" .
    echo ""
  fi
done