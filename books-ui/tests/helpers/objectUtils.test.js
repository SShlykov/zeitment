import {describe, test, expect} from "vitest";
import {reverseObject} from '@helpers/objectUtils'

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