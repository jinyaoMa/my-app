// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Batch Delete Group Batch delete group by ids. DELETE /api/group/batch/delete */
export async function groupDeleteBatchDelete(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/group/batch/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Batch Recover Group Batch recover group by ids. PUT /api/group/batch/recover */
export async function groupPutBatchRecover(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/group/batch/recover', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Delete Group Delete group by id. DELETE /api/group/delete/${param0} */
export async function groupDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/group/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Group Detail Get group detail by id, and users can include associated entities by passing includes. POST /api/group/detail/${param0} */
export async function groupPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupPostDetailParams,
  body: API.Include[],
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseGroupDetailBody>(`/api/group/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query Group List Query group and present results in list, and users can filter, sort and page data by passing criteria. POST /api/group/query */
export async function groupPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListGroupItemBody>('/api/group/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Recover Group Recover group by id. PUT /api/group/recover/${param0} */
export async function groupPutRecover(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupPutRecoverParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/group/recover/${param0}`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Group Recycle List Get deleted group and present results in list, and users can filter, sort and page data by passing criteria. POST /api/group/recycle */
export async function groupPostRecycle(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListGroupItemBody>('/api/group/recycle', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save Group Save group, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. PUT /api/group/save */
export async function groupPutSave(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.groupPutSaveParams,
  body: API.GroupSave,
  options?: { [key: string]: any }
) {
  return request<API.ResponseGroupSaveBody>('/api/group/save', {
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

/** Search Group Search group by given keyword and fields, and limit the result count. POST /api/group/search */
export async function groupPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ResponseListGroupSearchItemBody>('/api/group/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
