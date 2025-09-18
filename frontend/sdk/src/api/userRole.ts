// @ts-ignore
/* eslint-disable */
import { request } from '../utils'

/** detail user role Detail user role by id, and users can include associated entities by passing includes. 返回值: Error GET /api/user-role */
export async function userRoleDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.userRoleDetailParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/user-role', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** save user role Save user role, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error POST /api/user-role */
export async function userRoleSave(body: API.UserRoleSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/user-role', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** delete user role Delete user role by id. 返回值: Error DELETE /api/user-role */
export async function userRoleDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.userRoleDeleteParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/user-role', {
    method: 'DELETE',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** query user role list Query user role and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/user-role/query */
export async function userRoleQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/user-role/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
