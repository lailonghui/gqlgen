# gqlgen
GraphQL的前后端交互


前端请求方式：
fetch('http://localhost:8080/query',{
    method:'post',
   headers: {'Content-Type': 'application/json'},
    body: JSON.stringify({
    "operationName": "getEnterpriseList",
    "variables": {},
    "query": `query getEnterpriseList {
              EnterpriseInfoList {
                id
                enterprise_name
                enterprise_id
              }
            }
        `
})
}).then(res => res.json()).then(res => {
    console.log(res)
})

返回数据:
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
            {
                "id": 2419493307836007435,
                "enterprise_name": "阿里巴巴",
                "enterprise_id": ""
            },
            {
                "id": 2419494486435431436,
                "enterprise_name": "腾讯",
                "enterprise_id": "bu371o16a63nes7iuegg"
            }
        ]
    }
}

