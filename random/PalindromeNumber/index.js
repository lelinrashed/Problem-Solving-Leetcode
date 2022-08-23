const reverse = (num) => parseInt(String(num).split("").reverse().join(""), 10);

var isPalindrome = function (x) {
  if (x === reverse(x)) {
    return true;
  }
  return false;
};

console.log(isPalindrome(123));
