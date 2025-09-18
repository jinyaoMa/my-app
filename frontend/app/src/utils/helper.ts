/* eslint-disable @typescript-eslint/no-explicit-any */
export const convert2PathObject = (obj: any) => {
  const result: Record<string, any> = {}

  function recurse(current: { [x: string]: any }, path = '') {
    for (const key in current) {
      if (Object.prototype.hasOwnProperty.call(current, key)) {
        const newPath = path ? `${path}.${key}` : key
        if (typeof current[key] === 'object' && current[key] !== null) {
          recurse(current[key], newPath)
        } else {
          result[newPath] = current[key]
        }
      }
    }
  }

  recurse(obj)
  return result
}
