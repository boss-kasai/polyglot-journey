import { Hono } from 'hono'
import { fizzBuzz } from './fizzbuzz'

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

export default app
