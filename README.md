# Databricks Terraform Provider
## Getting Started
Please note that there is a Makefile which contains all the commands you would need to run this project.

This code base to contribute to requires the following software:

* golang 1.13.X

* python 3.6.X 

## Project Components
###High Level Databricks Client CR[U]D
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
* [ ] Create Token

* [ ] Read Token*

* [ ] Delete (Revoke) Token

#### Workspace (Notebooks) API
* [ ] Import/create Workspace Path (notebooks)

* [ ] Get Workspace Path (notebooks)

* [ ] Delete Workspace Path (notebooks)

#### SCIM API
##### Users
* [ ] Create User

* [ ] Update User

* [ ] Delete User

Note: For updating a user use Patch for entitlements, Put for everything else. We May want to deal with entitlements as a separate resource to make it easy to manage. Creating the user object and then mapping entitlements are two separate activities.

##### Groups
* [ ] Create Group

* [ ] Get Group*

* [ ] Delete Group

##### Group Member
* [ ] Create Member

* [ ] Get Member*

* [ ] Delete Member

##### Group entitlements
TBD!

#### MLFlow API
TODO!

#### Instance Profiles API (AWS Only)
* [ ] Create IAM Instance Profile Link

* [ ] Get IAM Instance Profile Link

* [ ] Delete IAM Instance Profile Link

#### DBFS
* [ ] Create Object from local file path (/dbfs/put)

* [ ] Get Object (metadata; /dbfs/get-status)

* [ ] Delete Object (metadata; /dbfs/get-status)

* For the Read operation there were no direct rest api calls for read so it is a list and search for the item and return the metadata from the list.


## Testing

###TODOs

* [ ] Integration tests should be run at a client level against both azure and aws to maintain sdk parity against both apis **(currently only on one cloud)**
* [ ] Terraform acceptance tests should be run against both aws and azure to maintain parity of provider between both cloud services **(currently only on one cloud)**