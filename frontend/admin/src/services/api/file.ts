// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Batch Delete File Batch delete file by ids. DELETE /api/file/batch/delete */
export async function fileDeleteBatchDelete(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/file/batch/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Batch Recover File Batch recover file by ids. PUT /api/file/batch/recover */
export async function filePutBatchRecover(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/file/batch/recover', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Get File Data Get file data by id. GET /api/file/data/${param0} */
export async function fileGetData(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileGetDataParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<any>(`/api/file/data/${param0}`, {
    method: 'GET',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Delete File Delete file by id. DELETE /api/file/delete/${param0} */
export async function fileDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/file/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get File Detail Get file detail by id, and users can include associated entities by passing includes. POST /api/file/detail/${param0} */
export async function filePostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.filePostDetailParams,
  body: API.Include[],
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseFileDetailBody>(`/api/file/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query File List Query file and present results in list, and users can filter, sort and page data by passing criteria. POST /api/file/query */
export async function filePostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListFileItemBody>('/api/file/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Recover File Recover file by id. PUT /api/file/recover/${param0} */
export async function filePutRecover(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.filePutRecoverParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/file/recover/${param0}`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get File Recycle List Get deleted file and present results in list, and users can filter, sort and page data by passing criteria. POST /api/file/recycle */
export async function filePostRecycle(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListFileItemBody>('/api/file/recycle', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save File Save file, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. PUT /api/file/save */
export async function filePutSave(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.filePutSaveParams,
  body: API.FileSave,
  options?: { [key: string]: any }
) {
  return request<API.ResponseFileSaveBody>('/api/file/save', {
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

/** Search File Search file by given keyword and fields, and limit the result count. POST /api/file/search */
export async function filePostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ResponseListFileSearchItemBody>('/api/file/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
