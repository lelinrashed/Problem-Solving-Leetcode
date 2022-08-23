/*
Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.

const I = 1;
const V = 5;
const X = 10;
const L = 50;
const C = 100;
const D = 500;
const M = 1000;
*/

let romanToInt = function (s) {
  let table = {
    I: 1,
    V: 5,
    X: 10,
    L: 50,
    C: 100,
    D: 500,
    M: 1000,
  };

  let result = 0;
  for (let i = 0; i < s.length; i++) {
    //if the next roman numeral is larger, then we know we have to subtract this number
    if (table[s[i]] < table[s[i + 1]]) {
      result -= table[s[i]];
    }
    //otherwise, add like normal.
    else {
      result += table[s[i]];
    }
  }

  return result;
};

console.log(romanToInt("MCMXCIV"));
