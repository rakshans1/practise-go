const numbers = function* () {
  for (let x = 2;; yield x++);
}

const sieve = function* (sequence) {
  const { value: prime, done } = sequence.next();
  if (done) return;

  yield prime;

  yield* sieve(function* () {
    for (let value of sequence) {
      if (value % prime) {
        yield value
      }
    }
  }())
}

const primeGen = sieve(numbers());

const primes = [...Array(10)].map(() => primeGen.next().value).concat('...').join(',')

console.log(primes);

