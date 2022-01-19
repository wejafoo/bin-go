#BinGo
A standardized minimalistic build/deployment utility built in Go with the intention of accelerating 
mifedom exploration.

Supports a small(but growing) list of popular languages/frameworks.

----
##Supported Environments

**Build targets:**
- Angular
- Docker
- Python <span style="color: red"> (Under construction)
- Nodejs(Typescript)
- Go

**Default deploy targets:**
- dev       - Target-specific runner (language-specific file watcher with hot reload)
- local     - Docker (local and/or remote)
- remote    - GCP    (local and/or remote, <span style="color: red"> Local K8s -- Under construction</span>)

----
## Developer Install

Clone the git repository and link project root to your path.

$  `git clone git@github.com:wejafoo/bin-go.git`

$  `cd bin-go && ln -s ~/bin  .`

$   `echo 'export PATH=~/bin:${PATH}' >> ~/.zshrc`

$   `go mod tidy`

$   `go build -race -o ~/bin`

----

## Description

This utility is a minimal build/test/deploy utility aspires to be the lowest common
denominator of the [Supported Environments](#supported-environments).

Once configured, builds a mife or mise, bundles its artifacts,
publishes an image, and deploys a container to the specified target
in a single command line that respects the environment with the following priority :

1.  **command line arguments**          -- _(`--TARGET_ALIAS=prod`)_
1.  **local environment**               -- _( local IDE / remote GCP web UI / local gcloud )_
1.  **remote deployment environment**   -- _( docker-compose.yml / GCP cloudbuild.json )_
1.  **remote deployment environment**   -- _( docker-compose.yml / GCP cloudbuild.json )_
1.  **bingo configuration file**        -- _(`<PROJECT ROOT>` / `.fd && <PROJECT ROOT>` / `.fd.<DEPLOYMENT DOMAIN>`)_
1.  **bingo defaults**                  -- _( global hard-coded rules based on language target / Docker image hard-coded defaults )_

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
    String  adc         FdBuildContext      -- REQUIRED - String indicating the location of default remote credentials
    String  context     FdBuildContext      -- REQUIRED - String indicating local(-local) or cloud(-remote) deploy
    String  init        FdInit              -- 
    String  nickname    FdService           -- Provides the route for mife
    String  route       FdRouteBase         -- Provides the route for mife
    String  repo        FdRepo              -- DEFAULT: $PWD - Working directory and/or Docker Compose service directive.
    String  site        FdTitle             -- Provides the pretty name for the service
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

From project root, run the following(assumes Docker is installed locally and working):

$   `bingo --local`


----

## Remote/Cloud

From project root, run the following(assumes 'gcloud' is installed locally and able to connect to the targeted cloud project):

$   `bingo -remote=stage`

----

## Housekeeping

For a variety of reasons, Docker tends to accumulate residue over time and eventually presents
unwanted side effects, so it might be a good idea to periodically run the following:

$ `bingo -clean`

With no additional arguments, this will prune the following local Docker artifacts:

    - Stopped containers
    - Unused images
    - Unused networks
    - Unused volumes

Although it may not clean up your Docker issue, it will likely return Gbs of unnecessarily hoarded disk in just a few seconds.

----

## Best Practice

#### FD config file

- FdTargetRealm - should always end with a '.'
- FdRouteBase - should always begin with a '/' and, if the mife contains sub-routes, end with a '/'


----

## Todos

1. **Todo**: Add documentation here that explains the default config file(s) and how they should be modified

1. **Todo**: Add automated test cases that build and deploy all project types following a commit to this repo
