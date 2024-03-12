import {describe, test, expect} from "vitest";
import {reverseObject, fetchParamsByDefaultObject} from '@helpers/objectUtils'

describe("tests objectReverse" , () => {
  test('reverse object', () => {
    const targetObject = {
      "id": 1,
      "name": "jsonplaceholder.typicode.com",
      "port": 9965,
      "proxy_ip": "",
      "site_ip": "siteIp"
    }
    const expectedObject = {
      1: "id",
      "jsonplaceholder.typicode.com": "name",
      9965: "port",
      "": "proxy_ip",
      "siteIp": "site_ip",
    }

    expect(reverseObject(targetObject)).toEqual(expectedObject)
  })
})

describe("tests of fetchParamsByDefaultObject" , () => {
  test('fetch params from object', () => {
    const targetObject = {
      foo: "bar",
      id: 1,
    }
    const defaultParams = {
      foo: "test",
      id: 0
    }

    expect(fetchParamsByDefaultObject(targetObject, defaultParams)).toEqual(targetObject)
  })

  test('fetch params from null', () => {
    const targetObject = null
    const defaultParams = {
      foo: "test",
      id: 0
    }

    expect(fetchParamsByDefaultObject(targetObject, defaultParams)).toEqual(defaultParams)
  })

  test('fetch params from empty object', () => {
    const targetObject = {}
    const defaultParams = {
      foo: "test",
      id: 0
    }

    expect(fetchParamsByDefaultObject(targetObject, defaultParams)).toEqual(defaultParams)
  })

  test('fetch params from object with one of two params', () => {
    const targetObject = {
      foo: "bar",
    }
    const defaultParams = {
      foo: "test",
      id: 0
    }
    const expectedObject = {
      foo: "bar",
      id: 0
    }

    expect(fetchParamsByDefaultObject(targetObject, defaultParams)).toEqual(expectedObject)
  })
})