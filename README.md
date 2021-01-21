# Golang CLI Project **Template**

![](https://res.cloudinary.com/digf90pwi/image/upload/v1581731174/1_8bPiDNL1K1ZdK9O_T5IVKw_xshtjh.png)

The template for golang project

## Setup Project

* [![](https://res.cloudinary.com/digf90pwi/image/upload/c_scale,r_14,w_98/a_0/v1581731363/%E6%8D%95%E8%8E%B7_iqiuwl.png)](https://github.com/Soontao/go-project-template/generate) to create new project 
* Change `LICENSE` if necessary
* Modify [info.repository_url](./chglog) as your own project url
* Update your repo url in `go.mod` file
* Write your application
* Run `go mod vendor` after development (remember `vendor` your deps after add any libraries)

## Setup CI

* Open [Circle CI](https://circleci.com/) and setup your project
* Add `Github Token (GITHUB_TOKEN)` to Circle CI Project Env for Release

## Release Version

* Run `./release.sh VERSION` like `./release.sh v0.0.1` to release a version