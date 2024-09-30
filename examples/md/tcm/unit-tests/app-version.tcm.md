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

```
cd packages/webapp
npx jest
```
