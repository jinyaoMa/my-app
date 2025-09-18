import axios, { type AxiosInstance, type AxiosRequestConfig } from 'axios'

const instance = axios.create({
  baseURL: '',
  timeout: 60000,
  withCredentials: true
})

export const request = async <T>(url: string, options: AxiosRequestConfig) => {
  const res = await instance.request<T>({ ...options, url })
  return res.data
}

request.instance = instance

const use: {
  useRequestInterceptor: typeof instance.interceptors.request.use
  useResponseInterceptor: typeof instance.interceptors.response.use
} = {
  useRequestInterceptor: (...args) => instance.interceptors.request.use(...args),
  useResponseInterceptor: (...args) => instance.interceptors.response.use(...args)
}

request.useRequestInterceptor = use.useRequestInterceptor
request.useResponseInterceptor = use.useResponseInterceptor

const crud = {
  get: <TRes>(url: string, options?: AxiosRequestConfig) => request<TRes>(url, { ...options, method: 'GET' }),
  post: <TRes>(url: string, data?: any, options?: AxiosRequestConfig) =>
    request<TRes>(url, { ...options, method: 'POST', data }),
  put: <TRes>(url: string, data?: any, options?: AxiosRequestConfig) =>
    request<TRes>(url, { ...options, method: 'PUT', data }),
  delete: <TRes>(url: string, data?: any, options?: AxiosRequestConfig) =>
    request<TRes>(url, { ...options, method: 'DELETE', data })
}

request.get = crud.get
request.post = crud.post
request.put = crud.put
request.delete = crud.delete

export default request as AxiosInstance &
  typeof use &
  typeof crud & {
    instance: AxiosInstance
  }
