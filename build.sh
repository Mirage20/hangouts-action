GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.webhookUrl=${GOOGLE_HANGOUTS_WEBHOOK_URL}" -o hangouts-action ./
docker build -t ${DOCKER_REPO}/hangouts-action:latest .
docker push ${DOCKER_REPO}/hangouts-action:latest
