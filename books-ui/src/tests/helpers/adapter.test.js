import { expect, test } from 'vitest'
import { convertObject, convertList } from '@helpers/adapter/adapter.js'

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

  const expectedObject = {
    id: 1,
    name2: "jsonplaceholder.typicode.com",
    port2: 9965,
    proxy_ip: "172.25.78.153",
    site_ip: "172.25.78.36",
    internet_uri: "localhost",
    description: "localhost",
    is_online: true,
  }

  const config = {
    id: "id",
    name: "name2",
    port: "port2",
    proxy_ip: "proxy_ip",
    site_ip: "site_ip",
    internet_uri: "internet_uri",
    description: "description",
    is_online: "is_online",
  }

  expect(convertObject(targetObject, {config})).toEqual(expectedObject)
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
    id: 1,
    name2: "jsonplaceholder.typicode.com",
    port2: 9965,
    proxy_ip: "172.25.78.153",
    site_ip: "172.25.78.36",
    internet_uri: "localhost",
    description: "localhost",
    is_online: true,
  }
  const selfObject2 = {
    id: 1,
    name2: "jsonplaceholder.typicode.com",
    port2: 9965,
    proxy_ip: "172.25.78.153",
    site_ip: "172.25.78.36",
    internet_uri: "localhost",
    description: "localhost",
    is_online: true,
  }
  const selfObject3 = {
    id: 1,
    name2: "jsonplaceholder.typicode.com",
    port2: 9965,
    proxy_ip: "172.25.78.153",
    site_ip: "172.25.78.36",
    internet_uri: "localhost",
    description: "localhost",
    is_online: true,
  }

  const selfList = [selfObject1, selfObject2, selfObject3]

  const config = {
    id: "id",
    name: "name2",
    port: "port2",
    proxy_ip: "proxy_ip",
    site_ip: "site_ip",
    internet_uri: "internet_uri",
    description: "description",
    is_online: "is_online",
  }

  expect(convertList(targetList, {config})).toEqual(selfList)
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

  const expectedObject1 = {
    id: 1,
    name2: "jsonplaceholder.typicode.com",
    port2: 9965,
    proxy_ip: "172.25.78.153",
    site_ip: "172.25.78.36",
    internet_uri: "localhost",
    description: "localhost",
    state: "active",
  }

  const expectedObject2 = {
    id: 1,
    name2: "jsonplaceholder.typicode.com",
    port2: 9965,
    proxy_ip: "172.25.78.153",
    site_ip: "172.25.78.36",
    internet_uri: "localhost",
    description: "localhost",
    state: "disabled",
  }


  const expectedList = [expectedObject1, expectedObject2]

  const config = {
    id: "id",
    name: "name2",
    port: "port2",
    proxy_ip: "proxy_ip",
    site_ip: "site_ip",
    internet_uri: "internet_uri",
    description: "description",
  }

  const callback = (el, {is_online}) => ({...el, state: is_online ? "active": "disabled"})

  expect(convertList(targetList, {config, callback})).toEqual(expectedList)
})