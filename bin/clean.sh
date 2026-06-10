#! /bin/zsh

set -x

rm -rf ~/Work/*/node_modules
rm -rf ~/Work/**/node_modules
rm -rf ~/Work/**/.terraform

rm -rf ~/Projects/*/node_modules
rm -rf ~/Projects/**/.terraform

rm -rf ~/Projects/**/.next
rm -rf ~/Work/**/.next

yarn cache clean
npm cache clean --force
pnpm cache delete

sudo rm -rf $GOPATH/pkg/

docker system prune -af
