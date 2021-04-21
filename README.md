# BinGo
A standardized build and deployment utility built in Go 
intended to accelerate exploration of your own mifedom.

Supports a small(but growing) list of popular application
frameworks/languages.

----
## Supported Environments

**Build targets:**
- Angular
- Docker
- Python <span style="color: red"> (Under construction)
- Nodejs(Typescript)
- Go

**Deploy targets:**
- Local runner  (langauage-specific, in-memory file watcher with hot reload)
- Docker        (Local/Remote)
- GCP           (Remote only, <span style="color: red"> Local K8s -- Under construction</span>)

----
## Developer Install

Clone the git repository and link project root to your path.

$  `git clone git@github.com:wejafoo/bin-zsh.git`

$  `cd bin-zsh && ln -s ~/bin  .`

$   `echo 'export PATH=~/bin:${PATH}' >> ~/.zshrc`

$   `go mod tidy`

$   `go build -race -o ~/bin`

----

## Description

This utility serves as a crude build/test/deploy for any of the [Supported Environments](#supported). 
When configured properly any supported environment should build a mife, bundle the artifacts,
bundle the artifacts in a deployable image, and finally deploy the app to the  
specified target with a single command line capable of leveraging any combination of
command line arguments, smart defaults, local config file, and/or ENV vars using the
following order of precedence:

1.  manual re-deployment of environment variables -- _( local IDE / remote GCP web UI / local gcloud )_
1.  runtime environment deployment overrides -- _( docker-compose.yml / GCP cloudbuild.json )_
1.  command line argument overrides -- (`--TARGET_ALIAS=prod`)
1.  current environment variable overrides -- (`. ~/.zshrc`/`export TARGET_ALIAS=prod`)
1.  flex deployment runtime configuration file -- (`<PROJECT ROOT>` / `.fd && <PROJECT ROOT>` / `.fd.<DEPLOYMENT DOMAIN>`)
1.  intelligent defaults -- _( global hard-coded rules based on language target / Docker image hard-coded defaults )_

The following itemizes the command line config arguments that can be applied at build time(second column) to the CLI or can 
be applied via config or ENV variable using the third column.

**Bools** default to `true` if the flag is added to the command, or using `--build=false` to negate the flag 

**Strings** expect an equal sign followed by double-quoted values.

By convention, the third column can be applied to the Env by prepended 
each camel-cased word boundary with an underscore and switching to all uppercase alphabetic characters 
(e.g. FdBuildContext = FD_BUILD_CONTEXT)

    Bool    build       FdBuild             -- DEFAULT: true - Turns on application-specific builds
    Bool    clean       FdClean             -- DEFAULT: false - Prunes all stopped containers and any unused image, network, or volume
    Bool    debug       FdDebug             -- Turns on detailed logging and enables a debugger if identified
    Bool    local       FdLocal             -- Identifies a build target as local(i.e. not remote)
    Bool    quiet       FdQuiet             -- Turns off all logging to STDOUT 
    Bool    remote      FdRemote            -- Identifies a build target as remote(i.e. not local)
    Bool    verbose     FdTest              -- Run test harness
    Bool    verbose     FdVerbose           -- Verbose execution output
    String  context     FdBuildContext      -- REQUIRED - Boolean that indicates local(-local) or cloud(-remote) deploy
    String  init        FdInit              -- 
    String  nickname    FdNickname          -- Provides the route for mife
    String  service     FdServiceName       -- DEFAULT: $PWD - Working directory and/or Docker Compose service directive.
    String  site        FdSiteNickname      -- Provides the route for mife
    String  alias       FdTargetAlias       -- Recognizable label added to viewable instance name
    String  domain      FdTargetDomain      -- The domain within which the target service will be mapped.
    String  image       FdTargetImageTag    -- The default tag of a newly minted build images.
    String  port        FdTargetLocalPort   -- The host port accessible by a user and mapped to the service port
    String  Log         FdTargetLogLevel    -- DEBUG, INFO, WARNING, ERROR, CRITICAL (default: INFO)
    String  pid         FdTargetProjectId   -- The project ID used for a cloud-based deployments.
    String  realm       FdTargetRealm       -- Prefix that, when prepended to root domain, serves as the app OAuth realm.
    String  port2       FdTargetRemotePort  -- The actual service port of a running container, rarely available to users.


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
