// @ts-ignore
/* eslint-disable */
import request from '@/utils/request'

/** Delete Role Permission Delete role permission by id. 返回值: Error DELETE /api/role-permission/delete/${param0} */
export async function rolePermissionDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.rolePermissionDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/role-permission/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Role Permission Detail Get role permission detail by id, and users can include associated entities by passing includes. 返回值: Error POST /api/role-permission/detail/${param0} */
export async function rolePermissionPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.rolePermissionPostDetailParams,
  body: any,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/role-permission/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query Role Permission List Query role permission and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/role-permission/query */
export async function rolePermissionPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/role-permission/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save Role Permission Save role permission, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error PUT /api/role-permission/save */
export async function rolePermissionPutSave(body: API.RolePermissionSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/role-permission/save', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Search Role Permission Search role permission by given keyword and fields, and limit the result count. 返回值: Error POST /api/role-permission/search */
export async function rolePermissionPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/role-permission/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
