// var twoSum = function (nums, target) {
//   const indices = new Map();

//   for (let index = 0; index < nums.length; index++) {
//     const complement = target - nums[index];
//     console.log("indices", indices);
//     console.log("complement", complement);
//     if (indices.has(complement)) {
//       console.log(indices.get(complement));
//       return [indices.get(complement), index];
//     }

//     indices.set(nums[index], index);
//   }
// };

// console.log(twoSum([2, 8, 7, 20], 9));

var twoSum = function (nums, target) {
  // Define map
  const testMap = new Map();

  for (let index = 0; index < nums.length; index++) {
    const element = nums[index];

    if (testMap.has(target - nums[index])) {
      return [testMap.get(target - nums[index]), index];
    }

    testMap.set(nums[index], index);
  }
};

console.log(twoSum([3, 6, 7, 2], 5));
