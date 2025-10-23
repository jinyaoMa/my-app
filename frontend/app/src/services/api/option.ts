// @ts-ignore
/* eslint-disable */
import request from '@/utils/request'

/** Delete Option Delete option by id. 返回值: Error DELETE /api/option/delete/${param0} */
export async function optionDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.optionDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/option/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Option Detail Get option detail by id, and users can include associated entities by passing includes. 返回值: Error POST /api/option/detail/${param0} */
export async function optionPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.optionPostDetailParams,
  body: any,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/option/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query Option List Query option and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/option/query */
export async function optionPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/option/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Recover Option Recover option by id. 返回值: Error PUT /api/option/recover/${param0} */
export async function optionPutRecover(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.optionPutRecoverParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ErrorModel>(`/api/option/recover/${param0}`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get Option Recycle List Get deleted option and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/option/recycle */
export async function optionPostRecycle(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/option/recycle', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save Option Save option, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error PUT /api/option/save */
export async function optionPutSave(body: API.OptionSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/option/save', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Search Option Search option by given keyword and fields, and limit the result count. 返回值: Error POST /api/option/search */
export async function optionPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/option/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
