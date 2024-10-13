---
id: TC-2
name: "200 concurrent users"
---

## Spec:
The application should be able to handle 200 concurrent requests in less than 2s each

## Steps:
1. Open JMeter
2. Make 100 concurrent requests to the homepage
3. Look at test results

## Expected Result:
- All 100 requets should be served successfully, within 2 seconds

    ![alt text](https://via.placeholder.com/300?text=test-screenshot-here)

- Server should not crash or anything