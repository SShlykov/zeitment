import { expect, test } from 'vitest'
import { convertObject, covertList } from './adapter'


test('parse object from api', () => {
  const targetObject = {
    "id": 1,
    "created_at": "2024-02-15T17:24:52.755254148+03:00",
    "updated_at": "2024-02-15T17:24:52.755254148+03:00",
    "deleted_at": null,
    "name": "jsonplaceholder.typicode.com",
    "port": 9965,
    "proxy_ip": "172.25.78.153",
    "site_ip": "172.25.78.36",
    "internet_uri": "localhost",
    "description": "localhost",
    "is_online": true
  }

  const selfObject = {
    "created_at": "2024-02-15T17:24:52.755254148+03:00",
    "description": "localhost",
    "device_ip": "172.25.78.36",
    "proxy_ip": "172.25.78.153",
    "is_online": true,
    "id": 1,
    "port2": 9965,
    "name": "jsonplaceholder.typicode.com",
    "internet_uri": "localhost",
  }

  const config = {
    port: "port2",
    site_ip: "device_ip"
  }

  const listExcluded = ['updated_at', 'updated_at', 'deleted_at']

  expect(convertObject(targetObject, config, {listExcluded})).toEqual(selfObject)
})

test('adapt list objects without callback', () => {
  const targetObject1 = {
    "id": 1,
    "created_at": "2024-02-15T17:24:52.755254148+03:00",
    "updated_at": "2024-02-15T17:24:52.755254148+03:00",
    "deleted_at": null,
    "name": "jsonplaceholder.typicode.com",
    "port": 9965,
    "proxy_ip": "172.25.78.153",
    "site_ip": "172.25.78.36",
    "internet_uri": "localhost",
    "description": "localhost",
    "is_online": true
  }
  const targetObject2 = {
    "id": 1,
    "created_at": "2024-02-15T17:24:52.755254148+03:00",
    "updated_at": "2024-02-15T17:24:52.755254148+03:00",
    "deleted_at": null,
    "name": "jsonplaceholder.typicode.com",
    "port": 9965,
    "proxy_ip": "172.25.78.153",
    "site_ip": "172.25.78.36",
    "internet_uri": "localhost",
    "description": "localhost",
    "is_online": true
  }
  const targetObject3 = {
    "id": 1,
    "created_at": "2024-02-15T17:24:52.755254148+03:00",
    "updated_at": "2024-02-15T17:24:52.755254148+03:00",
    "deleted_at": null,
    "name": "jsonplaceholder.typicode.com",
    "port": 9965,
    "proxy_ip": "172.25.78.153",
    "site_ip": "172.25.78.36",
    "internet_uri": "localhost",
    "description": "localhost",
    "is_online": true
  }

  const targetList = [targetObject1, targetObject2, targetObject3]

  const selfObject1 = {
    "created_at": "2024-02-15T17:24:52.755254148+03:00",
    "description": "localhost",
    "device_ip": "172.25.78.36",
    "proxy_ip": "172.25.78.153",
    "is_online": true,
    "id": 1,
    "port2": 9965,
    "name": "jsonplaceholder.typicode.com",
    "internet_uri": "localhost",
  }
  const selfObject2 = {
    "created_at": "2024-02-15T17:24:52.755254148+03:00",
    "description": "localhost",
    "device_ip": "172.25.78.36",
    "proxy_ip": "172.25.78.153",
    "is_online": true,
    "id": 1,
    "port2": 9965,
    "name": "jsonplaceholder.typicode.com",
    "internet_uri": "localhost",
  }
  const selfObject3 = {
    "created_at": "2024-02-15T17:24:52.755254148+03:00",
    "description": "localhost",
    "device_ip": "172.25.78.36",
    "proxy_ip": "172.25.78.153",
    "is_online": true,
    "id": 1,
    "port2": 9965,
    "name": "jsonplaceholder.typicode.com",
    "internet_uri": "localhost",
  }

  const selfList = [selfObject1, selfObject2, selfObject3]

  const config = {
    port: "port2",
    site_ip: "device_ip"
  }

  const listExcluded = ['updated_at', 'updated_at', 'deleted_at']

  expect(covertList(targetList, config, {listExcluded})).toEqual(selfList)
})

test('adapt list objects with callback', () => {
  const targetObject1 = {
    "id": 1,
    "created_at": "2024-02-15T17:24:52.755254148+03:00",
    "updated_at": "2024-02-15T17:24:52.755254148+03:00",
    "deleted_at": null,
    "name": "jsonplaceholder.typicode.com",
    "port": 9965,
    "proxy_ip": "172.25.78.153",
    "site_ip": "172.25.78.36",
    "internet_uri": "localhost",
    "description": "localhost",
    "is_online": true
  }
  const targetObject2 = {
    "id": 1,
    "created_at": "2024-02-15T17:24:52.755254148+03:00",
    "updated_at": "2024-02-15T17:24:52.755254148+03:00",
    "deleted_at": null,
    "name": "jsonplaceholder.typicode.com",
    "port": 9965,
    "proxy_ip": "172.25.78.153",
    "site_ip": "172.25.78.36",
    "internet_uri": "localhost",
    "description": "localhost",
    "is_online": false
  }

  const targetList = [targetObject1, targetObject2]

  const selfObject1 = {
    "created_at": "2024-02-15T17:24:52.755254148+03:00",
    "description": "localhost",
    "device_ip": "172.25.78.36",
    "proxy_ip": "172.25.78.153",
    "state": "active",
    "id": 1,
    "port2": 9965,
    "name": "jsonplaceholder.typicode.com",
    "internet_uri": "localhost",
  }
  const selfObject2 = {
    "created_at": "2024-02-15T17:24:52.755254148+03:00",
    "description": "localhost",
    "device_ip": "172.25.78.36",
    "proxy_ip": "172.25.78.153",
    "state": "disabled",
    "id": 1,
    "port2": 9965,
    "name": "jsonplaceholder.typicode.com",
    "internet_uri": "localhost",
  }


  const selfList = [selfObject1, selfObject2]

  const config = {
    port: "port2",
    site_ip: "device_ip"
  }

  const listExcluded = ['updated_at', 'updated_at', 'deleted_at', 'is_online']

  const callback = (el, {is_online}) => ({...el, state: is_online ? "active": "disabled"})

  expect(covertList(targetList, config, {listExcluded, callback})).toEqual(selfList)
})
