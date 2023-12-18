#go get github.com/pilu/fresh
sed -i -e "13 s/version: .*/version: $(date '+%Y.%m.%d.%H.%M')/" ./configs/development/app.yaml
sed -i -e "14 s/branch: .*/branch: $(git branch --show-current)[$(git rev-parse HEAD)]/" ./configs/development/app.yaml

rm ./configs/development/app.yaml-e
rm -rf ./logs/*
fresh runner.conf