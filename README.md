# Go

## Workspace is important

single workspace with below structure

```sh
export GOPATH='/home/abhishek/work/go'
#optional
export PATH=$PATH:$GOPATH/bin
```

folder structure very opiniated

```yaml
- go(workspace)
    - src
        - github.com
            - project-1
            - project-2
        - gitlab.com
    - pkg # for eg: go get github.com/aws-sdk-go/aws
        - mod
            - github.com
                - aws-sdk-go@v1.39.1
    - bin
```
