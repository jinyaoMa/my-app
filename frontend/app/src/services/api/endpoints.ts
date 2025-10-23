// @ts-ignore
/* eslint-disable */
import request from '@/utils/request'

/** List Endpoints List all registered endpoints, filtered by tag, operation id, etc. 返回值: Error GET /api/endpoints/list */
export async function endpointsGetList(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.endpointsGetListParams,
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

/** Get Tags of Endpoints Get all tags of registered endpoints. 返回值: Error GET /api/endpoints/tags */
export async function endpointsGetTags(options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/endpoints/tags', {
    method: 'GET',
    ...(options || {})
  })
}
