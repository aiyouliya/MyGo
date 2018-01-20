# MyGo
go study  note
git order
1	初始化git	Git init	
2	添加一个文件	Git add <file name >	
3	提交文件并添加一个注释	Git commit -m "注释名"	
4. 	查看日志	Git log	
5. 	逐行查看日志	Git log --pretty=oneline	
6.	退回上一个版本	Git reset --hard HEAD^	eg:  git reset --hard 3628164(版本号，位数越多越精准)			
7. 	退回后再找回	Git reflog查	查看历史命令，在命令中可以看到版本号信息
8	丢弃工作区的修改	Git checkout -- filename	让file变成上一次commit之前的版本(其实是用版本库里的版本替换工作区的版本，无论工作区是修改还是删除，都可以“一键还原”)
9	回退版本	Git reset	Git reset HEAD file “使用HEAD”表示使用最新版本
10	要关联一个远程库	git remote add origin git@github.com:aiyouliya/MyGo.git	Origin 远端服务器别名，也可以叫别的名字
			aiyouliya ：github中使用的路径名（账号）
			MyGo.git：在github的aiyouliya账号中添加的仓库
11	推送本地内容到远端仓库	git push -u origin master	第一次推送master分支时，加上了-u参数，Git不但会把本地的master分支内容推送的远程新的master分支，还会把本地的master分支和远程的master分支关联起来
12	把本地master分支的最新修改推送至GitHub	git push origin master	
13	从git端克隆到本地仓库	git clone git@github.com:aiyouliya/MyGo.git	
14	创建分支	Git branch dev	dev是创建 的分支名称
15	切换分支	Git checkout dev	切换到dev分支
16	创建并切换分支	Git checkout -b dev	创建并切换到dev分支
17	查看分支	Git branch	查看所有分支及当前所在分支
18	合并某个分支到当前分支	Git merge dev 	合并dev分支到当前分支
19	删除分支	Git branch -d dev	删除本地仓库的dev分支

注释：git解决conflict需要在merge后进行手工合并文件内容，并add该文件至当前版本，然后commit，从而达到所有版本一致。

