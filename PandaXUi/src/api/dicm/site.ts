import request from "@/utils/request";

// 查询机柜列表
export function listSites(query:any) {
    return request({
        url: '/dicm/sites/list',
        method: 'get',
        params: query
    })
}