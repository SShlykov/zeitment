import {is} from 'ramda'

/**
 *
 * @param obj
 */
const reverseObject = (obj) => {
  let newObject = {}
  for (const key in obj) {
    newObject[obj[key]] = key
  }
  return newObject
}

const fetchParamsByDefaultObject = (targetObject, defaultParams) => {
  if (!is(Object, targetObject)) {
    return defaultParams
  }
  let newObject = {}
  for (const key in defaultParams) {
    if (targetObject[key]) {
      newObject[key] = targetObject[key]
    } else {
      newObject[key] = defaultParams[key]
    }
  }
  return  newObject
}

export {reverseObject, fetchParamsByDefaultObject}