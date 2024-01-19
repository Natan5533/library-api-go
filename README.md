# library-api-go WIP 

TODO: context and arch example

## Run localy 
### Specifications:
```tool
 - Golang 1.21.0
 - Docker engine 
```
---
### In the project, run the following commands:

 ```shell
$ cd docker; docker compose up -d; cd ..
$ go run main.go 
```
---
### Make requests!
> Examples performed at the Postman.

[<img src="https://run.pstmn.io/button.svg" alt="Test In Postman" style="width: 100px; height: 25px;">](https://god.gw.postman.com/run-collection/21572659-5e64f842-2b0c-4e0f-835a-5e3f2a600d49?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D21572659-5e64f842-2b0c-4e0f-835a-5e3f2a600d49%26entityType%3Dcollection%26workspaceId%3D92295a17-efe7-41a0-bb69-ce7642a21c31)

#### Create a library: 
URL:

 `http://localhost:8080/api/v1/library`

JSON Body: 
``` JSON
{
        "name": "Kalunga",
        "address": "Guarapiranga Avenue"
}
```
Status Response:

`201`

Body Reponse: 
``` JSON
{"id":1}
```

#### Get by ID:
URL: 

`http://localhost:8080/api/v1/library/:id `

Path variable: `:id == 1`

Response status: `200`

Response body: 
> Authors were previously entered using another API route ;)
``` JSON
{
    "library": {
        "id": 1,
        "name": "Kalunga",
        "address": "Guarapiranga Avenue",
        "authors": [
            {
                "id": 1,
                "name": "Chico Moedas"
            },
            {
                "id": 2,
                "name": "Chiquinho"
            }
        ]
    }
}
```

---

> I'm too lazy to add the other routes so good luck discovering it, see Gin-debug output to verify all routes available.




### Draft
![image](https://github.com/Natan5533/library-api-go/assets/86797382/01f63370-756a-4779-ac10-efeb899fde1e)

![image](https://github.com/Natan5533/library-api-go/assets/86797382/6946344b-fbe1-470b-ad67-7c1d37de9645)

