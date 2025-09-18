// @ts-ignore
/* eslint-disable */
import { request } from '../utils'

/** detail role Detail role by id, and users can include associated entities by passing includes. 返回值: Error GET /api/role */
export async function roleDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.roleDetailParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/role', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** save role Save role, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error POST /api/role */
export async function roleSave(body: API.RoleSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/role', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** delete role Delete role by id. 返回值: Error DELETE /api/role */
export async function roleDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.roleDeleteParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/role', {
    method: 'DELETE',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** query role list Query role and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/role/query */
export async function roleQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/role/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
