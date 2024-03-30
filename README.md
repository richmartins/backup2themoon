# backup2themoon

Backup2themoon demonstrates a proof of concept for a backup strategy where a cloud containerized database is backed up to S3, ensuring independence from any specific cloud platform.

## Usage with [IK](https://infomaniak.com)

### Requirments

* [Infomaniak's public cloud instance of cloud computing service](https://www.infomaniak.com/en/hosting/public-cloud)

### Setup openstack cli

1. In public cloud console -> API Access -> Download Openstack RC file and cloud.yml
1. Put `cloud.yml` in -> $HOME/.config/openstack/config.yml
1. Install venv, e.g. `apt install python-venv`
1. activate venv `source venv/bin/activate`
1. `pip install openstackclient`
1. source RC file : `sh XXX-XXXXXXX-openrc.sh`

### Create S3 credential

1. test if connection is ok

        openstack --insecure server list

1. Enter openstack interactive cli

        openstack

1. create cred

        (openstack) ec2 credentials create

1. List credential

        (openstack) ec2 credentials list


### Connection with awscli

1. configure awscli with the S3 credentials created before

        aws configure
        
1.  List buckets

        aws --endpoint-url=https://s3.pub1.infomaniak.cloud s3api list-buckets