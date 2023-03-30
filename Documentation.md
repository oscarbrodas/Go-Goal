## Backend API
### User
- User: the structure of the User object. The variable names are case sensitive

  - ID: unique id of the user
  
  - CreatedAt: *do not worry about this* Time and date the user was created at
  
  - UpdatedAt: *do not worry about this* Time and date the user was updated at
  
  - DeletedAt: *do not worry about this* Time and date the user was deleted at
  
  - Username: username
  
  - FirstName: first name
  
  - LastName: last name
  
  - Email: the email is unique to each user
  
  - Password: password

  - XP: experience points (int)
  
- __CreateUser()__ - __POST__

  - Route: __/api/users__
  
  - Input: the body must contain user object
  
  - Output: "Successful", "ErrorExist", and "EmailExist" each as a bool detailing the result of the function

- __GetUser()__ - __GET__

  - Route: __/api/users?id=__
  
  - Input: no body is required. The URL must contain the id of the user you are trying to get, such as getting user with id 9, you need to use __"/api/users?id=9"__
  
  - Output: "ThisUser" as a user object, and "ErrorExist" as a bool

- __UpdateUsername__ - __PUT__

  - Route: __/api/users/{id}/username__
  
  - Input: The User ID (In the Route) to be updated and the new "Username" of type string
  
  - Output: "User" back with the updated change 

- __UpdateFirstname__ - __PUT__

  - Route: __/api/users/{id}/firstname__
  
  - Input: The User ID (In the Route) to be updated and "Firstname" of type string
  
  - Output: "User" back with the updated change 

- __UpdateLastname__ - __PUT__

  - Route: __/api/users/{id}/lastname__
  
  - Input: The User ID (In the Route) to be updated and "Lastname" of type string
  
  - Output: "User" back with the updated change 
  
- __UpdateEmail__ - __PUT__

  - Route: __/api/users/{id}/email__
  
  - Input: The User ID (In the Route) to be updated and "Email" of type string
  
  - Output: "User" back with the updated change 
  
 
- __UpdatePassword__ - __PUT__

  - Route: __/api/users/{id}password__
  
  - Input: The User ID (In the Route) to be updated and "Password" of type string
  
  - Output: "User" back with the updated change 
  
  - __AddXP__ - __PUT__

  - Route: __/api/users/{id}/xp__
  
  - Input: The User ID (In the Route) to be updated and "xp" of type int(32bit)
  
  - Output: "User" back with the updated change 
  
  
- __CheckLogin()__ - __GET__

  - Route: __/api/login/{email}/{password}__ 

  - Notes: Email and Password must be correct or the user object returned will be giberish

  - Input: "Email" and "Password" of type string passed in the route.

  - Output: "FindEmail" and "FindPassword" of type bool. "ThisUser" which is the user object

- __CheckUsername()__ - __GET__

  - Route: __/api/login/{username}__ 

  - Input: "Username" of type string passed in the route.

  - Output: "Exists" and "ValidName", boolean values to validate the "Username" with.


### Friend
- Friend: the structure of the Friend object. The variable names are case sensitive

  - ID: *do not worry about this* Unique id of the relationship
  
  - CreatedAt: *do not worry about this* Time and date the user was created at
  
  - UpdatedAt: *do not worry about this* Time and date the user was updated at
  
  - DeletedAt: *do not worry about this* Time and date the user was deleted at
  
  - User1: id of User1
  
  - User2: id of User2
  
  - WhoSent: int value. 0 means User1 and User2 are friends, 1 means User1 sent a request to User2, 2 means User2 sent a request to User1
  
- __GetAllFriends()__ - __GET__

  - Route: __/api/friends{id}__
  
  - Input: User ID is passed in the route for the user retrieving friends from.
  
  - Output: "IDs" as an array of ints which represent the IDs of the friends of this user. "ErrorExist" as bool
  
- __SendFriendRequest()__ - __PUT__

  - Notes: cannot sent a friend request if you already sent one, are already friends, or the other person sent you a request
  
  - Route: __/api/friends/sendFriendRequest/{sender}/{reciever}__
  
  - Input: Sender's ID and Reciever's ID passed in the route
  
  - Output: "Successful", "ErrorExist" as bools
  
- __GetOutgoingFriendRequests()__ - __GET__

  - Notes: gets the list of people the user sent a request to
  
  - Route: __/api/friends/getOutgoingFriendRequests/{id}__
  
  - Input: User ID passed in route.
  
  - Output: "IDs" as an array of user ids that you sent a request to. "ErrorExist" as a bool
  
- __GetIngoingFriendRequests()__ - __GET__

  - Notes: gets a list of people that sent the user a request
  
  - Route: __/api/friends/getIngoingFriendRequests/{id}__
  
  - Input: User ID passed in route.
  
  - Output: "IDs" as an array of user ids that sent a request to you. "ErrorExist" as a bool
  
- __AcceptFriendRequest()__ - __PUT__

  - Route: __/api/friends/acceptFriendRequest/{sender}/{reciever}__
  
  - Input: Sender/Reciever Pair of IDs passed in route.

  - Output: "Successful", "ErrorExist" as bools.

- __DeclineFriendRequest()__ - __DELETE__

  - Route: __/api/friends/declineFriendRequest/{sender}/{decliner}__
  
  - Input: Sender/Decliner pair of IDs passed in route.

  - Output: "Successful", "ErrorExist" as bools

- __RemoveFriend()__ - __DELETE__

  - Route: __/api/friends/removeFriend/{remover}/{friend}__
  
  - Input: Remover/Friend pair of IDs passed in route.

  - Output: "Successful", "ErrorExist" as bools

### Goal
- Goal: the structure of the Goal object. The variable names are case sensitive

  - ID: *do not worry about this* Unique id of the relationship
  
  - CreatedAt: *do not worry about this* Time and date the user was created at
  
  - UpdatedAt: *do not worry about this* Time and date the user was updated at
  
  - DeletedAt: *do not worry about this* Time and date the user was deleted at
  
  - Title: title of the goal (String)
  
  - Description: description of the goal (String)
  
  - UserID: id of user who has the goal
 
- Benchmark: the structure of the Benchmark object. The variable names are case sensitive

  - ID: *do not worry about this* Unique id of the relationship
  
  - CreatedAt: *do not worry about this* Time and date the user was created at
  
  - UpdatedAt: *do not worry about this* Time and date the user was updated at
  
  - DeletedAt: *do not worry about this* Time and date the user was deleted at
  
  - Description: description of the benchmark (String)
  
  - Completed: Is the benchmark completed? (Bool)

  - GoalID: ID of the goal the benchmark belongs to.

- __CreateGoal()__ - __POST__

  - Route: __/api/goal/{id}__
  
  - Input: Route holds user id and the body take the goal structure
  
  - Output: "Successful", "ErrorExist", and the created Goal. 

- __GetGoals()__ - __GET__

  - Route: __/api/goals/{id}__
  
  - Input: User ID is passed through the route.
  
  - Output: "Successful", "ErrorExist", and a list of goals is returned.

- __DeleteGoal()__ - __DELETE__

  - Route: __/api/goals/{id}__
  
  - Input: Goal ID is passed through the route.
  
  - Output: "Successful", "ErrorExist" are returned.

- __AddBenchmark()__ - __POST__

  - Route: __/goals/{id}/{goalID}__
  
  - Input: Goal ID is passed through the route. The body of request contains the input benchmark object with only the description required.
  
  - Output: "Successful", "ErrorExist" are returned.

- __GetBenchmarks()__ - __GET__

  - Route: __/goals/benchmarks/{goalID}__
  
  - Input: Goal ID is passed through the route.
  
  - Output: "Successful", "ErrorExist" are returned. "Benchmarks" is an array of benchmark objects

- __UpdateBenchmarkDescription()__ - __PUT__

  - Route: __/goals/benchmarks/description/{benchmarkID}__
  
  - Input: Goal ID is passed through the route.
  
  - Output: "Successful", "ErrorExist" are returned.

- __UpdateBenchmarkCompletion()__ - __PUT__

  - Route: __/goals/benchmarks/completion/{benchmarkID}__
  
  - Input: Goal ID is passed through the route.
  
  - Output: "Successful", "ErrorExist" are returned.

- __DeleteBenchmark()__ - __DELETE__

  - Route: __/goals/benchmarks/{benchmarkID}__
  
  - Input: Goal ID is passed through the route.
  
  - Output: "Successful", "ErrorExist" are returned.
