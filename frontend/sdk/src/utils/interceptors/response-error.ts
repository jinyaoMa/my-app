import type { AxiosError } from 'axios'

const errorMessages: Record<number, string> = {}

export const responseError = (error: AxiosError) => {
  const status = error.response?.status || -1
  return Promise.reject(error)
}
