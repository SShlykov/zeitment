/**
 * Каллбек изменнеия, ожидает что вернется измененный объект
 *
 * @callback adapterCallback
 * @param   {Object} modifiedObject - новый объект
 * @param   {Object} originalObject - изначальный объект
 * @returns {Object} измененный коллбэком объект
 */

/**
 *
 *  Если у newKey значение отстутсвует - null, '', false, 0, то oldKey удаляется из объекта
 *
 * @param {Object}                        targetObject
 * @param {Object}                        params
 * @param {Object}                        params.config     - oldKey: null // delete oldKey
 * @param {adapterCallback | undefined}   params.callback
 * @returns {Object}
 */
const convertObject = (targetObject, {config, callback = (v) => v}) => {
  let newObject = {}
  for (const key in config) {
    newObject[config[key]] = targetObject[key]
  }
  return callback(newObject, targetObject)
}

/**
 *
 * Если у newKey значение отстутсвует - null, '', false, 0, то oldKey удаляется из объекта
 *
 * @param {Array}                       targetList
 * @param {Object}                      options           - oldKey: null // delete oldKey
 * @param {Object}                      options.config    - oldKey: null // delete oldKey
 * @param {adapterCallback | undefined} options.callback
 * @returns {Object}
 */
const convertList = (targetList, {config, callback = (v) => v}) => {
  return targetList.map((targetObject) => convertObject(targetObject, {config, callback}))
}


export {convertList, convertObject}
