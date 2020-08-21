// 将数组按指定大小分块
export function chunk(arr = [], size = 1) {
  if (arr.length === 0) return [];
  return arr.reduce((total, currentValue) => {
    if (total[total.length - 1].length === size) {
      total.push([currentValue]);
    } else {
      total[total.length - 1].push(currentValue);
    }
    return total;
  }, [[]]);
}

export function generatePages(totalPage) {
  return new Array(totalPage).fill('').map((item, index) => index + 1);
}
