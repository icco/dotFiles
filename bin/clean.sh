#! /bin/zsh

set -x

rm -rf ~/Work/*/node_modules
rm -rf ~/Work/**/.terraform

rm -rf ~/Projects/*/node_modules
rm -rf ~/Projects/**/.terraform

sudo rm -rf $GOPATH/pkg/

docker system prune -af
