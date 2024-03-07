import {expect, test, describe} from 'vitest'
import AdapterOfBooks from '@/adapters/AdapterOfBooks.js'

describe('adapterOfBooks', () => {
  const uri = 'https://www.googleapis.com/books/v1/volumes?q=javascript';
  const adapter = new AdapterOfBooks(uri);

  test('adapterOfBooks is an instance of AdapterOfBooks', () => {
    expect(adapter).toBeInstanceOf(AdapterOfBooks)
  })
  test('adapterOfBooks is exist', () => {
    expect(adapter).toBeDefined()
  })
})