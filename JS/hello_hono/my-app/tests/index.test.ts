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
