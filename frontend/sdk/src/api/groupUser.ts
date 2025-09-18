// @ts-ignore
/* eslint-disable */
import { request } from '../utils'

/** detail group user Detail group user by id, and users can include associated entities by passing includes. 返回值: Error GET /api/group-user */
export async function groupUserDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupUserDetailParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/group-user', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** save group user Save group user, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error POST /api/group-user */
export async function groupUserSave(body: API.GroupUserSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/group-user', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** delete group user Delete group user by id. 返回值: Error DELETE /api/group-user */
export async function groupUserDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupUserDeleteParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/group-user', {
    method: 'DELETE',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** query group user list Query group user and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/group-user/query */
export async function groupUserQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/group-user/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
