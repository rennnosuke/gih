# gih
CLI tool for git remote repository hosting service. 

# Usage

# Setup

```
$ gih init
Input access token (current value: [ACCESS_TOKEN]) : xxx
Input repository path (current value: https://github.com/myorg/Hello-OctCat) : xxx
.config.json saved.
Git Hosting Service: GitHub
```

# Issue

## List issues
```
$ gih
ISSUEID       TITLE                DESCRIPTION                STATE    CREATED_AT
1     create .gitignore    Create .gitignore file...  opened
```

## Create Issue

#### Simple

`$ gih -c [TITLE] [DESCRIPTION]`

```
$ gih -c "create .gitignore" "Create .gitignore file for ignore some file."
create issue : 1 
[TITLE]
create .gitignore

[DESCRIPTION]
Create .gitignore file for ignore some file.
```

## Close Issue

`$ gih -d [ISSUE_NUMBER]`

```
$ gih -d 1 
closed issue : 1
[TITLE]
create .gitignore

[DESCRIPTION]
Create .gitignore file for ignore some file.
```

<details>
<summary>TBD functions</summary>

You can write your issue description as markdown in editor.

#### Edit with editor

`$ gih -c`

```
$ gih -c
### show editor (such as `vi`.)
```

###### in editor
```
[TITLE] 
title

[DESCRIPTION]
# Summary
Create .gitignore file for ignore some file.

# Description
Following folders/files should be contained in .gitignore but not yet.
- [] .idea
- [] .env
- [] config.yml
```


## Update Issue

`$ gih -u [ISSUE_NUMBER]`

```
$ gih -u 1
### show editor (such as `vi`.)
```


</details>

# Others

## Show Hosting Web Site

`$ gih -w`

```
$ gih -w
Open https://github.com/myorg/Hello-OctCat...
### go to web site (such as github).
```