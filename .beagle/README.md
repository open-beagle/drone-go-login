# version

<!-- https://github.com/drone/go-login -->

```bash
git remote add upstream git@github.com:drone/go-login.git

git fetch upstream

git merge v1.1.0
```



## dev

```bash
# 新建一个Tag
git tag v1.1.0-beagle.0

# 推送一个Tag ，-f 强制更新
git push -f origin v1.1.0-beagle.0

# 删除本地Tag
git tag -d v1.1.0-beagle.0
```

## realse

```bash
# 新建一个Tag
git tag v1.1.0-beagle

# 推送一个Tag ，-f 强制更新
git push -f origin v1.1.0-beagle

# 删除本地Tag
git tag -d v1.1.0-beagle
```
