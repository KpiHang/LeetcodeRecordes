# 通过输入id（可同时多个id），生成如下格式
# |[209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/) |中等|[Golang](./leetcode_recordes/209.长度最小的子数组.go)| |

import sys, os
import requests


# https://leetcode.cn/search/?q=3.%E6%97%A0%E9%87%8D%E5%A4%8D%E5%AD%97%E7%AC%A6%E7%9A%84%E6%9C%80%E9%95%BF%E5%AD%90%E4%B8%B2
# noj-go

def id2titl_en(title="3.无重复字符的最长子串"):
    """
    id 转换为英文题目, eg: 无重复字符的最长子串 --> longest-substring-without-repeating-characters

    """
    refere_url = "https://leetcode.cn/search/?q={}".format(title)
    url = "https://leetcode.cn/graphql/noj-go/"
    headers = {} # headers 可以为空；
    data = {
        "query": "\n    query problemsetQuestions($in: ProblemsetQuestionsInput!) {\n  problemsetQuestions(in: $in) {\n    hasMore\n    questions {\n      titleCn\n      titleSlug\n      title\n      frontendId\n      acRate\n      solutionNum\n      difficulty\n      userQuestionStatus\n    }\n  }\n}\n    ",
        "variables": {"in": {"query": title, "limit": 32, "offset": 0}},
        "operationName": "problemsetQuestions"
    }

    res = requests.post(url, headers, json=data)
    res_json = res.json()
    candidite_questions = res_json["data"]["problemsetQuestions"]["questions"]
    for q in candidite_questions:
        if q['frontendId'] == title.split(".")[0]:
            # q['titleSlug'] # 'longest-substring-without-repeating-characters'
            # {'titleCn': '无重复字符的最长子串', 'titleSlug': 'longest-substring-without-repeating-characters', 'title': 'Longest Substring Without Repeating Characters', 'frontendId': '3', 'acRate': 0.3933544158935547, 'solutionNum': 14657, 'difficulty': 'MEDIUM', 'userQuestionStatus': 'NOT_STARTED'}
            return {
                "titleCN": q['titleCn'],
                "titleSlug": q['titleSlug'],
                "frontendId": q['frontendId'],
                "difficulty": q['difficulty'].lower(),
            }

def printRow(infos, solutions):
    """
    infos: {'titleCN': '无重复字符的最长子串', 'titleSlug': 'longest-substring-without-repeating-characters', 'difficulty': 'medium',}
    solutions: ['go']

    print:
    |[209.长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/) |中等|[Golang](./leetcode_recordes/209.长度最小的子数组.go)| |
    |[{1}]({2})|{3}|{4 可能多个，go, py}| 备注信息，输出结果后，按需备注|
    """
    _1 = infos["frontendId"] + "." + infos['titleCN']
    _2 = "https://leetcode.cn/problems/" + infos["titleSlug"]
    if infos["difficulty"] == "easy":
        _3 = "简单"
    elif infos["difficulty"] == "medium":
        _3 = "中等"
    else:
        _3 = "困难"
    
    solutions_md = []
    for s in solutions:
        completed_file = "./leetcode_recordes/" + infos["frontendId"] + "." + infos["titleCN"] + "." + s
        if s == "go":
            language = "Golang"
        elif s == "py":
            language = "Python"
        else:
            language = s # 暂未支持其他语言
        solutions_md.append(f"[{language}]({completed_file})")
    _4 = ", ".join(solutions_md)

    row = [_1, _2, _3, _4]
    print("|[{}]({})|{}|{}| |".format(row[0], row[1], row[2], row[3]))

if __name__ == "__main__":
    """
    ids: 题号列表；
    title: id.中文题目；
    question: id.题目.py or .go ...
    completed_questions_dict: 可以通过id索引中文题目，也可以通过id.题目 索引题解语言；
    """
    ids = sys.argv[1:]
    if len(ids) == 0:
        sys.exit("Error: 请输入题号!")

    completed_questions = os.listdir("../leetcode_recordes/")
    completed_questions_dict = {}
    for title in completed_questions:
        tmp_ls = title.split(".")

        id = title.split(".")[0]
        completed_questions_dict[id] = tmp_ls[1]
        if (id + "." + tmp_ls[1]) not in completed_questions_dict: # 一些问题有多个语言的解法，后缀是list，eg py, go, cpp...
            completed_questions_dict[id + "." + tmp_ls[1]] = [tmp_ls[2]]
        else:
            completed_questions_dict[id + "." + tmp_ls[1]].append(tmp_ls[2])

    target_titles = []
    for id in ids:
        if id not in completed_questions_dict:
            sys.exit("Error: 题号{} 不存在，或未完成!".format(id))
        title = id + "." + completed_questions_dict[id]
        target_titles.append(title)
    
    print("处理题目：", target_titles)
    print("="*5, "处理中ing", "="*5)

    for title in target_titles:
        infos = id2titl_en(title)
        printRow(infos, solutions=completed_questions_dict[title])
