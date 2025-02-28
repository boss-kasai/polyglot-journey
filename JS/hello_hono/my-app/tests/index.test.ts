import { describe, it, expect } from 'bun:test'
import { app } from '../src/index'

describe('Index routes', () => {
  it('GET / should return status=200 and the expected text', async () => {
    // Hono の app.request() を用いて疑似的にリクエストを発行
    const res = await app.request('http://localhost/')
    expect(res.status).toBe(200)

    // レスポンス本文（テキスト）を確認
    const text = await res.text()
    expect(text).toBe('Hello Hono!')
  })

  it('GET /health should return status=200 and JSON body', async () => {
    // /hello エンドポイントをテスト
    const res = await app.request('http://localhost/health')
    expect(res.status).toBe(200)

    // レスポンス本文（JSON）を確認
    const body = await res.json()
    expect(body).toEqual({ status: 'ok' })
  })
})

describe('FizzBuzz route', () => {
  it('GET /fizzbuzz/:num should return the expected text', async () => {
    // /fizzbuzz/:num エンドポイントをテスト
    const res = await app.request('http://localhost/fizzbuzz/3')
    expect(res.status).toBe(200)

    // レスポンス本文（テキスト）を確認
    const text = await res.text()
    expect(text).toBe('Fizz')
  })

  it('GET /fizzbuzz/:num 数字に変換できない場合は400を返す ', async () => {
    // /fizzbuzz/:num エンドポイントをテスト
    const res = await app.request('http://localhost/fizzbuzz/abc')
    expect(res.status).toBe(400)

    // レスポンス本文（テキスト）を確認
    const text = await res.text()
    expect(text).toBe('Invalid number: abc')
  })
})