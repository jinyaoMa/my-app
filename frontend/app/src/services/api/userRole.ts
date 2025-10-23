// @ts-ignore
/* eslint-disable */
import request from '@/utils/request'

/** Delete User Role Delete user role by id. 返回值: Error DELETE /api/user-role/delete/${param0} */
export async function userRoleDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.userRoleDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/user-role/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get User Role Detail Get user role detail by id, and users can include associated entities by passing includes. 返回值: Error POST /api/user-role/detail/${param0} */
export async function userRolePostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.userRolePostDetailParams,
  body: any,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/user-role/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query User Role List Query user role and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/user-role/query */
export async function userRolePostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/user-role/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save User Role Save user role, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error PUT /api/user-role/save */
export async function userRolePutSave(body: API.UserRoleSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/user-role/save', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Search User Role Search user role by given keyword and fields, and limit the result count. 返回值: Error POST /api/user-role/search */
export async function userRolePostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/user-role/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
