function binaruSearch(
  arr: number[],
  value: number,
  floor: number,
  ceiling: number
): number {
  const mid = Math.floor((floor + ceiling) / 2);

  if (arr[mid] === value) {
    return mid;
  }

  if (floor > ceiling) return -1;

  if (arr[mid] > value) {
    return binaruSearch(arr, value, floor, mid - 1);
  } else {
    return binaruSearch(arr, value, mid + 1, ceiling);
  }
}

export function exponentialSearch(
  arr: number[],
  length: number,
  value: number
): number {
  if (arr[0] === value) {
    return 0;
  }
  let i = 1;
  while (i < length && arr[i] <= value) {
    i = i * 2;
  }
  return binaruSearch(arr, value, i / 2, Math.min(i, length));
}
const arr = [2, 3, 4, 10, 40, 65, 78, 100];
const value = 78;
const result = exponentialSearch(arr, arr.length, value);

console.log(result);
