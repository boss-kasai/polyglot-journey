import { Hono } from 'hono'
import { fizzBuzz } from './fizzbuzz'
import { getPrime } from './prime'
import { getFibonacci } from './fibonacci'

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

export default app
