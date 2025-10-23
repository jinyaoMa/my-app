// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** List Endpoints List all registered endpoints, filtered by tag, operation id, etc. GET /api/endpoints/list */
export async function endpointsGetList(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.endpointsGetListParams,
  options?: { [key: string]: any }
) {
  return request<API.ResponseListEndpointsItemBody>('/api/endpoints/list', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** Get Tags of Endpoints Get all tags of registered endpoints. GET /api/endpoints/tags */
export async function endpointsGetTags(options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/endpoints/tags', {
    method: 'GET',
    ...(options || {})
  })
}
