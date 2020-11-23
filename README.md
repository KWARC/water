# Water -- WissKI Automation Tool & Execution Robot

** This repository is a plan for how water will work. It hasn't been implemented yet **

This repository contains the source code of water. 
Water is a tool for automatically installing a maintaining a WissKI Factory or single WissKI installation. 

## Installation

1. Install [docker](https://docs.docker.com/install/) on the system you wish to use WissKI on. 
2. Either download a pre-built binary and put it in path, or run it via docker:

```docker run -ti --rm -v /var/run/docker.sock:/var/run/docker.sock kwarc/water```

## Reference

- ```water init /path/on/host``` -- Bootstrap a single-site or multi-site installation in `/path/to/host`
- ```water up``` -- Install or Update an existing installation
- ```water shell``` -- Open a shell inside the management container

## Internal notes

- In the initial version we depend on the external repository. 
- We run the setup script inside the water container, in a future version we will no longer do that and bundle dependencies accordingly.

### Water Init
- Create a container with host networking named `water`
- Should auto-start
- Mount in the docker socket and the host directory, store the mount in `info.kwarc.water.path`

### Water Up
- Clone the git repo
- Run the setup script
