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

export {reverseObject}