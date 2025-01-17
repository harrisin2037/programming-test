# Programming Test


## Question 1

Use the following undirected graph ­- nodes can be visited only once:

![List](assets/web/graph.png)

- a. Write a function that returns all the possible paths between A­-H.
- b. Write a function that returns the least number of hops (shortest path) between A­-H.


## Question 2

### API Specification
**Get list of people**
  * Method
    `GET`
  * Endpoint
    `https://api.json-generator.com/templates/Xp8zvwDP14dJ/data`
  * API Key: `v3srs6i1veetv3b2dolta9shrmttl72vnfzm220z` (Don't worry. It is a public key. If it doesn't works, do let us know)

### Example: Fetch From API by command line
```
 curl --request GET -H "Authorization: Bearer R4iN..." --url https://api.json-generator.com/templates/tAu-9/data
```

### User Requirements
- Retrieve list of people from the API
- Display list of people.
- Show details when user select an item in the list.
- Add marker on the map based on the provided latitude/longitude in `location`. 


### Technical Requirement
- Source code must be stored in a git repository (github or bitbucket)
- Must be a single-page application (SPA)
- Candidates are free to use any libraries
- Expect the code is **production ready**


### Wireframe
*For your reference only, you can have your own design and UX*


![List](assets/mobile/list.png)
![Map](assets/mobile/map.png)
