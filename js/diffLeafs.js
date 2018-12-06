// 链接：http://zhuanlan.zhihu.com/purerender/20346379

// 通过下面的 demo 可以清晰的描述传统 diff 算法的实现过程。

let result = [];

// 比较叶子节点
const diffLeafs = function(beforeLeaf, afterLeaf) {

  // 获取较大节点树的长度
  let count = Math.max(beforeLeaf.children.length, afterLeaf.children.length);

  // 循环遍历
  for (let i = 0; i < count; i++) {

    const beforeTag = beforeLeaf.children[i];
    const afterTag = afterLeaf.children[i];

    // 添加 afterTag 节点
    if (beforeTag === undefined) {
      result.push({type: "add", element: afterTag});
      // 删除 beforeTag 节点
    } else if (afterTag === undefined) {
      result.push({type: "remove", element: beforeTag});
      // 节点名改变时，删除 beforeTag 节点，添加 afterTag 节点
    } else if (beforeTag.tagName !== afterTag.tagName) {
      result.push({type: "remove", element: beforeTag});
      result.push({type: "add", element: afterTag});
      // 节点不变而内容改变时，改变节点
    } else if (beforeTag.innerHTML !== afterTag.innerHTML) {
      if (beforeTag.children.length === 0) {
        result.push({
          type: "changed",
          beforeElement: beforeTag,
          afterElement: afterTag,
          html: afterTag.innerHTML
        });
      } else {
        // 递归比较
        diffLeafs(beforeTag, afterTag);
      }
    }
  }
  return result;
}

// 传统 diff 算法的复杂度为 O(n^3)，显然这是无法满足性能要求的。
// React 通过制定大胆的策略，将 O(n^3) 复杂度的问题转换成 O(n) 复杂度的问题。


