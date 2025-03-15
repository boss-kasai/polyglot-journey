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

describe('POST /bmi', () => {
  it('should return 400 if height or weight is missing', async () => {
    // heightのみ渡してweightを渡さないケース
    const res = await app.request('http://localhost/bmi', {
      method: 'POST',
      body: JSON.stringify({ height: 170 }),
      headers: { 'Content-Type': 'application/json' }
    })

    expect(res.status).toBe(400)
    const text = await res.text()
    expect(text).toBe('Invalid request: height and weight are required')
  })

  it('should return correct BMI and category for valid data', async () => {
    // height=170, weight=65 の場合、BMIは約22.49となりカテゴリは "Normal weight"
    const res = await app.request('http://localhost/bmi', {
      method: 'POST',
      body: JSON.stringify({ height: 170, weight: 65 }),
      headers: { 'Content-Type': 'application/json' }
    })

    expect(res.status).toBe(200)
    const data = await res.json()
    // data = { bmi: 22.49, category: "Normal weight" } のイメージ
    expect(data.bmi).toBe(22.49)        // 小数点2桁四捨五入
    expect(data.category).toBe('Normal weight')
  })

  it('f64 error', async () => {
    // height=178.0, weight=63.72 の場合、BMIは約20.11となりカテゴリは "Normal weight"
    const res = await app.request('http://localhost/bmi', {
      method: 'POST',
      body: JSON.stringify({ height: 178.0, weight: 63.72 }),
      headers: { 'Content-Type': 'application/json' }
    })

    expect(res.status).toBe(200)
    const data = await res.json()
    expect(data.bmi).toBe(20.11)        // 小数点2桁四捨五入
    expect(data.category).toBe('Normal weight')
    })

  it('should classify BMI < 18.5 as Underweight', async () => {
    // height=170, weight=50 => BMI 約17.30
    const res = await app.request('http://localhost/bmi', {
      method: 'POST',
      body: JSON.stringify({ height: 170, weight: 50 }),
      headers: { 'Content-Type': 'application/json' }
    })

    expect(res.status).toBe(200)
    const data = await res.json()
    expect(data.bmi).toBe(17.3)
    expect(data.category).toBe('Underweight')
  })
})

describe('Prime route', () => {
  it ('GET /prime/:num should return the expected JSON', async () => {
    // /prime/:num エンドポイントをテスト
    const res = await app.request('http://localhost/prime/1')
    expect(res.status).toBe(200)
    // レスポンス本文（JSON）を確認
    const body = await res.json()
    expect(body).toEqual([2])
  })

  it('GET /prime/:num should return the expected JSON', async () => {
    // /prime/:num エンドポイントをテスト
    const res = await app.request('http://localhost/prime/2')
    expect(res.status).toBe(200)

    // レスポンス本文（JSON）を確認
    const body = await res.json()
    expect(body).toEqual([2, 3])
  })

  it('GET /prime/:num 数字に変換できない場合は400を返す ', async () => {
    // /prime/:num エンドポイントをテスト
    const res = await app.request('http://localhost/prime/abc')
    expect(res.status).toBe(400)

    // レスポンス本文（テキスト）を確認
    const text = await res.text()
    expect(text).toBe('Invalid number: abc')
  })
})

describe('Fibonacci route', () => {
  it('GET /fibonacci/:num should return the expected JSON', async () => {
    // /fibonacci/:num エンドポイントをテスト
    const res = await app.request('http://localhost/fibonacci/0')
    expect(res.status).toBe(200)

    // レスポンス本文（JSON）を確認
    const body = await res.json()
    expect(body).toEqual(0)
  })

  it('GET /fibonacci/:num should return the expected JSON', async () => {
    // /fibonacci/:num エンドポイントをテスト
    const res = await app.request('http://localhost/fibonacci/1')
    expect(res.status).toBe(200)

    // レスポンス本文（JSON）を確認
    const body = await res.json()
    expect(body).toEqual(1)
  })

  it('GET /fibonacci/:num should return the expected JSON', async () => {
    // /fibonacci/:num エンドポイントをテスト
    const res = await app.request('http://localhost/fibonacci/2')
    expect(res.status).toBe(200)

    // レスポンス本文（JSON）を確認
    const body = await res.json()
    expect(body).toEqual(1)
  })

  it('GET /fibonacci/:num should return the expected JSON', async () => {
    // /fibonacci/:num エンドポイントをテスト
    const res = await app.request('http://localhost/fibonacci/3')
    expect(res.status).toBe(200)

    // レスポンス本文（JSON）を確認
    const body = await res.json()
    expect(body).toEqual(2)
  })

  it('GET /fibonacci/:num should return the expected JSON', async () => {
    // /fibonacci/:num エンドポイントをテスト
    const res = await app.request('http://localhost/fibonacci/4')
    expect(res.status).toBe(200)

    // レスポンス本文（JSON）を確認
    const body = await res.json()
    expect(body).toEqual(3)
  })
})

describe('stringcount route', () => {
  it('POST /stringcount should return the expected JSON', async () => {
    // /stringcount エンドポイントをテスト
    const res = await app.request('http://localhost/stringcount', {
      method: 'POST',
      body: JSON.stringify({ text: 'hello' }),
      headers: { 'Content-Type': 'application/json' }
    })
    expect(res.status).toBe(200)

    // レスポンス本文（JSON）を確認
    const body = await res.json()
    expect(body).toEqual({ count: 5 })
  })
  it ('POST /stringcount 空欄を含む処理の確認', async () => {
    // /stringcount エンドポイントをテスト
    const res = await app.request('http://localhost/stringcount', {
      method: 'POST',
      body: JSON.stringify({ text: 'hello world' }),
      headers: { 'Content-Type': 'application/json' }
    })
    expect(res.status).toBe(200)

    // レスポンス本文（JSON）を確認
    const body = await res.json()
    expect(body).toEqual({ count: 11 })
  })
  it ('POST /stringcount 日本語の処理', async () => {
    // /stringcount エンドポイントをテスト
    const res = await app.request('http://localhost/stringcount', {
      method: 'POST',
      body: JSON.stringify({ text: 'こんにちは　世界' }),
      headers: { 'Content-Type': 'application/json' }
    })
    expect(res.status).toBe(200)

    // レスポンス本文（JSON）を確認
    const body = await res.json()
    expect(body).toEqual({ count: 8 })
  })
  it ('POST /stringcount 空欄処理の確認', async () => {
    // /stringcount エンドポイントをテスト
    const res = await app.request('http://localhost/stringcount', {
      method: 'POST',
      body: JSON.stringify({ text: '' }),
      headers: { 'Content-Type': 'application/json' }
    })
    expect(res.status).toBe(400)

    // レスポンス本文（JSON）を確認
    const text = await res.text()
    expect(text).toBe('Invalid input')
  })
  it ('POST /stringcount should return 400 if text is missing', async () => {
    // /stringcount エンドポイントをテスト
    const res = await app.request('http://localhost/stringcount', {
      method: 'POST',
      body: JSON.stringify({}),
      headers: { 'Content-Type': 'application/json' }
    })
    expect(res.status).toBe(400)

    // レスポンス本文（テキスト）を確認
    const text = await res.text()
    expect(text).toBe('Invalid input')
  })
})

describe('hasu route', () => {
  it('POST /hash should return the expected JSON', async () => {
    // /hasu エンドポイントをテスト
    const res = await app.request('http://localhost/hash', {
      method: 'POST',
      body: JSON.stringify({ text: 'password' }),
      headers: { 'Content-Type': 'application/json' }
    })
    expect(res.status).toBe(200)

    // レスポンス本文（JSON）を確認
    const body = await res.json()
    expect(body).toEqual({ hashed: '$2b$10$ABCDEFGHIJKLMNOPQRSTUOiAi7OcdE4zRCh6NcGWusEcNPtq6/w8.' })
  })
  it ('POST /hash should return 400 if text is missing', async () => {
    // /hasu エンドポイントをテスト
    const res = await app.request('http://localhost/hash', {
      method: 'POST',
      body: JSON.stringify({}),
      headers: { 'Content-Type': 'application/json' }
    })
    expect(res.status).toBe(400)
    // レスポンス本文（テキスト）を確認
    const text = await res.text()
    expect(text).toBe('Invalid input')
  })
})