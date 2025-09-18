// @ts-ignore
/* eslint-disable */
import { request } from '../utils'

/** list endpoints list all registered endpoints, filtered by tag, operation id, etc. 返回值: Error GET /api/endpoints/list */
export async function endpointsList(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.endpointsListParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/endpoints/list', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** list endpoints' tags 返回值: Error GET /api/endpoints/tags */
export async function endpointsTags(options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/endpoints/tags', {
    method: 'GET',
    ...(options || {})
  })
}
