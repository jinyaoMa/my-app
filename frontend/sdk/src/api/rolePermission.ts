// @ts-ignore
/* eslint-disable */
import { request } from '../utils'

/** detail role permission Detail role permission by id, and users can include associated entities by passing includes. 返回值: Error GET /api/role-permission */
export async function rolePermissionDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.rolePermissionDetailParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/role-permission', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** save role permission Save role permission, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error POST /api/role-permission */
export async function rolePermissionSave(body: API.RolePermissionSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/role-permission', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** delete role permission Delete role permission by id. 返回值: Error DELETE /api/role-permission */
export async function rolePermissionDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.rolePermissionDeleteParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/role-permission', {
    method: 'DELETE',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** query role permission list Query role permission and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/role-permission/query */
export async function rolePermissionQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/role-permission/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
