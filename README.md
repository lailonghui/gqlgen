# gqlgen
GraphQL的前后端交互:

##### 1.请求示例 Request Example

```js

 fetch("http://localhost:8080/query?account_id=1&access_token=1",{ 
    "method":"POST", 
    "headers": {"Content-Type": "application/json"}, 
    "body": JSON.stringify({ 
         "query": `
            query getDistrictVehicleList {
  	            DistrictVehicleList(page:1,num:10,sort:"newest"){
  			        district_name,
                     business_scope_info_list: {
					 business_scope,
					 number
                     }
		        }
            } 
        ` }) 
}).then(res => res.json()).then(res => console.log(res))
```

##### 2. 成功返回结果示例 Success Response

```json
{
    "data": {
        "EnterpriseInfoList": [
            {
                "id": 2419483492216734728,
                "enterprise_name": "辉哥科技有限公司",
                "enterprise_id": "bu36ng96a63jq17l7iq0"
            },
            {
                "id": 2419483726518944777,
                "enterprise_name": "鲜动智能科技有限公司",
                "enterprise_id": "bu36nn96a63rob5n98r0"
            },
            {
                "id": 2419492126619337738,
                "enterprise_name": "小灰灰科技有限公司",
                "enterprise_id": "bu36vhp6a63onj5l4l40"
            },
          
        ]
    }
}



```


