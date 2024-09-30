# Mini TCM - MD

This proposed interface to Mini TCM uses MD as the main file format for defining test cases


## Defining Test Cases

Test cases are defined in .tcm.md files in the "tcm" folder. You can use sub-folders, mini-tcm will search for any files that end in ".tcm.md".

Each .tcm.md file corresponds to a single test case. The contents of the file shoudl contain the test case steps, and expected results.

Each .tcm.md file should use yaml front-matter to define Mini TCM specifc information. Other than the front-matter, the md file content is completely free-form.

Example file: `100-concurrent-users.tcm.md`:
```
---
id: TC-1
kind: "jmeter"
name: "100 concurrent users"
frequency: 2
risk: 4
---

## Spec:
The application should be able to handle 100 concurrent requests in less than 2s each

## Steps:
1. Use JMeter to make 100 concurrent requests to the homepage

    ![alt text](https://via.placeholder.com/300?text=test-screenshot-here)

## Expected Result:
- All 100 requets should be served successfully, within 2 seconds

    ![alt text](https://via.placeholder.com/300?text=test-screenshot-here)

- Server should not crash or anything
```


Example file: `app-version.tcm.md`:
```
---
id: TC-2
kind: "jest"
name: "app-version"
frequency: 5
risk: 4
---

## Spec
The /health endpoint should return the app version, pulled from env variables

## Steps
Run via Jest

`cd packages/webapp && npx jest`

```