FORMAT: 1A

# Nano API

REST API provided by core service

# Group Resource Pool

## Compute Pool [/compute_pools/{pool}]

manage compute pool

+ Parameters
  + pool (string, optional) - pool name/id

### Get all pools in zone [GET /compute_pools/]

query all compute pool info in zone

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "name": "default",
          "enabled": true,
          "cells": 10,
          "storage": "some_ceph_cluster",
          "network": "nat_address_pool_1",
          "failover": true
        },
        {
          "name": "fast_instance",
          "enabled": true
          "cells": 12,
          "storage": "some_ceph_cluster",
          "network": "nat_address_pool_1",
          "failover": false
        }
      ]
    }
  ]

### Query Pool Info [GET]

query compute pool details

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "name": "default",
        "enabled": true,
        "cells": 10,
        "storage": "some_ceph_cluster",
        "network": "nat_address_pool_1",
        "failover": true
      }
    }
  ]

### Create Compute Pool [POST]

create new compute pool

+ Attributes
  + cells(array[string], optional) - list of initial cell names
  + network_mode (enum[string], optional) - plain/share/mono
  + storage (string, optional) - attached storage pool name
  + network (string, optional) - attached address pool name
  + failover (boolean, optional) - enable failover


+ Request create compute pool (application/json)

  [
    {
      "cells":["cell1", "cell2"],
      "network_mode": "share",
      "storage": "some_ceph_cluster",
      "network": "nat_address_pool_1",
      "failover": true
    }
  ]

### Delete Compute Pool [DELETE]

delete compute pool

### Modify Compute Pool [PUT]

modify compute pool config

+ Attributes

  + enable(boolean, optional) - set pool enable/disable
  + network_mode (enum[string], optional) - plain/share/mono
  + storage (string, optional) - attached storage pool name
  + network (string, optional) - attached address pool name

+ Request change compute pool (application/json)

  [
    {
      "failover": true,
      "storage": "nfs-poo1"
    }
  ]


## Compute Pool Cell [/compute_pool_cells/{pool}/{cell}]


+ Parameters

  + pool (string, optional) - pool name
  + cell (string, optional) - cell name

### Query unallocated cells [GET /compute_pool_cells/]

query cells not attached to any pool

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data": [
        {
          "name": "cell3",
          "address": "172.16.3.78",
          "enabled": true,
          "alive": true,
        },
        {
          "name": "cell4",
          "address": "172.16.3.79",
          "enabled": true,
          "alive": true,
        }
      ]
    }
  ]

### Query cells in pool [GET /compute_pool_cells/{pool}]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data": [
      {
        "name": "cell1",
        "address": "172.16.4.7",
        "enabled": true,
        "alive": true,
      },
      {
        "name": "cell2",
        "address": "172.16.7.3",
        "enabled": true,
        "alive": true,
      }
      ]
    }
  ]

### Add Pool Cell [POST]

Add {cell} to {pool}

### Remove Pool Cell [DELETE]

### Modify Pool Cell [PUT]

modify compute pool config

+ Attributes

  + enable(boolean, optional) - enable/disable cell

+ Request change compute pool (application/json)

  [
    {
      "enable": false
    }
  ]

### Get Cell Config [GET]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data": {
        "name": "cell1",
        "address": "172.16.4.7",
        "enabled": true,
        "alive": true,
        "storage": [
        {
          "name": "nfspool1",
          "attached": true
        },
        {
          "name": "cephpool1",
          "attached": false,
          "error": "not support yet"
        }
        ]
      }
    }
  ]

## Storage Pool [/storage_pools/{pool}]

### Query Storage Pool [GET /storage_pools/]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "name": "nfspool1",
          "type": "nfs",
          "host": "nfs.nano.com",
          "target": "/var/nano/pool1"
        },
        {
          "name": "cephpool1",
          "type": "ceph",
          "host": "ceph.nano.com",
          "target": "some_pool_in_ceph"
         }
      ]
    }
  ]


### Get Storage Pool Info [GET]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data": {
        "type": "nfs",
        "host": "nfs.nano.com",
        "target": "/var/nano/pool1"
      }
    }
  ]

### Create Storage Pool [POST]

Create a new stoarge pool

+ Attributes
  + type (string) - storage type
  + host (string, optional) - ip or hostname of pool target
  + target (string, optional) - pool targe path

+ Request create storage pool (application/json)

  [
    {
      "type": "nfs",
      "host": "nfs.nano.com",
      "target": "/var/nano/pool1"
    }
  ]

### Modify Storage Pool [PUT]

Modify stoarge pool, only empty pool allowed(no volumes available)

+ Attributes
  + type (string) - storage type
  + host (string, optional) - ip or hostname of pool target
  + target (string, optional) - pool targe path

+ Request create storage pool (application/json)

  [
    {
      "type": "nfs",
      "host": "another_nfs.nano.com",
      "target": "/var/nano/new_location"
    }
  ]

### Delete Storage Pool [DELETE]

## Compute Zone Status [/compute_zone_status/]

### query zone status [GET]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "name": "default",
        "pools": [1, 0]//[disabled, enabled]
        "cells": [0, 25], //[offline, online]
        "instances": [3, 100, 3, 1], //[stopped, running, lost, migrate]
        "cpu_usage": 156.45,
        "max_cpu": 320,
        "available_memory": 560,
        "max_memory": 960,
        "available_disk": 12457,
        "max_disk": 34000,
        "read_speed": 8634,
        "write_speed": 3673,
        "receive_speed": 7634,
        "send_speed": 2643,
        "start_time": "2018-01-02 15:04:05"
      }
    }
  ]

## Compute Pool Status [/compute_pool_status/]

+ Parameters
  + pool (string, optional) - target pool name

### query pools status in zone [GET]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[{
        "name": "default",
        "enabled": true,
        "cells": [0, 25], //[offline, online]
        "instances": [3, 100, 3, 1], //[stopped, running, lost, migrate]
        "cpu_usage": 156.45,
        "max_cpu": 320,
        "available_memory": 560,
        "max_memory": 960,
        "available_disk": 12457,
        "max_disk": 34000,
        "read_speed": 8634,
        "write_speed": 3673,
        "receive_speed": 7634,
        "send_speed": 2643
      },
      {
        "name": "pool2",
        "enabled": true,
        "cells": [0, 15], //[offline, online]
        "instances": [3, 50, 0, 0], //[stopped, running, lost, migrate]
        "cpu_usage": 156.45,
        "max_cpu": 320,
        "available_memory": 560,
        "max_memory": 960,
        "available_disk": 12457,
        "max_disk": 34000,
        "read_speed": 8634,
        "write_speed": 3673,
        "receive_speed": 7634,
        "send_speed": 2643
      }]
    }
  ]

### query compute pool status [GET /compute_pool_status/{pool}]


+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "name": "default",
        "enabled": true,
        "cells": [0, 25], //[offline, online]
        "instances": [3, 100, 3, 1], //[stopped, running, lost, migrate]
        "cpu_usage": 156.45,
        "max_cpu": 320,
        "available_memory": 560,
        "max_memory": 960,
        "available_disk": 12457,
        "max_disk": 34000,
        "read_speed": 8634,
        "write_speed": 3673,
        "receive_speed": 7634,
        "send_speed": 2643
      }
    }
  ]

## Compute Cell Status [/compute_cell_status/{pool}/{cell}]

+ Parameters
  + pool (string) - target pool name
  + cell (string, optional) - target cell name

### query compute cell status [GET /compute_cell_status/{pool}/]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "name": "cell_93e3de24f0a4",
          "address": "172.16.4.7",
          "enabled": true,
          "alive": true,
          "instances": [3, 15, 0, 0], //[stopped, running, lost, migrate]
          "cpu_usage": 156.45,
          "max_cpu": 320,
          "available_memory": 560,
          "max_memory": 960,
          "available_disk": 12457,
          "max_disk": 34000,
          "read_speed": 8634,
          "write_speed": 3673,
          "receive_speed": 7634,
          "send_speed": 2643
        },
        {
          "name": "cell_93e3de24f0a3",
          "address": "172.16.4.8",
          "enabled": true,
          "alive": true,
          "instances": [0, 12, 1, 0], //[stopped, running, lost, migrate]
          "cpu_usage": 156.45,
          "max_cpu": 320,
          "available_memory": 560,
          "max_memory": 960,
          "available_disk": 12457,
          "max_disk": 34000,
          "read_speed": 8634,
          "write_speed": 3673,
          "receive_speed": 7634,
          "send_speed": 2643
        }
      ]
    }
  ]

### get compute cell status [GET /compute_cell_status/{pool}/{cell}]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "name": "cell_93e3de24f0a4",
        "address": "172.16.4.7",
        "enabled": true,
        "alive": true,
        "instances": [3, 15, 0, 0], //[stopped, running, lost, migrate]
        "cpu_usage": 156.45,
        "max_cpu": 320,
        "available_memory": 560,
        "max_memory": 960,
        "available_disk": 12457,
        "max_disk": 34000,
        "read_speed": 8634,
        "write_speed": 3673,
        "receive_speed": 7634,
        "send_speed": 2643
      }
    }
  ]

## Instance Status [/instance_status/{pool}/{cell}/]

+ Parameters
  + pool (string) - target pool name
  + cell (string, optional) - target cell name

### Query instance in pool [GET /instance_status/{pool}/]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data": [{
        "name": "test01",
        "id": "df6723jhew67f3fdsf-fefew",
        "created": true,
        "running": true,
        "cores": 4,
        "memory": 5120,
        "disks": [4864000, 854365],
        "auto_start": true,
        "system": "linux",
        "display_protocol": "vnc",
        "monitor_secret": "abd",
        "internal":{
          "network_address": "172.18.6.7",
          "display_address": "172.18.5.3:5901",
          "allocated_address": "172.18.6.7"
        },
        "external":{
          "network_address": "202.3.1.34",
          "display_address": "202.3.1.34:5901"
        },
        "create_time": "2018-08-21 00:12:34",
        "media_attached": true,
        "media_source": "centos_7_x64_iso"
      },
      {
        "name": "test02",
        "id": "dr6ufh73dgjf3fdsf-fefew",
        "created": false,
        "progress": 46,
        "running": false,
        "cores": 4,
        "memory": 5120,
        "disks": [4864000, 854365],
        "auto_start": false,
        "system": "linux"
      }]
    }
  ]

### Query instance in cell [GET /instance_status/{pool}/{cell}]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data": [{
        "name": "test01",
        "id": "df6723jhew67f3fdsf-fefew",
        "created": true,
        "running": true,
        "cores": 4,
        "memory": 5120,
        "disks": [4864000, 854365],
        "auto_start": true,
        "system": "linux",
        "display_protocol": "vnc",
        "monitor_secret": "abd",
        "internal":{
          "network_address": "172.18.6.7",
          "display_address": "172.18.5.3:5901",
          "allocated_address": "172.18.6.7"
        },
        "external":{
          "network_address": "202.3.1.34",
          "display_address": "202.3.1.34:5901"
        },
        "create_time": "2018-08-21 00:12:34",
        "media_attached": true,
        "media_source": "centos_7_x64_iso"
      },
      {
        "name": "test02",
        "id": "dr6ufh73dgjf3fdsf-fefew",
        "created": false,
        "progress": 46,
        "running": false,
        "cores": 4,
        "memory": 5120,
        "disks": [4864000, 854365],
        "auto_start": false,
        "system": "windows"
      }]
    }
  ]

## Address Pool [/address_pools/{name}]

+ Parameters
  + name (string, optional) - pool name

### Query all address pools [GET /address_pools/]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "name": "address1",
          "gateway": "172.16.3.254",
          "addresses": 120,
          "allocated": 17
        },
        {
          "name": "address2",
          "gateway": "172.16.7.254",
          "addresses": 230,
          "allocated": 10
        }
      ]

    }
  ]

### Get address pool status [GET]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "gateway": "172.16.3.254",
        "dns": [
          "8.8.8.8",
          "10.10.10.10"
        ],
        "ranges": [
          {
            "start": "172.16.3.1",
            "end": "172.16.3.18",
            "netmask": "255.255.255.0",
          },
          {
            "start": "172.16.3.20",
            "end": "172.16.3.190",
            "netmask": "255.255.255.0",
          }
        ],
        "allocated": [
          {
              "address": "172.16.3.6",
              "instance": "abcdef-1234567890",
          },
          {
              "address": "172.16.3.17",
              "instance": "abcdef-1234567893",
          }
        ]
      }
    }
  ]

### Create new address pool [POST]

+ Request Add Address Range (application/json)

  + Attributes
    + gateway (string) - gateway bound with instance
    + dns (array[string]) - list of dns server bound with instance

  + Body

    {
      "gateway": "172.16.3.254",
      "dns": [
        "8.8.8.8",
        "10.10.10.10"
      ]
    }

### Modify address pool [PUT]

### Delete address pool [DELETE]

## Address Range [/address_pools/{name}/{type}/ranges/{start}]

+ Parameters
  + name (string) - name of address pool
  + type (enum[string]) - range type, 'external'/'internal'
  + start (string, optional) - start address of range

### Query Address Ranges [GET /address_pools/{name}/{type}/ranges/]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "start": "172.16.3.1",
          "end": "172.16.3.18",
          "netmask": "255.255.255.0"
        },
        {
          "start": "172.16.3.25",
          "end": "172.16.3.200",
          "netmask": "255.255.255.0"
        }
      ]
    }
  ]

### Get Address Ranges Status [GET]

+ Response 200 (application/json)

    [
      {
        "error_code": 0,
        "message": "",
        "data":{
          "start": "172.16.3.1",
          "end": "172.16.3.18",
          "netmask": "255.255.255.0",
          "allocated": [
            {
                "address": "172.16.3.6",
                "instance": "abcdef-1234567890",
            },
            {
                "address": "172.16.3.17",
                "instance": "abcdef-1234567893",
            }
          ]
        }
      }
    ]


### Add Address Range [POST]

+ Request Add Address Range (application/json)

  + Attributes
    + end (string) - end address of ip range
    + netmask (string) - netmask of ip range

  + Body

    {
      "end": "172.16.3.18",
      "netmask": "255.255.255.0"
    }

### Remove Address Range [DELETE]

# Group Virtual machines

## Guest Search [/guest_search/]

### Search Guests [GET]

+ Parameters

  + pool (string) - target pool name
  + cell (string, optional) - target cell
  + owner (string, optional) - guest owner
  + group (string, optional) - user group
  + status (number, optional) - running status
  + created (boolean, optional) - guest created

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
            "name": "guest1",
            "id": "df6723jhew67f3fdsf-fefew",
            "running": false,
            "owner": "admin",
            "group": "manager",
            "pool": "pool1",
            "cell": "cell_93e3de24f0a4",
            "cores": 4,
            "memory": 5120,
            "total_disk": 4864000,
            "disks": [482345, 487534],
            "auto_start": true,
            "system": "windows",
            "ethernet_address": "ed:35:3d:5a:4e:3f",
            "display_protocol": "vnc",
            "internal":{
              "network_address": "172.18.6.7",
              "display_address": "172.18.5.3:5901",
              "allocated_address": "172.18.6.7"
            },
            "external":{
              "network_address": "202.3.1.34",
              "display_address": "202.3.1.34:5901"
            },
            "create_time": "2018-08-21 00:12:34"
          }
        ],
      [
        {
            "name": "guest2",
            "id": "df6723jhew67f3fdsf-fefet",
            "running": false,
            "owner": "admin",
            "group": "manager",
            "pool": "pool1",
            "cell": "cell_93e3de24f0a4",
            "cores": 4,
            "memory": 5120,
            "total_disk": 4864000,
            "disks": [482345, 487534]
            "auto_start": true,
            "system": "linux",
            "ethernet_address": "ed:35:3d:5a:4e:3f",
            "display_protocol": "vnc",
            "internal":{
              "network_address": "172.18.6.7",
              "display_address": "172.18.5.3:5901",
              "allocated_address": "172.18.6.7"
            },
            "external":{
              "network_address": "202.3.1.34",
              "display_address": "202.3.1.34:5901"
            },
            "create_time": "2018-08-21 00:12:34"
          }
      ]
    }
  ]

## Guests [/guests/]

VM config/storage/network

### Get Guest Details [GET /guests/{id}]

+ Parameters
  + id (string, optional) - guest id


+ Response 201 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "name": "test01",
        "created": false
        "progress": 57 //limit to 100
      }
    }
  ]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "name": "test01",
        "created": true,
        "running": false
        "owner": "admin",
        "group": "manager",
        "pool": "pool1",
        "cell": "cell_93e3de24f0a4"
        "cores": 4,
        "memory": 5120,
        "total_disk": 4864000,
        "disks": [482345, 487534],
        "auto_start": true,
        "system": "linux",
        "ethernet_address": "ed:35:3d:5a:4e:3f"
        "display_protocol": "vnc",
        "monitor_secret": "abd",
        "internal":{
          "network_address": "172.18.6.7",
          "display_address": "172.18.5.3:5901",
          "allocated_address": "172.18.6.7"
        },
        "external":{
          "network_address": "202.3.1.34",
          "display_address": "202.3.1.34:5901"
        },
        "create_time": "2018-08-21 00:12:34"
      }
    }
  ]

### Create Guest [POST]

request create guest, return a new guest id for query after creating request submit

+ Request create guest (application/json)

  + Attributes
    + name (string) - display name
    + owner (string) - user id
    + group (string) - user group
    + pool (string) - compute pool name
    + cores (number) - vm cpu cores
    + memory (number) - vm memory
    + disks (array[number]) - [system_disk_size, data_disk_n_size...]
    + auto_start (boolean) - vm auto_start
    + system (string, optional) - system type, "linux", "windows", etc
    + network_address (string, optional) - optional specified ip address
    + ethernet_address (string, optional) - optional specified mac address
    + from_image (string, optional) - initial disk image id
    + ports (array[number], optional) - nat ports
    + cpu_priority (enum[number], optional) - 0~2, high ~ low
    + max_io (number, optional) - max iops
    + inbound (number, optional) - max inbound bandwidth
    + outbound (number, optional) - max outbound bandwidth
    + system_version (string, optional) - os version, 'centos7', 'win2012'
    + modules (array[string], optional) - installed guest module, "qemu"/"cloud-init"
    + cloud_init (object, optional) - ci params
      - root_enabled (boolean, optional) - allow root user login via SSH, enable in default
      - admin_name (string, optional) - specify admin name, using 'root' when omitted
      - admin_secret (string, optional) - password in plain text, nano generate a new one when omitted
      - data_path (string, optional) - data disk mount path, '/opt/data' in default

  + Body

    {
      "name": "some_instance",
      "owner": "admin",
      "group": "manager",
      "pool": "pool_1",
      "cores": 8,
      "memory": 4048,
      "disks": [4048, 38443],
      "auto_start": true,
      "system": "linux",
      "system_version": "centos7",
      "modules": ["qemu", "cloud-init"],
      "cloud_init":{
        "root_enabled": true,
        "admin_name": "nano",
        "admin_secret": "12345678",
        "data_path": "/opt/data"
      }
    }

+ Response 202 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "id": "dsfds8979847r3dsf-3r67",
      }
    }
  ]


### Delete Guest [DELETE /guests/{id}]

delete guest, release all resource back to pool

+ Request delete guest (application/json)

  + Attributes
    + force (boolean, optional) - stop running instance for deleting

  + Body

    {
        "force": true
    }

## Guest Cores [/guest/{id}/cores]

### Modify Cores [PUT]

+ Request modify cores (application/json)

  + Attributes
    + cores(number) - new cores
    + immediate (boolean, optional) - take effect when running

  + Body

    {
      "cores": 4,
      "immediate": false
    }

## Guest Memory [/guest/{id}/memory]

### Modify Memory [PUT]

+ Request modify memory (application/json)

  + Attributes
    + memory(number) - memory in bytes
    + immediate (boolean, optional) - take effect when running

  + Body

    {
      "memory": 4294967296,
      "immediate": false
    }

## Resize Guest Disk [/guest/{id}/disks/resize/{disk}]

+ Parameters
  + id (string) - guest id
  + disk (number) - index of target disk, system:0, data0:1, data1:2...

### Modify Disk Size [PUT]

+ Request modify disk (application/json)

  + Attributes
    + size(number) - size in bytes
    + immediate (boolean, optional) - take effect when running

  + Body

    {
      "size": 42949672960,
      "immediate": false
    }

## Shrink Guest Disk [/guest/{id}/disks/shrink/{disk}]

+ Parameters
  + id (string) - guest id
  + disk (number) - index of target disk, system:0, data0:1, data1:2...

### Shrink Disk [PUT]

+ Request shrink disk (application/json)

  + Attributes
    + immediate (boolean, optional) - take effect when running

  + Body

    {
      "immediate": false
    }

## Guest System [/guest/{id}/system/]

### Reset Guest System [PUT]

+ Request Reset Guest System (application/json)

  + Attributes
    + from_image (string) - id of new system image

  + Body

    {
      "from_image": "0123456789-abcdef"
    }

## Guest User Auth [/guest/{id}/auth]

### Modify User Password [PUT]

+ Request modify disk (application/json)

  + Attributes
    + password (string, optional) - new password, automated generate when omitted
    + user(number, optional) - username, automated choose username by system type when omitted

  + Body

    {
      "password": "asdb12356",
      "user": "root"
    }

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data": {
        "password": "1235656756",
        "user": "root"
      }
    }
  ]

### Get User Password [GET]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data": {
        "password": "1235656756",
        "user": "root"
      }
    }
  ]

## Guest Name [/guest/{id}/name/]

### Modify Guest Name [PUT]

modify guest name

+ Request Modify Name (application/json)

  + Attributes
    + name (string) - display name

  + Body

    {
      "name": "some_instance",
    }

## Instances [/instances/{id}]

VM running instance

+ Parameters
  + id (string) - instance id/ guest id

### Get running status [GET]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data": {
        "name": "test01",
        "created": true,
        "running": true,
        "owner": "admin",
        "group": "manager",
        "pool": "pool1",
        "cell": "cell_93e3de24f0a4",
        "cores": 4,
        "memory": 5120,
        "total_disk": 4864000,
        "disks": [482345, 487534]
        "auto_start": true,
        "ethernet_address": "ed:35:3d:5a:4e:3f",
        "display_protocol": "vnc",
        "internal":{
          "network_address": "172.18.6.7",
          "display_address": "172.18.5.3:5901",
          "allocated_address": "172.18.6.7"
        },
        "external":{
          "network_address": "202.3.1.34",
          "display_address": "202.3.1.34:5901"
        },
        "create_time": "2018-08-21 00:12:34",
        "media_attached": true,
        "media_source": "centos_7_x64_iso",
        "cpu_usage": 5.34,
        "memory_available": 1280,
        "disk_available": 4925700,
        "bytes_read": 3673,
        "bytes_written": 8634,
        "bytes_received": 2643,
        "bytes_sent": 7634
      }
    }
  ]

### Start Instance [POST]

start guest instance

+ Request start instance (application/json)

  + Attributes
    + from_media (boolean, optional) - boot from media
    + from_network (boolean, optional) - boot from network source/PXE
    + source (string, optional) - boot source id/URI

  + Body

    {
      "from_media": true,
      "source": "bd0fe127-a8db-4c89-98ea-9a5070b08aae"
    }

### Stop Instance [DELETE]

stop guest instance

+ Request stop instance (application/json)

  + Attributes
    + reboot (boolean) - shutdown or reboot
    + force(boolean) - force operate

  + Body

    {
      "reboot": false,
      "force": true
    }

## Instance Media [/instances/{id}/media]

### Insert Media [POST]

Insert media into a running instance

+ Request insert Media (application/json)

  + Attributes
    + source (string) - media id
    + type (int, optional) - media type, 0 = Media Image

  + Body

    {
        "source": "bd0fe127-a8db-4c89-98ea-9a5070b08aae",
        "type": 0
    }

### Eject Media [DELETE]

Eject media from a running instance

## Snapshots manage [/instances/{id}/snapshots/]

### Query snapshot tree [GET]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
      "origin":{
        "is_root": true
      },
      "snapshot1": {
        "backing": "origin"
      },
      "snapshot2": {
        "backing": "snapshot1"
      },
      "snapshot3": {
        "backing": "snapshot2",
        "is_current": true
      },
      "another_branch": {
        "backing": "origin"
      },
      "another_end": {
        "backing": "another_branch"
      }
      ]
    }
  ]

### Create new snapshot [POST]

creae a new snapshot

+ Request restore snapshot (application/json)

  + Attributes
    + name (string) - snapshot name
    + description (string, optional) - snapshot description

  + Body

    {
        "name": "snapshot1",
        "description": "this is a snapshot example"
    }

### Restore snapshot [PUT]

resume to a specified snapshot

+ Request restore snapshot (application/json)

  + Attributes
    + target (string) - snapshot name

  + Body

    {
        "target": "snapshot1"
    }


## Snapshot [/instances/{id}/snapshots/{name}]

### Get snapshot info [GET]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "running": false,
        "description": "this is snapshot before installing cloud-init",
        "create_time": "2018-08-21 12:34:56"
      }
    }
  ]

### Delete snapshot [DELETE]

## Monitor Channel [/monitor_channels/{id}]

create channel for monitor and control instance. Create a new temporary channel id 'ABC' using POST Method, then connect and consumed channel with websocket address 'ws://host/monitor_channels/ABC' using noVNC or other clients. Unused channel will be released if no connection established in time.

### Create New Channel [POST]

+ Request create guest (application/json)

  + Attributes
    + guest (string) - target guest id

  + Body

    {
      "guest": "some_guest_id",
    }

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
          "id": "channel-id-123",
          "protocol": "vnc",
          "username": "",
          "password": "some_secret",
        }
    }
  ]

# Group Image

## Media Image [/media_images/{id}]

+ Parameters
  + id (string) - image id

### Query Image list [Get]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "name": "centos7_64_minimal",
          "id": "sdf83kjfe-23r4wedf",
          "description": "some desc",
          "create_time": "2018-08-21 00:12:34",
          "modify_time": "2018-08-21 00:13:34",
          "size": 4872334659735,
          "tags": ["linux", "64bit", "centos"]
        },
        {
          "name": "win7_home_64",
          "id": "sdf83kjfe-23r4wertytdf",
          "description": "win desktop",
          "create_time": "2018-08-21 00:12:34",
          "modify_time": "2018-08-21 00:13:34",
          "size": 4872334659774,
          "tags": ["windows", "64bit", "windows7"]
        }
      ]
    }
  ]

### Create Media Image [POST]

+ Request Create Image (application/json)

  + Attributes

    + name (string) - image name
    + owner (string) - image creater
    + group (string) - image group
    + description (string, optional) - image description
    + tags (array[string], optional) - image tags

  + Body

    {
      "name": "centos7_64_minimal",
      "owner": "admin",
      "group": "manager",
      "description": "some desc",
      "tags": ["linux", "64bit", "centos"]
    }

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "id": "sdf83kjfe-23r4wedf",
        }
      ]
    }
  ]

### Modify Media Image [PUT]

+ Request Modify Media Image (application/json)

  + Attributes

    + name (string, optional) - image name
    + owner (string, optional) - image creater
    + group (string, optional) - image group
    + description (string, optional) - image description
    + tags (array[string], optional) - image tags

  + Body

    {
      "name": "centos7_64_minimal",
      "owner": "admin",
      "group": "manager",
      "description": "some desc",
      "tags": ["linux", "64bit", "centos"]
    }

### Detele Media Image [DELETE]

## Media Image File [/media_image_files/{id}]

+ Parameters
  + id (string) - image id

### Download image file [GET]

### Upload image file [POST]

+ Request upload (multipart/form-data)

  + Attributes
    + image (binary) - binary upload image data

  + Body

## Disk Image Search [/disk_image_search]

### Search Disk Image [GET]

+ Request query disk images (application/json)

  + Parameters

    + group (string) - image group
    + tags (array[string], optional) - image tags

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "name": "centos7_64_minimal",
          "id": "sdf83kjfe-23r4wedf",
          "description": "some desc",
          "create_time": "2018-08-21 00:12:34",
          "modify_time": "2018-08-21 00:13:34",
          "size": 4872334659735,
          "tags": ["linux", "64bit", "centos"]
        },
        {
          "name": "win7_home_64",
          "id": "sdf83kjfe-23r4wertytdf",
          "description": "win desktop",
          "create_time": "2018-08-21 00:12:34",
          "modify_time": "2018-08-21 00:13:34",
          "size": 4872334659774,
          "tags": ["windows", "64bit", "windows7"]
        }
      ]
    }
  ]

## Disk Image [/disk_images/{id}]

+ Parameters

  + id (string, optional) - disk image id

### Query Disk Image [GET /disk_images/{id}]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data": {
          "name": "centos7_64_minimal",
          "created": false,
          "progress": 46,
          "description": "some desc",
          "tags": ["linux", "64bit", "centos"]
        }
    }
  ]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data": {
          "name": "centos7_64_minimal",
          "created": true,
          "description": "some desc",
          "create_time": "2018-08-21 00:12:34",
          "modify_time": "2018-08-21 00:13:34",
          "size": 4872334659735,
          "tags": ["linux", "64bit", "centos"]
        }
    }
  ]



### Create Disk Image [POST]

+ Request Create Image (application/json)

  + Attributes

    + name (string) - image name
    + guest (string) - source guest id
    + owner (string) - image creater
    + group (string) - image group
    + description (string, optional) - image description
    + tags (array[string], optional) - image tags

  + Body

    {
      "name": "centos7_64_minimal",
      "guest": "dir723rgfyu-rre67grg-efw",
      "owner": "admin",
      "group": "manager",
      "description": "some desc",
      "tags": ["linux", "64bit", "centos"]
    }

+ Response 202 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "id": "sdf83kjfe-23r4wedf",
        }
      ]
    }
  ]

### Modify Disk Image [PUT]

+ Request Modify Disk Image (application/json)

  + Attributes

    + name (string, optional) - image name
    + owner (string, optional) - image creater
    + group (string, optional) - image group
    + description (string, optional) - image description
    + tags (array[string], optional) - image tags

  + Body

    {
      "name": "centos7_64_minimal",
      "owner": "admin",
      "group": "manager",
      "description": "some desc",
      "tags": ["linux", "64bit", "centos"]
    }

### Delete Disk Image [DELETE]

# System Maintain

## Migration [/migrations/{id}]

+ Parameters

  + id (string, optional) - migration task id

### Start Migration [POST]

+ Request Start Migration (application/json)

  + Attributes

    + source_pool (string) - pool of instance host
    + source_cell (string) - cell of instance host
    + target_pool (string, optional) - target pool you want to migrate
    + target_cell (string, optional) - target cell you want to migrate
    + instances (array[string], optional) - UUID list of instances need to migrate

  + Body

    {
      "source_pool": "default",
      "source_cell": "cell1",
      "target_pool": "another-pool",
      "target_cell": "new-cell",
      "instances": ["jhfguf85-34uyf-4t4tghjh", "dfjui-euy37-dyuy3t5qn"]
    }

+ Response 202 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "id": "abcdef-1234567890"
        }
      ]
    }
  ]

### Query Migration [GET]

+ Response 201 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "finished": false,
        "progress": 57 //limit to 100
      }
    }
  ]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "finished": true
      }
    }
  ]

### List All Migration in progress [GET /migrations/]


+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
            "id": "1234567890-abcdef",
            "finished": false,
            "progress": 57 //limit to 100
        },
        {
            "id": "1234567890-abcdef",
            "finished": true
        }
      ]
    }
  ]

## Batch Creating Guest [/batch/create_guest/{id}]

+ Parameters
  + id (string, optional) - id of batch task

### Start Creation [POST]

+ Request Start Batch Creating (application/json)

  + Attributes
    + name_rule: 'order' (enum[string]) - rule of generating guest name
      + Members
        + 'order' - incremental digital by order, etc 'xxx_0', 'xxx_1'
        + 'MAC' - by guest MAC address
        + 'address' - by assigned guest address
    + name_prefix (string, optional) - guest name prefix
    + owner (string) - user id
    + group (string) - user group
    + pool (string) - compute pool name
    + count (number) - instance count
    + cores (number) - vm cpu cores
    + memory (number) - vm memory
    + disks (array[number]) - [system_disk_size, data_disk_n_size...]
    + auto_start (boolean) - vm auto_start
    + from_image (string, optional) - initial disk image id
    + system_version (string, optional) - os version, 'centos7', 'win2012'
    + modules (array[string], optional) - installed guest module, "qemu"/"cloud-init"
    + cloud_init (object, optional) - ci params
      - root_enabled (boolean, optional) - allow root user login via SSH, enable in default
      - admin_name (string, optional) - specify admin name, using 'root' when omitted
      - admin_secret (string, optional) - password in plain text, nano generate a new one when omitted
      - data_path (string, optional) - data disk mount path, '/opt/data' in default

  + Body

    {
      "name_rule": "order",
      "name_prefix": "test",
      "owner": "admin",
      "group": "manager",
      "pool": "pool_1",
      "count": 3,
      "cores": 8,
      "memory": 4048,
      "disks": [4048, 38443],
      "auto_start": true,
      "system": "linux",
      "system_version": "centos7",
      "modules": ["qemu", "cloud-init"],
      "cloud_init":{
        "root_enabled": true,
        "admin_name": "nano",
        "admin_secret": "12345678",
        "data_path": "/opt/data"
      }
    }

+ Response 202 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "id": "dsfds8979847r3dsf-3r67",
      }
    }
  ]

### Query Batch Status [GET]

+ Response 202 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "name": "test_1",
          "id": "0123456789-abcdef-1",
          "progress": 45,
          "status": "creating"
        },
        {
          "name": "test_2",
          "id": "0123456789-abcdef-2",
          "status": "fail"
          "error": "not enough space"
        },
        {
          "name": "test_3",
          "id": "0123456789-abcdef-3",
          "status": "created"
        }
      ]
    }
  ]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "name": "test_1",
          "id": "0123456789-abcdef-1",
          "status": "created"
        },
        {
          "name": "test_2",
          "id": "0123456789-abcdef-2",
          "status": "fail",
          "error": "not enough space"
        },
        {
          "name": "test_3",
          "id": "0123456789-abcdef-3",
          "status": "created"
        }
      ]
    }
  ]

## Batch Deleting Guest [/batch/delete_guest/{id}]

+ Parameters
  + id (string, optional) - id of batch task

### Start Batch Deleting [POST]

+ Request Start Batch Deleting (application/json)

  + Attributes
    + guest (array[string]) - list of target guest UUID
    + force (boolean, optional) - stop running instance for deleting

  + Body

    {
      "guest": [
        "0123456789-abcdef-1",
        "0123456789-abcdef-2",
        "0123456789-abcdef-3"
      ],
      "force": false
    }

+ Response 202 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "id": "dsfds8979847r3dsf-3r67",
      }
    }
  ]

### Query Batch Status [GET]

+ Response 202 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "id": "0123456789-abcdef-1",
          "status": "deleted"
        },
        {
          "id": "0123456789-abcdef-2",
          "status": "fail",
          "error": "invalid guest"
        },
        {
          "id": "0123456789-abcdef-3",
          "status": "deleting"
        }
      ]
    }
  ]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
      {
        "id": "0123456789-abcdef-1",
        "status": "deleted"
      },
      {
        "id": "0123456789-abcdef-2",
        "status": "fail",
        "error": "invalid guest"
      },
      {
        "id": "0123456789-abcdef-3",
        "status": "deleted"
      }
      ]
    }
  ]

# Group FrontEnd Management

## Role [/roles/{role_name}]

+ Parameters
  + role_name (string, optional) - name of user role

### List Roles [GET /roles/]

list all use roles

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        "maintainer",
        "auditor",
        "user"
      ]
    }
  ]

### Get Role [GET]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "menu": ["dashboard", "compute_pool", "address_pool", "instance"],
        "resource": []
      }
    }
  ]

### Add Role [POST]

add new to system

+ Request Add Role (application/json)

  + Attributes
    + menu (array[string]) - menu list role can access

  + Body

    {
      "menu": ["dashboard", "compute_pool", "address_pool", "instance"]
    }


### Modify Role [PUT]

modify current authentication

+ Request Modify Role (application/json)

  + Attributes
    + menu (array[string]) - menu list role can access

  + Body

    {
      "menu": ["dashboard", "compute_pool", "address_pool", "instance"]
    }

### Remove Role [DELETE]

## User Group [/user_groups/{group_name}]

+ Parameters
  + group_name (string, optional) - name of user group

### List Groups [GET /user_groups/]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "name": "manager_group",
          "display": "Group of Manager",
          "member": 3
        },
        {
          "name": "auditor_group",
          "display": "Group of Auditor",
          "member": 1
        },
        {
          "name": "user_group",
          "display": "Group of User",
          "member": 10
        }
      ]
    }
  ]

### Get Group [GET]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "name": "manager_group",
        "display": "Group of Manager",
        "role": ["manager", "user"],
        "member": [
          "admin",
          "nano",
          "akumas"
        ]
      }
    }
  ]

### Add Group [POST]

+ Request Add Group (application/json)

  + Attributes
    + display (string) - display name of group
    + role (array[string], optional) - list of role name

  + Body

    {
      "display": "Group of Manager",
      "role": ["manager", "user"]
    }

### Modify Group [PUT]

+ Request Modify Group (application/json)

  + Attributes
    + display (string, optional) - display name of group
    + role (array[string], optional) - list of role name

  + Body

    {
      "display": "Group of Manager",
      "role": ["manager", "user"]
    }

### Remove Group [DELETE]

## Group Member [/user_groups/{group_name}/members/{user_name}]

+ Parameters
  + group_name (string) - name of user group
  + user_name (string, optional) - user name

### List Member [GET /user_groups/{group_name}/members/]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        "admin",
        "nano",
        "akumas"
      ]
    }
  ]

### Add Member [POST]

### Remove Member [DELETE]

## User [/users/{user_name}]

+ Parameters
  + user_name (string, optional) - user name

### List Users [GET /users/]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        "admin",
        "nano",
        "akumas"
      ]
    }
  ]

### Get User Info [GET]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "nick": "AK",
        "mail": "a@b.com"
      }
    }
  ]

### Create New User [POST]

+ Request Create New user (application/json)

  + Attributes
    + nick (string, optional) - nick name for display
    + mail (string, optional) - user email
    + password (string) - user password

  + Body

    {
      "nick": "AK",
      "mail": "a@b.com",
      "password": "abcdefg"
    }

### Modify User [PUT]

+ Request Modify user (application/json)

  + Attributes
    + nick (string, optional) - nick name for display
    + mail (string, optional) - user email

  + Body

    {
      "nick": "AK",
      "mail": "a@b.com"
    }

### Delete User [DELETE]

## User Password [/users/{user_name}/password/]

+ Parameters
  + user_name (string, optional) - user name

### Modify Password [PUT]

+ Request Modify Password (application/json)

  + Attributes
    + old (string) - old password for verify
    + new (string) - new password

  + Body

    {
      "old": "123456",
      "new": "abcdef"
    }

## User Search [/user_search/]

### Search User [GET]

+ Parameters

  + group (string, optional) - user group

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        "nano",
        "alex",
        "sam"
        ]
      }
    ]

## Sessions [/sessions/{session_id}]

Manage authenticated session

+ Parameters
  + session_id (string, optional) - session ID

### List session [GET /sessions/]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        "bsu6384js3=",
        "bsu633jtq84",
        "bsu638skjfu"
      ]
    }
  ]


### Auth Session [POST /sessions/]

+ Request Auth New Session (application/json)

  + Attributes
    + user (string) - user name
    + password (string) - new password
    + nonce (string) - nonce for crypt

  + Body

    {
      "user": "alex",
      "password": "abcdef",
      "nonce": "1234567890"
    }

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "session": "bsu6384js3=",
        "timeout": 120
      }
    }
  ]

### Update Session [PUT]

### Obtain Session Authentication [GET]

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":{
        "menu": ["dashboard", "compute_pool", "address_pool", "instance"],
        "resource": [],
        "user": "nano"
      }
    }
  ]

## Logs [/logs/]

### Query Log Entry [GET]

+ Parameters

  + limit (number) - max entries returned
  + start (number, optional) - start offset
  + after (string, optional) - query logs after this date
  + before (string, optional) - query logs before this date

+ Response 200 (application/json)

  [
    {
      "error_code": 0,
      "message": "",
      "data":[
        {
          "time": "2018-12-01 14:57:39",
          "content": "admin create new user example"
        },
        {
          "time": "2018-12-03 14:27:19",
          "content": "nano.sabrina create new guest nano.test1"
        }

      ]
    }
  ]

### Add Log Entry [POST]

+ Request Add Log (application/json)

  + Attributes
    + format (string, optional) - format of content
    + content (string) - log content

  + Body

    {
      "format": "string",
      "content": "here is a new log for somebody creating a new instance"
    }
