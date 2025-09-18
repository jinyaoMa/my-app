// @ts-ignore
/* eslint-disable */
import { request } from '../utils'

/** detail permission Detail permission by id, and users can include associated entities by passing includes. 返回值: Error GET /api/permission */
export async function permissionDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.permissionDetailParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/permission', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** save permission Save permission, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error POST /api/permission */
export async function permissionSave(body: API.PermissionSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/permission', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** delete permission Delete permission by id. 返回值: Error DELETE /api/permission */
export async function permissionDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.permissionDeleteParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/permission', {
    method: 'DELETE',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** query permission list Query permission and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/permission/query */
export async function permissionQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/permission/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
