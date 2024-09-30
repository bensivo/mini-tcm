# Mini TCM - Overview

This document explains what Mini TCM is (high level), and why it exists.


## What is Test Case Management?

Generally, Test Case Management (TCM) describes the systems and processes used to define test cases for an application, and manage them throughout the software development life cycle.

Why do we need TCM? Well, anyone who's worked on any application of significant size will be able to relate to the following problems:

- To thouroughly test an application, we use a variety of testing frameworks (thing unit testing, e2e testing, load / stress testing, etc.)
- Each of these tools has their own syntax and location for defining test cases themselves.
- Each of these tools usually generates test reports different formats.
- Some test cases, like UAT and manual testing, don't have a specific framework for defining and executing test cases.


These problems are what leads to Test Case Management (TCM), and specifically TCM tools. A TCM tool is a centralized location for defining, tracking, and executing test cases, especially when those test cases are executed using a variety of technologies and techniques. 


## Why Mini TCM?

When I was searching for TCM tools, I found tons of fully-managed, SaaS solutions that were intended for large enterprises and dedicated QA teams. These all had fancy UI's, Enterprise features, and a monthly subscription.

However, I was just a single developer, working on a small personal project. I wanted a TCM tool to manage my test cases, but I didn't want a montly subscription. Additionally, I wanted to keep everything about my project IN my project. Today, we use IaC tools like Terraform to define infrastructure as code, we have javadocs, godocs, jsdocs to define our documentation as code. But curiously, test case management has not adopted this code-first approach.

I wanted to find a TCM solution that gave me all the features of a modern TCM platform, but without a monthly subscription, and just using git as my database. That's Mini TCM.

