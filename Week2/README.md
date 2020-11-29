学习笔记

## sort

- sort 
string sort :https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
golang doc: https://pkg.go.dev/sort?utm_source=gopls#example-Slice

## 树

- 搜索二叉树

左子树上所有的节点都小于它的根结点

右子树上所有节点都大于它的根结点

左右子树分别都是二叉查找树（重复性）

操作

1.查找

```shell
logn
```

2.插入

```shell
logn，查找插入节点，最后遍历到的节点为空，就是该插入的位置
```

3.删除

```shell
1.叶子结点直接删除
2.非叶子结点删除后，取小离它最近的结点（即左子树的最右结点，右子树的最左结点）代替当前位置
```

二叉搜索树的中序遍历为升序

## 堆

* 什么是堆？
  heap：可以迅速找到一堆元素中的最值的数据结构。根结点最大的堆为大顶堆，根结点最小就是小顶堆

* 堆的实现以及复杂度

  ![61727CD7-1F47-4394-AA93-5F8419F8BEE4](/var/folders/gz/8hvg10jd5w31vdl67c3ts5y80000gn/T/com.yinxiang.Mac/com.yinxiang.Mac/WebKitDnD.Q2c9E5/61727CD7-1F47-4394-AA93-5F8419F8BEE4.png)

  堆是一个抽象的概念，它有着常用的如下实现：
  面试中我们常用二叉堆，虽然性能不是最优，但是实现简单,二叉堆（大顶）满足如下性质
  1.是一颗完全树
  2.任意节点的值总是大于等于其子节点
  实现：
  1.因为是一颗完全树所以完全可使用一个一维数组实现
  2.索引为i的左孩子的索引是 2*i+1，右孩纸是2*i+2，父节点是 (i-1)/2

* 堆的常见操作

  1. find max
     	A[0]

  2. insert（push）

  ```shell
  - 添加到堆的尾部
  - 依次上升，如果大于父节点则交换，直到父节点大于它
  ```

  3. pop

  ```shell
  - 将堆顶pop、堆尾元素赋值给堆顶，堆size减一
  - 将堆顶元素下沉，每次将元素和较健壮的子节点交换，直到大于每个子节点
  ```

     

     