// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max'

/** Batch Delete File Category Batch delete file category by ids. DELETE /api/file-category/batch/delete */
export async function fileCategoryDeleteBatchDelete(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/file-category/batch/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Batch Recover File Category Batch recover file category by ids. PUT /api/file-category/batch/recover */
export async function fileCategoryPutBatchRecover(body: string[], options?: { [key: string]: any }) {
  return request<API.ResponseListStringBody>('/api/file-category/batch/recover', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Delete File Category Delete file category by id. DELETE /api/file-category/delete/${param0} */
export async function fileCategoryDeleteDelete(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileCategoryDeleteDeleteParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/file-category/delete/${param0}`, {
    method: 'DELETE',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get File Category Detail Get file category detail by id, and users can include associated entities by passing includes. POST /api/file-category/detail/${param0} */
export async function fileCategoryPostDetail(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileCategoryPostDetailParams,
  body: API.Include[],
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseFileCategoryDetailBody>(`/api/file-category/detail/${param0}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    params: { ...queryParams },
    data: body,
    ...(options || {})
  })
}

/** Query File Category List Query file category and present results in list, and users can filter, sort and page data by passing criteria. POST /api/file-category/query */
export async function fileCategoryPostQuery(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListFileCategoryItemBody>('/api/file-category/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Recover File Category Recover file category by id. PUT /api/file-category/recover/${param0} */
export async function fileCategoryPutRecover(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileCategoryPutRecoverParams,
  options?: { [key: string]: any }
) {
  const { id: param0, ...queryParams } = params
  return request<API.ResponseStringBody>(`/api/file-category/recover/${param0}`, {
    method: 'PUT',
    params: { ...queryParams },
    ...(options || {})
  })
}

/** Get File Category Recycle List Get deleted file category and present results in list, and users can filter, sort and page data by passing criteria. POST /api/file-category/recycle */
export async function fileCategoryPostRecycle(body: API.Criteria, options?: { [key: string]: any }) {
  return request<API.ResponseListFileCategoryItemBody>('/api/file-category/recycle', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}

/** Save File Category Save file category, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated. PUT /api/file-category/save */
export async function fileCategoryPutSave(
  // 叠加生成的Param类型 (非body参数swagger默认没有生成对象)
  params: API.fileCategoryPutSaveParams,
  body: API.FileCategorySave,
  options?: { [key: string]: any }
) {
  return request<API.ResponseFileCategorySaveBody>('/api/file-category/save', {
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

/** Search File Category Search file category by given keyword and fields, and limit the result count. POST /api/file-category/search */
export async function fileCategoryPostSearch(body: API.CrudSearchRequest, options?: { [key: string]: any }) {
  return request<API.ResponseListFileCategorySearchItemBody>('/api/file-category/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    data: body,
    ...(options || {})
  })
}
