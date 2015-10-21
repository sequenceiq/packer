Testing AWS build process is slow. Lets say you just want to test some post-processor,
or test the build process, which calls `packer build`

This is a modified **amazon-ebs** builder, which doesn’t really builds the images,
just returns fake ami ids.

## Installation

The plugin binary can be downloaded from [github release](https://github.com/sequenceiq/packer/releases/)


You just have to copy the binary beside the **packer** binary:

```
curl -L https://github.com/sequenceiq/packer/releases/download/v0.8.7/packer-builder-amazon-ebs-mock-$(uname).tgz | tar -xzv -C $(dirname $(which packer))
```

## Usage

The only change you need is to **modify the builder type**.
```
  "builders":[
  {
    "name": "my-aws-image",
    "type": "amazon-ebs-mock",
    "mock": "true",
    "ami_name": "delme-docker-{{ user `yum_version_docker` }}-redhat71{{ user `namesuffix` }}",
    "region": "eu-west-1",
    "ami_regions": ["ap-southeast-1","ap-southeast-2","eu-central-1","ap-northeast-1","us-east-1","sa-east-1","us-west-1","us-west-2"],
    
  }
```
## Configuration

The following config prameters are used:

- **mock** : Any value will turn the buider into mock mode. No real AWS api calls are made.
- **region**: Where to build the AMI
- **ami-regions**: An array of regions to copy the AMI (mocked)

## Artifact

The plugin will return an [Artifact](https://godoc.org/github.com/mitchellh/packer/packer#Artifact) with the following properties:

- **String**(): Human readable form of the image
- **State**(): its a generic map, but in case of an ebs-instance builder:
    - **atlas.artifact.metadata**:
        - region.eu-west-1: ami-123456
        - region.eu-central-1: ami-111111

In a post-processor you might get it as a map[interface{}]interface{} because of the RPC magic.

## Hacking

I didn’t create a separate repo for the Mock plugin, but rather made changes in the forked packe repo.

You need golang installed. than you can set up the dev env as:

```
go get -d github.com/mitchellh/packer
cd $GOPATH/src/github.com/mitchellh/packer
git remote add sequenceiq https://github.com/sequenceiq/packer.git
git fetch sequenceiq
git checkout awsmock
## this may take a couple of minutes, be patient
go get ./...
```

### Building the binary

Make targets are added to help dev tasks:
```
make build-mock
```

### Installing

You can install the binary to : `~/.packer.d/plugins/`, so [packer will find it](https://www.packer.io/docs/extend/plugins.html).
```
make install-mock
```

### Releasing

Thats a self note for the binary maintainer:

```
make gh-release
```
