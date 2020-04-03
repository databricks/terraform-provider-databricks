# Databricks Terraform Provider

## Table Of Contents
  * [Getting Started](#getting-started)
  * [Project Components](#project-components)
    + [High Level Databricks Client CR[U]D](#high-level-databricks-client-cr-u-d)
      - [Clusters](#clusters)
      - [Libraries](#libraries)
      - [Jobs API](#jobs-api)
      - [Secrets](#secrets)
        * [Secret Scope](#secret-scope)
        * [Secret](#secret)
        * [Secret ACL](#secret-acl)
      - [Token API](#token-api)
      - [Workspace (Notebooks) API](#workspace--notebooks--api)
      - [SCIM API](#scim-api)
        * [Users](#users)
        * [Groups](#groups)
        * [Group Member](#group-member)
        * [Group entitlements](#group-entitlements)
      - [MLFlow API](#mlflow-api)
      - [Instance Profiles API (AWS Only)](#instance-profiles-api--aws-only-)
      - [DBFS](#dbfs)
    + [Databricks terraform provider resources](#databricks-terraform-provider-resources)
  * [Testing](#testing)
  * [Project Support](#project-support)
  * [Building the Project](#building-the-project)
  * [Deploying / Installing the Project](#deploying---installing-the-project)
  * [Releasing the Project](#releasing-the-project)
  * [Using the Project](#using-the-project)
  
  
## Getting Started
Please note that there is a Makefile which contains all the commands you would need to run this project.

This code base to contribute to requires the following software:

* golang 1.13.X

* python 3.6.X 

## Project Components

### High Level Databricks Client CR[U]D
The client folder contains all the code for the golang sdk for the Databricks REST API. It is kept separate from the databricks terraform provider with its own unit/integration tests. The client so far supports the following:

#### Clusters
TODO!

#### Libraries
* [ ] Create Library Installation

* [ ] Get Library Installation (cluster-status api)

* [ ] Delete Library Installation

> Note: All library manipulation for clusters must be performed when the cluster is in a running state.

#### Jobs API
TODO!

Instance Pools
* [x] Create Instance Pools

* [x] Read Instance Pools

* [x] Update Instance Pools

* [x] Delete Instance Pools

#### Secrets
##### Secret Scope
* [x] Create Secret Scope

* [x] Read Secret Scope*

* [x] Delete Secret Scope

##### Secret
* [x] Create Secret

* [x] Read Secret*

* [x] Delete Secret

Note: For reading a secret it will only fetch the metadata so the value of the secret is only accessible via the databricks runtime in databricks.

##### Secret ACL
* [x] Create Secret ACL

* [x] Read Secret ACL*

* [x] Delete Secret ACL

#### Token API
* [x] Create Token

* [x] Read Token*

* [x] Delete (Revoke) Token

#### Workspace (Notebooks) API
* [ ] Import/create Workspace Path (notebooks)

* [ ] Get Workspace Path (notebooks)

* [ ] Delete Workspace Path (notebooks)

#### SCIM API
##### Users
* [x] Create User

* [x] Update User

* [x] Delete User

Note: For updating a user use Patch for entitlements, Put for everything else. We May want to deal with entitlements as a separate resource to make it easy to manage. Creating the user object and then mapping entitlements are two separate activities.

##### Groups
* [x] Create Group

* [x] Get Group*

* [x] Delete Group

##### Group Member
* [x] Create Member

* [x] Get Member*

* [x] Delete Member

##### Group entitlements
TBD!

#### MLFlow API
TBD!

#### Instance Profiles API (AWS Only)
* [ ] Create IAM Instance Profile Link

* [ ] Get IAM Instance Profile Link

* [ ] Delete IAM Instance Profile Link

#### DBFS
* [ ] Create Object from local file path (/dbfs/put)

* [ ] Get Object (metadata; /dbfs/get-status)

* [ ] Delete Object (metadata; /dbfs/get-status)

* For the Read operation there were no direct rest api calls for read so it is a list and search for the item and return the metadata from the list.


### Databricks terraform provider resources

* [x] Instance pools
* [x] Scim groups
* [x] Scim users
* [x] Secret scopes
* [x] Secrets
* [x] Secret acls
* [ ] Clusters
* [ ] Cluster Policies
* [ ] Entitlements
* [ ] IAM Instance Profiles

## Testing

###TODOs

* [ ] Integration tests should be run at a client level against both azure and aws to maintain sdk parity against both apis **(currently only on one cloud)**
* [ ] Terraform acceptance tests should be run against both aws and azure to maintain parity of provider between both cloud services **(currently only on one cloud)**

## Project Support
Please note that all projects in the /databrickslabs github account are provided for your exploration only, and are not formally supported by Databricks with Service Level Agreements (SLAs).  They are provided AS-IS and we do not make any guarantees of any kind.  Please do not submit a support ticket relating to any issues arising from the use of these projects.

Any issues discovered through the use of this project should be filed as GitHub Issues on the Repo.  They will be reviewed as time permits, but there are no formal SLAs for support.


## Building the Project
Instructions for how to build the project

## Deploying / Installing the Project
Instructions for how to deploy the project, or install it

## Releasing the Project
Instructions for how to release a version of the project

## Using the Project
Simple examples on how to use the project
