// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Delete User Password Delete user password by id. DELETE /api/user-password/delete/${param0} */
export async function userPasswordDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.userPasswordDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/user-password/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get User Password Detail Get user password detail by id, and users can include associated entities by passing includes. POST /api/user-password/detail/${param0} */
export async function userPasswordPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.userPasswordPostDetailParams,
  body: API.Include[],
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseUserPasswordDetailBody>(`/api/user-password/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query User Password List Query user password and present results in list, and users can filter, sort and page data by passing criteria. POST /api/user-password/query */
export async function userPasswordPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListUserPasswordItemBody>('/api/user-password/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save User Password Save user password, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. PUT /api/user-password/save */
export async function userPasswordPutSave(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.userPasswordPutSaveParams,
  body: API.UserPasswordSave,
  options?: { [key: string]: any }
) {
  return request<API.ResponseUserPasswordSaveBody>('/api/user-password/save', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    params: {
      ...params
    },
    data: body,
    ...(options || {})
  })
}

/** Search User Password Search user password by given keyword and fields, and limit the result count. POST /api/user-password/search */
export async function userPasswordPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ResponseListUserPasswordSearchItemBody>('/api/user-password/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
