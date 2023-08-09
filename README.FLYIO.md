# Deployment
When deploying on Fly.io, you need to set the `CSRF_KEY` value. This can be done by running the following, with a 32 character CSRF key:

```bash
flyctl secrets set CSRF_KEY=<your key here>
```

## Environment
You need to update the `fly.toml` app name to something unique for your account.

## Deploying
The [./Taskfile.yml](./Taskfile.yml) has a `deploy` task that will build and deploy the app to Fly.io. You can run it with:

```bash
task deploy:flyio
```