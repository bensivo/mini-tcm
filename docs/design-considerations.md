# Mini TCM - Design Considerations

This doc outlines the main factors that will go in to the design of Mini TCM itself.


## Considerations

### 🔍 Test Case Authoring Interface
How will developers and qa engineers actually write out test cases? 

Mini TCM needs to support many different kinds of test cases, with the most complex (in definition) being manual / uat tests. These might require actual screenshots, step-by-step walkthroughs, and detailed, formatted instructions on how to execute them.

Automated tests don't really need a step-by-step instructions, but they do need to be registered / tracked.


Options:
- A web UI. Users could run a local "mini-tcm" server, then edit and browse test cases in their web browser
- Directly in files. If our storage format (below) is 100% human readable, users could directly edit test cases in these files. This would require a storage format that is both machine and human readable.
- An SDK. We could just pick a language, then get users to define test cases using a library built for that language. This makes it easier to enforce the structure of test cases, but may be too technical for a test-case interface. 

### 🔍 Test Case Storage Format
How will test cases be stored? 

One of the main points of Mini TCM is to use the Git repo itself as a database for TCM. This means we have to use a text-first storage format, but one that also supports rich media (for manual test execution steps).

Options:
- Markdown. It's both text-first (and therefore could be reviewed in an MR), and supports media. The downside of MD is it's much less machine-readable. If we build a script for processing the test cases, we'd have to get very clever with how to parse the markdown files as structured content.

- Yaml. It's both human-readable (to a degree), and machine-readable. However, its media support is lacking. 


### 🔍 Test Execution Tracking
How will test executions be tracked?

The traditional process for using a TCM during a regression is as follows:
1. A "test execution" is created, signalling the start of a regression
2. All of the test cases are copied into the test execution
3. Testers use the TCM tool as a tracking device, executing test cases, and labelling them as "Pass" or "Fail" in the TCM.
4. When everything is finished, a report is generated, consolidating the results of all the test cases. Based on this, the team will decide whether the regression overall passed or failed.

This process breaks down somewhat for a code-first testing tool. If the main database for test cases is the git repo itself, then should test reports also be tracked in the git repo?

Options:
- Treat it like automated testing tools today. The test suite just stores the steps for executing tests in code, and reports are generated when tests are executed. This report could in-theory get checked in, but is more frequently pushed as an artifact to the CI build, or uploaded to an MR. 

- Add a "mini-tcm-server", which can store test definitions, and also track executions. This also gives all users a place to centrally track their test executions. However, it would require a big architectural lift, and also requiring a deployed server to use Mini TCM breaks the "Mini" part of it. 



