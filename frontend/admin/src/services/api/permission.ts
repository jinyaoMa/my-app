// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Batch Delete Permission Batch delete permission by ids. DELETE /api/permission/batch/delete */
export async function permissionDeleteBatchDelete(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/permission/batch/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Batch Recover Permission Batch recover permission by ids. PUT /api/permission/batch/recover */
export async function permissionPutBatchRecover(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/permission/batch/recover', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Delete Permission Delete permission by id. DELETE /api/permission/delete/${param0} */
export async function permissionDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.permissionDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/permission/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Permission Detail Get permission detail by id, and users can include associated entities by passing includes. POST /api/permission/detail/${param0} */
export async function permissionPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.permissionPostDetailParams,
  body: API.Include[],
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponsePermissionDetailBody>(`/api/permission/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Permission Has Field Value Check Permission if the given field with the given value exists. GET /api/permission/has */
export async function permissionGetHas(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.permissionGetHasParams,
  options?: { [key: string]: any }
) {
  return request<API.ResponseBoolBody>('/api/permission/has', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** Query Permission List Query permission and present results in list, and users can filter, sort and page data by passing criteria. POST /api/permission/query */
export async function permissionPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListPermissionItemBody>('/api/permission/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Recover Permission Recover permission by id. PUT /api/permission/recover/${param0} */
export async function permissionPutRecover(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.permissionPutRecoverParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/permission/recover/${param0}`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Permission Recycle List Get deleted permission and present results in list, and users can filter, sort and page data by passing criteria. POST /api/permission/recycle */
export async function permissionPostRecycle(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListPermissionItemBody>('/api/permission/recycle', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save Permission Save permission, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. PUT /api/permission/save */
export async function permissionPutSave(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.permissionPutSaveParams,
  body: API.PermissionSave,
  options?: { [key: string]: any }
) {
  return request<API.ResponsePermissionSaveBody>('/api/permission/save', {
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

/** Search Permission Search permission by given keyword and fields, and limit the result count. POST /api/permission/search */
export async function permissionPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ResponseListPermissionSearchItemBody>('/api/permission/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
