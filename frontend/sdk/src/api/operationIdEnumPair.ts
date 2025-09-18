// @ts-ignore
/* eslint-disable */
import { request } from '../utils'

/** list operation id enum pair list all registered operation id enum pair, filtered by operation id. 返回值: Error GET /api/operation-id-enum-pair/list */
export async function operationIdEnumPairList(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.operationIdEnumPairListParams,
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
