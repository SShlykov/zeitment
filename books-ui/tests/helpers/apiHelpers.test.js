import {expect, test, describe} from 'vitest'
import {post, get, put, remove} from '@helpers/apiHelpers.js'

describe('tests of apiHelpers', () => {
  test('post is a function', () => {
    expect(post).toBeInstanceOf(Function)
  })
  test('get is a function', () => {
    expect(get).toBeInstanceOf(Function)
  })
  test('put is a function', () => {
    expect(put).toBeInstanceOf(Function)
  })
  test('remove is a function', () => {
    expect(remove).toBeInstanceOf(Function)
  })
})