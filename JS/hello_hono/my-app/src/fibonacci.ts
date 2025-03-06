// n番目のフィボナッチ数を求める関数
export function getFibonacci(n: number): number {
  if (n <= 0) {
    return 0;
  }
  if (n === 1) {
    return 1;
  }
  return getFibonacci(n - 1) + getFibonacci(n - 2);
}