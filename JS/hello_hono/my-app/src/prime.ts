// 素数を判定する関数
export function isPrime(num: number): boolean {
  if (num < 2) {
    return false;
  }
  for (let i = 2; i <= Math.sqrt(num); i++) {
    if (num % i === 0) {
      return false;
    }
  }
  return true;
}

// n番目までの素数を列挙する関数
export function getPrime(num: number): number[] {
  const primes = [];
  for (let i = 2; primes.length < num; i++) {
    if (isPrime(i)) {
      primes.push(i);
    }
  }
  return primes;
}