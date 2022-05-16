# stillasTracker

stillasTracker is a software solution which provides MB-Stillas with an API, a database and two-front end solutions.
The readme will serve as endpoint documentation for the stillastracker API as well as a brief user guide for both of the front end solutions in the repository

## Endpoint-Documentation:
The following endpoints are available in the current version of the API:

### Scaffolding (stillasdel) endpoints

The scaffolding endpoint has GET, POST, PUT and DELETE functionality. Get requests are used 
get information regarding scaffolding parts, you can get scaffolding parts based on the queries listed
in the table below.

**GET requests**

The GET requests fetches scaffolding parts from the database
```
/stillastracking/v1/api/unit?type={:type}&id={:id}
/stillastracking/v1/api/unit?type={:type}
/stillastracking/v1/api/unit
```
**POST and PUT requests**

The post and put requests are used to add in new scaffolding parts or update them.
```
/stillastracking/v1/api/unit

Example body:
{
    "id": "AB23GW",
    "type": "Flooring",
    "project": "CCHamar",
    "batteryLevel": 100
}
```

**DELETE**

The delete request deletes scaffolding units from the database
```
/stillastracking/v1/api/unit

Example body:
[
    {
        "id": "4A6352"
    },
    {
        "id": "56GWAG"
    },
    {
        "id": "8GTW21"
    },
    {
        "id": "LPW123"
    }
]
```


### Project (byggeprosjekt) endpoints
This endpoint handles all data regarding projects in the database.

**GET**

These endpoints are used to fetch information regarding projects from the database. Adding scaffolding=true 
after the endpoint will list all the scaffolding parts associated with the project. 
```
/stillastracking/v1/api/project?id={:id}&scaffolding={:scaffolding}
/stillastracking/v1/api/project?name={:name}&scaffolding={:scaffolding}
/stillastracking/v1/api/project?id={:id}
/stillastracking/v1/api/project?name={:name}
/stillastracking/v1/api/project?id={:id}
/stillastracking/v1/api/project?scaffolding={:scaffolding}
/stillastracking/v1/api/project
/stillastracking/v1/api/storage
```

**POST and PUT**

The following endpoint is used to add in new scaffoldingparts

```
/stillastracking/v1/api/project

Example body
{
    "projectID":2321112,
    "projectName":"MBStillas",
    "latitude":60.79077759591496,
    "longitude":10.683249543160402,
    "state":"Active",
    "size":322,
    "period":{
        "startDate":"25-04-2022",
        "endDate":"30-04-2022"
        },
    "customer":{
        "name":"Martin Ivers",
        "number":98435621,
        "email":"martin@mail.no"
        },
    "address":{
        "street":"Halsetsvea 40",
        "zipcode":"2323",
        "municipality":"Ingeberg",
        "county":"Innlandet"
        },
        "geofence":{
            "w-position":{"latitude":-73.98326396942211,"longitude":40.69287790858968},
            "x-position":{"latitude":-73.98387551307742,"longitude":40.6922433936175},
            "y-position":{"latitude":-73.98255586624245,"longitude":40.691999347788055},
            "z-position":{"latitude":-73.98124694824298,"longitude":40.69267453906423}
        }
    }
    
/stillastracking/v1/api/project/scaffold

Example body:
{
    "toProjectID":755,
    "fromProjectID":12,
    "scaffold":[{
        "type":"Bunnskrue",
        "quantity":1
    },]
}
```

**DELETE** 

The endpoint is used to delete projects

```
/stillastracking/v1/api/project

Example body:
[
    {
        "id": 430
    },
    {
        "id": 420
    }
]
```


### Profile (bruker) endpoints

The endpoint handles user creation,updates and removals. 

**GET**

The endpoints below fetches users from the database
```
/stillastracking/v1/api/user?id={:id}
/stillastracking/v1/api/user?role={:role}
/stillastracking/v1/api/user
```

**POST or PUT** 

The endpoint creates or updates users
```
/stillastracking/v1/api/user

Example body:
{
    "employeeID": 232,
    "name": "Ola Nordmann",
    "dateOfBirth": "01.04.1988",
    "role": "Storage",
    "admin": true
}
```

**DELETE**

The endpoint deletes users from the database
```
/stillastracking/v1/api/user

Example body:
[    
    {"id" : "12521"},
    {"id" : "12521"},
    {"id" : "12521"},
    {"id" : "12521"},
]
```

### Gateway endpoints
The endpoint handles all management of BLE Gateways in the database

**GET**

The following endpoints can be used to fetch gateways from the database
```
/stillastracking/v1/api/gateway?id={:id}
/stillastracking/v1/api/gateway?projectName={:projectName}
/stillastracking/v1/api/gateway?projectID={:projectID}
/stillastracking/v1/api/gateway
```

**POST or PUT**

The endpoint creates and updates gateways
```
/stillastracking/v1/api/gateway

Example body:
{
    "Status": true
    "gatewayID": "34AB954B54E4"
    "latitude": 59.911491
    "longitude": 10.757933
    "projectID": 4
    "projectName": "CCHamar"
}
```

**DELETE**

The endpoint deletes gateways 
```
/stillastracking/v1/api/gateway

Example body:
[    
    {"id" : "34AB954B54E4"},
    {"id" : "34AB954BABEE4"},
    {"id" : "34TQWD21SDAE4"},
    {"id" : "1241WADQWDQW4"},
]
```
