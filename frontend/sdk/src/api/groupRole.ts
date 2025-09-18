// @ts-ignore
/* eslint-disable */
import { request } from '../utils'

/** detail group role Detail group role by id, and users can include associated entities by passing includes. 返回值: Error GET /api/group-role */
export async function groupRoleDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupRoleDetailParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/group-role', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** save group role Save group role, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error POST /api/group-role */
export async function groupRoleSave(body: API.GroupRoleSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/group-role', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** delete group role Delete group role by id. 返回值: Error DELETE /api/group-role */
export async function groupRoleDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupRoleDeleteParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/group-role', {
    method: 'DELETE',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** query group role list Query group role and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/group-role/query */
export async function groupRoleQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/group-role/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
