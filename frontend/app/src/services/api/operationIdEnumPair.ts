// @ts-ignore
/* eslint-disable */
import request from '@/utils/request'

/** Get Pair List From Flag Get operation id enum pair list from flag (base64). 返回值: Error POST /api/operation-id-enum-pair/from-flag */
export async function operationIdEnumPairPostFromFlag(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.operationIdEnumPairPostFromFlagParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/operation-id-enum-pair/from-flag', {
    method: 'POST',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** Get Pair List Get operation id enum pair list, filtered by operation id. 返回值: Error GET /api/operation-id-enum-pair/list */
export async function operationIdEnumPairGetList(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.operationIdEnumPairGetListParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/operation-id-enum-pair/list', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** Convert Enums to Flag Convert enums of operation ids/roles to base64 flag. 返回值: Error POST /api/operation-id-enum-pair/to-flag */
export async function operationIdEnumPairPostToFlag(body: any, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/operation-id-enum-pair/to-flag', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
