/**
 *
 * @param {Object}   targetObject
 * @param {Object}   config - oldKey: newKey
 * @param {Object}   options
 * @param {Array}    options.listExcluded
 * @returns {Object}
 */
const convertObject = (_targetObject, config, {listExcluded = []}) => {
  const targetObject = {..._targetObject}
  let configList = Object.entries(config)

  let newObject = configList.reduce((mutationObject, [oldKey, newKey]) => {
    mutationObject[newKey] = mutationObject[oldKey]
    delete mutationObject[oldKey]
    return mutationObject
  }, targetObject)

  listExcluded.forEach((removedKey) => {
    delete newObject[removedKey]
  })

  return newObject
}

/**
 *
 * @param {Array}   targetList
 * @param {Object}   config - oldKey: newKey
 * @param {Object}   options
 * @param {Array}    options.listExcluded
 * @param {Function}    options.callback
 * @returns {Object}
 */
const covertList = (targetList, config, {listExcluded = [], callback = (v) => v}) => {
  return targetList.map((targetObject) => callback(convertObject(targetObject, config, {listExcluded}), targetObject))
}


export {covertList, convertObject}
