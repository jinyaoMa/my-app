// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Batch Delete File Extension Batch delete file extension by ids. DELETE /api/file-extension/batch/delete */
export async function fileExtensionDeleteBatchDelete(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/file-extension/batch/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Batch Recover File Extension Batch recover file extension by ids. PUT /api/file-extension/batch/recover */
export async function fileExtensionPutBatchRecover(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/file-extension/batch/recover', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Delete File Extension Delete file extension by id. DELETE /api/file-extension/delete/${param0} */
export async function fileExtensionDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileExtensionDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/file-extension/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get File Extension Detail Get file extension detail by id, and users can include associated entities by passing includes. POST /api/file-extension/detail/${param0} */
export async function fileExtensionPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileExtensionPostDetailParams,
  body: API.Include[],
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseFileExtensionDetailBody>(`/api/file-extension/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query File Extension List Query file extension and present results in list, and users can filter, sort and page data by passing criteria. POST /api/file-extension/query */
export async function fileExtensionPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListFileExtensionItemBody>('/api/file-extension/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Recover File Extension Recover file extension by id. PUT /api/file-extension/recover/${param0} */
export async function fileExtensionPutRecover(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileExtensionPutRecoverParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/file-extension/recover/${param0}`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get File Extension Recycle List Get deleted file extension and present results in list, and users can filter, sort and page data by passing criteria. POST /api/file-extension/recycle */
export async function fileExtensionPostRecycle(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListFileExtensionItemBody>('/api/file-extension/recycle', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save File Extension Save file extension, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. PUT /api/file-extension/save */
export async function fileExtensionPutSave(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileExtensionPutSaveParams,
  body: API.FileExtensionSave,
  options?: { [key: string]: any }
) {
  return request<API.ResponseFileExtensionSaveBody>('/api/file-extension/save', {
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

/** Search File Extension Search file extension by given keyword and fields, and limit the result count. POST /api/file-extension/search */
export async function fileExtensionPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ResponseListFileExtensionSearchItemBody>('/api/file-extension/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
