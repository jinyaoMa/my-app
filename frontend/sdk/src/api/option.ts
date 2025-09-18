// @ts-ignore
/* eslint-disable */
import { request } from '../utils'

/** detail option Detail option by id, and users can include associated entities by passing includes. 返回值: Error GET /api/option */
export async function optionDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.optionDetailParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/option', {
    method: 'GET',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** save option Save option, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. 返回值: Error POST /api/option */
export async function optionSave(body: API.OptionSave, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/option', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** delete option Delete option by id. 返回值: Error DELETE /api/option */
export async function optionDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.optionDeleteParams,
  options?: { [key: string]: any }
) {
  return request<API.ErrorModel>('/api/option', {
    method: 'DELETE',
    params: {
      ...params
    },
    ...(options || {})
  })
}

/** query option list Query option and present results in list, and users can filter, sort and page data by passing criteria. 返回值: Error POST /api/option/query */
export async function optionQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ErrorModel>('/api/option/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
