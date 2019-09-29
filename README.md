# Hangouts Chat Action

Send Pull request notifications to your Google Hangouts Chat room

## Usage

### Build and Push with your Hangouts Chats Webhook Url

```bash
export GOOGLE_HANGOUTS_WEBHOOK_URL=<webhook-url>
DOCKER_REPO=<your-repo> ./build.sh
```

### Workflow `notify.yaml`
```yaml
name: Notify Pull Request
on: [pull_request]
jobs:
  hangouts:
    name: Hangouts
    runs-on: ubuntu-18.04
    steps:
      - name: Send Message
        uses: docker://<your-repo>/hangouts-action:latest
        env:
              GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```



