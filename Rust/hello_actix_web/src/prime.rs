// n番目までの素数を返す
use std::f64;

pub fn get_prime(n: usize) -> Vec<u32> {
    let mut primes = vec![];
    let mut num = 2;

    while primes.len() < n {
        // √num を事前に計算 (u32にキャストする)
        let limit = (num as f64).sqrt() as u32;

        // “p <= limit” のみに絞って試し割り
        let is_prime = primes
            .iter()
            .take_while(|&&p| p <= limit)
            .all(|&p| num % p != 0);

        if is_prime {
            primes.push(num);
        }
        num += 1;
    }

    primes
}

/// エラトステネスの篩で、`limit`以下の素数を返す
fn sieve_of_eratosthenes(limit: usize) -> Vec<usize> {
    let mut sieve = vec![true; limit + 1]; // インデックス=数値, true=候補
    sieve[0] = false;
    if limit >= 1 {
        sieve[1] = false;
    }

    let mut p = 2;
    while p * p <= limit {
        if sieve[p] {
            // pの倍数を false に
            let mut multiple = p * p;
            while multiple <= limit {
                sieve[multiple] = false;
                multiple += p;
            }
        }
        p += 1;
    }

    // true なインデックスだけを集めて返す
    sieve
        .iter()
        .enumerate()
        .filter(|&(_, &is_prime)| is_prime)
        .map(|(num, _)| num)
        .collect()
}

pub fn nth_primes(n: usize) -> Vec<usize> {
    // n番目の素数を知りたい → だいたい n log n くらいを見積もる
    // （n=1や小さいときの補正も必要）
    if n < 2 {
        // 特殊対応: n=1 の場合, 2が最初の素数, みたいな
        return vec![2];
    }

    let n_f64 = n as f64;
    // 素数近似式: n (ln n + ln ln n), n>=6 ならOK
    let mut upper = (n_f64 * (n_f64.ln() + n_f64.ln().ln())) as usize;
    // 安全マージンを少し足しておく
    upper = upper.max(10).max(n * 12); // 適当な保険

    // 1) エラトステネスの篩
    let primes = sieve_of_eratosthenes(upper);

    // 2) primesに n個以上あれば、その先頭n個を返す
    if primes.len() >= n {
        primes[..n].to_vec()
    } else {
        // 万が一足りなければ、もうちょっと範囲を広げて再試行するなど
        // ここでは適当に2倍して再試行
        let bigger_primes = sieve_of_eratosthenes(upper * 2);
        bigger_primes[..n].to_vec()
    }
}
