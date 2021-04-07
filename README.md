# BinGo

**NOTE:** If go.mod has changed `go mod vendor` should be rerun.
# bin-zsh

A set of zsh scripts that standardize build and deployment across a growing list of application frameworks and languages.

----

## Currently Supported Environments

**Build targets:**
- Angular
- Docker
- Python
- Nodejs(Typescript)
- Go

**Deploy targets:**
- Local build target runner(in-memory w/ file watcher and hot reload)
- Docker(Local/Remote)
- GCP

----

## Developer Install

Clone the git repository and link project root to your path.

$  `git clone git@github.com:wejafoo/bin-zsh.git`

$  `cd bin-zsh && ln -s ~/bin  .`

$   `echo 'export PATH=~/bin:${PATH}' >> ~/.zshrc`
    
----
## Description

This utility serves as a simplistic way to build/test/deploy any of the [Supported Environments]() with a single command that uses
quasi-intelligent defaults based on the following order of precedence:

1.  manual re-deployment of environment variables ---- _( local IDE / remote GCP web UI / local gcloud )_
1.  runtime environment deployment overrides ---- _( docker-compose.yml / GCP cloudbuild.json )_
1.  command line argument overrides ---- (`--TARGET_ALIAS=prod`)
1.  current environment variable overrides ---- (`. ~/.zshrc`/`export TARGET_ALIAS=prod`)
1.  flex deployment runtime configuration file ---- (`<PROJECT ROOT>` / `.fd && <PROJECT ROOT>` / `.fd.<DEPLOYMENT DOMAIN>`)
1.  intelligent defaults ---- _( global hard-coded rules based on language target / Docker image hard-coded defaults )_


----

## Local/Docker

From one of the aforementioned build targets project root, run the following(assumes Docker is installed locally and working):

$   `bingo --local`


----

## Remote/Cloud

From one of the aforementioned build targets project root, run the following(assumes 'gcloud' is installed locally and able to connect to the targeted cloud project):

$   `bingo -remote=stage`

----

## Housekeeping

Periodically, run this to ensure local Docker container/image residue is not accumulating.  There are several circumstances that can be responsible for this that these scripts may
not otherwise be able to address.

$ `bingo -clean`

Running the above command with no arguments will prune the following local Docker artifacts:
- Stopped containers
- Unused images
- Unused networks
- Unused volumes

----

## Todos

1. **Todo**: Add documentation here that explains the default config file(s) and how they should be modified

1. **Todo**: Add automated test cases that build and deploy all project types following a commit to this repo
