// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Batch Delete Option Batch delete option by ids. DELETE /api/option/batch/delete */
export async function optionDeleteBatchDelete(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/option/batch/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Batch Recover Option Batch recover option by ids. PUT /api/option/batch/recover */
export async function optionPutBatchRecover(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/option/batch/recover', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Delete Option Delete option by id. DELETE /api/option/delete/${param0} */
export async function optionDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.optionDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/option/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Option Detail Get option detail by id, and users can include associated entities by passing includes. POST /api/option/detail/${param0} */
export async function optionPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.optionPostDetailParams,
  body: API.Include[],
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseOptionDetailBody>(`/api/option/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query Option List Query option and present results in list, and users can filter, sort and page data by passing criteria. POST /api/option/query */
export async function optionPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListOptionItemBody>('/api/option/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Recover Option Recover option by id. PUT /api/option/recover/${param0} */
export async function optionPutRecover(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.optionPutRecoverParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/option/recover/${param0}`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Option Recycle List Get deleted option and present results in list, and users can filter, sort and page data by passing criteria. POST /api/option/recycle */
export async function optionPostRecycle(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListOptionItemBody>('/api/option/recycle', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save Option Save option, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. PUT /api/option/save */
export async function optionPutSave(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.optionPutSaveParams,
  body: API.OptionSave,
  options?: { [key: string]: any }
) {
  return request<API.ResponseOptionSaveBody>('/api/option/save', {
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

/** Search Option Search option by given keyword and fields, and limit the result count. POST /api/option/search */
export async function optionPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ResponseListOptionSearchItemBody>('/api/option/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
