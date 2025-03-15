import { Hono } from 'hono'
import { fizzBuzz } from './fizzbuzz'
import { getPrime } from './prime'
import { getFibonacci } from './fibonacci'
import { getHash } from './hash'

export const app = new Hono()

app.get('/', (c) => {
  return c.text('Hello Hono!')
})

app.get('/health', (c) => {
  return c.json({ status: 'ok' })
})

app.get('/fizzbuzz/:num', (c) => {
  const rawNum = c.req.param('num')
  const num = parseInt(rawNum)
  if (isNaN(num)) {
    return c.text(`Invalid number: ${rawNum}`, 400)
  }
  const result = fizzBuzz(num)
  return c.text(result)
})

app.post('/bmi', async (c) => {
  const { height, weight } = await c.req.json()
  if (!height || !weight) {
    return c.text('Invalid request: height and weight are required', 400)
  }
  const heightInMeters = height / 100
  const bmi = weight / (heightInMeters * heightInMeters)
  const roundedBmi = Math.round(bmi * 100) / 100
  let category = ''

  if (roundedBmi < 18.5) {
    category = 'Underweight'
  } else if (roundedBmi < 24.9) {
    category = 'Normal weight'
  } else if (roundedBmi < 29.9) {
    category = 'Overweight'
  } else {
    category = 'Obesity'
  }

  return c.json({ bmi: roundedBmi, category })
})

app.get(`prime/:num`, (c) => {
  const rawNum = c.req.param('num')
  const num = parseInt(rawNum)
  if (isNaN(num)) {
    return c.text(`Invalid number: ${rawNum}`, 400)
  }
  const primes = getPrime(num)
  return c.json(primes)
})

app.get('/fibonacci/:num', (c) => {
  const rawNum = c.req.param('num')
  const num = parseInt(rawNum)
  if (isNaN(num)) {
    return c.text(`Invalid number: ${rawNum}`, 400)
  }
  const fib = getFibonacci(num)
  return c.json(fib)
})

app.post('/loan', async (c) => {
  // リクエスト JSON: { monthlyPayment, years, annualRate }
  const { monthlyPayment, years, annualRate } = await c.req.json<{
    monthlyPayment?: number
    years?: number
    annualRate?: number
  }>()

  // バリデーションチェック
  if (
    typeof monthlyPayment !== 'number' ||
    typeof years !== 'number' ||
    typeof annualRate !== 'number'
  ) {
    return c.text('Invalid input', 400)
  }
  // 月々の返済額が0以下の場合はエラー
  if (monthlyPayment <= 0) {
    return c.text('Monthly payment should be greater than 0', 400)
  }

  // 計算
  const n = years * 12
  const i = annualRate / 12
  if (i === 0) {
    return c.text('Annual rate should not be 0', 400)
  }

  const numerator = 1 - Math.pow(1 + i, -n)
  const principal = monthlyPayment * (numerator / i)

  // 結果を返す (整数に四捨五入するか、少数第2位まで表示するかなど、仕様に合わせて)
  return c.json({ principal })
})

app.post('/stringcount', async (c) => {
  const { text } = await c.req.json<{ text?: string }>()
  if (!text) {
    return c.text('Invalid input', 400)
  }
  const count = text.length
  return c.json({ count })
})

app.post('/hash', async (c) => {
  // リクエスト JSON: { text }
  const { text } = await c.req.json<{ text?: string }>()
  if (!text) {
    return c.text('Invalid input', 400)
  }
  const hashed = await getHash(text)
  return c.json({ hashed })
})


export default app
