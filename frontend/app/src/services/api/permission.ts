// @ts-ignore
/* eslint-disable */
import request from '@/utils/request'

/** Delete Permission Delete permission by id. 返回值: Error DELETE /api/permission/delete/${param0} */
export async function permissionDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.permissionDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/permission/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Permission Detail Get permission detail by id, and users can include associated entities by passing includes. 返回值: Error POST /api/permission/detail/${param0} */
export async function permissionPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.permissionPostDetailParams,
  body: any,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/permission/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query Permission List Query permission and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/permission/query */
export async function permissionPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/permission/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Recover Permission Recover permission by id. 返回值: Error PUT /api/permission/recover/${param0} */
export async function permissionPutRecover(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.permissionPutRecoverParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/permission/recover/${param0}`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Permission Recycle List Get deleted permission and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/permission/recycle */
export async function permissionPostRecycle(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/permission/recycle', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save Permission Save permission, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error PUT /api/permission/save */
export async function permissionPutSave(body: API.PermissionSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/permission/save', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Search Permission Search permission by given keyword and fields, and limit the result count. 返回值: Error POST /api/permission/search */
export async function permissionPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/permission/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
