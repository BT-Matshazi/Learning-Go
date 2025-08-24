"use strict";

// Declare int array
let numbers = [
  55, 10, 305, 220, 175, 90, 330, 25, 145, 250, 70, 195, 385, 40, 275, 130, 350,
  115, 5, 365, 180, 290, 135, 50, 215, 395, 160, 80, 225, 300, 120, 20, 245,
  375, 110, 60, 320, 165, 210, 285, 35, 150, 400, 140, 270, 95, 340, 15, 200,
  370, 105, 290, 45, 185, 310, 100, 280, 235, 365, 125, 55, 240, 390, 360, 230,
  85, 260, 325, 375, 190, 345, 255, 330, 170, 205, 315, 75, 300, 365, 150,
];

// Declare function bubbleSort
const bubbleSort = (arr) => {
  for (let i = 0; i < arr.length - 1; i++) {
    // Outer loop
    for (let j = 0; j < arr.length - 1 - i; j++) {
      // Inner loop optimized
      if (arr[j] > arr[j + 1]) {
        //left element is greater than right
        let temp = arr[j];
        arr[j] = arr[j + 1];
        arr[j + 1] = temp;
      }
    }
  }
  console.log(arr);
};

bubbleSort(numbers);
