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