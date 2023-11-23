# 力扣刷题记录
 
## 学习资料

代码随想录：https://programmercarl.com/

灵茶山艾府，灵神
- Start: 20230906
- 课程表：https://github.com/EndlessCheng/codeforces-go/tree/master/leetcode
- 教学视频：https://www.bilibili.com/video/BV1Qg411q7ia/

## 题目分类
> 跟着卡哥（代码随想录，程序员卡尔）刷完一遍，一段时间没刷遗忘了不少；但在这期间几乎算裸考的，蓝桥杯省一和国三。考试前一个月一道题也没碰过，抱着试一试的心理，结果还不错。坚持刷题很有帮助！
> 
>对我而言，做过的题会遗忘，坚持刷题是一方面，回顾也很重要。特别是某一类型的题目，快速捡起来看曾经分好类做好笔记的刷题记录很有用。这也是这个Repo的意义，加油！题目分类，参考灵神。

### 相向双指针1

**left++, right--**
- 作用在有序列表，一左一右，相向而行；
- 因为有序，大小确定信息。

拓展，三数之和：
- fixed + min1 + min2 > target; break;
- fixed + max1 + max2 < target; 下一个fixed congtinue;

|  题目   |难度| 题解  | 备注 |
|  ----  | ---- |----  | ----|
| [167.两数之和 II - 输入有序数组](https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/) | 中等 | [Golang](./leetcode_recordes/167.两数之和-ii-输入有序数组.go) |      |
| [15.三数之和](https://leetcode.cn/problems/3sum/description/) | 中等 | [Golang](./leetcode_recordes/15.三数之和.go)| 两个优化、先167 |
|[16.最接近的三数之和](https://leetcode.cn/problems/3sum-closest/description/)|中等|[Golang](./leetcode_recordes/16.最接近的三数之和.go)| |
|[18.四数之和](https://leetcode.cn/problems/4sum/)|中等|[Golang](./leetcode_recordes/18.四数之和.go)| |


### 相向双指针2

**left++, right--**
- 列表无序，也不能排序；

|  题目   |难度| 题解  | 备注 |
|  ----  | ---- |----  | ----|
| [11.盛最多水的容器](https://leetcode.cn/problems/container-with-most-water/description/) |中等 |[Golang](./leetcode_recordes/11.盛最多水的容器.go)| |
|[42.接雨水](https://leetcode.cn/problems/trapping-rain-water/description/) |困难|[Golang](./leetcode_recordes/42.接雨水.go)|他法：前后缀分解|


### 同向双指针-滑动窗口
笔记：https://www.yuque.com/cnbetter/fsa89u/egm0agwrpf00qqqo?singleDoc#

|  题目   |难度| 题解  | 备注 |
|  ----  | ---- |----  | ----|
| [3.无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/) |中等 |[Python](./leetcode_recordes/3.无重复字符的最长子串.py)| 寻找最长 |
|[209.长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/) |中等|[Golang](./leetcode_recordes/209.长度最小的子数组.go)| 寻找最短 |
|[1004.最大连续1的个数 III](https://leetcode.cn/problems/max-consecutive-ones-iii)|中等| [Python](./leetcode_recordes/1004.最大连续-1-的个数-iii.py) |UseM 1AC|




## Todo
- ✅ ~~通过题号，生成markdown表格格式，eg：`|[209.长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/) |中等|[Golang](.leetcode_recordes/209.长度最小的子数组.go)| |`~~
    - 使用方式: `python id2row.py 3 11`
        ```
        |[3.无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters)|中等|[Python](./leetcode_recordes/3.无重复字符的最长子串.py)| |
        |[11.盛最多水的容器](https://leetcode.cn/problems/container-with-most-water/description/)|中等|[Golang](./leetcode_recordes/11.盛最多水的容器.go)| |
        ```