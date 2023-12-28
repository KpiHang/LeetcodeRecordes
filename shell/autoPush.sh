# 1、cd 到上一层目录 Shell在哪执行，位置就在哪。
# currentDir=$(dirname "$PWD") 
# cd ${currentDir}

# 2、获取当前的分支
currentBranch=$(git symbolic-ref --short HEAD)

pwd
# 4、提交操作
git add .

# 查看文件变更；
git status 

# 3、获取提交的信息
echo "\033[32m——-----请输入提交信息：🙃——-----\033[0m"
read commitInfo

git commit -m "$commitInfo"
echo "\033[32m——-----git 提交完毕🙃—-----\033[0m"

# 挂代理，加速；
export http_proxy=http://127.0.0.1:7890
export https_proxy=$http_proxy

# 5、推送代码
git push -u origin $currentBranch
echo "\033[32m——-----git 推送完毕🙃—-----\033[0m"