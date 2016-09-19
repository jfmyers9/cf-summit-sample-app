# cf-summit-sample-app

This is a sample application used to demonstrate the differences in user experience for pushing buildpack and docker applications on Cloud Foundry.

## Deploying

### Buildpacks

In order to push this application using the golang buildpack, run the following from the root of the repository:

```bash
cf push app-name
```

### Docker

In order to push this application uisng docker, run the rollowing from the root of the repository:

```bash
docker build . -t DOCKER_ORG/DOCKER_REPO
docker push DOCKER_ORG/DOCKER_REPO
cf push app-name -o DOCKER_ORG/DOCKER_REPO
```

Replace `DOCKER_ORG` and `DOCKER_REPO` with the docker organization and repository of your docker image.
