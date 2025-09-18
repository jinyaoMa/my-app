// @ts-ignore
/* eslint-disable */
import { request } from '../utils'

/** detail user Detail user by id, and users can include associated entities by passing includes. 返回值: Error GET /api/user */
export async function userDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.userDetailParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/user', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** save user Save user, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error POST /api/user */
export async function userSave(body: API.UserSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/user', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** delete user Delete user by id. 返回值: Error DELETE /api/user */
export async function userDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.userDeleteParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/user', {
    method: 'DELETE',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** query user list Query user and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/user/query */
export async function userQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/user/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
